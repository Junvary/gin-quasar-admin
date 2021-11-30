import { boot } from 'quasar/wrappers'
import axios from 'axios'
import { Notify, Dialog } from 'quasar'
// import { GetToken } from 'src/utils/cookies'

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)

const api = axios.create(
    {
        baseURL: process.env.API,
        timeout: 40000,
        withCredentials: false
    }
)

export default boot(({ app, router, store }) => {
    // 请求拦截
    api.interceptors.request.use(res => {
        const token = store.getters['user/token']
        res.headers = {
            'Content-Type': 'application/json;charset=utf-8',
            'Gqa-Token': token,
        }
        return res
    }, error => {
        Notify.create({
            type: 'negative',
            message: error,
        })
        return Promise.reject(error)
    })
    // 响应拦截
    api.interceptors.response.use(response => {
        const responseData = response.data
        const { code } = responseData
        if (code === 1) {
            return response.data
        } else {
            switch (code) {
                case 0:
                    if (responseData.data && responseData.data.reload) {
                        Dialog.create({
                            title: "身份鉴别失败！",
                            message: response.data.message || '你的身份鉴别已过期，请退出系统重新登录！',
                            persistent: true,
                            ok: {
                                push: true,
                                color: 'negative',
                                label: "重新登录"
                            },
                        }).onOk(() => {
                            store.dispatch('user/HandleLogout')
                            router.push({ name: 'login' })
                        })
                    } else {
                        Notify.create({
                            type: 'negative',
                            message: response.data.message || '操作失败！',
                        })
                        return response.data
                    }
                default:
                    Notify.create({
                        type: 'negative',
                        message: response.data.message || '操作失败！',
                    })
                    return response.data
            }
        }
    }, error => {
        // 500的情况，比如后台是初始化的，但前台还有token，多出现在开发时
        if (error + '' === 'Error: Request failed with status code 500') {
            Dialog.create({
                title: "抱歉！",
                message: '数据异常，请退出系统重新登录！',
                persistent: true,
                ok: {
                    push: true,
                    color: 'negative',
                    label: "重新登录"
                },
            }).onOk(() => {
                store.dispatch('user/HandleLogout')
                router.push({ name: 'login' })
            })
        }
        // 超时
        if (error + '' === 'Error: timeout of 15000ms exceeded') {
            Notify.create({
                type: 'negative',
                message: '后台响应超时！',
            })
        }
        // 网络错误情况，比如后台没有对应的接口
        if (error + '' === 'Error: Network Error') {
            router.push({ name: 'notFound' })
        } else if (error.response && error.response.status === 404) {
            console.log('请求地址不存在 [' + error.response.request.responseURL + ']')
            // Notify.create({
            //     type: 'negative',
            //     message: '请求地址不存在 [' + error.response.request.responseURL + ']',
            // })
        }

        return Promise.reject(error)
    })
    // for use inside Vue files (Options API) through this.$axios and this.$api

    app.config.globalProperties.$axios = axios
    // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
    //       so you won't necessarily have to import axios in each vue file

    app.config.globalProperties.$api = api
    // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
    //       so you can easily perform requests against your app's API
})

export { api }

import { boot } from 'quasar/wrappers'
import axios from 'axios'
import { Notify, Cookies } from 'quasar'
import Store from 'src/store'
import Router from 'src/router'

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)
const api = axios.create(
    {
        baseURL: 'http://localhost:8080',
        timeout: 15000
    }
)

// 请求拦截
api.interceptors.request.use(res => {
    const token = Cookies.get('gqa-token')
    res.headers = {
        'Content-Type': 'application/json',
        'gqa-token': token,
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
            case 401:
                postAction(refreshTokenUrl, {
                    refresh: Cookies.get('gqa-token')
                }).then(res => {
                    Cookies.set('gqa-token', res.access)
                    // window.location.reload()
                }).catch(error => {
                    Cookies.remove('gqa-token')
                    Notify.create({
                        type: 'negative',
                        message: '获取数据错误，请重新登录',
                    })
                    Router.push({ name: 'login' })
                })
            case 400:
                Notify.create({
                    type: 'negative',
                    message: response.data.message || '数据异常！',
                })
                return response.data
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
        Notify.create({
            type: 'negative',
            message: '数据异常，自动退出登录！',
        })
        Store().dispatch('user/HandleLogout')
        Router.push({ name: 'login' })
    }
    // 网络错误情况，比如后台没有对应的接口
    if (error + '' === 'Error: Network Error') {
        Router.push({ name: 'notFound' })
    } else if (error.response && error.response.status === 404) {
        Notify.create({
            type: 'negative',
            message: '请求地址不存在 [' + error.response.request.responseURL + ']',
        })
    }

    return Promise.reject(error)
})

export default boot(({ app, router }) => {
    // for use inside Vue files (Options API) through this.$axios and this.$api

    app.config.globalProperties.$axios = axios
    // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
    //       so you won't necessarily have to import axios in each vue file

    app.config.globalProperties.$api = api
    // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
    //       so you can easily perform requests against your app's API
})

export { api }

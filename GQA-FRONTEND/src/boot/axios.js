import { boot } from 'quasar/wrappers'
import axios from 'axios'
import { Notify, Cookies } from 'quasar'

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
                    router.push({ name: 'login' })
                })
            case 400:
                Notify.create({
                    type: 'negative',
                    message: response.data.message || '数据异常',
                })
                return response.data
            default:
                Notify.create({
                    type: 'negative',
                    message: response.data.message || '数据异常',
                })
                return response.data
        }
    }
}, error => {
    if (error + '' === 'Error: Network Error') {
        router.push({ name: 'notFound' })
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

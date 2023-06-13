import { boot } from 'quasar/wrappers'
import axios from 'axios'
import { Notify, Dialog, Cookies } from 'quasar'
import { i18n } from './i18n'
import { useUserStore } from 'src/stores/user'
import { DemoMode, ForbiddenUrl } from 'src/config/config'

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)

const api = axios.create({
    baseURL: process.env.API,
    timeout: 120000,
    withCredentials: false
})

export default boot(({ app, router }) => {
    const userStore = useUserStore()

    api.interceptors.request.use(request => {
        const token = userStore.GetToken()
        request.headers = {
            'Content-Type': 'application/json;charset=utf-8',
            'Gqa-Token': token,
            'Gqa-Lang': Cookies.get("gqa-language") || "zh-CN"
        }
        if (DemoMode && ForbiddenUrl.some(item => item === request.url)) {
            Notify.create({
                type: 'negative',
                message: i18n.global.t('DemoMode')
            })
            return
        }
        return request
    }, error => {
        Notify.create({
            type: 'negative',
            message: error,
        })
        return Promise.reject(error)
    })

    api.interceptors.response.use(response => {
        // If the ExpiresAt of the JWT has expired,
        // but the RefreshAt has not expired, 
        // the background will insert a Gqa Refresh Token in the headers, 
        // which will be saved here to form a token replacement logic
        if (response.headers['gqa-refresh-token'] && response.data.data.refresh) {
            userStore.SetToken(response.headers['gqa-refresh-token'])
            // store.dispatch('user/SetToken', response.headers['gqa-refresh-token'])
            Notify.create({
                type: 'positive',
                message: i18n.global.t('Refresh') + 'Token' + i18n.global.t('Success'),
            })
            return api(response.config)
        }
        const responseData = response.data
        const { code } = responseData
        if (code === 1) {
            return response.data
        } else {
            switch (code) {
                case 0:
                    if (responseData.data && responseData.data.reload) {
                        Dialog.create({
                            title: i18n.global.t('Authentication') + i18n.global.t('Failed'),
                            message: response.data.message || i18n.global.t('Please') + i18n.global.t('Relogin'),
                            persistent: true,
                            ok: {
                                push: true,
                                color: 'negative',
                                label: i18n.global.t('Relogin')
                            },
                        }).onOk(() => {
                            userStore.HandleLogout()
                            router.push({ name: 'login' })
                        })
                    } else {
                        Notify.create({
                            type: 'negative',
                            message: response.data.message || i18n.global.t('Operation') + i18n.global.t('Failed'),
                        })
                        return response.data
                    }
                default:
                    return response.data
            }
        }
    }, error => {
        // 500
        if (error + '' === 'Error: Request failed with status code 500') {
            Dialog.create({
                title: i18n.global.t('Error'),
                message: i18n.global.t('Data') + i18n.global.t('Exception') + ',' + i18n.global.t('Please') + i18n.global.t('Relogin'),
                persistent: true,
                ok: {
                    push: true,
                    color: 'negative',
                    label: i18n.global.t('Logout')
                },
            }).onOk(() => {
                userStore.HandleLogout()
                router.push({ name: 'login' })
            })
        }
        // timeout
        if (error + '' === 'Error: timeout of 40000ms exceeded') {
            Notify.create({
                type: 'negative',
                message: i18n.global.t('Operation') + i18n.global.t('Timeout')
            })
        }
        // network error
        if (error + '' === 'Error: Network Error') {
            router.push({ name: 'notFound' })
        } else if (error.response && error.response.status === 404) {
            Notify.create({
                type: 'negative',
                message: i18n.global.t('Request') + i18n.global.t('Address') + i18n.global.t('NotFound') + ' ' + error.response.request.responseURL,
            })
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
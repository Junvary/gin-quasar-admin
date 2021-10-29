import { boot } from 'quasar/wrappers'
import { LoadingBar, Loading, QSpinnerGears, Cookies } from 'quasar'
import { Allowlist } from 'src/settings'

LoadingBar.setDefaults({
    color: 'red',
    size: '4px',
    position: 'top'
})

function startLoading() {
    Loading.show({
        // spinner: QSpinnerGears,
        message: '系统努力加载中...'
    })
    LoadingBar.start()
}

function stopLoading() {
    Loading.hide()
    LoadingBar.stop()
}

export default boot(({ router, store }) => {
    router.beforeEach((to, from, next) => {
        startLoading()
        if (Cookies.get('gqa-token')) {
            if (Allowlist.indexOf(to.path) !== -1) {
                next({ path: '/' })
                stopLoading()
            } else {
                if (!store.getters['permission/userMenu'].length) {
                    store.dispatch('permission/GetUserMenu').then(res => {
                        // 在vue-router4中，addRoutes被废弃，改为了addRoute，循环调用
                        // 动态添加鉴权路由表
                        if (res) {
                            res.forEach(item => {
                                router.addRoute(item)
                            })
                            next({ ...to, replace: true })
                        } else {
                            store.dispatch('user/HandleLogout')
                            next({ path: '/', replace: true })
                        }
                    })
                } else {
                    next()
                }
                stopLoading()
            }
        } else {
            if (Allowlist.indexOf(to.path) !== -1) {
                next()
                stopLoading()
            } else {
                next(`/login?redirect=${to.fullPath}`)
                stopLoading()
            }
        }
    })
    router.afterEach(() => {
        stopLoading()
    })
})
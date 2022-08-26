import { boot } from 'quasar/wrappers'
import { LoadingBar, Loading, QSpinnerGears } from 'quasar'
// import { Allowlist } from 'src/settings'
// import { GetToken } from 'src/utils/cookies'
import { useUserStore } from 'src/stores/user'
import { usePermissionStore } from 'src/stores/permission'
import useCommon from 'src/composables/useCommon'
import { i18n } from './i18n'

LoadingBar.setDefaults({
    color: 'red',
    size: '4px',
    position: 'top'
})

function startLoading() {
    Loading.show({
        // spinner: QSpinnerGears,
        message: i18n.global.t('System') + i18n.global.t('Loading')
    })
    LoadingBar.start()
}

function stopLoading() {
    Loading.hide()
    LoadingBar.stop()
}

export default boot(({ router, store }) => {
    router.beforeEach((to, from, next) => {
        const userStore = useUserStore()
        const permissionStore = usePermissionStore()
        startLoading()
        const token = userStore.GetToken()
        const { AllowList } = useCommon()
        if (token) {
            if (AllowList.indexOf(to.path) !== -1) {
                next({ path: '/' })
                stopLoading()
            } else {
                if (!permissionStore.userMenu.length) {
                    permissionStore.GetUserMenu().then(res => {
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
            if (AllowList.indexOf(to.path) !== -1) {
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
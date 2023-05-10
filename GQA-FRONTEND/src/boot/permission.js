import { boot } from 'quasar/wrappers'
import { LoadingBar, Loading, QSpinnerGears } from 'quasar'
import { useUserStore } from 'src/stores/user'
import { usePermissionStore } from 'src/stores/permission'
import useConfig from 'src/composables/useConfig'

import { i18n } from './i18n'

LoadingBar.setDefaults({
    color: 'red',
    size: '4px',
    position: 'top'
})

function startLoading() {
    Loading.show({
        spinner: QSpinnerGears,
        message: i18n.global.t('System') + i18n.global.t('Loading')
    })
    LoadingBar.start()
}

function stopLoading() {
    Loading.hide()
    LoadingBar.stop()
}

export default boot(({ router }) => {
    router.beforeEach(async (to, from, next) => {
        const userStore = useUserStore()
        const permissionStore = usePermissionStore()
        startLoading()
        const token = userStore.GetToken()
        const { AllowList } = useConfig()
        if (token) {
            if (AllowList.indexOf(to.path) !== -1) {
                next({ path: '/' })
                stopLoading()
            } else {
                if (!permissionStore.userMenu.length) {
                    const res = await permissionStore.GetUserMenu()
                    if (res && res.length) {
                        res.forEach(item => {
                            router.addRoute(item)
                        })
                        next({ ...to, replace: true })
                    } else {
                        stopLoading()
                        userStore.HandleLogout()
                        next({ path: '/', replace: true })
                    }
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
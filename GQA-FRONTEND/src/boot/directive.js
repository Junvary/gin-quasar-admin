import { boot } from 'quasar/wrappers'
import { usePermissionStore } from 'src/stores/permission'

export default boot(({ app }) => {
    const permissionStore = usePermissionStore()
    app.directive('has', {
        mounted(el, binding, vnode) {
            const hasPermission = permissionStore.userButton.some(item => item === binding.value)
            if (!hasPermission) {
                el.parentNode.removeChild(el)
            }
        }
    })
})
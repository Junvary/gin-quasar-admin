import { boot } from 'quasar/wrappers'
import { version } from 'vue'

export default boot(({ app }) => {
    app.config.globalProperties.$vueVersion = version
})
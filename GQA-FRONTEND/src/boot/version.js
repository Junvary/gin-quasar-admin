import { boot } from 'quasar/wrappers'
import { version } from 'vue'
import { Quasar } from 'quasar'

export default boot(({ app }) => {
    app.config.globalProperties.$vueVersion = version
    app.config.globalProperties.$quasarVersion = Quasar.version
})
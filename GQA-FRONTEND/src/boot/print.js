import { boot } from 'quasar/wrappers'
import print from 'vue3-print-nb'

export default boot(({ app }) => {
    app.use(print)
})
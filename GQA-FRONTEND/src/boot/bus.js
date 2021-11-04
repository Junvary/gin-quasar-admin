import { boot } from 'quasar/wrappers'
import mitt from 'mitt'



export default boot(({ app }) => {
    app.config.globalProperties.$bus = new mitt()
})

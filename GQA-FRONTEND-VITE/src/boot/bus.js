import { boot } from 'quasar/wrappers'
import mitt from 'mitt'

export const emitter = mitt()



export default boot(({ app }) => {
    app.config.globalProperties.$bus = new mitt()
})

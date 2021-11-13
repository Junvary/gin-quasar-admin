import { createI18n } from 'vue-i18n'
import messages from 'src/i18n'

export default ({ app }) => {
    // Create I18n instance
    const i18n = createI18n({
        locale: 'en-US',
        messages
    })

    // Tell app to use the I18n instance
    app.use(i18n)
}
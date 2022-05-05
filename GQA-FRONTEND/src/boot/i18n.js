import { createI18n } from 'vue-i18n'
import messages from 'src/i18n'
import { useUserStore } from 'src/stores/user'

export const userStore = useUserStore()

export const i18n = createI18n({
    locale: userStore.GetLanguage(),
    fallbackLocale: 'zh-CN',
    messages,
    silentTranslationWarn: true,
    silentFallbackWarn: true
})

export default ({ app, store }) => {
    // Tell app to use the I18n instance
    app.use(i18n)
}
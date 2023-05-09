import { createI18n } from 'vue-i18n'
import { Translator } from 'src/i18n'
import { useSettingStore } from 'src/stores/setting'

export const settingStore = useSettingStore()

export const i18n = createI18n({
    legacy: false,
    locale: settingStore.GetLanguage(),
    fallbackLocale: 'zh-CN',
    messages: Translator(),
    // silentTranslationWarn: true,
    // silentFallbackWarn: true,
    globalInjection: true,
    missingWarn: false,
    fallbackWarn: false,
})

export default ({ app, store }) => {
    // Tell app to use the I18n instance
    app.use(i18n)
}
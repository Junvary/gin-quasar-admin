import { computed, onMounted, watch } from 'vue';
import { useStorageStore } from 'src/stores/storage';
import { useSettingStore } from 'src/stores/setting';
// import { GqaFrontendDefault } from 'src/settings'
import { useI18n } from 'vue-i18n';
import { Quasar } from 'quasar'
import useCommon from 'src/composables/useCommon';

// 动态更改网站标题和favicon，在MainLayout中调用
export default function useDocument() {
    const storageStore = useStorageStore();
    const settingStore = useSettingStore();
    const gqaFrontend = computed(() => storageStore.GetGqaFrontend());
    const language = computed(() => settingStore.GetLanguage());
    const { GqaFrontendDefault } = useCommon()

    watch(gqaFrontend, (newValue) => {
        document.title = newValue.subTitle
        createLink()
    })

    onMounted(async () => {
        document.title = gqaFrontend.value.subTitle || GqaFrontendDefault.subTitle
        createLink()
        const { locale } = useI18n({ useScope: 'global' })
        locale.value = language.value
        await import(
            /* webpackInclude: /(ru|zh-CN|en-US)\.js$/ */
            'quasar/lang/' + language.value
        ).then((lang) => {
            // ! NOTICE ssrContext param:
            Quasar.lang.set(lang.default)
        })
    })
    const createLink = () => {
        const toDelete = document.getElementsByName('gqa-link-href')
        if (toDelete && toDelete.length) {
            document.getElementsByTagName('head')[0].removeChild(toDelete[0])
        }
        const gqaLink = document.createElement('link')
        gqaLink.type = 'image/ico'
        gqaLink.rel = 'icon'
        gqaLink.setAttribute('name', 'gqa-link-href')
        if (gqaFrontend.value.favicon && gqaFrontend.value.favicon !== '') {
            const favicon = process.env.API + gqaFrontend.value.favicon.substring(11)
            gqaLink.href = favicon
        } else {
            gqaLink.href = 'favicon.ico'
        }
        document.getElementsByTagName('head')[0].appendChild(gqaLink)
    }
    return {
        createLink
    }
}

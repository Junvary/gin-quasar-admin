import { computed, onMounted, watch } from 'vue';
import { useStorageStore } from 'src/stores/storage';
import { useUserStore } from 'src/stores/user';
// import { GqaFrontendDefault } from 'src/settings'
import { useI18n } from 'vue-i18n';
import { Quasar } from 'quasar'
import useCommon from 'src/composables/useCommon';

// 动态更改网站标题和favicon，在MainLayout中调用
export default function useDocument() {
    const storageStore = useStorageStore();
    const userStore = useUserStore();
    const gqaFrontend = computed(() => storageStore.GetGqaFrontend());
    const language = computed(() => userStore.GetLanguage());
    const { GqaFrontendDefault } = useCommon()
    const langList = import.meta.glob('../../node_modules/quasar/lang/*.mjs')

    watch(gqaFrontend, (newValue) => {
        document.title = newValue.subTitle
        createLink()
    })

    onMounted(async () => {
        document.title = gqaFrontend.value.subTitle || GqaFrontendDefault.subTitle
        createLink()
        const { locale } = useI18n({ useScope: 'global' })
        locale.value = language.value

        try {
            await langList[`../../node_modules/quasar/lang/${language.value}.mjs`]().then((lang) => {
                Quasar.lang.set(lang.default)
            })
        } catch (err) {
            console.log(err)
        }
        // await import(
        //     /* webpackInclude: /(zh-CN|en-US)\.js$/ */
        //     'quasar/lang/' + language.value
        // ).then((lang) => {
        //     // ! NOTICE ssrContext param:
        //     Quasar.lang.set(lang.default)
        // })
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
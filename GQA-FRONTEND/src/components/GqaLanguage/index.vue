<template>
    <q-select v-model="lang" :options="langOptions" :label="$t('SwitchLanguage')" dense borderless emit-value
        map-options options-dense @update:model-value="changeLang" style="width: 100%" />
</template>

<script>
import { useQuasar } from 'quasar'
import languages from 'quasar/lang/index.json'
import { ref, watch, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { useStore } from 'vuex'

const appLanguages = languages.filter((lang) => ['zh-CN', 'en-US'].includes(lang.isoName))

const langOptions = appLanguages.map((lang) => ({
    label: lang.nativeName,
    value: lang.isoName,
}))

export default {
    setup() {
        const $q = useQuasar()
        const lang = ref($q.lang.isoName)
        const { locale } = useI18n({ useScope: 'global' })
        const store = useStore()
        onMounted(() => {
            lang.value = store.getters['user/language']
        })

        watch(lang, (val) => {
            // dynamic import, so loading on demand only
            import(
                /* webpackInclude: /(zh-CN|en-US)\.js$/ */
                'quasar/lang/' + val
            ).then((lang) => {
                $q.lang.set(lang.default)
                locale.value = val
                store.dispatch('user/ChangeLanguage', val)
            })
        })

        return {
            lang,
            langOptions,
            changeLang(val) {
                location.reload()
            },
        }
    },
}
</script>
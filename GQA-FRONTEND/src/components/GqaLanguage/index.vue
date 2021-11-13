<template>
    <q-select v-model="lang" :options="langOptions" :label="$t('ComponentSelectLanguage')" dense borderless emit-value map-options options-dense
        style="min-width: 150px" />
</template>

<script>
import { useQuasar } from 'quasar'
import languages from 'quasar/lang/index.json'
import { ref, watch } from 'vue'

const appLanguages = languages.filter((lang) => ['de', 'en-US', 'zh-CN'].includes(lang.isoName))

const langOptions = appLanguages.map((lang) => ({
    label: lang.nativeName,
    value: lang.isoName,
}))

export default {
    setup() {
        const $q = useQuasar()
        const lang = ref($q.lang.isoName)

        watch(lang, (val) => {
            // dynamic import, so loading on demand only
            import(
                /* webpackInclude: /(de|en-US)\.js$/ */
                'quasar/lang/' + val
            ).then((lang) => {
                $q.lang.set(lang.default)
            })
        })

        return {
            lang,
            langOptions,
        }
    },
}
</script>

<template>
    <q-select v-model="lang" :options="langOptions" label="Languages" dense emit-value map-options options-dense
        @update:model-value="changeLang" style="width: 100%" />
</template>

<script setup>
import { useQuasar } from 'quasar'
import languages from 'quasar/lang/index.json'
import { ref, watch, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { useSettingStore } from 'src/stores/setting'

const settingStore = useSettingStore()
const appLanguages = languages.filter((lang) => ['ru', 'zh-CN', 'en-US'].includes(lang.isoName))

const langOptions = appLanguages.map((lang) => ({
    label: lang.nativeName,
    value: lang.isoName,
}))
const $q = useQuasar()
const lang = ref($q.lang.isoName)
const { locale } = useI18n({ useScope: 'global' })

onMounted(() => {
    lang.value = settingStore.GetLanguage()
})
watch(lang, (val) => {
    // dynamic import, so loading on demand only
    import(
        /* webpackInclude: /(zh-CN|en-US|ru)\.js$/ */
        'quasar/lang/' + val
    ).then((lang) => {
        $q.lang.set(lang.default)
        locale.value = val
        settingStore.ChangeLanguage(val)
    })
})
const changeLang = () => { }
</script>

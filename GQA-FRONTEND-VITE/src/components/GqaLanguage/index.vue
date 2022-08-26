<template>
    <q-select v-model="lang" :options="langOptions" :label="$t('Switch') + $t('Language')" dense borderless emit-value
        map-options options-dense @update:model-value="changeLang" style="width: 100%" />
</template>

<script setup>
import { useQuasar } from 'quasar'
import languages from 'quasar/lang/index.json'
import { ref, watch, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { useUserStore } from 'src/stores/user'

const userStore = useUserStore()
const appLanguages = languages.filter((lang) => ['zh-CN', 'en-US'].includes(lang.isoName))
const langList = import.meta.glob('../../../node_modules/quasar/lang/*.mjs')

const langOptions = appLanguages.map((lang) => ({
    label: lang.nativeName,
    value: lang.isoName,
}))
const $q = useQuasar()
const lang = ref($q.lang.isoName)
const { locale } = useI18n({ useScope: 'global' })

onMounted(() => {
    lang.value = userStore.GetLanguage()
})
watch(lang, (val) => {
    try {
        langList[`../../../node_modules/quasar/lang/${val}.mjs`]().then((lang) => {
            $q.lang.set(lang.default)
            locale.value = val
            userStore.ChangeLanguage(val)
        })
    } catch (err) {
        console.log(err)
    }
})
const changeLang = () => {}
</script>

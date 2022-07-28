<template>
    <q-layout style="overflow-x: hidden">
        <q-page-container>
            <component v-if="loginLayoutReady" :is="loginLayout" :pluginComponent="pluginComponent"
                :pluginCurrent="pluginCurrent" />
        </q-page-container>
    </q-layout>
</template>

<script setup>
import useCommon from 'src/composables/useCommon'
import { computed, onBeforeMount, ref, markRaw, defineAsyncComponent } from 'vue';
import indexSimple from './indexSimple.vue';
import indexComplex from './indexComplex.vue';
import { postAction } from 'src/api/manage'
import { useStorageStore } from 'src/stores/storage'
import { useRouter } from 'vue-router';
import { useQuasar } from 'quasar';
import { useI18n } from 'vue-i18n';
import useDocument from 'src/composables/useDocument'

useDocument();
const $q = useQuasar();
const router = useRouter();
const { t } = useI18n();
const { gqaFrontend } = useCommon()
const storageStore = useStorageStore()
const pluginCurrent = ref(null);
const pluginComponent = ref(null);

const loginLayout = computed(() => {
    if (gqaFrontend.value.loginLayoutStyle && gqaFrontend.value.loginLayoutStyle === 'simple') {
        return indexSimple
    } else if (gqaFrontend.value.loginLayoutStyle && gqaFrontend.value.loginLayoutStyle === 'complex') {
        return indexComplex
    } else {
        return indexSimple
    }
})

const loginLayoutReady = ref(false);

onBeforeMount(() => {
    checkDb()
    console.info('欢迎使用Gin-Quasar-Admin!')
    console.info('项目地址: https://github.com/Junvary/gin-quasar-admin ')
    console.info('欢迎交流, 感谢Star!')
})

const checkDb = async () => {
    const res = await postAction('public/check-db')
    if (res.code === 1) {
        storageStore.SetGqaGoVersion(res.data.go_version)
        storageStore.SetGqaGinVersion(res.data.gin_version)
        storageStore.SetGqaPluginList(res.data.plugin_list)
        if (res.data.need_init === false) {
            loginLayoutReady.value = true
            await storageStore.SetGqaDict()
            await storageStore.SetGqaFrontend()
            await storageStore.SetGqaBackend()
            pluginCurrent.value = res.data.plugin_login_layout
            if (pluginCurrent.value) {
                try {
                    const pluginCode = pluginCurrent.value.slice(7)
                    pluginComponent.value = markRaw(defineAsyncComponent(() => import(`src/plugins/${pluginCode}/LoginLayout/index.vue`)))
                } catch (error) {
                    $q.notify({
                        type: 'negative',
                        message: t('LoginLayoutNotSupport'),
                    })
                }
            }
        }
        if (res.data.need_init === true) {
            loginLayoutReady.value = false
            $q.notify({
                type: 'warning',
                message: t('Database') + t('Need') + t('Init'),
            })
            router.push({ path: '/init-db' })
        }
    }
}

</script>

<style lang="scss" scoped>
</style>
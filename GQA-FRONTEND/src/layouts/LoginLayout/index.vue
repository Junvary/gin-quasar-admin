<template>
    <q-layout style="overflow-x: hidden">
        <q-page-container>
            <component :is="loginLayout" :pluginComponent="pluginComponent" :pluginCurrent="pluginCurrent" />
        </q-page-container>
    </q-layout>
</template>

<script setup>
import useFrontendBackend from 'src/composables/useFrontendBackend'
import { computed, onBeforeMount, ref, markRaw, defineAsyncComponent } from 'vue';
import indexSimple from './indexSimple.vue';
import indexComplex from './indexComplex.vue';
import { postAction } from 'src/api/manage'
import { useStorageStore } from 'src/stores/storage'
import { useRouter } from 'vue-router';
import { useQuasar } from 'quasar';
import { useI18n } from 'vue-i18n';

const $q = useQuasar();
const router = useRouter();
const { t } = useI18n();
const { gqaFrontend } = useFrontendBackend()
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

onBeforeMount(() => {
    checkDb()
    console.info('欢迎使用Gin-Quasar-Admin!')
    console.info('项目地址: https://github.com/Junvary/gin-quasar-admin ')
    console.info('欢迎交流, 感谢Star!')
})

const checkDb = async () => {
    const res = await postAction('public/check-db')
    if (res.code === 1) {
        storageStore.SetGqaGoVersion(res.data.goVersion)
        storageStore.SetGqaGinVersion(res.data.ginVersion)
        storageStore.SetGqaPluginList(res.data.pluginList)
        if (res.data.needInit === false) {
            await storageStore.SetGqaDict()
            await storageStore.SetGqaFrontend()
            await storageStore.SetGqaBackend()
            pluginCurrent.value = res.data.pluginLoginLayout
            if (pluginCurrent.value) {
                try {
                    pluginComponent.value = markRaw(defineAsyncComponent(() => import(`src/layouts/LoginLayout/${this.pluginCurrent}/index.vue`)))
                } catch (error) {
                    $q.notify({
                        type: 'negative',
                        message: t('LoginLayoutWizardInitDBAfterText2'),
                    })
                }
            }
        }
        if (res.data.needInit === true) {
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
<template>
    <q-layout style="overflow-x: hidden">
        <q-page-container>
            <div class="gqa-login-layout-img-container" v-if="dbNeedInit">
                <InitDbView @initDbSuccess="checkDb" />
            </div>
            <component v-else :is="loginLayout" :pluginComponent="pluginComponent" :pluginCurrent="pluginCurrent"
                :dbNeedInit="dbNeedInit" @initDbSuccess="checkDb" />
        </q-page-container>
    </q-layout>
</template>

<script setup>
import useCommon from 'src/composables/useCommon'
import { computed, onBeforeMount, ref, markRaw, defineAsyncComponent } from 'vue'
import SimpleView from './SimpleView/index.vue'
import ComplexView from './ComplexView/index.vue'
import InitDbView from './InitDbView/index.vue'
import { postAction } from 'src/api/manage'
import { useStorageStore } from 'src/stores/storage'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import useDocument from 'src/composables/useDocument'

useDocument()
const $q = useQuasar()
const { t } = useI18n()
const { gqaFrontend, gqaLogo } = useCommon()
const storageStore = useStorageStore()
const pluginCurrent = ref(null)
const pluginComponent = ref(null)
const dbNeedInit = ref(true)

const loginLayout = computed(() => {
    if (gqaFrontend.value.loginLayoutStyle && gqaFrontend.value.loginLayoutStyle === 'displayStyle_simple') {
        return SimpleView
    } else if (gqaFrontend.value.loginLayoutStyle && gqaFrontend.value.loginLayoutStyle === 'displayStyle_complex') {
        return ComplexView
    } else {
        return SimpleView
    }
})

onBeforeMount(() => {
    checkDb()
    gqaLogo()
})

const checkDb = async () => {
    dbNeedInit.value = true
    const res = await postAction('public/check-db')
    if (res.code === 1) {
        storageStore.SetGqaGoVersion(res.data.go_version)
        storageStore.SetGqaGinVersion(res.data.gin_version)
        storageStore.SetGqaPluginList(res.data.plugin_list)
        if (res.data.need_init === false) {
            dbNeedInit.value = false
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
            dbNeedInit.value = true
            $q.notify({
                type: 'warning',
                message: t('Database') + t('Need') + t('Init'),
            })
        }
    }
}
</script>

<template>
    <q-layout style="overflow-x: hidden">
        <q-page-container>
            <page-banner :checkDbStatus="!dbOkStatus" ref="pageBanner" />
            <!-- <GqaScrollDown id="login-layout-scroll-down" class="login-layout-scroll-down"
                v-if="pluginCurrent && pluginComponent && showScrollDown" @click="scrollDown" /> -->
            <component v-if="pluginCurrent && pluginComponent" :key="pluginCurrent" :is="pluginComponent" />
            <q-card v-else class="row items-center justify-center" style="padding: 20px 0;">
                <q-btn color="primary" :label="$t('LoginLayoutWithoutPlugin')"></q-btn>
            </q-card>
            <page-footer style="z-index: 0" />
            <q-page-scroller position="bottom-right" :scroll-offset="150" :offset="[3, 18]">
                <q-btn dense fab push icon="keyboard_arrow_up" color="primary" />
            </q-page-scroller>
        </q-page-container>
        <InitDbDialog ref="initDbDialog" @dbOk="dbOk" />
    </q-layout>
</template>

<script setup>
import { ref, onMounted, onUnmounted, markRaw, defineAsyncComponent } from 'vue';
import { postAction } from 'src/api/manage'
import PageBanner from './PageBanner.vue'
import PageFooter from './PageFooter.vue'
import InitDbDialog from './InitDbDialog.vue'
import GqaScrollDown from 'src/components/GqaScrollDown/index.vue'
import { useStorageStore } from 'src/stores/storage'
import { useQuasar } from 'quasar';
import { useI18n } from 'vue-i18n';

const $q = useQuasar();
const dbOkStatus = ref(false);
const pluginCurrent = ref(null);
const pluginComponent = ref(null);
const showScrollDown = ref(true)
const storageStore = useStorageStore()
const { t } = useI18n();
const initDbDialog = ref(null);
const pageBanner = ref(null);

onMounted(() => {
    checkDb()
    window.addEventListener('scroll', documentTop())
})
onUnmounted(() => {
    window.removeEventListener('scroll', documentTop())
})

const scrollDown = () => {
    document.getElementById('login-layout-scroll-down').nextElementSibling.scrollIntoView({
        behavior: 'smooth',
    })
}

const dbOk = (params) => {
    dbOkStatus.value = params
    pageBanner.value.showLoginForm()
}

const documentTop = () => {
    let top = document.documentElement.scrollTop
    if (top > 200) {
        showScrollDown.value = false
    } else {
        showScrollDown.value = true
    }
}

const checkDb = async () => {
    dbOkStatus.value = false
    const res = await postAction('public/check-db')
    if (res.code === 1) {
        storageStore.SetGqaGoVersion(res.data.goVersion)
        storageStore.SetGqaGinVersion(res.data.ginVersion)
        storageStore.SetGqaPluginList(res.data.pluginList)
        if (res.data.needInit === false) {
            await storageStore.SetGqaDict()
            await storageStore.SetGqaFrontend()
            await storageStore.SetGqaBackend()
            dbOkStatus.value = true
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
                message: t('DbNeedInit'),
            })
            initDbDialog.value.initDbVisible = true
        }
    }
}

</script>

<style lang="scss" scoped>
.login-layout-scroll-down {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: -60px;
    z-index: 100;
}
</style>
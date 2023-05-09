<template>
    <q-layout style="overflow: hidden">
        <q-page-container>
            <ComplexView v-if="loginLayout === ComplexView && !dbNeedInit" :pluginComponent="pluginComponent"
                :pluginCurrent="pluginCurrent" />
            <div v-else :class="changeContainerImg === '' ? 'gqa-login-layout-img-container' : ''"
                :style="changeContainerImg === '' ? { background: $q.dark.isActive ? 'black' : '#e3f4fa' } : changeContainerImg">
                <div class="main-title-logo row justify-center items-center">
                    <gqa-avatar class="gin-quasar-admin-logo" :src="gqaFrontend.logo" size="45px" />
                    <span class="text-weight-bold text-h4" style="margin-left:8px">
                        {{ gqaFrontend.mainTitle }}
                    </span>
                </div>
                <figure class="primary-ball" />
                <figure class="positive-ball" />
                <figure class="warning-ball" />
                <figure class="negative-ball" />
                <q-card bordered class="init-login-card shadow-15" style="border-radius: 20px;" :class="darkThemeLoginCard">
                    <InitDbView @initDbSuccess="checkDb" v-if="dbNeedInit" />
                    <SimpleView v-else />
                </q-card>
                <audio ref="loginBgm" src="http://music.163.com/song/media/outer/url?id=1479760027.mp3" loop></audio>
                <div class="power-show">
                    Powered by Gin-Quasar-Admin
                    {{ gqaVersion }}
                </div>
                <div class="version-git-show">
                    <q-btn flat>
                        {{ $t('Version') }}{{ $t('Info') }}
                        <gqa-version-menu />
                    </q-btn>
                    <q-btn flat label="Github" @click="openLink('https://github.com/Junvary/gin-quasar-admin')" />
                    <q-btn flat label="Gitee" @click="openLink('https://gitee.com/junvary/gin-quasar-admin')" />
                </div>
                <div class="language-show">
                    <GqaLanguage style="width: 100%;" />
                </div>
                <div class="dark-theme-show">
                    <q-btn :icon="playFlag ? 'music_off' : 'music_note'" round flat @click="playBgm"></q-btn>
                    <DarkTheme />
                </div>
            </div>
        </q-page-container>
    </q-layout>
</template>

<script setup>
import useCommon from 'src/composables/useCommon'
import useConfig from 'src/composables/useConfig'
import { computed, onBeforeMount, ref, markRaw, defineAsyncComponent } from 'vue'
import SimpleView from './SimpleView/index.vue'
import ComplexView from './ComplexView/index.vue'
import InitDbView from './InitDbView/index.vue'
import { postAction } from 'src/api/manage'
import { useStorageStore } from 'src/stores/storage'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import GqaLanguage from 'src/components/GqaLanguage/index.vue'
import config from '../../../package.json'
import DarkTheme from 'src/components/GqaTheme/DarkTheme.vue';
import useTheme from 'src/composables/useTheme';

const { darkThemeLoginCard } = useTheme()
const gqaVersion = config.version
const $q = useQuasar()
const { t } = useI18n()
const { gqaFrontend } = useCommon()
const { GqaConsoleLogo } = useConfig()
const storageStore = useStorageStore()
const pluginCurrent = ref(null)
const pluginComponent = ref(null)
const dbNeedInit = ref(false)

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
    GqaConsoleLogo()
})

const getResultKey = (fileKey, pluginCode) => {
    const resultKey = fileKey.filter(key => {
        return key.split('/')[2] === pluginCode
    })
    return resultKey
}

const checkDb = async () => {
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
                    const pluginCodeComponent = getResultKey(Object.keys(pluginsFile), pluginCode)
                    pluginComponent.value = markRaw(defineAsyncComponent(() => pluginsFile[pluginCodeComponent[0]]))
                    // pluginComponent.value = markRaw(defineAsyncComponent(() => import(`src/plugins/${pluginCode}/LoginLayout/index.vue`)))
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
                message: res.message,
            })
        }
    }
}
const openLink = (url) => {
    window.open(url)
}
const changeContainerImg = computed(() => {
    if (bannerImage.value === '') {
        return ''
    } else {
        return {
            background: 'url(' + bannerImage.value + ')',
            backgroundRepeat: "no",
            backgroundPosition: 'center',
            backgroundSize: "cover",
            overflow: 'hidden',
            height: '100vh',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
        }
    }

})
const bannerImage = computed(() => {
    if (gqaFrontend.value.bannerImage && gqaFrontend.value.bannerImage.substring(0, 11) === 'gqa-upload:') {
        return process.env.API + gqaFrontend.value.bannerImage.substring(11)
    }
    return ''
})

const loginBgm = ref(null)
const playFlag = ref(false)
const playBgm = () => {
    if (playFlag.value) {
        loginBgm.value.pause()
        playFlag.value = false
    } else {
        loginBgm.value.play()
        playFlag.value = true
    }

}
</script>

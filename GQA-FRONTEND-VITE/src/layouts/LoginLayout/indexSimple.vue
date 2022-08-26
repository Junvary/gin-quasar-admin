<template>
    <div class="gin-quasar-admin-init-db-video-container">
        <div :style="fixStyle" class="filter">
            <div class="row" style="height: 100vh">
                <transition appear enter-active-class="animated slideInLeft" leave-active-class="animated slideOutLeft">
                    <div class="col column justify-center items-center text-center" style="margin-bottom: 100px;">
                        <GqaAvatar class="gin-quasar-admin-logo" :src="gqaFrontend.logo" size="100px" />
                        <br />
                        <span class="text-weight-bold text-h3 text-white">
                            {{ gqaFrontend.subTitle }}
                        </span>
                        <span class="text-white text-subtitle1" style="margin-top: 20px">
                            {{ gqaFrontend.webDescribe }}
                        </span>
                        <span class="q-gutter-md">
                            <q-btn push glossy color="primary"
                                @click="openLink('https://github.com/Junvary/gin-quasar-admin')">
                                Github
                            </q-btn>
                            <q-btn push glossy color="primary">
                                {{ $t('Version') }}{{ $t('Info') }}
                                <GqaVersion />
                            </q-btn>
                            <q-btn push glossy color="primary"
                                @click="openLink('https://gitee.com/junvary/gin-quasar-admin')">
                                Gitee
                            </q-btn>
                        </span>
                    </div>
                </transition>
                <transition appear enter-active-class="animated slideInRight"
                    leave-active-class="animated slideOutRight">
                    <div class="col row column justify-center items-center" style="margin-bottom: 80px;">
                        <q-card style="width: 60%; background: rgba(255, 255, 255, 0.5); padding: 10px;">
                            <LoginForm />
                        </q-card>
                    </div>
                </transition>
            </div>
        </div>
        <video :style="fixStyle" autoplay loop muted class="fillWidth" @canplay="canplay">
            <source :src="ramdomSrc" type="video/mp4" />
            浏览器不支持 video 标签，建议升级浏览器。
            <source :src="ramdomSrc" type="video/webm" />
            浏览器不支持 video 标签，建议升级浏览器。
        </video>
        <div class="gin-quasar-admin-init-db-page" v-if="!videoCanPlay"></div>
    </div>

</template>

<script setup>
import { onMounted } from 'vue';
import useCommon from 'src/composables/useCommon'
import GqaVersion from 'src/components/GqaVersion/index.vue';
import GqaAvatar from 'src/components/GqaAvatar/index.vue';
import LoginForm from './LoginForm.vue';
import useVideoContainer from 'src/composables/useVideoContainer'

const { resizeWindow, ramdomSrc, videoCanPlay, canplay, fixStyle } = useVideoContainer()
const { gqaFrontend, openLink } = useCommon()

onMounted(() => {
    resizeWindow()
})
</script>
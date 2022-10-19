<template>
    <div :style="changeContainerImg">
        <div class="row" style="height: 100vh">
            <transition appear enter-active-class="animated slideInLeft" leave-active-class="animated slideOutLeft"
                v-if="$q.screen.gt.xs">
                <div class="col column justify-center items-center text-center" style="margin-bottom: 100px;">
                    <GqaAvatar class="gin-quasar-admin-logo" :src="gqaFrontend.logo" size="100px" />
                    <br />
                    <span class="text-weight-bold text-h3 text-white">
                        {{ gqaFrontend.subTitle }}
                    </span>

                    <div class="row justify-center" style="width: 70%;">
                        <q-chip v-for="(item, index) in pluginList" class="glossy" color="primary" text-color="white"
                            style="cursor: pointer;">
                            {{item.plugin_name}}
                            <q-tooltip>
                                <q-badge>
                                    {{item.plugin_version}}
                                </q-badge>
                                {{item.plugin_memo}}
                            </q-tooltip>
                        </q-chip>
                    </div>

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
            <transition appear enter-active-class="animated slideInRight" leave-active-class="animated slideOutRight"
                :class="`${$q.screen.gt.xs ? 'col-6' : 'col'}`">
                <div class="col row column justify-center items-center" style="margin-bottom: 80px;">
                    <q-card style="width: 60%; background: rgba(255, 255, 255, 0.6); padding: 10px;">
                        <LoginForm />
                    </q-card>
                </div>
            </transition>
        </div>
    </div>
</template>

<script setup>
import useCommon from 'src/composables/useCommon'
import { computed } from 'vue';
import GqaAvatar from 'src/components/GqaAvatar/index.vue'
import GqaVersion from 'src/components/GqaVersion/index.vue'
import { useStorageStore } from 'src/stores/storage'
import LoginForm from '../LoginForm.vue'

const { gqaFrontend, openLink } = useCommon()
const storageStore = useStorageStore()

const pluginList = computed(() => {
    return storageStore.GetGqaPluginList()
})

const changeContainerImg = computed(() => {
    if (bannerImage.value === '') {
        return {
            background: 'url("/img/sky.jpg")',
            backgroundRepeat: "no",
            backgroundPosition: 'center',
            backgroundSize: "cover",
            overflow: "hidden",
            height: "100vh",
        }
    } else {
        return {
            background: 'url(' + bannerImage.value + ')',
            backgroundRepeat: "no",
            backgroundPosition: 'center',
            backgroundSize: "cover",
            overflow: "hidden",
            height: "100vh",
        }
    }

})
const bannerImage = computed(() => {
    if (gqaFrontend.value.bannerImage && gqaFrontend.value.bannerImage.substring(0, 11) === 'gqa-upload:') {
        return process.env.API + gqaFrontend.value.bannerImage.substring(11)
    }
    return ''
})
</script>

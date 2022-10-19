<template>
    <div class="row" style="height: 100vh">
        <transition appear enter-active-class="animated slideInLeft" leave-active-class="animated slideOutLeft"
            v-if="$q.screen.gt.xs">
            <div class="col column justify-center items-center text-center" style="margin-bottom: 100px;"
                transition-show="jump-down" transition-hide="jump-up">
                <q-avatar class="gin-quasar-admin-logo" size="100px">
                    <img src="gqa128.png" />
                </q-avatar><br />
                <span class="text-weight-bold text-h3 text-white">
                    {{ $t('WelcomeTo') }}<br />
                    Gin-Quasar-Admin<sup>v2</sup>
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
                    {{ t('InitDbHelp1') }}
                </span>
                <span class="text-white text-subtitle1" style="margin-bottom: 20px">
                    {{ t('InitDbHelp2') }}
                </span>

                <span class="q-gutter-md">
                    <q-btn push glossy color="primary" @click="openLink('https://github.com/Junvary/gin-quasar-admin')">
                        Github
                    </q-btn>
                    <q-btn push glossy color="primary">
                        {{ $t('Version') }}{{ $t('Info') }}
                        <GqaVersion />
                    </q-btn>
                    <q-btn push glossy color="primary" @click="openLink('https://gitee.com/junvary/gin-quasar-admin')">
                        Gitee
                    </q-btn>
                </span>
            </div>
        </transition>
        <transition appear enter-active-class="animated slideInRight" leave-active-class="animated slideOutRight"
            :class="`${$q.screen.gt.xs ? 'col-6' : 'col'}`">
            <div class="col row column justify-center items-center" style="margin-bottom: 80px;">
                <q-card style="width: 80%; background: rgba(255, 255, 255, 0.5);">
                    <q-toolbar>
                        <q-avatar class="gin-quasar-admin-logo">
                            <img src="gqa128.png" />
                        </q-avatar>

                        <q-toolbar-title class="row items-center">
                            <span class="text-weight-bold">
                                {{ $t('Init') }}
                                {{ $t('Database') }}
                            </span>
                            <q-space />
                            <GqaLanguage style="width: 20%" />
                        </q-toolbar-title>
                    </q-toolbar>

                    <q-card-section>
                        <q-form class="text-center gqa-form" @submit="onInitDb">
                            <div class="q-gutter-y-md column">
                                <div class="row q-gutter-md">
                                    <q-input class="col" outlined bottom-slots v-model.trim="form.db_type"
                                        :label="$t('Database') + $t('Type')" disable
                                        :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                                    </q-input>
                                    <q-input class="col" outlined v-model.trim="form.db_host"
                                        :label="$t('Database') + $t('Address')"
                                        :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                                    </q-input>
                                    <q-input class="col" outlined v-model.trim="form.db_port"
                                        :label="$t('Database') + $t('Port')"
                                        :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                                    </q-input>
                                </div>

                                <div class="row q-gutter-md">
                                    <q-input class="col" outlined v-model.trim="form.db_schema"
                                        :label="$t('Database') + $t('Name')"
                                        :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                                    </q-input>
                                    <q-input class="col" outlined v-model.trim="form.db_user"
                                        :label="$t('Database') + $t('Username')"
                                        :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                                    </q-input>
                                    <q-input class="col" outlined v-model.trim="form.db_password" autofocus
                                        :label="$t('Database') + $t('Password')"
                                        :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                                    </q-input>
                                </div>
                            </div>
                            <q-btn type="submit" color="primary" :label="$t('Start') + $t('Init')"
                                :loading="initLoading" />
                        </q-form>
                    </q-card-section>
                </q-card>
            </div>
        </transition>
    </div>
</template>


<script setup>
import { onMounted, ref, computed } from 'vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import GqaLanguage from 'src/components/GqaLanguage/index.vue'
import GqaVersion from 'src/components/GqaVersion/index.vue'
import { useStorageStore } from 'src/stores/storage'
import { postAction } from 'src/api/manage'
import { useRouter } from 'vue-router'
import useDocument from 'src/composables/useDocument'

useDocument()
const $q = useQuasar()
const { t } = useI18n()
const router = useRouter()
const storageStore = useStorageStore()
const initLoading = ref(false)
const form = ref({
    db_type: 'Mysql',
    db_host: '127.0.0.1',
    db_port: '3306',
    db_schema: 'gin-quasar-admin',
    db_user: 'root',
    db_password: '',
})

const emit = defineEmits(['initDbSuccess'])

const onInitDb = () => {
    initLoading.value = true
    postAction('public/init-db', form.value)
        .then((res) => {
            if (res.code === 1) {
                $q.notify({
                    type: 'positive',
                    message: t('Database') + t('Init') + t('Success'),
                })
                storageStore.SetGqaDict()
                storageStore.SetGqaFrontend()
                storageStore.SetGqaBackend()
                emit('initDbSuccess')
            }
        })
        .finally(() => {
            initLoading.value = false
        })
}
const openLink = (url) => {
    window.open(url)
}

onMounted(() => {
    checkDb()
})

const pluginList = computed(() => {
    return storageStore.GetGqaPluginList()
})

const checkDb = async () => {
    const res = await postAction('public/check-db')
    if (res.code === 1) {
        if (res.data.need_init === false) {
            router.push({ path: '/' })
        }
    }
}
</script>
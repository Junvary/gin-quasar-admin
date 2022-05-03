<template>
    <q-dialog v-model="initDbVisible" persistent>
        <q-card style="width: 100%; max-width: 70vw">
            <q-toolbar>
                <q-avatar class="gin-quasar-admin-logo">
                    <img src="gqa128.png" />
                </q-avatar>

                <q-toolbar-title class="row items-center">
                    <span class="text-weight-bold">
                        {{ $t('WelcomeTo') }}
                        Gin-Quasar-Admin
                        <q-badge align="top">v2</q-badge>
                    </span>
                    <q-space />
                    <GqaLanguage style="width: 15%" />
                    <q-btn dense push rounded glossy color="primary"
                        @click="openLink('https://gitee.com/junvary/gin-quasar-admin')">
                        Gitee
                    </q-btn>
                    <q-btn dense push rounded glossy color="primary" style="margin: 0 5px"
                        @click="openLink('https://github.com/Junvary/gin-quasar-admin')">
                        Github
                    </q-btn>

                    <q-btn dense push rounded glossy color="primary">
                        {{ $t('Version') }}
                        {{ $t('Info') }}
                        <GqaVersion />
                    </q-btn>
                </q-toolbar-title>
            </q-toolbar>

            <GqaPluginList />

            <q-card-section>
                <q-form class="text-center" @submit="onInitDb">

                    <div class="q-gutter-y-md column">
                        <span class="col text-negative text-subtitle1">
                            {{ t('InitDbHelp') }}
                        </span>
                        <div class="row q-gutter-md">
                            <q-input class="col" outlined bottom-slots v-model.trim="form.db_type"
                                :label="$t('Database') + $t('Type')" disable
                                :rules="[(val) => (val && val.length > 0) || $t('FixForm')]">
                            </q-input>
                            <q-input class="col" outlined v-model.trim="form.db_host"
                                :label="$t('Database') + $t('Address')"
                                :rules="[(val) => (val && val.length > 0) || $t('FixForm')]">
                            </q-input>
                            <q-input class="col" outlined v-model.trim="form.db_port"
                                :label="$t('Database') + $t('Port')"
                                :rules="[(val) => (val && val.length > 0) || $t('FixForm')]">
                            </q-input>
                        </div>

                        <div class="row q-gutter-md">
                            <q-input class="col" outlined v-model.trim="form.db_schema"
                                :label="$t('Database') + $t('Name')"
                                :rules="[(val) => (val && val.length > 0) || $t('FixForm')]">
                            </q-input>
                            <q-input class="col" outlined v-model.trim="form.db_user"
                                :label="$t('Database') + $t('Username')"
                                :rules="[(val) => (val && val.length > 0) || $t('FixForm')]">
                            </q-input>
                            <q-input class="col" outlined v-model.trim="form.db_password" autofocus
                                :label="$t('Database') + $t('Password')"
                                :rules="[(val) => (val && val.length > 0) || $t('FixForm')]">
                            </q-input>
                        </div>
                    </div>
                    <q-btn type="submit" color="primary" :label="$t('Start') + $t('Init')" />
                    <q-inner-loading :showing="initLoading">
                        <q-spinner-gears size="50px" color="primary" />
                    </q-inner-loading>
                </q-form>
            </q-card-section>
        </q-card>
    </q-dialog>
</template>

<script setup>
import { ref } from 'vue';
import { useQuasar } from 'quasar';
import { useI18n } from 'vue-i18n';
import GqaLanguage from 'src/components/GqaLanguage/index.vue';
import GqaVersion from 'src/components/GqaVersion/index.vue';
import GqaPluginList from 'src/components/GqaPluginList/index.vue';
import { useStorageStore } from 'src/stores/storage'
import { postAction } from 'src/api/manage';

const $q = useQuasar();
const { t } = useI18n();
const storageStore = useStorageStore()
const initDbVisible = ref(false);
const initLoading = ref(false);
const emit = defineEmits(['dbOk'])
const form = ref({
    db_type: 'Mysql',
    db_host: '127.0.0.1',
    db_port: '3306',
    db_schema: 'gin-quasar-admin',
    db_user: 'root',
    db_password: '',
});

defineExpose({
    initDbVisible
})

const onInitDb = () => {
    initLoading.value = true
    postAction('public/init-db', form.value).then((res) => {
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: t('Database') + t('Init') + t('Success'),
            })
            storageStore.SetGqaDict()
            storageStore.SetGqaFrontend()
            storageStore.SetGqaBackend()
            emit('dbOk', true)
            initDbVisible.value = false
        }
    }).finally(() => {
        initLoading.value = false
    })
}
const openLink = (url) => {
    window.open(url)
}
</script>
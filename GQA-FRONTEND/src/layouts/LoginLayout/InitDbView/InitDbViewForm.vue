<template>
    <transition appear enter-active-class="animated slideInRight" leave-active-class="animated slideOutRight">
        <div class="col row column justify-center items-center">
            <q-card flat style="background: rgba(0, 0, 0, 0);">
                <q-toolbar>
                    <q-avatar class="gin-quasar-admin-logo">
                        <img src="gqa128.png" />
                    </q-avatar>
                    <q-toolbar-title class="row items-center">
                        <span class="text-weight-bold">
                            {{ $t('Init') }}
                            {{ $t('Database') }}
                        </span>
                    </q-toolbar-title>
                </q-toolbar>

                <q-card-section>
                    <q-form class="text-center gqa-form" @submit="onInitDb">
                        <div class="q-gutter-y-md column">
                            <div class="row q-gutter-md">
                                <q-input class="col" outlined bottom-slots v-model.trim="form.db_type"
                                    :label="$t('Database') + $t('Type')" disable no-error-icon
                                    :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                                </q-input>
                                <q-input class="col" outlined v-model.trim="form.db_host"
                                    :label="$t('Database') + $t('Address')" no-error-icon
                                    :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                                </q-input>
                                <q-input class="col" outlined v-model.trim="form.db_port"
                                    :label="$t('Database') + $t('Port')" no-error-icon
                                    :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                                </q-input>
                            </div>
                            <div class="row q-gutter-md">
                                <q-input class="col" outlined v-model.trim="form.db_schema"
                                    :label="$t('Database') + $t('Name')" no-error-icon
                                    :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                                </q-input>
                                <q-input class="col" outlined v-model.trim="form.db_user"
                                    :label="$t('Database') + $t('Username')" no-error-icon
                                    :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                                </q-input>
                                <q-input class="col" outlined v-model.trim="form.db_password" autofocus
                                    :label="$t('Database') + $t('Password')" type="password" no-error-icon
                                    :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                                </q-input>
                            </div>
                        </div>
                        <q-btn type="submit" color="primary" :label="$t('Start') + $t('Init')" :loading="initLoading" />
                    </q-form>
                </q-card-section>
            </q-card>
        </div>
    </transition>
</template>

<script setup>
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { useStorageStore } from 'src/stores/storage'

const { t } = useI18n()
const $q = useQuasar()
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

</script>
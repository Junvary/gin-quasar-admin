<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{ $t('User') }}:
                    {{ recordDetail.value.nickname ? recordDetail.value.nickname :
                    recordDetail.value.real_name ? recordDetail.value.real_name : ""
                    }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <div class="row">
                        <q-file class="col" v-model="avatarFile" :label="$t('Avatar')" max-files="1"
                            @rejected="rejected" :accept="gqaBackend.avatarExt"
                            :max-file-size="gqaBackend.avatarMaxSize * 1024 * 1024">
                            <template v-slot:prepend>
                                <GqaAvatar :src="recordDetail.value.avatar" />
                            </template>
                            <template v-slot:after v-if="avatarFile">
                                <q-btn dense flat color="primary" icon="cloud_upload" @click="handleUpload" />
                            </template>
                        </q-file>
                        <q-input class="col" v-model="recordDetail.value.username" :label="$t('Username')" lazy-rules
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" :disable="formType === 'edit'" />
                        <q-input class="col" v-model="recordDetail.value.nickname" :label="$t('Nickname')" />
                        <q-input class="col" v-model="recordDetail.value.real_name" :label="$t('RealName')" lazy-rules
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" :disable="formType === 'edit'" />
                    </div>
                    <!-- <div class="row">
                        <q-field class="col" :label="$t('Role')">
                            <template v-slot:control>
                                <span v-for="(item, index) in recordDetail.value.role" :key="index">
                                    {{ item.role_name }};&nbsp;
                                </span>
                            </template>
                        </q-field>
                        <q-field class="col" :label="$t('Dept')">
                            <template v-slot:control>
                                <span v-for="(item, index) in recordDetail.value.dept" :key="index">
                                    {{ item.dept_name }};&nbsp;
                                </span>
                            </template>
                        </q-field>
                    </div> -->
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.mobile" :label="$t('Mobile')" />
                        <q-input class="col" v-model="recordDetail.value.email" :label="$t('Email')" />
                    </div>

                    <div class="row">
                        <q-field dense class="col" :label="$t('Gender')" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="recordDetail.value.gender" :options="dictOptions.gender"
                                    color="primary" inline>
                                </q-option-group>
                            </template>
                        </q-field>
                        <q-field dense class="col" :label="$t('Status')" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="recordDetail.value.status" :options="dictOptions.onOff"
                                    color="primary" inline :disable="recordDetail.value.username === 'admin'">
                                </q-option-group>
                            </template>
                        </q-field>
                    </div>
                    <q-input v-model="recordDetail.value.memo" type="textarea" :label="$t('Memo')" />
                </q-form>
            </q-card-section>

            <q-separator />

            <q-card-actions align="right">
                <q-btn :label="$t('Save')" color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('Cancel')" color="negative" v-close-popup />
            </q-card-actions>

            <q-inner-loading :showing="loading">
                <q-spinner-gears size="50px" color="primary" />
            </q-inner-loading>
        </q-card>
    </q-dialog>
</template>

<script setup>
import useRecordDetail from 'src/composables/useRecordDetail'
import { postAction } from 'src/api/manage'
import { useStorageStore } from 'src/stores/storage'
import { ref, computed } from 'vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'

const $q = useQuasar()
const { t } = useI18n()
const storageStore = useStorageStore()
const gqaBackend = computed(() => storageStore.GetGqaBackend())
const emit = defineEmits(['handleFinish'])
const url = {
    add: 'user/add-user',
    edit: 'user/edit-user',
    queryById: 'user/query-user-by-id',
    avatarUrl: 'upload/upload-avatar',
}
const {
    GqaAvatar,
    dictOptions,
    formType,
    formTypeName,
    recordDetail,
    recordDetailVisible,
    loading,
    show,
    recordDetailForm,
    handleAddOrEidt
} = useRecordDetail(url, emit)

defineExpose({
    show,
    formType
})

const avatarFile = ref(null)
const handleUpload = () => {
    if (!avatarFile.value) {
        $q.notify({
            type: 'negative',
            message: t('PleaseSelectFile'),
        })
        return
    }
    let form = new FormData()
    form.append('file', avatarFile.value)
    postAction(url.avatarUrl, form).then((res) => {
        if (res.code === 1) {
            recordDetail.value.avatar = res.data.records
            avatarFile.value = null
            $q.notify({
                type: 'positive',
                message: t('UploadSuccess'),
            })
        }
    })
}
const rejected = (rejectedEntries) => {
    $q.notify({
        type: 'negative',
        message: t('FileRejected'),
    })
}
</script>

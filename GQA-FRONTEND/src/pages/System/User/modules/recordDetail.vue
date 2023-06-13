<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 1200px; max-width: 80vw;">
            <q-card-section class="text-h6">
                {{ formTypeName }} {{ $t('User') }}:
                {{ recordDetail.value.nickname ? recordDetail.value.nickname :
                    recordDetail.value.real_name ? recordDetail.value.real_name : ""
                }}
            </q-card-section>
            <gqa-form-top :recordDetail="recordDetail" style="margin: 0 16px;"></gqa-form-top>
            <div class="row">
                <q-card-section class="q-gutter-md col-8">
                    <q-form ref="recordDetailForm">
                        <div class="row q-gutter-md">
                            <q-file outlined hint="" class="col" v-model="avatarFile" :label="$t('Avatar')" max-files="1"
                                @rejected="rejected" :accept="gqaBackend.avatarExt"
                                :max-file-size="gqaBackend.avatarMaxSize * 1024 * 1024">
                                <template v-slot:prepend>
                                    <gqa-avatar :src="recordDetail.value.avatar" />
                                </template>
                                <template v-slot:after v-if="avatarFile">
                                    <q-btn dense flat color="primary" icon="cloud_upload" @click="handleUpload" />
                                </template>
                            </q-file>
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.username"
                                :label="$t('Username')" lazy-rules
                                :rules="[val => val && val.length > 0 || $t('NeedInput')]" :disable="formType === 'edit'" />
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.real_name"
                                :label="$t('RealName')" lazy-rules
                                :rules="[val => val && val.length > 0 || $t('NeedInput')]" :disable="formType === 'edit'" />
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.nickname"
                                :label="$t('Nickname')" />
                        </div>
                        <div class="row q-gutter-md">
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.mobile"
                                :label="$t('Mobile')" />
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.email" :label="$t('Email')" />
                        </div>
                        <div class="row q-gutter-md">
                            <q-field outlined hint="" dense class="col" :label="$t('Gender')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="recordDetail.value.gender" :options="dictOptions.gender"
                                        color="primary" inline>
                                        <template v-slot:label="opt">
                                            <div class="row items-center">
                                                <span>{{ $t(opt.label) }}</span>
                                            </div>
                                        </template>
                                    </q-option-group>
                                </template>
                            </q-field>
                            <q-field outlined hint="" dense class="col" :label="$t('Status')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="recordDetail.value.status" :options="dictOptions.onOff"
                                        color="primary" inline :disable="recordDetail.value.username === 'admin'">
                                        <template v-slot:label="opt">
                                            <div class="row items-center">
                                                <span>{{ $t(opt.label) }}</span>
                                            </div>
                                        </template>
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>
                        <q-input outlined hint="" v-model="recordDetail.value.memo" type="textarea" :label="$t('Memo')" />
                    </q-form>
                </q-card-section>
                <q-card-section class="q-gutter-md col">
                    <q-field outlined hint="" dense :label="$t('Select') + $t('Dept')" stack-label>
                        <template v-slot:control>
                            <q-tree :nodes="deptList" v-model:ticked="selectDept" @update:ticked="handleSelectDept"
                                label-key="dept_name" node-key="dept_code" default-expand-all tick-strategy="strict" />
                        </template>
                    </q-field>
                </q-card-section>
            </div>

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
import { ref, computed, toRefs } from 'vue'
import XEUtils from 'xe-utils'

const storageStore = useStorageStore()
const gqaBackend = computed(() => storageStore.GetGqaBackend())
const emit = defineEmits(['handleFinish'])
const url = {
    add: 'user/add-user',
    edit: 'user/edit-user',
    queryById: 'user/query-user-by-id',
    avatarUrl: 'upload/upload-avatar',
    deptListUrl: 'dept/get-dept-list'
}
const {
    $q,
    t,
    dictOptions,
    formType,
    formTypeName,
    recordDetail,
    recordDetailVisible,
    loading,
    // show,
    // handleQueryById,
    recordDetailForm,
    handleAddOrEidt
} = useRecordDetail(url, emit)

const show = (row) => {
    loading.value = true
    recordDetail.value = {}
    recordDetailVisible.value = true
    if (row && row.id) {
        handleQueryById(row.id)
    } else {
        selectDept.value = []
        recordDetail.value = {}
        recordDetailVisible.value = true
        loading.value = false
    }
}

const handleQueryById = async (id) => {
    postAction(url.queryById, {
        id,
    }).then(res => {
        if (res.code === 1) {
            selectDept.value = []
            recordDetail.value = res.data.records
            for (let d of res.data.records.dept) {
                selectDept.value.push(d.dept_code)
            }

        }
    }).finally(() => {
        loading.value = false
    })
}

defineExpose({
    show,
    formType
})

const props = defineProps({
    deptList: {
        type: Array,
        required: false,
        default: () => [],
    },
})
const { deptList } = toRefs(props)

const avatarFile = ref(null)
const handleUpload = () => {
    if (!avatarFile.value) {
        $q.notify({
            type: 'negative',
            message: t('Please') + t('Select') + t('File'),
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
                message: t('Upload') + t('Success'),
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

const selectDept = ref([])
const handleSelectDept = (target) => {
    recordDetail.value.dept = []
    for (let t of target) {
        const tt = XEUtils.filterTree(deptList.value, item => item.dept_code === t)
        recordDetail.value.dept.push(tt[0])
    }
}
</script>

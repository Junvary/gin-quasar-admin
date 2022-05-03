<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{ $t('Notice') }}:
                    {{ recordDetail.value.notice_title }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <q-input v-model="recordDetail.value.notice_title" :label="$t('Title')"
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                    <q-select v-model="recordDetail.value.notice_type" :options="dictOptions.noticeType" emit-value
                        map-options :rules="[val => val && val.length > 0 || $t('NeedInput')]"
                        :label="$t('Notice') + $t('Type')" />
                    <q-input v-model="recordDetail.value.notice_content" type="textarea" :label="$t('Content')" />

                    <q-field :label="$t('SendTo')" stack-label>
                        <template v-slot:control>
                            <q-option-group :options="noticeToUserTypeOption" name="noticeToUserType"
                                v-model="recordDetail.value.notice_to_user_type" inline
                                @update:model-value="changenoticeToUserType" />
                        </template>
                    </q-field>
                    <q-select v-if="recordDetail.value.notice_to_user_type === 'some' && changeTableData.value.length"
                        v-model="selectUser" :options="changeTableData.value" multiple clearable emit-value map-options
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]" :label="$t('User')" />
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
import useTableData from 'src/composables/useTableData'
import GqaShowName from 'src/components/GqaShowName'
import { postAction } from 'src/api/manage'
import { useStorageStore } from 'src/stores/storage'
import { ref, computed } from 'vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import GqaSeleteUser from 'src/components/GqaSeleteUser'

const $q = useQuasar()
const { t } = useI18n()
const emit = defineEmits(['handleFinish'])
const url = {
    list: 'user/get-user-list',
    add: 'notice/add-notice',
    edit: 'notice/edit-notice',
    queryById: 'notice/query-notice-by-id',
}
const {
    dictOptions,
    showDateTime,
    formType,
    formTypeName,
    recordDetail,
    recordDetailVisible,
    loading,
    // show,
    handleQueryById,
    recordDetailForm,
    // handleAddOrEidt
} = useRecordDetail(url, emit)
const {
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    GqaAvatar,
    // loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    onRequest,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

const show = (row) => {
    loading.value = true
    recordDetail.value = {}
    recordDetailVisible.value = true
    if (row && row.noticeToUserType === 'some') {
        changenoticeToUserType('some')
    }
    if (row && row.id) {
        handleQueryById(row.id)
    } else {
        recordDetail.value = {}
        recordDetailVisible.value = true
        loading.value = false
    }
}
defineExpose({
    show,
    formType,
    recordDetail
})

const noticeToUserTypeOption = computed(() => {
    return [
        { label: t('All') + t('User'), value: 'all' },
        { label: t('Some') + t('User'), value: 'some' },
    ]
})
const changeTableData = computed(() => {
    if (tableData.value.length) {
        for (let i of tableData.value) {
            i.label = i.nickname ? i.nickname : i.realName ? i.realName : i.username
            i.value = i.username
        }
        return tableData.value
    }
    return []
})
const selectUser = ref([])


const changenoticeToUserType = (val) => {
    if (val === 'some') {
        selectUser.value = []
        onRequest({
            pagination: pagination.value,
            queryParams: queryParams.value
        }).then(() => {
            if (recordDetail.value.notice_to_user_type === 'some') {
                for (let u of recordDetail.value.notice_to_user) {
                    selectUser.value.push(u.toUser)
                }
            }
        })
    }
}
const handleAddOrEidt = async () => {
    if (!recordDetail.value.notice_to_user_type) {
        $q.notify({
            type: 'negative',
            message: t('FixForm') + ': ' + t('SendTo'),
        })
        return
    }
    const success = await recordDetailForm.value.validate()
    if (success) {
        if (recordDetail.value.noticeToUserType === 'some') {
            recordDetail.value.notice_to_user = selectUser.value
        } else {
            recordDetail.value.notice_to_user = []
        }
        if (formType.value === 'edit') {
            if (url === undefined || !url.edit) {
                $q.notify({
                    type: 'negative',
                    message: "请先配置url",
                })
                return
            }
            recordDetail.value.notice_to_user = []
            selectUser.value.forEach((item) => {
                recordDetail.value.notice_to_user.push({
                    noticeId: recordDetail.value.notice_id,
                    toUser: item,
                })
            })
            const res = await postAction(url.edit, recordDetail.value)
            if (res.code === 1) {
                $q.notify({
                    type: 'positive',
                    message: res.message,
                })
                recordDetailVisible.value = false
            }
        } else if (formType.value === 'add') {
            if (url === undefined || !url.add) {
                $q.notify({
                    type: 'negative',
                    message: "请先配置url",
                })
                return
            }
            const res = await postAction(url.add, recordDetail.value)
            if (res.code === 1) {
                $q.notify({
                    type: 'positive',
                    message: res.message,
                })
                recordDetailVisible.value = false
            }
        } else {
            $q.notify({
                type: 'negative',
                message: t('CanNotAddOrEdit'),
            })
        }
        emit('handleFinish')
    } else {
        $q.notify({
            type: 'negative',
            message: t('FixForm'),
        })
    }
}
</script>

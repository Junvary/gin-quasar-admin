<template>
    <q-dialog v-model="recordDetailVisible" position="top" @hide="hide">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ recordDetail.value.notice_title }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section style="max-height: 80vh" class="scroll q-gutter-md">
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <q-input v-model="recordDetail.value.notice_title" :label="$t('Title')"
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]" readonly />

                    <q-select v-model="recordDetail.value.notice_type" :options="dictOptions.noticeType" emit-value
                        map-options :rules="[val => val && val.length > 0 || $t('NeedInput')]"
                        :label="$t('Notice') + $t('Type')" readonly />

                    <q-input v-model="recordDetail.value.notice_content" type="textarea" :label="$t('Content')"
                        readonly />
                </q-form>
            </q-card-section>

            <q-inner-loading :showing="loading">
                <q-spinner-gears size="50px" color="primary" />
            </q-inner-loading>
        </q-card>
    </q-dialog>
</template>


<script setup>
import useRecordDetail from 'src/composables/useRecordDetail'

const emit = defineEmits(['handleFinish', 'hide'])
const url = {
    queryById: 'notice/query-notice-read-by-id',
}
const {
    bus,
    dictOptions,
    formType,
    recordDetail,
    recordDetailVisible,
    loading,
    show,
    recordDetailForm,
} = useRecordDetail(url, emit)

defineExpose({
    show,
    formType,
    recordDetail
})
const hide = () => {
    bus.emit('noticeGetTableData')
    emit('hide')
}
</script>

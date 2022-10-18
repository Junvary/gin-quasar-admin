<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} 测试数据
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.column_1" label="第1列" lazy-rules
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        <q-input class="col" v-model="recordDetail.value.column_2" label="第2列" lazy-rules
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        <q-input class="col" v-model="recordDetail.value.column_3" label="第3列" lazy-rules
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        <q-input class="col" v-model="recordDetail.value.column_4" label="第4列" lazy-rules
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        <q-input class="col" v-model="recordDetail.value.column_5" label="第5列" lazy-rules
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                    </div>
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

const emit = defineEmits(['handleFinish'])
const url = {
    add: 'plugin-example/add-test-data',
    edit: 'plugin-example/edit-test-data',
    queryById: 'plugin-example/query-test-data-by-id',
}
const {
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
</script>

<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section class="text-h6">
                {{ formTypeName }} {{ $t('Role') }}:
                {{ recordDetail.value.role_name }}
            </q-card-section>
            <q-separator />
            <q-card-section>
                <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                <q-form ref="recordDetailForm">
                    <div class="row q-gutter-md">
                        <q-input outlined hint="" class="col" v-model.number="recordDetail.value.sort" type="number"
                            :rules="[val => val >= 1 || $t('SortRule')]" :label="$t('Sort')" />
                        <q-input outlined hint="" class="col" v-model="recordDetail.value.role_code"
                            :label="($t('Role') + $t('Code'))" :disable="formType === 'edit'"
                            :rules="[val => val && val.length > 0 || $t('NeetInput')]" />
                        <q-input outlined hint="" class="col" v-model="recordDetail.value.role_name"
                            :label="($t('Role') + $t('Name'))" :rules="[val => val && val.length > 0 || $t('NeetInput')]" />
                    </div>
                    <q-input outlined hint="" v-model="recordDetail.value.memo" type="textarea" :label="$t('Memo')" />
                </q-form>
            </q-card-section>
            <q-separator />
            <q-card-actions align="right">
                <q-btn :label="$t('Save') + formTypeName" color="primary" @click="handleAddOrEidt" />
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
    add: 'role/add-role',
    edit: 'role/edit-role',
    queryById: 'role/query-role-by-id',
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
    formType,
    recordDetail
})
</script>

<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }}Obtain:
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.category" label="分类"
                            :rules="[val => val && val.length > 0 || '必须输入分类']" />
                        <q-input class="col" v-model="recordDetail.value.code" label="成就代码"
                            :rules="[val => val && val.length > 0 || '必须输入成就代码']" />
                        <q-input class="col" v-model="recordDetail.value.name" label="成就名"
                            :rules="[val => val && val.length > 0 || '必须输入成就名']" />
                    </div>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.memo" type="textarea" label="成就描述" />
                    </div>
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
    add: 'plugin-achievement/add-category',
    edit: 'plugin-achievement/edit-category',
    queryById: 'plugin-achievement/query-category-by-id',
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

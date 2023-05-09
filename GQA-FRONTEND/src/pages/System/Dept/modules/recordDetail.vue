<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section class="text-h6">
                {{ formTypeName }} {{ $t('Dept') }}:
                {{ recordDetail.value.dept_name }}
            </q-card-section>
            <q-separator />
            <q-card-section>
                <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                <q-form ref="recordDetailForm">
                    <div class="row q-gutter-md">
                        <q-input outlined hint="" class="col" :label="$t('Dept') + $t('Code')"
                            v-model="recordDetail.value.dept_code"
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        <q-input outlined hint="" class="col" :label="$t('Parent') + $t('Dept') + $t('Code')"
                            v-model="recordDetail.value.parent_code" />
                    </div>
                    <div class="row q-gutter-md">
                        <q-input outlined hint="" class="col" :label="$t('Dept') + $t('Name')"
                            v-model="recordDetail.value.dept_name"
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        <GqaSeleteUser outlined hint="" class="col" :label="$t('Leader')"
                            v-model:selectUser="recordDetail.value.leader_user"
                            v-model:selectUsername="recordDetail.value.leader" selection="single" />
                    </div>
                    <div class="row q-gutter-md">
                        <q-input outlined hint="" class="col" v-model.number="recordDetail.value.sort" type="number"
                            :rules="[val => val >= 1 || $t('SortRule')]" :label="$t('Sort')" />
                        <q-field dense outlined hint="" class="col" :label="$t('Status')" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="recordDetail.value.status" :options="dictOptions.onOff"
                                    color="primary" inline>
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
import GqaSeleteUser from 'src/components/GqaSeleteUser/index.vue'

const emit = defineEmits(['handleFinish'])
const url = {
    list: 'dept/get-dept-list',
    add: 'dept/add-dept',
    edit: 'dept/edit-dept',
    queryById: 'dept/query-dept-by-id',
}
const {
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
    formType,
    recordDetail
})
</script>

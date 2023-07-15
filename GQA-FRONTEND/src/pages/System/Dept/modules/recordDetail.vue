<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 1200px; max-width: 80vw;">
            <q-card-section class="text-h6">
                {{ formTypeName }} {{ $t('Dept') }}:
                {{ recordDetail.value.dept_name }}
            </q-card-section>
            <gqa-form-top :recordDetail="recordDetail" style="margin: 0 16px;"></gqa-form-top>
            <div class="row">
                <q-card-section class="q-gutter-md col-8">
                    <!-- <gqa-form-top :recordDetail="recordDetail"></gqa-form-top> -->
                    <q-form ref="recordDetailForm">
                        <div class="row q-gutter-md">
                            <q-input outlined hint="" class="col" :label="$t('Dept') + $t('Code')"
                                v-model="recordDetail.value.dept_code" :disable="formType === 'edit'"
                                :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                            <q-input outlined hint="" class="col" :label="$t('Dept') + $t('Name')"
                                v-model="recordDetail.value.dept_name"
                                :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                            <q-input outlined hint="" class="col" :label="$t('Parent') + $t('Dept') + $t('Code')"
                                v-model="recordDetail.value.parent_code" readonly />
                        </div>
                        <div class="row q-gutter-md">
                            <q-input outlined hint="" class="col" v-model.number="recordDetail.value.sort" type="number"
                                :rules="[val => val >= 1 || $t('SortRule')]" :label="$t('Sort')" />
                            <GqaSeleteUser outlined hint="" class="col" :label="$t('Leader')"
                                v-model:selectUser="recordDetail.value.leader_user"
                                v-model:selectUsername="recordDetail.value.leader" selection="single" />
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
                <q-card-section class="q-gutter-md col">
                    <q-field outlined hint="" dense :label="$t('Parent') + $t('Dept')" stack-label
                        :disable="!recordDetail.value.parent_code">
                        <template v-slot:control>
                            <q-tree :nodes="deptTree" v-model:selected="recordDetail.value.parent_code" style="width: 100%;"
                                selected-color="primary" label-key="dept_name" node-key="dept_code" default-expand-all
                                @update:selected="handleSelectParentDept" />
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
import GqaSeleteUser from 'src/components/GqaSeleteUser/index.vue'
import { toRefs } from 'vue'

const emit = defineEmits(['handleFinish'])
const url = {
    list: 'dept/get-dept-list',
    add: 'dept/add-dept',
    edit: 'dept/edit-dept',
    queryById: 'dept/query-dept-by-id',
}

const props = defineProps({
    deptTree: {
        type: Array,
        required: false,
        default: () => [],
    },
})
const { deptTree } = toRefs(props)

const {
    $q,
    t,
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

const handleSelectParentDept = (target) => {
    if (target === recordDetail.value.dept_code) {
        recordDetail.value.parent_code = ''
        $q.notify({ type: 'warning', message: t('ParentNotThis') })
    }
}
</script>

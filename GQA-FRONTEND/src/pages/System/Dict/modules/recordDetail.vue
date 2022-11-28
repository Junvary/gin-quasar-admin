<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 1000px; max-width: 80vw;">
            <q-card-section class="text-h6">
                {{ formTypeName }}{{ $t('Dict') }}:
                {{ recordDetail.value.dict_label }}
                <div class="text-subtitle2 text-deep-orange" v-if="recordDetail.value.parent_code">
                    {{ $t('DictHelp') }} , {{ $t('WithoutOther') }} '_'
                </div>
            </q-card-section>
            <q-separator />
            <q-card-section>
                <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                <q-form ref="recordDetailForm">
                    <div class="row q-gutter-md">
                        <q-input outlined hint="" class="col" v-model.number="recordDetail.value.sort" type="number"
                            :rules="[val => val >= 1 || $t('SortRule')]" :label="$t('Sort')" />
                        <template v-if="recordDetail.value.parent_code">
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.parent_code"
                                :label="$t('Parent') + $t('Code')" lazy-rules disable
                                :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        </template>
                    </div>
                    <div class="row q-gutter-md">
                        <q-input outlined hint="" class="col" v-model="recordDetail.value.dict_label"
                            :label="$t('Dict') + $t('Name')" lazy-rules
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        <template v-if="recordDetail.value.parent_code">
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.dict_code" lazy-rules
                                label-slot :rules="[val =>
                                val && val.length > 0
                                && val.indexOf(recordDetail.value.parent_code + '_') === 0
                                && val.slice(recordDetail.value.parent_code.length + 1).indexOf('_') === -1
                                || $t('FixForm')]">
                                <template v-slot:label>
                                    <span class="text-weight-bold text-deep-orange">
                                        {{ $t('Dict') }}{{ $t('Code') }}
                                        ( {{ $t('StartWith', { name: recordDetail.value.parent_code + '_ ' }) }},
                                        {{ $t('WithoutOther') }} _ )
                                    </span>
                                </template>
                            </q-input>
                        </template>
                        <!-- val && val.length > 0 && val.indexOf(recordDetail.value.parent_code) !== 0        || $t('NeedInput')-->
                        <template v-else>
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.dict_code"
                                :label="$t('Dict') + $t('Code')" lazy-rules
                                :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        </template>
                    </div>
                    <div class="row q-gutter-md">
                        <q-input outlined hint="" class="col" v-model="recordDetail.value.dict_ext_1" label="Ext1"
                            lazy-rules />
                        <q-input outlined hint="" class="col" v-model="recordDetail.value.dict_ext_2" label="Ext2"
                            lazy-rules />
                        <q-input outlined hint="" class="col" v-model="recordDetail.value.dict_ext_3" label="Ext3"
                            lazy-rules />
                        <q-input outlined hint="" class="col" v-model="recordDetail.value.dict_ext_4" label="Ext4"
                            lazy-rules />
                        <q-input outlined hint="" class="col" v-model="recordDetail.value.dict_ext_5" label="Ext5"
                            lazy-rules />
                    </div>
                    <div class="row q-gutter-md">
                        <q-field outlined hint="" class="col" :label="$t('Status')" stack-label>
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

const emit = defineEmits(['handleFinish'])
const url = {
    add: 'dict/add-dict',
    edit: 'dict/edit-dict',
    queryById: 'dict/query-dict-by-id',
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

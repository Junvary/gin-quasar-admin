<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{ $t('Dict') }}:
                    {{ recordDetail.value.dict_name }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section>
                <q-form ref="addOrEditForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <div class="row">
                        <q-input class="col" v-model.number="recordDetail.value.sort" type="number"
                            :rules="[val => val >= 1 || $t('SortRule')]" :label="$t('Sort')" />
                        <q-input class="col" v-model="recordDetail.value.dict_code" :label="$t('Dict') + $t('Code')"
                            lazy-rules :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        <q-input class="col" v-model="recordDetail.value.dict_label" :label="$t('Dict') + $t('Name')"
                            lazy-rules :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        <q-input class="col" v-model="recordDetail.value.parent_code"
                            :label="$t('Parent') + $t('Dict') + $t('Code')" />
                    </div>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.dict_ext_1" label="Ext1" lazy-rules />
                        <q-input class="col" v-model="recordDetail.value.dict_ext_2" label="Ext2" lazy-rules />
                        <q-input class="col" v-model="recordDetail.value.dict_ext_3" label="Ext3" lazy-rules />
                        <q-input class="col" v-model="recordDetail.value.dict_ext_4" label="Ext4" lazy-rules />
                        <q-input class="col" v-model="recordDetail.value.dict_ext_5" label="Ext5" lazy-rules />
                    </div>
                    <div class="row">
                        <q-field class="col" :label="$t('Status')" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="recordDetail.value.status" :options="dictOptions.statusOnOff"
                                    color="primary" inline>
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
import GqaAvatar from 'src/components/GqaAvatar'
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
    add: 'dict/add-dict',
    edit: 'dict/edit-dict',
    queryById: 'dict/query-dict-by-id',
}
const {
    dictOptions,
    showDateTime,
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

<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{ $t('Role') }}:
                    {{ recordDetail.value.role_name }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <div class="row">
                        <q-input class="col" v-model.number="recordDetail.value.sort" type="number"
                            :rules="[val => val >= 1 || '排序必须大于0']" label="排序" />
                        <q-input class="col" v-model="recordDetail.value.role_code" label="角色编码(英文)"
                            :rules="[val => val && val.length > 0 || '必须输入角色编码']" />
                        <q-input class="col" v-model="recordDetail.value.role_name" label="角色名(中文)"
                            :rules="[val => val && val.length > 0 || '必须输入角色名']" />
                    </div>
                    <!-- <div class="row">
                        <q-input class="col" v-model.number="recordDetail.value.sort" type="number"
                            :rules="[val => val >= 1 || '排序必须大于0']" label="排序" />
                        <q-field class="col" dense label="是否启用" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="recordDetail.value.status" :options="dictOptions.onOff"
                                    color="primary" inline>
                                </q-option-group>
                            </template>
                        </q-field>
                    </div> -->
                    <q-input v-model="recordDetail.value.memo" type="textarea" :label="$t('Memo')" />
                </q-form>
            </q-card-section>

            <q-separator />

            <q-card-actions align="right">
                <q-btn :label="'保存' + formTypeName" color="primary" @click="handleAddOrEidt" />
                <q-btn label="取消" color="negative" v-close-popup />
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
import { ref, computed } from 'vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import GqaSeleteUser from 'src/components/GqaSeleteUser'

const $q = useQuasar()
const { t } = useI18n()
const emit = defineEmits(['handleFinish'])
const url = {
    add: 'role/add-role',
    edit: 'role/edit-role',
    queryById: 'role/query-role-by-id',
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

<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{ $t('Frontend') }} {{ $t('Config') }}:
                    {{ recordDetail.value.config_item }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <div class="row">
                        <q-input class="col" v-model.number="recordDetail.value.sort" type="number"
                            :rules="[val => val >= 1 || '排序必须大于0']" label="排序" />
                        <q-input class="col" v-model="recordDetail.value.memo" label="描述" />
                    </div>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.config_item" label="字段名（英）"
                            :rules="[val => val && val.length > 0 || '必须输入字段名']" />

                        <q-input class="col" v-model="recordDetail.value.item_default" label="默认值"
                            :rules="[val => val && val.length > 0 || '必须输入默认值']" />
                    </div>
                    <q-field label="是否启用" stack-label>
                        <template v-slot:control>
                            <q-option-group v-model="recordDetail.value.status" :options="dictOptions.statusOnOff"
                                color="primary" inline>
                            </q-option-group>
                        </template>
                    </q-field>
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
    add: 'config-frontend/config-frontend-add',
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

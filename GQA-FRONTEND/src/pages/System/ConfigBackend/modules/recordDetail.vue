<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{ $t('Backend') }} {{ $t('Config') }}:
                    {{ recordDetail.value.config_item }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <div class="row">
                        <q-input class="col" v-model.number="recordDetail.value.sort" type="number"
                            :rules="[val => val >= 1 || $t('SortRule')]" :label="$t('Sort')" />
                        <q-input class="col" v-model="recordDetail.value.memo" :label="$t('Memo')" />
                    </div>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.config_item" :label="$t('Column')"
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" />

                        <q-input class="col" v-model="recordDetail.value.item_default" :label="$t('Default')"
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                    </div>
                    <q-field :label="$t('Status')" stack-label>
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
    add: 'config-backend/config-backend-add',
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

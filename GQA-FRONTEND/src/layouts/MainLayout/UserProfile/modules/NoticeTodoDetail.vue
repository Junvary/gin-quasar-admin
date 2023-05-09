<template>
    <q-dialog v-model="recordDetailVisible" position="top" @hide="hide">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{ $t('Todo') }}:
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <q-input v-model="recordDetail.value.todo_detail" type="textarea" :label="$t('Detail')"
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]" />

                    <q-field :label="$t('Done')" stack-label>
                        <template v-slot:control>
                            <q-option-group v-model="recordDetail.value.todo_status" :options="dictOptions.yesNo"
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
    add: 'todo/add-todo',
    edit: 'todo/edit-todo',
    queryById: 'todo/query-todo-by-id',
}
const {
    bus,
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
const hide = () => {
    bus.emit('noticeGetTableData')
}
</script>

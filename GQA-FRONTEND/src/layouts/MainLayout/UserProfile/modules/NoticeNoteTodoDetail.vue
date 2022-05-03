<template>
    <q-dialog v-model="recordDetailVisible" position="top" @hide="hide">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{ $t('NoteTodo') }}:
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
                            <q-option-group v-model="recordDetail.value.todo_status" :options="dictOptions.statusYesNo"
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
import { emitter } from 'src/boot/bus'

const $q = useQuasar()
const { t } = useI18n()
const emit = defineEmits(['handleFinish'])
const url = {
    add: 'note-todo/add-note-todo',
    edit: 'note-todo/edit-note-todo',
    queryById: 'note-todo/query-note-todo-by-id',
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
const hide = () => {
    emitter.emit('noticeGetTableData')
}
</script>

<template>
    <q-dialog v-model="recordDetailVisible">
        <q-card style="width: 900px; max-width: 80vw;">
            <div class="row justify-between items-center">
                <q-card-section>
                    <div class="text-h6">
                        {{ $t('Register') + $t('Button') }}:
                        {{ recordDetail.value.title ? $t(recordDetail.value.title) : '' }}
                    </div>
                </q-card-section>
            </div>

            <q-separator />

            <q-card-section>
                <q-table dense row-key="button_code" separator="cell" :rows="recordDetail.value.button"
                    :columns="columns" v-model:pagination="pagination" :rows-per-page-options="pageOptions"
                    :loading="loading">

                    <template v-slot:top>
                        <q-btn round dense @click="addButton" icon="add" />
                    </template>

                    <template v-slot:header-cell-button_name="props">
                        <q-th :props="props">
                            {{ props.col.label }}
                            <q-icon name="edit" size="1.3em" />
                        </q-th>
                    </template>
                    <template v-slot:body-cell-button_name="props">
                        <q-td key="button_name" :props="props">
                            {{ props.row.button_name }}
                            <q-popup-edit v-model="props.row.button_name" buttons v-slot="scope">
                                <q-input v-model="scope.value" dense autofocus counter />
                            </q-popup-edit>
                        </q-td>
                    </template>

                    <template v-slot:header-cell-button_code="props">
                        <q-th :props="props">
                            {{ props.col.label }}
                            <q-icon name="edit" size="1.3em" />
                        </q-th>
                    </template>
                    <template v-slot:body-cell-button_code="props">
                        <q-td key="button_code" :props="props">
                            {{ props.row.button_code }}
                            <q-popup-edit v-model="props.row.button_code" buttons v-slot="scope">
                                {{ $t('Button') + $t('Code') + ' ' + $t('StartWith', {
                                        name: recordDetail.value.name +
                                            ':'
                                    })
                                }}
                                <q-input v-model="scope.value" dense autofocus counter />
                            </q-popup-edit>
                        </q-td>
                    </template>

                    <template v-slot:body-cell-actions="props">
                        <q-td :props="props">
                            <div class="q-gutter-xs">
                                <q-btn dense color="negative" @click="deleteButton(props.row)" :label="$t('Delete')" />
                            </div>
                        </q-td>
                    </template>

                </q-table>
            </q-card-section>
            <q-card-actions align="right">
                <q-btn :label="$t('Save')" color="primary" @click="editButton" />
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
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { postAction } from 'src/api/manage';
import { useQuasar } from 'quasar';

const $q = useQuasar()
const { t } = useI18n()
const emit = defineEmits(['handleFinish'])
const pagination = ref({
    sortBy: 'button_code',
    descending: false,
    page: 1,
    rowsPerPage: 10,
})
const pageOptions = ref([10, 30, 50, 100])
const url = {
    edit: 'menu/edit-menu',
    queryById: 'menu/query-menu-by-id',
}
const {
    recordDetail,
    recordDetailVisible,
    loading,
    show,
} = useRecordDetail(url, emit)

const columns = computed(() => {
    return [
        { name: 'menu_name', align: 'center', label: t('Menu') + t('Name'), field: 'menu_name' },
        { name: 'button_name', align: 'center', label: t('Button') + t('Name'), field: 'button_name' },
        { name: 'button_code', align: 'center', label: t('Button') + t('Code'), field: 'button_code' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})

defineExpose({
    show,
    recordDetail
})

const editButton = async () => {
    const res = await postAction(url.edit, recordDetail.value)
    if (res.code === 1) {
        $q.notify({
            type: 'positive',
            message: res.message,
        })
        show({
            id: recordDetail.value.id
        })
    }
}

const deleteButton = (row) => {
    recordDetail.value.button = recordDetail.value.button.filter(item => item.button_code !== row.button_code)
}
const addButton = () => {
    recordDetail.value.button.push({
        menu_name: recordDetail.value.name,
        button_name: '',
        button_code: ''
    })
}
</script>
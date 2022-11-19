<template>
    <q-page padding class="q-gutter-y-md">
        <div class="row q-gutter-x-md items-center">
            <q-input style="width: 20%" v-model="queryParams.config_item" :label="$t('Config') + $t('Name')" />
            <q-input style="width: 20%" v-model="queryParams.memo" :label="$t('Config') + $t('Memo')" />
            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Config')"
                    v-has="'config-backend:add'" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:header-cell-item_custom="props">
                <q-th :props="props">
                    {{ props.col.label }}
                    <q-icon name="edit" size="1.3em" />
                </q-th>
            </template>

            <template v-slot:body-cell-item_custom="props">
                <q-td :props="props">
                    {{ props.row.item_custom }}
                    <q-popup-edit v-model="props.row.item_custom" :class="darkThemeSelect">
                        <template v-slot="scope">
                            {{ $t('Customize') + ' ' + props.row.config_item }}
                            <q-input v-model="props.row.item_custom" dense autofocus clearable
                                @keyup.enter="scope.set" />
                        </template>
                    </q-popup-edit>
                </q-td>
            </template>

            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <GqaDictShow :dictCode="props.row.status" />
                </q-td>
            </template>

            <template v-slot:body-cell-stable="props">
                <q-td :props="props">
                    <GqaDictShow :dictCode="props.row.stable" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="handleSave(props.row)" :label="$t('Save')"
                            v-has="'config-backend:save'" />
                        <q-btn color="warning" @click="handleReset(props.row)" :label="$t('Reset')"
                            v-has="'config-backend:reset'" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')"
                            v-has="'config-backend:delete'" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { useStorageStore } from 'src/stores/storage'
import recordDetail from './modules/recordDetail'
import useTheme from 'src/composables/useTheme';

const $q = useQuasar()
const { t } = useI18n()
const storageStore = useStorageStore()
const { darkThemeSelect } = useTheme()
const url = {
    list: 'config-backend/get-config-backend-list',
    edit: 'config-backend/edit-config-backend',
    delete: 'config-backend/delete-config-backend-by-id',
}
const columns = computed(() => {
    return [
        { name: 'sort', align: 'center', label: t('Sort'), field: 'sort' },
        { name: 'config_item', align: 'center', label: t('Config') + t('Item'), field: 'config_item' },
        { name: 'memo', align: 'center', label: t('Memo'), field: 'memo' },
        { name: 'item_default', align: 'center', label: t('Default'), field: 'item_default' },
        { name: 'item_custom', align: 'center', label: t('Custom'), field: 'item_custom' },
        { name: 'status', align: 'center', label: t('Status'), field: 'status' },
        { name: 'stable', align: 'center', label: t('Stable'), field: 'stable' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    onRequest,
    getTableData,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    getTableData()
})

const handleReset = (row) => {
    row.item_custom = ''
}

const handleSave = async (row) => {
    const res = await postAction(url.edit, row)
    if (res.code === 1) {
        $q.notify({
            type: 'positive',
            message: res.message,
        })
        storageStore.SetGqaBackend()
    }
}
</script>

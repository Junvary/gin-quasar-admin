<template>
    <div>
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
            <template v-slot:top>
                <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('NoteTodo')" />
            </template>
            <template v-slot:body-cell-todo_status="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusYesNo" :dictCode="props.row.todo_status" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="showEditForm(props.row)" :label="$t('Edit')" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <NoticeNoteTodoDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </div>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { DictOptions } from 'src/utils/dict'
import { FormatDateTime } from 'src/utils/date'
import NoticeNoteTodoDetail from './NoticeNoteTodoDetail'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'note-todo/get-note-todo-list',
}
const columns = computed(() => {
    return [
        { name: 'id', align: 'center', label: 'ID', field: 'id' },
        { name: 'todo_detail', align: 'center', label: t('Detail'), field: 'todo_detail' },
        { name: 'todo_status', align: 'center', label: t('Done'), field: 'todo_status' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    GqaAvatar,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    onRequest,
    getTableData,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    pagination.value.sortBy = 'created_at'
    getTableData()
})
</script>

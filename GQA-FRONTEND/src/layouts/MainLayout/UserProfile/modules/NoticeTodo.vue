<template>
    <div>
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest" style="max-height: 70vh;">
            <template v-slot:top>
                <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Todo')" />
            </template>
            <template v-slot:body-cell-todo_detail="props">
                <q-td :props="props" class="ellipsis" style="max-width: 400px;">
                    {{ props.row.todo_detail }}
                </q-td>
            </template>
            <template v-slot:body-cell-todo_status="props">
                <q-td :props="props">
                    <GqaDictShow :dictCode="props.row.todo_status" />
                </q-td>
            </template>

            <template v-slot:body-cell-created_at="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.created_at) }}
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
        <NoticeTodoDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </div>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import NoticeTodoDetail from './NoticeTodoDetail.vue'

const { t } = useI18n()
const url = {
    list: 'todo/get-todo-list',
}
const columns = computed(() => {
    return [
        { name: 'id', align: 'center', label: 'ID', field: 'id' },
        { name: 'todo_detail', align: 'center', label: t('Detail'), field: 'todo_detail' },
        { name: 'todo_status', align: 'center', label: t('Done'), field: 'todo_status' },
        { name: 'created_at', align: 'center', label: t('CreatedAt'), field: 'created_at' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    pagination,
    pageOptions,
    GqaDictShow,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    onRequest,
    getTableData,
    handleFinish,
    handleDelete,
    showDateTime,
} = useTableData(url)

onMounted(() => {
    pagination.value.sortBy = 'created_at'
    pagination.value.descending = true
    getTableData()
})
</script>

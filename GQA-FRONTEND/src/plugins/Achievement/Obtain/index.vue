<template>
    <q-page padding>
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
            <template v-slot:body-cell-created_by="props">
                <q-td :props="props">
                    <GqaShowName :customNameObject="props.row.created_by_user" />
                </q-td>
            </template>
            <template v-slot:body-cell-created_at="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.created_at) }}
                </q-td>
            </template>
        </q-table>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, onMounted } from 'vue'
import recordDetail from './modules/recordDetail.vue'
import GqaShowName from 'src/components/GqaShowName/index.vue'

const url = {
    list: 'plugin-achievement/get-obtain-list',
    delete: 'plugin-achievement/delete-obtain-by-id',
}
const columns = computed(() => {
    return [
        { name: 'category_code', align: 'center', label: '成就', field: 'category_code' },
        { name: 'created_by', align: 'center', label: '获得人', field: 'created_by' },
        { name: 'created_at', align: 'center', label: '获得时间', field: 'created_at' },
    ]
})
const {
    showDateTime,
    pagination,
    pageOptions,
    loading,
    tableData,
    recordDetailDialog,
    onRequest,
    getTableData,
    handleFinish,
} = useTableData(url)

onMounted(() => {
    pagination.value.sortBy = 'created_at'
    pagination.value.descending = true
    getTableData()
})
</script>

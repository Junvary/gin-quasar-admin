<template>
    <q-page padding>
        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.column_1" label="第1列" />
            <q-btn color="primary" @click="handleSearch" label="搜索" />
            <q-btn color="primary" @click="resetSearch" label="重置" />
        </div>
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <div class="row q-gutter-x-md">
                    <q-btn color="primary" @click="showAddForm()" label="新增测试数据" />
                    <q-btn color="primary" @click="downloadTemplate()" label="下载模板" />
                    <q-btn color="primary" @click="exportData()" label="导出查询数据" />
                </div>
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
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
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script setup>
import { downloadAction } from 'src/api/manage';
import useTableData from 'src/composables/useTableData'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import recordDetail from './modules/recordDetail'
import { FormatDateTimeShort } from 'src/utils/date'

const { t } = useI18n()
const url = {
    list: 'plugin-example/get-test-data-list',
    delete: 'plugin-example/delete-test-data-by-id',
    export: 'plugin-example/export-test-data',
    template: 'plugin-example/download-template-test-data'
}
const columns = computed(() => {
    return [
        { name: 'column1', align: 'center', label: '第1列', field: 'column_1' },
        { name: 'column2', align: 'center', label: '第2列', field: 'column_2' },
        { name: 'column3', align: 'center', label: '第3列', field: 'column_3' },
        { name: 'column4', align: 'center', label: '第4列', field: 'column_4' },
        { name: 'column5', align: 'center', label: '第5列', field: 'column_5' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    pagination,
    pageOptions,
    queryParams,
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
    pagination.value.descending = true
    getTableData()
})

const exportData = () => {
    const timeStamp = Date.now()
    const params = {}
    params.sort_by = pagination.value.sortBy
    params.desc = pagination.value.descending
    const allParams = Object.assign({}, params, queryParams.value)
    downloadAction(url.export, allParams, "GqaExport-" + FormatDateTimeShort(timeStamp) + '.xlsx')
}

const downloadTemplate = () => {
    const filename = 'TestDataTemplate.xlsx'
    downloadAction(url.template, {
        filename: filename
    }, filename)
}
</script>

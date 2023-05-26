<template>
    <q-page padding>
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" label="新增{{ .PluginModel.ModelName }}" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="showEditForm(props.row)" label="编辑" />
                        <q-btn color="negative" @click="handleDelete(props.row)" label="删除" />
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
import recordDetail from './modules/recordDetail'
import GqaShowName from 'src/components/GqaShowName'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'plugin-{{ .PluginCode }}/get-{{ .PluginModel.ModelName }}-list',
    delete: 'plugin-{{ .PluginCode }}/delete-{{ .PluginModel.ModelName }}-by-id',
}
const columns = computed(() => {
    return [
        {{ range .PluginModel.ColumnList }}
        { name: '{{ .ColumnName }}', align: 'center', label: '{{ .ColumnName }}', field: '{{ .ColumnName }}' },
        {{ end }}
        { name: 'actions', align: 'center', label: '操作', field: 'actions' },
    ]
})
const {
    showDateTime,
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
    pagination.value.descending = true
    getTableData()
})
</script>

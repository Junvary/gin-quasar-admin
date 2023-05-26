<template>
    <q-page padding>
        <q-card flat>
            <q-card-section class="row q-gutter-x-md items-center">
                <!-- <q-input outlined dense style="width: 20%" v-model="queryParams.notice_title" :label="$t('Title')" />
                <q-select outlined dense style="width: 20%" v-model="queryParams.notice_type"
                    :options="dictOptions.noticeType" emit-value map-options :label="$t('Notice') + $t('Type')" />
                <q-select outlined dense style="width: 20%" v-model="queryParams.notice_sent" :options="dictOptions.yesNo"
                    emit-value map-options :label="$t('Sent')"
                    :option-label="opt => Object(opt) === opt && 'label' in opt ? $t(opt.label) : '- Null -'" />
                <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
                <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" /> -->
            </q-card-section>
            <q-card-section>
                <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                    :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
                    <template v-slot:top="props">
                        <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Plugin')"
                            v-has="'genPlugin:add'" />
                        <q-space />
                        <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                            @click="props.toggleFullscreen" class="q-ml-md" />
                    </template>
                    <template v-slot:body-cell-created_at="props">
                        <q-td :props="props">
                            {{ showDateTime(props.row.created_at) }}
                        </q-td>
                    </template>
                    <template v-slot:body-cell-actions="props">
                        <q-td :props="props" class="q-gutter-x-xs">
                            <q-btn flat dense rounded icon="download" color="warning" @click="HandleDownload(props.row)"
                                v-has="'genPlugin:download'">
                                <q-tooltip>
                                    {{ $t('Download') }}
                                </q-tooltip>
                            </q-btn>
                            <q-btn flat dense rounded icon="delete_outline" color="negative"
                                @click="handleDelete(props.row)" v-has="'genPlugin:delete'">
                                <q-tooltip>
                                    {{ $t('Delete') }}
                                </q-tooltip>
                            </q-btn>
                        </q-td>
                    </template>
                </q-table>
            </q-card-section>
        </q-card>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { postBlobAction } from 'src/api/manage'
import { computed, onMounted } from 'vue'
import recordDetail from './modules/recordDetail.vue'

const url = {
    list: 'gen-plugin/get-gen-plugin-list',
    download: 'gen-plugin/download-gen-plugin-by-id',
    delete: 'gen-plugin/delete-gen-plugin-by-id',
}
const columns = computed(() => {
    return [
        { name: 'id', align: 'center', label: 'ID', field: 'id' },
        { name: 'plugin_sort', align: 'center', label: t('Plugin') + t('Sort'), field: 'plugin_sort' },
        { name: 'plugin_code', align: 'center', label: t('Plugin') + t('Code'), field: 'plugin_code' },
        { name: 'plugin_name', align: 'center', label: t('Plugin') + t('Name'), field: 'plugin_name' },
        { name: 'plugin_file', align: 'center', label: t('Plugin') + t('File'), field: 'plugin_file' },
        { name: 'created_at', align: 'center', label: t('CreatedAt'), field: 'created_at' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    $q,
    t,
    showDateTime,
    dictOptions,
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

onMounted(async () => {
    pagination.value.sortBy = 'id'
    pagination.value.descending = true
    getTableData()
})

const HandleDownload = async (row) => {
    const data = await postBlobAction(url.download, { id: row.id })
    $q.notify({
        type: 'positive',
        message: t('Gen') + t('Success'),
    })
    const blob = new Blob([data])
    const fileName = row.plugin_file.split('/').slice(-1)[0]
    if ('download' in document.createElement('a')) {
        // 不是IE浏览器
        const url = window.URL.createObjectURL(blob)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = url
        link.setAttribute('download', fileName)
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        window.URL.revokeObjectURL(url)
    } else {
        // IE 10+
        window.navigator.msSaveBlob(blob, fileName)
    }

}

</script>

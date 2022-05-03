<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.operationUsername" :label="$t('User')" />

            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-created_at="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.created_at) }}
                </q-td>
            </template>

        </q-table>
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { DictOptions } from 'src/utils/dict'
import { FormatDateTime } from 'src/utils/date'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'log/get-log-operation-list',
    delete: 'log/delete-log-operation-by-id',
}
const columns = computed(() => {
    return [
        { name: 'id', align: 'center', label: 'ID', field: 'id' },
        { name: 'operation_username', align: 'center', label: t('User'), field: 'operation_username' },
        { name: 'operation_ip', align: 'center', label: 'IP', field: 'operation_ip' },
        { name: 'operation_method', align: 'center', label: t('Method'), field: 'operation_method' },
        { name: 'operation_api', align: 'center', label: t('Api'), field: 'operation_api' },
        { name: 'operation_status', align: 'center', label: t('Status'), field: 'operation_status' },
        { name: 'created_at', align: 'center', label: t('CreatedAt'), field: 'created_at' },
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
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

const dictOptions = ref({})
onMounted(async () => {
    dictOptions.value = await DictOptions()
    pagination.value.sortBy = 'created_at'
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    })
})

const showDateTime = computed(() => {
    return (datetime) => {
        return FormatDateTime(datetime)
    }
})
</script>

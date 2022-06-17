<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.notice_title" :label="$t('Title')" />
            <q-select style="width: 20%" v-model="queryParams.notice_type" :options="dictOptions.noticeType" emit-value
                map-options :label="$t('Notice') + $t('Type')" />
            <q-select style="width: 20%" v-model="queryParams.notice_sent" :options="dictOptions.statusYesNo" emit-value
                map-options :label="$t('Sent')" />
            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Notice')" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-notice_type="props">
                <q-td :props="props">
                    <GqaDictShow dictName="noticeType" :dictCode="props.row.notice_type" />
                </q-td>
            </template>

            <template v-slot:body-cell-notice_sent="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusYesNo" :dictCode="props.row.notice_sent" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="warning" @click="sendMessage(props.row)" :label="$t('Send')"
                            v-if="props.row.notice_sent === 'no'" />
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
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import recordDetail from './modules/recordDetail'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'notice/get-notice-list',
    delete: 'notice/delete-notice-by-id',
    send: 'notice/send-notice',
}
const columns = computed(() => {
    return [
        { name: 'id', align: 'center', label: 'ID', field: 'id' },
        { name: 'notice_title', align: 'center', label: t('Title'), field: 'notice_title' },
        { name: 'notice_type', align: 'center', label: t('Type'), field: 'notice_type' },
        { name: 'notice_sent', align: 'center', label: t('Sent'), field: 'notice_sent' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    dictOptions,
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
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

onMounted(async () => {
    pagination.value.sortBy = 'created_at'
    pagination.value.descending = true
    getTableData()
})

const sendMessage = (row) => {
    $q.dialog({
        title: t('Confirm') + t('Send'),
        message: t('Confirm') + t('Send'),
        cancel: true,
        persistent: true,
    }).onOk(async () => {
        const res = await postAction(url.send, row)
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: t('Send') + ' ' + t('Success'),
            })
            getTableData()
        }
    })
}
</script>

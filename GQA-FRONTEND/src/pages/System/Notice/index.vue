<template>
    <q-page padding>
        <q-card flat>
            <q-card-section class="row q-gutter-x-md items-center">
                <q-input outlined dense style="width: 20%" v-model="queryParams.notice_title" :label="$t('Title')" />
                <q-select outlined dense style="width: 20%" v-model="queryParams.notice_type"
                    :options="dictOptions.noticeType" emit-value map-options :label="$t('Notice') + $t('Type')" />
                <q-select outlined dense style="width: 20%" v-model="queryParams.notice_sent"
                    :options="dictOptions.yesNo" emit-value map-options :label="$t('Sent')"
                    :option-label="opt => Object(opt) === opt && 'label' in opt ? $t(opt.label) : '- Null -'" />
                <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
                <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
            </q-card-section>
            <q-card-section>
                <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns"
                    v-model:pagination="pagination" :rows-per-page-options="pageOptions" :loading="loading"
                    @request="onRequest">
                    <template v-slot:top="props">
                        <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Notice')"
                            v-has="'notice:add'" />
                        <q-space />
                        <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                            @click="props.toggleFullscreen" class="q-ml-md" />
                    </template>
                    <template v-slot:body-cell-notice_type="props">
                        <q-td :props="props">
                            <GqaDictShow :dictCode="props.row.notice_type" />
                        </q-td>
                    </template>
                    <template v-slot:body-cell-notice_sent="props">
                        <q-td :props="props">
                            <GqaDictShow :dictCode="props.row.notice_sent" />
                        </q-td>
                    </template>
                    <template v-slot:body-cell-actions="props">
                        <q-td :props="props">
                            <div class="q-gutter-md">
                                <q-btn flat dense icon="send" color="warning" @click="sendMessage(props.row)"
                                    :label="$t('Send')" v-has="'notice:send'"
                                    v-if="props.row.notice_sent === 'yesNo_no'" />
                                <q-btn flat dense icon="eva-edit-2-outline" color="primary"
                                    @click="showEditForm(props.row)" :label="$t('Edit')" v-has="'notice:edit'" />
                                <q-btn flat dense icon="delete_outline" color="negative"
                                    @click="handleDelete(props.row)" :label="$t('Delete')" v-has="'notice:delete'" />
                            </div>
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

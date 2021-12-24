<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.noticeTitle" :label="$t('Title')" />
            <q-select style="width: 20%" v-model="queryParams.noticeType" :options="dictOptions.noticeType" emit-value
                map-options :label="$t('Notice') + $t('Type')" />
            <q-select style="width: 20%" v-model="queryParams.noticeSent" :options="dictOptions.statusYesNo" emit-value
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

            <template v-slot:body-cell-noticeType="props">
                <q-td :props="props">
                    <GqaDictShow dictName="noticeType" :dictCode="props.row.noticeType" />
                </q-td>
            </template>

            <template v-slot:body-cell-noticeSent="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusYesNo" :dictCode="props.row.noticeSent" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="warning" @click="sendMessage(props.row)" :label="$t('Send')"
                            v-if="props.row.noticeSent === 'no'" />
                        <q-btn color="primary" @click="showEditForm(props.row)" :label="$t('Edit')" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <add-or-edit-dialog ref="addOrEditDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import addOrEditDialog from './modules/addOrEditDialog'
import GqaDictShow from 'src/components/GqaDictShow'
import { DictOptions } from 'src/utils/dict'
import { postAction } from 'src/api/manage'

export default {
    name: 'Notice',
    mixins: [tableDataMixin],
    components: {
        addOrEditDialog,
        GqaDictShow,
    },
    computed: {
        columns() {
            return [
                { name: 'id', align: 'center', label: 'ID', field: 'id' },
                { name: 'noticeTitle', align: 'center', label: this.$t('Title'), field: 'noticeTitle' },
                { name: 'noticeType', align: 'center', label: this.$t('Type'), field: 'noticeType' },
                { name: 'noticeSent', align: 'center', label: this.$t('Sent'), field: 'noticeSent' },
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ]
        },
    },
    data() {
        return {
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 10,
            },
            url: {
                list: 'notice/notice-list',
                delete: 'notice/notice-delete',
                send: 'notice/notice-send',
            },
            dictOptions: {},
        }
    },
    async created() {
        this.getTableData()
        this.dictOptions = await DictOptions()
    },
    methods: {
        sendMessage(row) {
            this.$q
                .dialog({
                    title: this.$t('Confirm') + ' ' + this.$t('Send'),
                    message: this.$t('ConfirmSend'),
                    cancel: true,
                    persistent: true,
                })
                .onOk(async () => {
                    const res = await postAction(this.url.send, row)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: this.$t('Send') + ' ' + this.$t('Success'),
                        })
                        this.getTableData()
                    }
                })
        },
    },
}
</script>

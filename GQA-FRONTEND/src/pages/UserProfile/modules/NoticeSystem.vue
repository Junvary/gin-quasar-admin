<template>
    <div>
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
            <template v-slot:body-cell-noticeType="props">
                <q-td :props="props">
                    <GqaDictShow dictName="noticeType" :dictCode="props.row.noticeType" />
                </q-td>
            </template>

            <template v-slot:body-cell-noticeRead="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusYesNo"
                        :dictCode="props.row.noticeToUser.filter(item => item.toUser === username)[0].userRead" />
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
                        <q-btn color="primary" @click="readNotice(props.row)" :label="$t('Read')" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <NoticeDetail ref="noticeDetail" @hide="hide" />
    </div>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { mapGetters } from 'vuex'
import GqaDictShow from 'src/components/GqaDictShow'
import NoticeDetail from './NoticeDetail'

export default {
    name: 'NoticeSystem',
    mixins: [tableDataMixin],
    components: {
        GqaDictShow,
        NoticeDetail,
    },
    computed: {
        ...mapGetters({
            username: 'user/username',
        }),
        columns() {
            return [
                { name: 'id', align: 'center', label: 'ID', field: 'id' },
                { name: 'noticeTitle', align: 'center', label: this.$t('Title'), field: 'noticeTitle' },
                { name: 'noticeType', align: 'center', label: this.$t('Type'), field: 'noticeType' },
                { name: 'noticeRead', align: 'center', label: this.$t('Read') + this.$t('Status'), field: 'noticeRead' },
                { name: 'noticeSent', align: 'center', label: this.$t('Sent'), field: 'noticeSent' },
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ]
        },
    },
    data() {
        return {
            queryParams: {
                noticeType: 'system',
                noticeSent: 'yes',
                noticeToUser: String(this.username),
            },
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 10,
            },
            url: {
                list: 'notice/notice-list',
            },
        }
    },
    mounted() {
        this.queryParams.noticeToUser = String(this.username)
        this.getTableData()
    },
    methods: {
        readNotice(row) {
            this.$refs.noticeDetail.show(row)
        },
        hide() {
            this.getTableData()
        },
    },
}
</script>
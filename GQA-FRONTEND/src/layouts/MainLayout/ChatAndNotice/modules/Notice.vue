<template>
    <q-btn dense round glossy push color="primary" icon="notifications">
        <q-badge color="negative" floating v-if="tableData.length + todoNoteData.length">
            {{ tableData.length + todoNoteData.length }}
        </q-badge>
        <q-menu>
            <q-card>
                <q-tabs v-model="noticeType" dense class="text-grey" active-color="primary" indicator-color="primary"
                    align="justify" narrow-indicator style="padding: 10px">
                    <q-tab name="system" :label="$t('NoticeSystem')">
                        <q-badge color="negative" floating v-if="systemData.length">
                            {{ systemData.length }}
                        </q-badge>
                    </q-tab>
                    <q-tab name="message" :label="$t('NoticeMessage')">
                        <q-badge color="negative" floating v-if="messageData.length">
                            {{ messageData.length }}
                        </q-badge>
                    </q-tab>

                    <q-tab name="todoNote" :label="$t('TodoNote')">
                        <q-badge color="negative" floating v-if="todoNoteData.length">
                            {{ todoNoteData.length }}
                        </q-badge>
                    </q-tab>
                </q-tabs>

                <q-separator />

                <q-tab-panels v-model="noticeType" animated>
                    <q-tab-panel style="padding: 0" name="system">
                        <NoticeSystem :systemData="systemData" />
                    </q-tab-panel>

                    <q-tab-panel style="padding: 0" name="message">
                        <NoticeMessage :messageData="messageData" />
                    </q-tab-panel>

                    <q-tab-panel style="padding: 0" name="todoNote">
                        <NoticeTodoNote :todoNoteData="todoNoteData" />
                    </q-tab-panel>
                </q-tab-panels>
            </q-card>

        </q-menu>
    </q-btn>

</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import NoticeSystem from './NoticeSystem'
import NoticeMessage from './NoticeMessage'
import NoticeTodoNote from './NoticeTodoNote'
import { mapGetters } from 'vuex'
import { postAction } from 'src/api/manage'

export default {
    name: 'Notice',
    mixins: [tableDataMixin],
    components: {
        NoticeSystem,
        NoticeMessage,
        NoticeTodoNote,
    },
    computed: {
        ...mapGetters({
            username: 'user/username',
        }),
        systemData() {
            return this.tableData.filter((item) => item.noticeType === 'system')
        },
        messageData() {
            return this.tableData.filter((item) => item.noticeType === 'message')
        },
    },
    data() {
        return {
            noticeType: 'system',
            queryParams: {
                noticeRead: 'no',
                noticeSent: 'yes',
                noticeToUser: String(this.username),
            },
            todoQueryParams: {
                todoStatus: 'no',
            },
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 9999,
            },
            url: {
                list: 'notice/notice-list',
                todoNoteList: 'todo-note/todo-note-list',
            },
            todoNoteData: [],
        }
    },
    mounted() {
        this.$bus.on('noticeGetTableData', () => {
            this.getTableData()
            this.getTodoNoteData({ pagination: this.pagination })
        })
        this.queryParams.noticeToUser = String(this.username)
        console.log(typeof this.queryParams.noticeToUser, this.queryParams.noticeToUser)
        this.getTableData()
        this.getTodoNoteData({ pagination: this.pagination })
    },
    methods: {
        async getTodoNoteData(props) {
            if (this.url === undefined || !this.url.todoNoteList) {
                this.$q.notify({
                    type: 'negative',
                    message: this.$t('UrlNotConfig'),
                })
                return
            }
            this.todoNoteData = []
            // 组装分页和过滤条件
            const params = {}
            params.sortBy = props.pagination.sortBy
            params.desc = props.pagination.descending
            params.page = props.pagination.page
            params.pageSize = props.pagination.rowsPerPage
            const allParams = Object.assign({}, params, this.todoQueryParams)
            // 带参数请求数据
            await postAction(this.url.todoNoteList, allParams)
                .then((res) => {
                    if (res.code === 1) {
                        // 最终要把分页给同步掉
                        this.pagination = props.pagination
                        this.todoNoteData = res.data.records
                    }
                })
                .finally(() => {
                    this.loading = false
                })
        },
    },
}
</script>

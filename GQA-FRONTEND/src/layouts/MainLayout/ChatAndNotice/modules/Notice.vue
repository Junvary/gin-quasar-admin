<template>
    <q-btn dense round glossy push color="primary" icon="notifications">
        <q-badge color="negative" floating>
            {{ tableData.length + 4 }}
        </q-badge>
        <q-menu>
            <q-card>
                <q-tabs v-model="messageType" dense class="text-grey" active-color="primary" indicator-color="primary"
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

                    <q-tab name="todo" :label="$t('NoticeToDo')">
                        <q-badge color="negative" floating>4</q-badge>
                    </q-tab>
                </q-tabs>

                <q-separator />

                <q-tab-panels v-model="messageType" animated>
                    <q-tab-panel style="padding: 0" name="system">
                        <NoticeSystem :systemData="systemData" />
                    </q-tab-panel>

                    <q-tab-panel style="padding: 0" name="message">
                        <NoticeMessage :messageData="messageData" />
                    </q-tab-panel>

                    <q-tab-panel style="padding: 0" name="todo">
                        <q-list bordered separator style="min-width: 300px">
                            <q-item clickable v-ripple v-for="(item, index) in todo" :key="index">
                                <q-item-section avatar>
                                    <q-icon color="primary" name="list" />
                                </q-item-section>

                                <q-item-section>
                                    {{item.title}}
                                </q-item-section>
                            </q-item>
                        </q-list>
                        <q-item clickable v-ripple class="text-center">
                            <q-item-section>{{ $t('CheckAll') }}</q-item-section>
                        </q-item>
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
import { mapGetters } from 'vuex'

export default {
    name: 'Notice',
    mixins: [tableDataMixin],
    components: {
        NoticeSystem,
        NoticeMessage,
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
            messageType: 'system',
            queryParams: {
                noticeRead: 'no',
                noticeSent: 'yes',
                noticeToUser: this.username,
            },
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 9999,
            },
            url: {
                list: 'notice/notice-list',
            },
            todo: [],
        }
    },
    mounted() {
        this.$bus.on('noticeGetTableData', () => {
            this.getTableData()
        })
        this.queryParams.noticeToUser = this.username
        this.getTableData()
        const t = {
            title: this.$t('NoticeToDoNew'),
        }

        this.todo.push(t, t, t, t)
    },
}
</script>

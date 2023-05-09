<template>
    <q-btn dense round flat icon="ion-ios-notifications-outline">
        <q-tooltip>
            {{ $t('Notice') }}
        </q-tooltip>
        <q-badge color="negative" floating rounded v-if="systemNum + messageNum + todoNum">
            {{ (systemNum + messageNum + todoNum) > 99
                ? "99+" :
                (systemNum + messageNum + todoNum)
            }}
        </q-badge>
        <q-menu anchor="bottom start" self="top middle">
            <q-card style="max-width: 400px">
                <q-tabs v-model="noticeType" dense class="text-grey" active-color="primary" indicator-color="primary"
                    align="justify" narrow-indicator style="padding: 10px">
                    <q-tab name="system" :label="$t('NoticeSystem')">
                        <q-badge color="negative" floating v-if="systemNum">
                            {{ systemNum > 99 ? '99+' : systemNum }}
                        </q-badge>
                    </q-tab>
                    <q-tab name="message" :label="$t('NoticeMessage')">
                        <q-badge color="negative" floating v-if="messageNum">
                            {{ messageNum > 99 ? '99+' : messageNum }}
                        </q-badge>
                    </q-tab>

                    <q-tab name="todo" :label="$t('Todo')">
                        <q-badge color="negative" floating v-if="todoNum">
                            {{ todoNum > 99 ? '99+' : todoNum }}
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

                    <q-tab-panel style="padding: 0" name="todo">
                        <NoticeTodo :todoData="todoData" />
                    </q-tab-panel>
                </q-tab-panels>
            </q-card>
        </q-menu>
    </q-btn>
</template>

<script setup>
import { onMounted, ref, inject, computed } from 'vue'
import NoticeSystem from './NoticeSystem.vue'
import NoticeMessage from './NoticeMessage.vue'
import NoticeTodo from './NoticeTodo.vue'
import { postAction } from 'src/api/manage';
import { useUserStore } from 'src/stores/user'

const noticeType = ref('system')
const userStore = useUserStore()
const username = computed(() => userStore.GetUsername())
const bus = inject('bus')
const url = {
    list: 'notice/get-notice-list',
    todo: 'todo/get-todo-list'
}

const getNoticeAndTodoData = () => {
    getNoticeSystem()
    getNoticeMessage()
    getNoticeTodo()
}

const systemData = ref([])
const systemNum = ref(0)
const getNoticeSystem = () => {
    postAction(url.list, {
        notice_read: 'yesNo_no',
        notice_sent: 'yesNo_yes',
        notice_to_user: String(username.value),
        notice_type: 'noticeType_system',
        sort_by: 'created_at',
        desc: true,
        page: 1,
        page_size: 10
    }).then(res => {
        if (res.code === 1) {
            systemData.value = res.data.records
            systemNum.value = res.data.total
        }
    })
}

const messageData = ref([])
const messageNum = ref(0)
const getNoticeMessage = () => {
    postAction(url.list, {
        notice_read: 'yesNo_no',
        notice_sent: 'yesNo_yes',
        notice_to_user: String(username.value),
        notice_type: 'noticeType_message',
        sort_by: 'created_at',
        desc: true,
        page: 1,
        page_size: 10
    }).then(res => {
        if (res.code === 1) {
            messageData.value = res.data.records
            messageNum.value = res.data.total
        }
    })
}

const todoData = ref([])
const todoNum = ref(0)
const getNoticeTodo = () => {
    postAction(url.todo, {
        todo_status: 'yesNo_no',
        sort_by: 'created_at',
        desc: true,
        page: 1,
        page_size: 10
    }).then(res => {
        if (res.code === 1) {
            todoData.value = res.data.records
            todoNum.value = res.data.total
        }
    })
}

onMounted(() => {
    bus.on('noticeGetTableData', () => {
        getNoticeAndTodoData()
    })
    getNoticeAndTodoData()
})

defineExpose({
    getNoticeAndTodoData
})
</script>

<style lang="scss" scoped>
.q-badge--floating {
    top: -1px;
    right: -4px;
}
</style>

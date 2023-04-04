<template>
    <q-btn dense round flat icon="notifications_none">
        <q-badge color="negative" floating rounded v-if="systemNum + messageNum + noteTodoNum">
            {{ (systemNum + messageNum + noteTodoNum) > 99
                ? "99+" :
                (systemNum + messageNum + noteTodoNum)
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

                    <q-tab name="noteTodo" :label="$t('NoteTodo')">
                        <q-badge color="negative" floating v-if="noteTodoNum">
                            {{ noteTodoNum > 99 ? '99+' : noteTodoNum }}
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

                    <q-tab-panel style="padding: 0" name="noteTodo">
                        <NoticeNoteTodo :noteTodoData="noteTodoData" />
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
import NoticeNoteTodo from './NoticeNoteTodo.vue'
import { postAction } from 'src/api/manage';
import { useUserStore } from 'src/stores/user'

const noticeType = ref('system')
const userStore = useUserStore()
const username = computed(() => userStore.GetUsername())
const bus = inject('bus')
const url = {
    list: 'notice/get-notice-list',
    todo: 'note-todo/get-note-todo-list'
}

const getNoticeData = () => {
    getNoticeSystem()
    getNoticeMessage()
    getNoticeNoteTodo()
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

const noteTodoData = ref([])
const noteTodoNum = ref(0)
const getNoticeNoteTodo = () => {
    postAction(url.todo, {
        todo_status: 'yesNo_no',
        sort_by: 'created_at',
        desc: true,
        page: 1,
        page_size: 10
    }).then(res => {
        if (res.code === 1) {
            noteTodoData.value = res.data.records
            noteTodoNum.value = res.data.total
        }
    })
}

onMounted(() => {
    bus.on('noticeGetTableData', () => {
        getNoticeData()
    })
    getNoticeData()
})

defineExpose({
    getNoticeData
})

</script>

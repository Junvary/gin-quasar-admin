<template>
    <q-btn dense round flat icon="notifications_none">
        <q-badge color="negative" floating rounded v-if="tableData.length + noteTodoData.length">
            {{ (pagination.rowsNumber + noteTodoData.length) > 99
                    ? "99+" :
                    (pagination.rowsNumber + noteTodoData.length)
            }}
        </q-badge>
        <q-menu anchor="bottom start" self="top middle">
            <q-card style="max-width: 400px">
                <q-tabs v-model="noticeType" dense class="text-grey" active-color="primary" indicator-color="primary"
                    align="justify" narrow-indicator style="padding: 10px">
                    <q-tab name="system" :label="$t('NoticeSystem')">
                        <q-badge color="negative" floating v-if="systemData.length">
                            {{ systemData.length == 10 ? "10+" : systemData.length }}
                        </q-badge>
                    </q-tab>
                    <q-tab name="message" :label="$t('NoticeMessage')">
                        <q-badge color="negative" floating v-if="messageData.length">
                            {{ messageData.length == 10 ? "10+" : messageData.length }}
                        </q-badge>
                    </q-tab>

                    <q-tab name="noteTodo" :label="$t('NoteTodo')">
                        <q-badge color="negative" floating v-if="noteTodoData.length">
                            {{ noteTodoData.length == 10 ? "10+" : noteTodoData.length }}
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
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref, inject } from 'vue'
import { useI18n } from 'vue-i18n'
import NoticeSystem from './NoticeSystem.vue'
import NoticeMessage from './NoticeMessage.vue'
import NoticeNoteTodo from './NoticeNoteTodo.vue'
import { useUserStore } from 'src/stores/user'

const bus = inject('bus')
const $q = useQuasar()
const { t } = useI18n()
const userStore = useUserStore()
const url = {
    list: 'notice/get-notice-list',
    noteTodoList: 'note-todo/get-note-todo-list',
}
const {
    pagination,
    queryParams,
    loading,
    tableData,
    getTableData,
} = useTableData(url)

const username = computed(() => userStore.GetUsername())
const systemData = computed(() => tableData.value.filter((item) => item.notice_type === 'noticeType_system'))
const messageData = computed(() => tableData.value.filter((item) => item.notice_type === 'noticeType_message'))

const noticeType = ref('system')

onMounted(() => {
    queryParams.value = {
        notice_read: 'yesNo_no',
        notice_sent: 'yesNo_yes',
        notice_to_user: String(username.value),
    }
    pagination.value.sortBy = 'created_at'
    pagination.value.descending = true
    bus.on('noticeGetTableData', () => {
        getTableData()
        getNoteTodoData({ pagination: pagination.value })
    })
    queryParams.value.notice_to_user = String(username.value)
    getTableData()
    getNoteTodoData({ pagination: pagination.value })
})

const todoQueryParams = {
    todo_status: 'yesNo_no',
}

defineExpose({
    getTableData
})

const noteTodoData = ref([])

const getNoteTodoData = async (props) => {
    if (url === undefined || !url.noteTodoList) {
        $q.notify({
            type: 'negative',
            message: t('UrlNotConfig'),
        })
        return
    }
    noteTodoData.value = []
    const params = {}
    params.sort_by = props.pagination.sortBy
    params.desc = props.pagination.descending
    params.page = props.pagination.page
    params.page_size = props.pagination.rowsPerPage
    const allParams = Object.assign({}, params, todoQueryParams)
    await postAction(url.noteTodoList, allParams).then((res) => {
        if (res.code === 1) {
            pagination.value = props.pagination
            noteTodoData.value = res.data.records
        }
    }).finally(() => {
        loading.value = false
    })
}
</script>

<template>
    <q-tabs v-model="messageType" dense class="text-grey" active-color="primary" indicator-color="primary" align="justify"
        narrow-indicator style="width: 100%">
        <q-tab name="noticeType_system" :label="$t('NoticeSystem')" />
        <q-tab name="noticeType_message" :label="$t('NoticeMessage')" />
        <q-tab name="todo" :label="$t('Todo')" />
    </q-tabs>
    <q-tab-panels v-model="messageType" animated style="width: 100%">
        <q-tab-panel name="noticeType_system">
            <NoticeSystem v-if="messageType === 'noticeType_system'" />
        </q-tab-panel>

        <q-tab-panel name="noticeType_message">
            <NoticeMessage v-if="messageType === 'noticeType_message'" />
        </q-tab-panel>

        <q-tab-panel name="todo">
            <NoticeTodo v-if="messageType === 'todo'" />
        </q-tab-panel>
    </q-tab-panels>
</template>

<script setup>
import NoticeSystem from './NoticeSystem.vue'
import NoticeMessage from './NoticeMessage.vue'
import NoticeTodo from './NoticeTodo.vue'
import { ref } from 'vue';

const messageType = ref('noticeType_system')
const changeMessageType = (type) => {
    messageType.value = type
}
defineExpose({
    changeMessageType
})
</script>

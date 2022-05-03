<template>
    <q-tabs v-model="messageType" dense class="text-grey" active-color="primary" indicator-color="primary"
        align="justify" narrow-indicator style="width: 100%">
        <q-tab name="system" :label="$t('NoticeSystem')" />
        <q-tab name="message" :label="$t('NoticeMessage')" />
        <q-tab name="noteTodo" :label="$t('NoteTodo')" />
    </q-tabs>
    <q-tab-panels v-model="messageType" animated style="width: 100%">
        <q-tab-panel name="system">
            <NoticeSystem v-if="messageType === 'system'" />
        </q-tab-panel>

        <q-tab-panel name="message">
            <NoticeMessage v-if="messageType === 'message'" />
        </q-tab-panel>

        <q-tab-panel name="noteTodo">
            <NoticeNoteTodo v-if="messageType === 'noteTodo'" />
        </q-tab-panel>
    </q-tab-panels>
</template>

<script setup>
import NoticeSystem from './NoticeSystem'
import NoticeMessage from './NoticeMessage'
import NoticeNoteTodo from './NoticeNoteTodo'
import { ref } from 'vue';

const messageType = ref('system')
const changeMessageType = (type) => {
    messageType.value = type
}
defineExpose({
    changeMessageType
})
</script>

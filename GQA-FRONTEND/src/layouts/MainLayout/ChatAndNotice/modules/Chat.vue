<template>
    <q-btn dense round flat icon="mdi-chat-processing-outline" @click="showChat">
        <q-badge color="negative" floating v-if="badgeCount">
            {{ badgeCount }}
        </q-badge>
        <q-tooltip>
            {{ $t('ChatRoom') }}
        </q-tooltip>
    </q-btn>
    <ChatDialog ref="chatDialog" :oldMessage="oldMessage" @onSendMessage="onSendMessage" @changeShow="changeShow" />
</template>

<script setup>
import { ref, toRefs } from 'vue';
import ChatDialog from './ChatDialog.vue'

const props = defineProps({
    oldMessage: {
        type: Array,
        required: true,
        default: () => [],
    },
})
const { oldMessage } = toRefs(props)
const badgeCount = ref(0)
const chatDialog = ref(null)
const showChat = () => {
    chatDialog.value.show()
    badgeCount.value = 0
}
const emit = defineEmits(['sendMessage', 'changeChatDialogShow'])
const changeShow = (event) => {
    emit('changeChatDialogShow', event)
}
const onSendMessage = (event) => {
    emit('sendMessage', event)
}
const receiveMessage = (count) => {
    badgeCount.value += count
}
defineExpose({
    receiveMessage
})
</script>

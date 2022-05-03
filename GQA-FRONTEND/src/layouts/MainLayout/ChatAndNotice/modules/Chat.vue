<template>
    <div>
        <q-btn dense round glossy push color="primary" icon="groups" @click="showChat">
            <q-badge color="negative" floating v-if="badgeCount">
                {{ badgeCount }}
            </q-badge>
        </q-btn>
        <ChatDialog ref="chatDialog" :oldMessage="oldMessage" @sendMessage="sendMessage" @changeShow="changeShow" />
    </div>
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
    chatDialog.show()
    badgeCount.value = 0
}
const emit = defineEmits(['sendMessage'])
const changeShow = (event) => {
    emit('changeChatDialogShow', event)
}
const sendMessage = (event) => {
    emit('sendMessage', event)
}
const receiveMessage = (count) => {
    badgeCount.value += count
}
defineExpose({
    receiveMessage
})
</script>

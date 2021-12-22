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

<script>
import ChatDialog from './ChatDialog.vue'

export default {
    name: 'Chat',
    components: {
        ChatDialog,
    },
    props: {
        oldMessage: {
            type: Array,
            required: true,
            default: () => [],
        },
    },
    data() {
        return {
            badgeCount: 0,
        }
    },
    methods: {
        showChat() {
            this.$refs.chatDialog.show()
            this.badgeCount = 0
        },
        changeShow(event) {
            this.$emit('changeChatDialogShow', event)
        },
        sendMessage(event) {
            this.$emit('sendMessage', event)
        },
        receiveMessage(count) {
            this.badgeCount += count
        },
    },
}
</script>

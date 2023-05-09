<template>
    <div class="row q-gutter-x-sm" style="margin-left: 0;">
        <Chat ref="moduleChat" :oldMessage="chatOldMessage" @changeChatDialogShow="changeChatDialogShow"
            @sendMessage="sendMessage" />
        <Notice ref="moduleNotice" />
    </div>
</template>

<script setup>
import Chat from './modules/Chat.vue'
import Notice from './modules/Notice.vue'
import { useUserStore } from 'src/stores/user'
import { computed, onMounted, ref, onUnmounted } from 'vue';
import { useQuasar } from 'quasar';
import useConfig from 'src/composables/useConfig';

const { GqaDefaultUsername } = useConfig()
const $q = useQuasar()
const userStore = useUserStore()
const username = computed(() => userStore.GetUsername())
const myName = computed(() => {
    const nickname = $q.cookies.get('gqa-nickname')
    const realName = $q.cookies.get('gqa-realName')
    if (nickname) {
        return nickname
    } else if (realName) {
        return realName
    } else {
        return GqaDefaultUsername
    }
})
const ws = ref(null)
const lockReconnect = ref(false)
const chatDialogShow = ref(false)
const chatOldMessage = ref([])
const timer = ref(null)

onMounted(() => {
    initWebSocket()
})
onUnmounted(() => {
    websocketOnclose()
})

const changeChatDialogShow = (event) => {
    chatDialogShow.value = event
}
const initWebSocket = () => {
    ws.value = new WebSocket(process.env.API.replace('https://', 'wss://').replace('http://', 'ws://') + 'public/ws/' + username)
    ws.value.onopen = websocketOnopen
    ws.value.onerror = websocketOnerror
    ws.value.onmessage = websocketOnmessage
    ws.value.onclose = websocketOnclose
}
const websocketOnopen = () => {
    console.log('Gin-Quasar-Admin: WebSocket Connection Succeeded!')
}
const websocketOnerror = () => {
    console.log('Gin-Quasar-Admin: WebSocket Connection Error!')
    reconnect()
}
const moduleChat = ref(null)
const moduleNotice = ref(null)
const websocketOnmessage = (e) => {
    let newData = JSON.parse(e.data)
    if (newData.message_type === 'chat') {
        // chat message
        newData.text = [newData.text]
        if (newData.name === myName.value) {
            newData.sent = true
        } else {
            newData.sent = false
        }
        chatOldMessage.value.push(newData)
        if (!chatDialogShow.value) {
            moduleChat.value.receiveMessage(1)
        }
    } else {
        moduleNotice.value.getNoticeAndTodoData()
    }
}
const websocketOnclose = (e) => {
    if (e && e.code) {
        console.log('Gin-Quasar-Admin: WebSocket Connection Closed:', e, e.code)
        reconnect()
    } else {
        console.log('Gin-Quasar-Admin: WebSocket Connection Closed')
        clearInterval(timer.value)
        timer.value = null
    }
}
const reconnect = () => {
    if (lockReconnect.value) return
    lockReconnect.value = true
    timer.value = setTimeout(() => {
        console.log('Gin-Quasar-Admin: WebSocket Reconnecting...')
        initWebSocket()
        lockReconnect.value = false
    }, 5000)
}
const sendMessage = (event) => {
    try {
        ws.value.send(JSON.stringify(event))
    } catch (error) {
        console.log(error)
    }
}
</script>

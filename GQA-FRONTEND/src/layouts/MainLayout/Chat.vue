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
    computed: {
        myName() {
            const nickname = this.$q.cookies.get('gqa-nickname')
            const realName = this.$q.cookies.get('gqa-realName')
            if (nickname) {
                return nickname
            } else if (realName) {
                return realName
            } else {
                return GqaUsername
            }
        },
    },
    data() {
        return {
            websock: null,
            lockReconnect: false,
            oldMessage: [],
            badgeCount: 0,
            chatDialogShow: false,
        }
    },
    mounted() {
        this.initWebSocket()
    },
    methods: {
        showChat() {
            this.$refs.chatDialog.show()
            this.badgeCount = 0
        },
        changeShow(event) {
            this.chatDialogShow = event
        },
        initWebSocket() {
            this.websock = new WebSocket('ws://127.0.0.1:8888/public/ws')
            this.websock.onopen = this.websocketOnopen
            this.websock.onerror = this.websocketOnerror
            this.websock.onmessage = this.websocketOnmessage
            this.websock.onclose = this.websocketOnclose
        },
        websocketOnopen: function () {
            console.log('Gin-Quasar-Admin: WebSocket连接成功!')
        },
        websocketOnerror: function (e) {
            console.log(e)
            console.log('Gin-Quasar-Admin: WebSocket连接发生错误，开始重新连接!')
            this.reconnect()
        },
        websocketOnmessage: function (e) {
            let newData = JSON.parse(e.data)
            newData.text = [newData.text]
            if (newData.name === this.myName) {
                newData.sent = true
            } else {
                newData.sent = false
            }
            this.oldMessage.push(newData)
            if (!this.chatDialogShow) {
                this.badgeCount += 1
            }
        },
        websocketOnclose: function (e) {
            console.log('Gin-Quasar-Admin: WebSocket连接已关闭 (' + e + ')')
            console.log(e)
            if (e) {
                console.log('Gin-Quasar-Admin: WebSocket连接已关闭 (' + e.code + ')')
            }
        },
        reconnect() {
            var that = this
            if (that.lockReconnect) return
            that.lockReconnect = true
            //没连接上会一直重连，设置延迟避免请求过多
            setTimeout(function () {
                console.info('Gin-Quasar-Admin: WebSocket尝试重连...')
                that.initWebSocket()
                that.lockReconnect = false
            }, 5000)
        },
        sendMessage(event) {
            try {
                this.websock.send(JSON.stringify(event))
            } catch (error) {
                console.log(error)
            }
        },
    },
}
</script>

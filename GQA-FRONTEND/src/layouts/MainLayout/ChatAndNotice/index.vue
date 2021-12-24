<template>
    <div class="row">
        <Chat style="margin: 0 5px" ref="moduleChat" :oldMessage="chatOldMessage"
            @changeChatDialogShow="changeChatDialogShow" @sendMessage="sendMessage" />
        <Notice style="margin: 0 5px" ref="moduleNotice" />
    </div>
</template>

<script>
import Chat from './modules/Chat'
import Notice from './modules/Notice'
import { GqaUsername } from 'src/settings'
import { mapGetters } from 'vuex'

export default {
    name: 'ChatAndNotice',
    components: {
        Chat,
        Notice,
    },
    computed: {
        ...mapGetters({
            username: 'user/username',
        }),
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
            chatDialogShow: false,
            chatOldMessage: [],
        }
    },
    mounted() {
        this.initWebSocket()
    },
    beforeUnmount() {
        this.websocketOnclose()
    },
    methods: {
        changeChatDialogShow(event) {
            this.chatDialogShow = event
        },
        initWebSocket() {
            this.websock = new WebSocket('ws://127.0.0.1:8888/public/ws/' + this.username)
            this.websock.onopen = this.websocketOnopen
            this.websock.onerror = this.websocketOnerror
            this.websock.onmessage = this.websocketOnmessage
            this.websock.onclose = this.websocketOnclose
        },
        websocketOnopen() {
            console.log('Gin-Quasar-Admin: WebSocket连接成功!')
        },
        websocketOnerror() {
            console.log('Gin-Quasar-Admin: WebSocket连接发生错误!')
            this.reconnect()
        },
        websocketOnmessage(e) {
            let newData = JSON.parse(e.data)
            console.log('Gin-Quasar-Admin: 收到新消息', e.data)
            if (newData.messageType === 'chat') {
                // 聊天信息
                newData.text = [newData.text]
                if (newData.name === this.myName) {
                    newData.sent = true
                } else {
                    newData.sent = false
                }
                this.chatOldMessage.push(newData)
                if (!this.chatDialogShow) {
                    this.$refs.moduleChat.receiveMessage(1)
                }
            } else {
                this.$refs.moduleNotice.getTableData()
            }
        },
        websocketOnclose(e) {
            if (e.code) {
                console.log('Gin-Quasar-Admin: WebSocket连接已关闭:', e.code)
            } else {
                console.log('Gin-Quasar-Admin: WebSocket连接已关闭:', e)
            }
            this.reconnect()
        },
        reconnect() {
            const that = this
            if (that.lockReconnect) return
            that.lockReconnect = true
            //没连接上会一直重连，设置延迟避免请求过多
            setTimeout(function () {
                console.log('Gin-Quasar-Admin: WebSocket尝试重连...')
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
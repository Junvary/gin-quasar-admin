<template>
    <q-dialog persistent v-model="chatDialogVisible" transition-hide="scale" @before-show="beforeShow"
        @before-hide="beforeHide">
        <q-card bordered style="width: 750px; max-width: 45vw;">
            <q-bar class="bg-primary text-white">
                <q-space />
                <q-btn dense flat icon="close" v-close-popup />
            </q-bar>
            <q-card-section horizontal style="height: 50vh">
                <q-card-section class="col-4" style="padding: 0">
                    <q-scroll-area ref="userScroll" visible style="height: 50vh; width: 100%">
                        <q-list>
                            <q-item clickable v-ripple v-for="(item, index) in tableData" :key="index">
                                <q-item-section avatar>
                                    <GqaAvatar :src="item.avatar" />
                                </q-item-section>
                                <q-item-section>
                                    <q-item-label>
                                        <GqaShowName :customNameObject="item" />
                                    </q-item-label>
                                    <q-item-label caption>
                                        xxxxxx???
                                    </q-item-label>
                                </q-item-section>
                            </q-item>
                        </q-list>
                    </q-scroll-area>
                    <q-inner-loading :showing="loading">
                        <q-spinner-gears size="50px" color="primary" />
                    </q-inner-loading>
                </q-card-section>

                <q-separator vertical inset />

                <q-card-section style="width: 100%; padding: 0">
                    <q-scroll-area ref="messageScroll" visible style="height: 50vh; width: 100%">
                        <div style="margin-right: 15px; margin-left: 5px" v-if="oldMessage.length">
                            <q-chat-message v-for="(item, index) in oldMessage" :key="index" :name="item.name"
                                :avatar="item.avatar" :text="item.text" :sent="item.sent" :stamp="item.stamp" />
                        </div>
                    </q-scroll-area>
                </q-card-section>
            </q-card-section>

            <q-separator />

            <q-card-actions>
                <q-form ref="newMessageForm" style="width: 100%">
                    <q-input v-model="newMessage" outlined type="textarea"
                        :rules="[ val => val && val.length > 0 || $t('NeedInput') ]">
                        <template v-slot:after>
                            <q-btn color="primary" round dense flat icon="send" @click="sendMessage" />
                        </template>
                    </q-input>
                </q-form>
            </q-card-actions>
        </q-card>
    </q-dialog>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import GqaAvatar from 'src/components/GqaAvatar'
import GqaShowName from 'src/components/GqaShowName'
import { GqaLoginAvatar, GqaUsername } from 'src/settings'

export default {
    name: 'ChatDialog',
    mixins: [tableDataMixin],
    props: {
        oldMessage: {
            type: Array,
            required: false,
            default: () => [],
        },
    },
    components: {
        GqaAvatar,
        GqaShowName,
    },
    computed: {
        myAvatar() {
            const cookieAvatar = this.$q.cookies.get('gqa-avatar')
            // 登录用户的头像
            if (cookieAvatar) {
                // 如果Cookies里拿到头像，就显示
                if (cookieAvatar.substring(0, 4) === 'http') {
                    // 头像为链接
                    return cookieAvatar
                } else if (cookieAvatar.substring(0, 11) === 'gqa-upload:') {
                    // 头像为上传
                    return process.env.API + cookieAvatar.substring(11)
                }
                return cookieAvatar
            } else {
                // 否则显示默认头像
                return GqaLoginAvatar
            }
        },
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
    watch: {
        oldMessage: {
            handler() {
                // 只有在Dialog打开的时候滚动到最底部
                if (this.chatDialogVisible) {
                    this.$nextTick(() => {
                        this.$refs.messageScroll.setScrollPosition('vertical', this.oldMessage.length * 75, 200)
                    })
                }
            },
            deep: true,
        },
    },
    data() {
        return {
            chatDialogVisible: false,
            url: {
                list: 'user/user-list',
            },
            newMessage: '',
        }
    },
    methods: {
        show() {
            this.chatDialogVisible = true
            this.getTableData()
            this.$nextTick(() => {
                // userList不要滚动到最底部了，保持在最上面
                // this.$refs.userScroll.setScrollPosition('vertical', this.tableData.length * 55)
                this.$refs.messageScroll.setScrollPosition('vertical', this.oldMessage.length * 75)
            })
        },
        beforeShow() {
            this.$emit('changeShow', true)
        },
        beforeHide() {
            this.$emit('changeShow', false)
        },
        async sendMessage() {
            const success = await this.$refs.newMessageForm.validate()
            if (success) {
                this.$emit('sendMessage', {
                    name: this.myName,
                    avatar: this.myAvatar,
                    text: this.newMessage,
                    messageType: 'chat',
                })
                this.newMessage = ''
                this.$refs.newMessageForm.reset()
                this.$refs.newMessageForm.focus()
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: this.$t('FixForm'),
                })
            }
        },
    },
}
</script>
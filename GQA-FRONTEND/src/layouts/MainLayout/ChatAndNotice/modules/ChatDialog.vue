<template>
    <q-dialog persistent v-model="chatDialogVisible" transition-hide="scale" @before-show="beforeShow"
        @before-hide="beforeHide">
        <q-card bordered style="width: 750px; max-width: 45vw;">
            <q-bar class="bg-primary text-white">
                {{ gqaFrontend.subTitle }}
                {{ $t('ChatRoom') }}
                <q-space />
                <q-btn dense flat icon="close" v-close-popup />
            </q-bar>
            <q-card-section horizontal style="height: 50vh">
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
                <q-form ref="newMessageForm" style="width: 100%" class="gqa-form">
                    <q-input v-model="newMessage" outlined type="textarea"
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]">
                        <template v-slot:after>
                            <q-btn color="primary" round dense flat icon="send" @click="sendMessage" />
                        </template>
                    </q-input>
                </q-form>
            </q-card-actions>
        </q-card>
    </q-dialog>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref, watch, nextTick, toRefs } from 'vue'
import { useI18n } from 'vue-i18n'
import { DictOptions } from 'src/utils/dict'
import { FormatDateTime } from 'src/utils/date'
import GqaShowName from 'src/components/GqaShowName'
// import { GqaDefaultAvatar, GqaDefaultUsername } from 'src/settings'
import useCommon from 'src/composables/useCommon';

const { GqaDefaultUsername, GqaDefaultAvatar } = useCommon()
const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'user/get-user-list',
}
const props = defineProps({
    oldMessage: {
        type: Array,
        required: false,
        default: () => [],
    }
})
const { oldMessage } = toRefs(props)
const {
    gqaFrontend,
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    GqaAvatar,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    onRequest,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

const myAvatar = computed(() => {
    const cookieAvatar = $q.cookies.get('gqa-avatar')
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
        return GqaDefaultAvatar
    }
})
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
const chatDialogVisible = ref(false)
const messageScroll = ref(null)
watch(oldMessage.value, () => {
    // 只有在Dialog打开的时候滚动到最底部
    if (chatDialogVisible.value) {
        nextTick(() => {
            messageScroll.value.setScrollPosition('vertical', oldMessage.value.length * 75, 200)
        })
    }
})
const newMessage = ref('')
const show = () => {
    chatDialogVisible.value = true
    // this.getTableData()
    nextTick(() => {
        // userList不要滚动到最底部了，保持在最上面
        // this.$refs.userScroll.setScrollPosition('vertical', this.tableData.length * 55)
        messageScroll.value.setScrollPosition('vertical', oldMessage.value.length * 75)
    })
}
defineExpose({
    show
})
const emit = defineEmits(['onSendMessage', 'changeShow', 'changeMessage'])
const beforeShow = () => {
    emit('changeShow', true)
}
const beforeHide = () => {
    emit('changeShow', false)
}
const newMessageForm = ref(null)
const sendMessage = async () => {
    const success = await newMessageForm.value.validate()
    if (success) {
        emit('onSendMessage', {
            name: myName.value,
            avatar: myAvatar.value,
            text: newMessage.value,
            message_type: 'chat',
        })
        newMessage.value = ''
        newMessageForm.value.reset()
        newMessageForm.value.focus()
    } else {
        $q.notify({
            type: 'negative',
            message: t('FixForm'),
        })
    }
}
</script>

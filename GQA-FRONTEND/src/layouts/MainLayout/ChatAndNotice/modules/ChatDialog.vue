<template>
    <q-dialog persistent v-model="chatDialogVisible" transition-hide="scale" @before-show="beforeShow"
        @before-hide="beforeHide">
        <q-card bordered style="width: 800px; max-width: 50vw;">
            <q-toolbar>
                <q-toolbar-title>
                    {{ gqaFrontend.subTitle }}
                    {{ $t('ChatRoom') }}
                </q-toolbar-title>
                <q-btn dense flat icon="close" v-close-popup />
            </q-toolbar>
            <q-separator />
            <q-card-section horizontal>
                <q-card-section class="col-4" style="padding: 0">
                    <q-scroll-area ref="userScroll" visible style="height: 50vh; width: 100%">
                        <q-list>
                            <q-item clickable v-ripple v-for="(item, index) in tableData" :key="index">
                                <q-item-section avatar>
                                    <gqa-avatar :src="item.user.avatar" />
                                </q-item-section>
                                <q-item-section>
                                    <q-item-label>
                                        <GqaShowName :customNameObject="item.user" />
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
                    <q-toolbar>
                        <q-form ref="newMessageForm" class="gqa-form" style="width: 100%">
                            <q-input class="col-auto" dense v-model="newMessage" rounded outlined
                                @keyup.enter.stop="sendMessage" :rules="[val => val && val.length > 0 || $t('NeedInput')]"
                                :label="$t('Content')">
                                <template v-slot:append>
                                    <q-btn flat rounded icon="send" @click="sendMessage"></q-btn>
                                </template>
                            </q-input>
                        </q-form>
                    </q-toolbar>
                </q-card-section>
            </q-card-section>
        </q-card>
    </q-dialog>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { computed, onMounted, ref, watch, nextTick, toRefs } from 'vue'
import { useI18n } from 'vue-i18n'
import useConfig from 'src/composables/useConfig';

const { GqaDefaultUsername, GqaDefaultAvatar } = useConfig()
const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'user-online/get-user-online-list',
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
    GqaShowName,
    loading,
    tableData,
    getTableData,
} = useTableData(url)

onMounted(() => {
    pagination.value.sortBy = 'username'
    getTableData()
})

const myAvatar = computed(() => {
    const cookieAvatar = $q.cookies.get('gqa-avatar')
    if (cookieAvatar) {
        if (cookieAvatar.substring(0, 4) === 'http') {
            return cookieAvatar
        } else if (cookieAvatar.substring(0, 11) === 'gqa-upload:') {
            return process.env.API + cookieAvatar.substring(11)
        }
        return cookieAvatar
    } else {
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
    // Scroll to the bottom only when Dialog is open
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
        // Don't scroll to the bottom of userList, keep it at the top
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

<template>
    <div>
        <q-list bordered separator style="min-width: 300px">
            <q-item clickable v-ripple v-for="(item, index) in messageData" :key="index" @click="toNoticeDetail(item)">
                <q-item-section avatar>
                    <q-icon color="primary" name="message" />
                </q-item-section>

                <q-item-section>
                    <q-item-label>
                        {{ item.notice_title }}
                    </q-item-label>
                    <q-item-label caption>
                        {{ showDateTime(item.created_at) }}
                    </q-item-label>
                </q-item-section>
            </q-item>
        </q-list>
        <q-item clickable v-ripple class="text-center" @click="toUserProfile">
            <q-item-section>
                {{ $t('ViewAll') }}
            </q-item-section>
        </q-item>

        <UserProfile ref="userProfile" />
        <NoticeDetail ref="noticeDetail" />
    </div>
</template>

<script setup>
import UserProfile from 'src/layouts/MainLayout/UserProfile/index.vue'
import NoticeDetail from 'src/layouts/MainLayout/UserProfile/modules/NoticeDetail.vue'
import { ref, toRefs } from 'vue';
import useCommon from 'src/composables/useCommon'

const { showDateTime } = useCommon()

const props = defineProps({
    messageData: {
        type: Array,
        required: false,
        default: () => [],
    },
})
const { messageData } = toRefs(props)

const userProfile = ref(null)
const toUserProfile = () => {
    userProfile.value.show('noticeType_message')
}

const noticeDetail = ref(null)
const toNoticeDetail = (item) => {
    noticeDetail.value.show(item)
}
</script>
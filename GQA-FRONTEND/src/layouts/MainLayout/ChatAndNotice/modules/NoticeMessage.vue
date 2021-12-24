<template>
    <div>
        <q-list bordered separator style="min-width: 300px">
            <q-item clickable v-ripple v-for="(item, index) in messageData" :key="index" @click="toNoticeDetail(item)">
                <q-item-section avatar>
                    <q-icon color="primary" name="message" />
                </q-item-section>

                <q-item-section>
                    {{ item.noticeTitle }}
                </q-item-section>
            </q-item>
        </q-list>
        <q-item clickable v-ripple class="text-center" @click="toUserProfile">
            <q-item-section>
                {{ $t('CheckAll') }}
            </q-item-section>
        </q-item>

        <UserProfile ref="userProfile" />
        <NoticeDetail ref="noticeDetail" />
    </div>
</template>

<script>
import UserProfile from 'src/pages/UserProfile'
import NoticeDetail from 'src/pages/UserProfile/modules/NoticeDetail.vue'

export default {
    name: 'NoticeMessage',
    props: {
        messageData: {
            type: Array,
            required: false,
            default: () => [],
        },
    },
    components: {
        UserProfile,
        NoticeDetail,
    },
    methods: {
        toUserProfile() {
            this.$refs.userProfile.show('message')
        },
        toNoticeDetail(item) {
            this.$refs.noticeDetail.show(item)
        },
    },
}
</script>
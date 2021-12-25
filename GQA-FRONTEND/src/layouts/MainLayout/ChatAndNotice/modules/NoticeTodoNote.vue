<template>
    <div>
        <q-list bordered separator style="min-width: 300px">
            <q-item clickable v-ripple v-for="(item, index) in todoNoteData.slice(0, 5)" :key="index"
                @click="toTodoDetail(item)">
                <q-item-section avatar>
                    <q-icon color="primary" name="notifications" />
                </q-item-section>

                <q-item-section>
                    {{ item.todoDetail }}
                </q-item-section>
            </q-item>
        </q-list>
        <q-item clickable v-ripple class="text-center" @click="toUserProfile">
            <q-item-section>
                {{ $t('CheckAll') }}
            </q-item-section>
        </q-item>

        <UserProfile ref="userProfile" />
        <NoticeTodoNoteDetail ref="noticeTodoNoteDetail" />
    </div>
</template>

<script>
import UserProfile from 'src/pages/UserProfile'
import NoticeTodoNoteDetail from 'src/pages/UserProfile/modules/NoticeTodoNoteDetail.vue'

export default {
    name: 'NoticeTodoNote',
    props: {
        todoNoteData: {
            type: Array,
            required: false,
            default: () => [],
        },
    },
    components: {
        UserProfile,
        NoticeTodoNoteDetail,
    },
    methods: {
        toUserProfile() {
            this.$refs.userProfile.show('todoNote')
        },
        toTodoDetail(item) {
            this.$refs.noticeTodoNoteDetail.formType = 'edit'
            this.$refs.noticeTodoNoteDetail.show(item)
        },
    },
}
</script>
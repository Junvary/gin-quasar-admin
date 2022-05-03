<template>
    <div>
        <q-list bordered separator style="min-width: 300px">
            <q-item clickable v-for="(item, index) in noteTodoData.slice(0, 5)" :key="index"
                @click="toTodoDetail(item)">
                <q-item-section avatar>
                    <q-icon color="primary" name="notifications" />
                </q-item-section>

                <q-item-section>
                    {{ item.todo_detail }}
                </q-item-section>
            </q-item>
        </q-list>
        <q-item clickable class="text-center" @click="toUserProfile">
            <q-item-section>
                {{ $t('CheckAll') }}
            </q-item-section>
        </q-item>

        <UserProfile ref="userProfile" />
        <NoticeNoteTodoDetail ref="noticeNoteTodoDetail" />
    </div>
</template>

<script setup>
import UserProfile from 'src/layouts/MainLayout/UserProfile/index.vue'
import NoticeNoteTodoDetail from 'src/layouts/MainLayout/UserProfile/modules/NoticeNoteTodoDetail.vue'
import { ref, toRefs } from 'vue';

const props = defineProps({
    noteTodoData: {
        type: Array,
        required: false,
        default: () => [],
    },
})
const { noteTodoData } = toRefs(props)

const userProfile = ref(null)
const toUserProfile = () => {
    userProfile.value.show('noteTodo')
}
const noticeNoteTodoDetail = ref(null)
const toTodoDetail = (item) => {
    noticeNoteTodoDetail.value.formType = 'edit'
    noticeNoteTodoDetail.value.show(item)
}
</script>

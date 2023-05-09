<template>
    <div>
        <q-list bordered separator style="min-width: 300px">
            <q-item clickable v-ripple v-for="(item, index) in todoData" :key="index" @click="toTodoDetail(item)">
                <q-item-section avatar>
                    <q-icon color="primary" name="notifications" />
                </q-item-section>

                <q-item-section>
                    <q-item-label>
                        {{ item.todo_detail }}
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
        <NoticeTodoDetail ref="noticeTodoDetail" />
    </div>
</template>

<script setup>
import UserProfile from 'src/layouts/MainLayout/UserProfile/index.vue'
import NoticeTodoDetail from 'src/layouts/MainLayout/UserProfile/modules/NoticeTodoDetail.vue'
import { ref, toRefs } from 'vue';
import useCommon from 'src/composables/useCommon'

const { showDateTime } = useCommon()
const props = defineProps({
    todoData: {
        type: Array,
        required: false,
        default: () => [],
    },
})
const { todoData } = toRefs(props)

const userProfile = ref(null)
const toUserProfile = () => {
    userProfile.value.show('todo')
}
const noticeTodoDetail = ref(null)
const toTodoDetail = (item) => {
    noticeTodoDetail.value.formType = 'edit'
    noticeTodoDetail.value.show(item)
}
</script>
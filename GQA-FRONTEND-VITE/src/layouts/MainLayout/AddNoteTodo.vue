<template>
    <q-btn dense round flat icon="edit_note" @click="addNoteTodo">
        <q-tooltip>
            {{ $t('Add') + ' ' + $t('NoteTodo') }}
        </q-tooltip>
    </q-btn>
    <NoticeNoteTodoDetail ref="noticeNoteTodoDetail" />
</template>

<script setup>
import { ref } from 'vue';
import NoticeNoteTodoDetail from './UserProfile/modules/NoticeNoteTodoDetail.vue'

const noticeNoteTodoDetail = ref(null);
const addNoteTodo = () => {
    noticeNoteTodoDetail.value.formType = 'add'
    noticeNoteTodoDetail.value.show({})
}

</script>
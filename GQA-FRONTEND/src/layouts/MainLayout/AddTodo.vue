<template>
    <q-btn dense round flat icon="las la-user-edit" @click="addTodo">
        <q-tooltip>
            {{ $t('Add') + ' ' + $t('Todo') }}
        </q-tooltip>
    </q-btn>
    <NoticeTodoDetail ref="noticeTodoDetail" />
</template>

<script setup>
import { ref } from 'vue';
import NoticeTodoDetail from './UserProfile/modules/NoticeTodoDetail.vue'

const noticeTodoDetail = ref(null);
const addTodo = () => {
    noticeTodoDetail.value.formType = 'add'
    noticeTodoDetail.value.show({})
}

</script>
<template>
    <q-dialog v-model="loginVisible" style="width: 100%;">
        <q-card>
            <LoginForm />
        </q-card>
    </q-dialog>
</template>

<script setup>
import LoginForm from './LoginForm.vue';
import { ref } from 'vue';

const loginVisible = ref(false);
const show = () => {
    loginVisible.value = true
}

defineExpose({
    show
})
</script>

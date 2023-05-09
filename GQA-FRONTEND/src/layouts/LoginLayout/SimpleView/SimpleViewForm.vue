<template>
    <transition appear enter-active-class="animated slideInLeft" leave-active-class="animated slideOutLeft">
        <div class="col row column justify-center items-center">
            <q-card flat style="background: rgba(0, 0, 0, 0); padding: 10px;">
                <LoginForm />
            </q-card>
        </div>
    </transition>
</template>

<script setup>
import LoginForm from '../LoginForm.vue'
</script>
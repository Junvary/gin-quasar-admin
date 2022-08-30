<template>
    <div class="gqa-login-layout-img-container">
        <LoginForm v-if="!dbNeedInit" />
        <InitDbForm v-if="dbNeedInit" @initDbSuccess="checkDb" />
    </div>
</template>

<script setup>
import { defineProps, toRefs, defineEmits } from 'vue'
import LoginForm from './LoginForm.vue'
import InitDbForm from './InitDbForm.vue'

const props = defineProps({
    dbNeedInit: {
        type: Boolean,
        default: false,
    },
})
const { dbNeedInit } = toRefs(props)

const emit = defineEmits(['initDbSuccess'])
const checkDb = () => {
    emit('initDbSuccess')
}
</script>

<template>
    <q-page padding>
        <GqaEditor v-model="text" />
    </q-page>
</template>

<script setup>
import { ref } from 'vue';
import GqaEditor from 'src/components/GqaEditor/index.vue'

const text = ref('<p>Hello Gin-Quasar-Admin</p>')
</script>
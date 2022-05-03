<template>
    <div class="row justify-center items-center ">
        {{ gqaFrontend.subTitle }}
    </div>
</template>

<script setup>
import { computed } from 'vue';
import { useStorageStore } from 'src/stores/storage'
const storageStore = useStorageStore()
const gqaFrontend = computed(() => storageStore.GetGqaFrontend())
</script>

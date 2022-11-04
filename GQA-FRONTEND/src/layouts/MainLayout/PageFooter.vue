<template>
    <div class="row justify-center items-center" :class="darkTheme">
        {{ gqaFrontend.subTitle }}
    </div>
</template>

<script setup>
import { computed } from 'vue';
import { useStorageStore } from 'src/stores/storage'
import useTheme from 'src/composables/useTheme';

const { darkTheme } = useTheme()
const storageStore = useStorageStore()
const gqaFrontend = computed(() => storageStore.GetGqaFrontend())
</script>

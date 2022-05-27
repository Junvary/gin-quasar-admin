<template>
    <div class="row justify-center items-center" :class="darkTheme">
        {{ gqaFrontend.subTitle }}
    </div>
</template>

<script setup>
import { computed } from 'vue';
import { useStorageStore } from 'src/stores/storage'
import useDarkTheme from 'src/composables/useDarkTheme';

const { darkTheme } = useDarkTheme()
const storageStore = useStorageStore()
const gqaFrontend = computed(() => storageStore.GetGqaFrontend())
</script>

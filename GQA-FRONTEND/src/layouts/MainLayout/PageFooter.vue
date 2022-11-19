<template>
    <div class="row justify-center items-end gqa-footer" :style="{ color: $q.dark.isActive ? '#fff' : '#000' }">
        {{ gqaFrontend.subTitle }}
        is powered by&nbsp;
        <a href="https://github.com/Junvary/gin-quasar-admin" target="_blank" style="text-decoration: none"
            :style="{ color: $q.dark.isActive ? '#fff' : '#000' }">
            Gin-Quasar-Admin
            {{ gqaVersion }}
        </a>
    </div>
</template>

<script setup>
import { computed } from 'vue';
import { useStorageStore } from 'src/stores/storage'
import useTheme from 'src/composables/useTheme';
import config from '../../../package.json'

const { darkTheme } = useTheme()
const storageStore = useStorageStore()
const gqaFrontend = computed(() => storageStore.GetGqaFrontend())
const gqaVersion = config.version
</script>

<style lang="scss" scoped>
.gqa-footer {
    height: 20px;
    font-size: 10px;
    background: rgba($color: #000000, $alpha: 0) !important;
}
</style>
<template>
    <q-toggle v-model="sidebarDense" />
</template>

<script setup>
import { onMounted, ref, watch } from 'vue';
import { useSettingStore } from 'src/stores/setting';

const sidebarDense = ref(false)
const settingStore = useSettingStore()

onMounted(() => {
    sidebarDense.value = settingStore.GetSidebarDense()
})
watch(sidebarDense, (val) => {
    settingStore.SetSidebarDense(val)
})
</script>
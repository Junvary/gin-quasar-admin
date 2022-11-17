<template>
    <q-toggle v-model="darkTheme" icon="eva-moon-outline" unchecked-icon="eva-sun-outline"
        :color="darkTheme ? 'black' : ''" />
</template>

<script setup>
import { useQuasar } from 'quasar'
import { onMounted, ref, watch } from 'vue';
import { useSettingStore } from 'src/stores/setting';

const darkTheme = ref(false)
const $q = useQuasar()
const settingStore = useSettingStore()

onMounted(() => {
    darkTheme.value = settingStore.GetDarkTheme()
})
watch(darkTheme, (val) => {
    settingStore.SetDarkTheme(val)
    $q.dark.set(darkTheme.value)
})
</script>
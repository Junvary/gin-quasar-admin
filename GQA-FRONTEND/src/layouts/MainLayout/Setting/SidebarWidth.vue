<template>
    <q-field :label="$t('Sidebar') + $t('Width') + ' ' + drawerWidth" stack-label>
        <template v-slot:control>
            <q-slider v-model="drawerWidth" :min="200" :max="300" @change="setDrawerWidth" />
        </template>
    </q-field>
</template>

<script setup>
import { useSettingStore } from 'src/stores/setting';
import { ref, onMounted } from 'vue';

const settingStore = useSettingStore()

const drawerWidth = ref(220)
const setDrawerWidth = () => {
    settingStore.SetSidebarWidth(drawerWidth.value)
}

onMounted(() => {
    drawerWidth.value = settingStore.GetSidebarWidth()
})
</script>
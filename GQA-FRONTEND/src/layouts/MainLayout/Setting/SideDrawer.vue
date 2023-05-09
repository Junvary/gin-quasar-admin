<template>
    <q-field :label="$t('SideDrawer') + $t('Width') + ' ' + drawerWidth" stack-label>
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
    settingStore.SetSideDrawerWidth(drawerWidth.value)
}

onMounted(() => {
    drawerWidth.value = settingStore.GetSideDrawerWidth()
})
</script>
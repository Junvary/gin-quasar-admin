<template>
    <q-field :label="$t('SideDrawer') + $t('Width') + ' ' + drawerWidth" stack-label>
        <template v-slot:control>
            <q-slider v-model="drawerWidth" :min="200" :max="300" @change="setDrawerWidth" />
        </template>
    </q-field>
    <q-field :label="$t('SideDrawer') + $t('Top')" stack-label>
        <template v-slot:control>
            <q-btn-toggle v-model="drawerTop" toggle-color="primary" :options="topOption"
                @update:model-value="setDrawerTop" />
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

const drawerTop = ref('h')
const topOption = [
    { label: 'h', value: 'h' },
    { label: 'l', value: 'l' },
]
const setDrawerTop = () => {
    settingStore.SetSideDrawerTop(drawerTop.value)
}

const drawerBottom = ref('l')
const bottomOption = [
    { label: 'l', value: 'l' },
    { label: 'f', value: 'f' },
]
const setDrawerBottom = () => {
    settingStore.SetSideDrawerBottom(drawerBottom.value)
}

onMounted(() => {
    drawerWidth.value = settingStore.GetSideDrawerWidth()
    drawerTop.value = settingStore.GetSideDrawerTop()
    drawerBottom.value = settingStore.GetSideDrawerBottom()
})
</script>
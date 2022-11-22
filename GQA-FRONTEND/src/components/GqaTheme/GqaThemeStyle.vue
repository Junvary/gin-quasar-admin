<template>
    <div class="row q-gutter-md">
        <q-btn v-for="(value, key, index) in themeStyleList" :key="index" size="90px" @click="setThemeStyle(key)">
            <q-icon name="check_box" size="40px" style="top: 2px; right: 2px" color="negative"
                class="absolute all-pointer-events" v-if="themeStyle === key" />
            <q-avatar size="160px">
                <img :src="value">
            </q-avatar>
            <div class="absolute-bottom text-subtitle1 text-center bg-grey">
                {{ key }}{{ $t('Style') }}
            </div>
        </q-btn>
    </div>
</template>

<script setup>
import { computed } from 'vue';
import { useSettingStore } from 'src/stores/setting';
import { setCssVar, Cookies } from 'quasar';
import useConfig from 'src/composables/useConfig';

const { ThemeStyleQuasar, ThemeStyleElement, ThemeStyleAnt } = useConfig()
const themeStyleList = {
    'Gin-Quasar-Admin': 'icons/gqa128.png',
    'Quasar': 'icon/quasar.svg',
    'Element Plus': 'icon/element-plus.svg',
    'Ant Design Vue': 'icon/ant-design-vue.svg',
}

const settingStore = useSettingStore()
const themeStyle = computed(() => settingStore.GetThemeStyle())
const setThemeStyle = (style) => {
    switch (style) {
        case 'Gin-Quasar-Admin':
            for (let q in ThemeStyleQuasar) {
                setBrand(q, ThemeStyleQuasar[q])
            }
            break
        case 'Quasar':
            for (let q in ThemeStyleQuasar) {
                setBrand(q, ThemeStyleQuasar[q])
            }
            break
        case 'Element Plus':
            for (let e in ThemeStyleElement) {
                setBrand(e, ThemeStyleElement[e])
            }
            break
        case 'Ant Design Vue':
            for (let a in ThemeStyleAnt) {
                setBrand(a, ThemeStyleAnt[a])
            }
            break
        default:
            for (let q in ThemeStyleQuasar) {
                setBrand(q, ThemeStyleQuasar[q])
            }
    }
    settingStore.SetThemeStyle(style)
}

const setBrand = (type, color) => {
    Cookies.set('gqa-theme-' + type, color)
    setCssVar(type, color)
}

</script>
<template>
    <!-- 本页面根据需求维护 -->
    <q-layout style="overflow-x: hidden">
        <q-page-container>
            <page-banner />
            <component v-if="pluginCurrent && pluginComponent" :key="pluginCurrent" :is="pluginComponent" />
            <q-card v-else class="row items-center justify-center" style="padding: 20px 0;">
                <q-btn color="primary" :label="$t('LoginLayoutWithoutPlugin')"></q-btn>
            </q-card>
            <page-footer />
            <q-page-scroller position="bottom-right" :scroll-offset="150" :offset="[3, 18]">
                <q-btn dense fab push icon="keyboard_arrow_up" color="primary" />
            </q-page-scroller>
        </q-page-container>
    </q-layout>
</template>

<script setup>
import { ref, onUnmounted, onBeforeMount, toRefs } from 'vue';
import PageBanner from './PageBanner.vue'
import PageFooter from './PageFooter.vue'

const props = defineProps({
    pluginComponent: {
        type: Object,
        default: () => { },
    },
    pluginCurrent: {
        type: String,
        default: '',
    },
})
const { pluginComponent } = toRefs(props)

const showScrollDown = ref(true)

onBeforeMount(() => {
    window.addEventListener('scroll', documentTop())
})
onUnmounted(() => {
    window.removeEventListener('scroll', documentTop())
})

const documentTop = () => {
    let top = document.documentElement.scrollTop
    if (top > 200) {
        showScrollDown.value = false
    } else {
        showScrollDown.value = true
    }
}
</script>

<style lang="scss" scoped>
.login-layout-scroll-down {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: -60px;
    z-index: 100;
}
</style>
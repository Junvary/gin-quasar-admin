<template>
    <q-layout style="overflow-x: hidden">
        <q-page-container>
            <page-banner />
            <!-- <GqaScrollDown id="login-layout-scroll-down" class="login-layout-scroll-down"
                v-if="pluginCurrent && pluginComponent && showScrollDown" @click="scrollDown" /> -->
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
import GqaScrollDown from 'src/components/GqaScrollDown/index.vue'

const props = defineProps({
    pluginComponent: {
        type: String,
        default: '',
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

const scrollDown = () => {
    document.getElementById('login-layout-scroll-down').nextElementSibling.scrollIntoView({
        behavior: 'smooth',
    })
}

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
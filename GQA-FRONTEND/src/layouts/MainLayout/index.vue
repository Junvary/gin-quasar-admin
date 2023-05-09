<template>
    <q-layout view="hHh LpR lFf" :style="{ backgroundColor: $q.dark.isActive ? '#1d1d1d' : '#f0f2f5' }"
        style="overflow-x: hidden;">
        <q-header reveal bordered :class="darkTheme" style="height: 40px;">
            <q-toolbar class="min-height-40">
                <q-btn dense round flat :icon="toggleLeftDrawer ? 'eva-arrowhead-left' : 'eva-arrowhead-right'"
                    @click="toggleLeftDrawer = !toggleLeftDrawer" />

                <gqa-avatar class="gin-quasar-admin-logo" :src="gqaFrontend.logo" @mouseenter="startCheck"
                    @mouseleave="stopCheck" />

                <q-toolbar-title shrink class="text-bold text-italic cursor-pointer" style="padding: 0 5px;">
                    {{ gqaFrontend.subTitle }}
                </q-toolbar-title>

                <q-separator vertical inset spaced />

                <ChoosePlugin @changePlugin="changeTop" />

                <q-breadcrumbs v-if="findCurrentTopMenu" style="margin-left: 20px;">
                    <q-breadcrumbs-el :label="selectOptionLabel(findCurrentTopMenu)" :icon="findCurrentTopMenu?.icon"
                        :class="darkTheme" />
                    <q-breadcrumbs-el :label="selectRouteLabel(route)" :icon="route.meta.icon" />
                </q-breadcrumbs>

                <q-space />

                <div class="q-gutter-sm row items-center no-wrap">
                    <SearchMenu />
                    <Fullscreen />
                    <!-- <ChoosePlugin @changePlugin="changeTop" /> -->
                    <ChatAndNotice />
                    <AddTodo />
                    <GitLink v-if="gqaFrontend.showGit === 'yesNo_yes'" />
                    <UserMenu @showProfile="$refs.userProfile.show()" />
                    <Setting />
                </div>
            </q-toolbar>
        </q-header>

        <q-drawer bordered v-model="toggleLeftDrawer" show-if-above :width="drawerWidth" :mini="miniState"
            :mini-to-overlay="miniStateOut ? true : false" @mouseover="miniStateMouseover" @mouseout="miniStateMouseout">
            <SideBarLeft :topMenuChildren="topMenuChildren" />
            <div class="q-mini-drawer-hide absolute" style="bottom: 15px; right: -17px">
                <q-btn dense round :icon="miniStateOut ? 'chevron_right' : 'chevron_left'"
                    @click="miniStateOut = !miniStateOut" />
            </div>
        </q-drawer>

        <q-page-container style="overflow: hidden;">
            <router-view v-slot="{ Component }">
                <component :is="Component" />
            </router-view>

            <q-page-sticky expand position="top">
                <TabMenu />
            </q-page-sticky>

            <q-page-sticky position="bottom-right" :offset="fabPos">
                <q-page-scroller position="bottom-right" :scroll-offset="150" :offset="[0, -70]">
                    <q-btn dense glossy rounded icon="keyboard_arrow_up" v-touch-pan.prevent.mouse="moveFab" />
                </q-page-scroller>
            </q-page-sticky>
            <PageFooter />
        </q-page-container>

        <UserProfile ref="userProfile" />
        <AchievementDialog ref="achievementDialog" />
    </q-layout>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue';
import { useUserStore } from 'src/stores/user';
import { usePermissionStore } from 'src/stores/permission';
import { useStorageStore } from 'src/stores/storage';
import { useSettingStore } from 'src/stores/setting'
import useTheme from 'src/composables/useTheme';
import SideBarLeft from './SideBarLeft/index.vue'
import TabMenu from './TabMenu.vue'
import SearchMenu from './SearchMenu.vue'
import Fullscreen from './Fullscreen.vue'
import ChoosePlugin from './ChoosePlugin.vue'
import ChatAndNotice from './ChatAndNotice/index.vue'
import GitLink from './GitLink.vue'
import UserMenu from './UserMenu.vue'
import Setting from './Setting/index.vue'
import PageFooter from './PageFooter.vue'
import UserProfile from './UserProfile/index.vue'
import AddTodo from './AddTodo.vue';
import { useRoute } from 'vue-router';
import { useQuasar } from 'quasar';
import useCommon from 'src/composables/useCommon'
import XEUtils from 'xe-utils'
import { postAction } from 'src/api/manage';
import AchievementDialog from 'src/plugins/Achievement/AchievementDialog.vue';

const $q = useQuasar();
const { darkTheme } = useTheme()
const route = useRoute();
const userStore = useUserStore();
const storageStore = useStorageStore();
const permissionStore = usePermissionStore();
const settingStore = useSettingStore();
const toggleLeftDrawer = ref(false);
const topMenuChildren = ref({});
const currentTopMenu = ref('');
const fabPos = ref([3, 80]);
const userProfile = ref(null);
// The top navigation bar generally controls whether it is mini mode
const miniStateOut = ref(false)
// Control mini mode
const miniState = ref(false)
// Monitor the overall transformation of mini mode
watch(miniStateOut, (newVlaue) => {
    miniState.value = newVlaue
})
// If the overall mode is mini, perform mouse entry conversion
const miniStateMouseover = () => {
    if (miniStateOut.value === true) {
        miniState.value = false
    }
}
// If the overall mode is mini, perform mouse out conversion
const miniStateMouseout = () => {
    if (miniStateOut.value === true) {
        miniState.value = true
    }
}

const { selectOptionLabel, selectRouteLabel } = useCommon()
const gqaFrontend = computed(() => storageStore.GetGqaFrontend())
const drawerWidth = computed(() => settingStore.GetSideDrawerWidth())

const changeTop = (childrenMenu) => {
    topMenuChildren.value = childrenMenu.children
}

onMounted(() => {
    $q.dark.set(settingStore.GetDarkTheme())
    currentTopMenu.value = findCurrentTopMenu.value?.name
    topMenuChildren.value = topMenu.value.filter(item => item.name === currentTopMenu.value)[0]?.children
})
watch(route, () => {
    currentTopMenu.value = findCurrentTopMenu.value?.name
    topMenuChildren.value = topMenu.value.filter(item => item.name === currentTopMenu.value)[0]?.children
})

const topMenu = computed(() => {
    return permissionStore.topMenu
})

const findCurrentTopMenu = computed(() => {
    const name = route.name
    return XEUtils.searchTree(topMenu.value, item => item.name === name)[0]
})

const moveFab = (ev) => {
    fabPos.value = [fabPos.value[0] - ev.delta.x, fabPos.value[1] - ev.delta.y]
}

const timer = ref(null)

const achievementDialog = ref(null)
const startCheck = () => {
    timer.value = setTimeout(() => {
        postAction('/plugin-achievement/obtain-find', {
            category_code: 'QiYu-Roll-Code',
            username: userStore.GetUsername()
        }).then(res => {
            if (res.code === 1 /* 为了演示效果注释掉了后面 && res.data?.get */) {
                achievementDialog.value.show({
                    category: "奇遇",
                    name: "摇滚吧，少年！"
                })
            }
        })
    }, 1000)
}
const stopCheck = () => {
    clearTimeout(timer.value)
}
</script>

<style lang="scss" scoped>
.min-height-40 {
    min-height: 40px;
}
</style>
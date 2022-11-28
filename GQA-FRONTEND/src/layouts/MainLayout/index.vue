<template>
    <q-layout :view="layoutView" :style="{ backgroundColor: $q.dark.isActive ? '#1d1d1d' : '#f0f2f5' }"
        style="overflow-x: hidden;">
        <q-header reveal bordered :class="darkTheme" style="height: 50px; border-bottom: 1px solid white;">
            <q-toolbar>
                <q-btn dense round flat icon="eva-menu" @click="toggleLeftDrawer = !toggleLeftDrawer" />

                <q-btn dense round flat :icon="miniStateOut ? 'eva-arrowhead-right' : 'eva-arrowhead-left'"
                    @click="miniStateOut = !miniStateOut" />

                <GqaAvatar class="gin-quasar-admin-logo" :src="gqaFrontend.logo" @mouseenter="startCheck"
                    @mouseleave="stopCheck" />

                <q-toolbar-title shrink class="text-bold text-italic cursor-pointer" style="padding: 0 5px;">
                    {{ gqaFrontend.subTitle }}
                </q-toolbar-title>

                <q-select v-if="layoutView.split(' ')[0] === 'hHh'" id="menuSelect" dense borderless
                    v-model="currentTopMenu" :options="topMenu" map-options option-value="name"
                    @update:model-value="changeTop" style="margin-left: 20px;" dark
                    :option-label="opt => Object(opt) === opt && 'title' in opt ? $t(opt.title) : opt.title">
                    <template v-slot:prepend>
                        <q-icon name="ion-md-apps" :class="darkTheme" />
                    </template>
                </q-select>

                <q-breadcrumbs v-if="findCurrentTopMenu" style="margin-left: 20px;">
                    <q-breadcrumbs-el :label="$t(findCurrentTopMenu?.title)" :icon="findCurrentTopMenu?.icon"
                        :class="darkTheme" />
                    <q-breadcrumbs-el :label="$t(route.meta.title)" :icon="route.meta.icon" />
                </q-breadcrumbs>

                <q-space />

                <div class="q-gutter-sm row items-center no-wrap">
                    <Fullscreen />
                    <ChatAndNotice />
                    <AddNoteTodo />
                    <GitLink v-if="gqaFrontend.showGit === 'yesNo_yes'" />
                    <UserMenu @showProfile="$refs.userProfile.show()" />
                    <Setting />
                </div>
            </q-toolbar>
        </q-header>

        <q-drawer bordered v-model="toggleLeftDrawer" show-if-above :width="drawerWidth" :mini="miniState"
            :mini-to-overlay="miniStateOut ? true : false" @mouseover="miniStateMouseover"
            @mouseout="miniStateMouseout">
            <SideBarLeft :topMenuChildren="topMenuChildren">
                <q-select v-if="layoutView.split(' ')[0] === 'lHh'" id="menuSelect" v-model="currentTopMenu"
                    :options="topMenu" map-options option-value="name" @update:model-value="changeTop" filled
                    :dark="(themeStyle === 'Gin-Quasar-Admin' || themeStyle === 'Quasar') ? false : true" borderless
                    :option-label="opt => Object(opt) === opt && 'title' in opt ? $t(opt.title) : opt.title">
                    <template v-slot:prepend>
                        <q-icon name="ion-md-apps" style="margin-right: 17px;" size="30px" />
                    </template>
                </q-select>
            </SideBarLeft>
        </q-drawer>

        <q-page-container style="overflow: hidden;">
            <router-view v-slot="{ Component }">
                <component :is="Component" />
            </router-view>

            <q-page-sticky expand position="top">
                <TabMenu style="border-bottom: 1px solid #e0e0e0;" />
            </q-page-sticky>

            <q-page-sticky position="bottom-right" :offset="fabPos">
                <q-page-scroller position="bottom-right" :scroll-offset="150" :offset="[0, -70]">
                    <q-btn dense glossy rounded icon="keyboard_arrow_up" v-touch-pan.prevent.mouse="moveFab" />
                </q-page-scroller>
            </q-page-sticky>
            <PageFooter />
        </q-page-container>

        <!-- <q-footer reveal>
            <PageFooter />
        </q-footer> -->

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
import Fullscreen from './Fullscreen.vue'
import ChatAndNotice from './ChatAndNotice/index.vue'
import GitLink from './GitLink.vue'
import UserMenu from './UserMenu.vue'
import Setting from './Setting/index.vue'
import PageFooter from './PageFooter.vue'
import GqaAvatar from 'src/components/GqaAvatar/index.vue'
import UserProfile from './UserProfile/index.vue'
import AddNoteTodo from './AddNoteTodo.vue';
import { useRoute } from 'vue-router';
import useDocument from 'src/composables/useDocument'
import { useQuasar } from 'quasar';
import XEUtils from 'xe-utils'
import { postAction } from 'src/api/manage';
import AchievementDialog from 'src/plugins/Achievement/AchievementDialog.vue';
// 动态更改网站标题和favicon
useDocument()

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
// 顶部导航栏总体控制是否mini模式
const miniStateOut = ref(false)
// 控制mini模式
const miniState = ref(false)
// 监听总体变换mini模式
watch(miniStateOut, (newVlaue) => {
    miniState.value = newVlaue
})
// 如果总体是mini模式，进行鼠标进入转换
const miniStateMouseover = () => {
    if (miniStateOut.value === true) {
        miniState.value = false
    }
}
// 如果总体是mini模式，进行鼠标移出转换
const miniStateMouseout = () => {
    if (miniStateOut.value === true) {
        miniState.value = true
    }
}

const gqaFrontend = computed(() => storageStore.GetGqaFrontend())
const drawerWidth = computed(() => settingStore.GetSideDrawerWidth())
const layoutView = computed(() => settingStore.GetLayoutView())
const themeStyle = computed(() => settingStore.GetThemeStyle())

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

watch(currentTopMenu, () => {
    if (currentTopMenu !== "") {
        let menuSelect = document.getElementById("menuSelect")
        let menuSpan = menuSelect.querySelector("span")
        // menuSpan.style.fontWeight = "bold"
        menuSpan.style.animation = 'headShake'
        menuSpan.style.animationDuration = "2s"
        setTimeout(() => {
            menuSpan.style.fontWeight = ""
            menuSpan.style.animation = ""
            menuSpan.style.animationDuration = ""
        }, 2000)
    }
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
.hhhh {
    color: red !important;
}
</style>
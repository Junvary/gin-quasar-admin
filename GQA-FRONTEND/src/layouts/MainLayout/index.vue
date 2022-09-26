<template>
    <q-layout view="hHh LpR lFr">
        <q-header reveal elevated :class="darkTheme">
            <q-toolbar>
                <q-btn dense round flat icon="menu" aria-label="Menu" @click="toggleLeftDrawer = !toggleLeftDrawer" />

                <GqaAvatar class="gin-quasar-admin-logo" :src="gqaFrontend.logo" style="margin-left: 5px;"
                    @mouseenter="startCheck" @mouseleave="stopCheck" />

                <q-toolbar-title shrink class="text-bold text-italic cursor-pointer" style="padding: 0 5px;">
                    {{ gqaFrontend.subTitle }}
                </q-toolbar-title>

                <q-select id="menuSelect" dense borderless v-model="currentTopMenu" :options="topMenu" map-options
                    option-value="name" @update:model-value="changeTop" style="margin-left: 20px;"
                    :option-label="opt => Object(opt) === opt && 'title' in opt ? t(opt.title) : opt.title" />

                <q-space />

                <div class="q-gutter-sm row items-center no-wrap">
                    <Fullscreen />
                    <ChatAndNotice />
                    <AddNoteTodo />
                    <Setting />
                    <GitLink v-if="gqaFrontend.showGit === 'yes'" />
                    <UserMenu @showProfile="$refs.userProfile.show()" />
                </div>
            </q-toolbar>

            <!-- header下面的标签页 -->
            <div class="row bg-white">
                <TabMenu />
            </div>
        </q-header>

        <q-drawer elevated v-model="toggleLeftDrawer" show-if-above bordered content-class="bg-grey-1">
            <SideBarLeft :topMenuChildren="topMenuChildren" />
        </q-drawer>

        <q-page-container style="overflow-x: hidden;">
            <router-view v-slot="{ Component }">
                <transition appear enter-active-class="animated fadeIn" leave-active-class="animated fadeOut">
                    <component :is="Component" />
                </transition>
            </router-view>

            <q-page-sticky position="bottom-right" :offset="fabPos" class="column">
                <q-page-scroller position="bottom-right" :scroll-offset="150" :offset="[0, -80]">
                    <q-btn push fab glossy rounded icon="keyboard_arrow_up" color="primary"
                        v-touch-pan.prevent.mouse="moveFab" />
                </q-page-scroller>
            </q-page-sticky>
        </q-page-container>

        <q-footer reveal elevated>
            <PageFooter />
        </q-footer>

        <UserProfile ref="userProfile" />
        <AchievementDialog ref="achievementDialog" />
    </q-layout>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue';
import { useUserStore } from 'src/stores/user';
import { usePermissionStore } from 'src/stores/permission';
import { useStorageStore } from 'src/stores/storage';
import useDarkTheme from 'src/composables/useDarkTheme';
import SideBarLeft from './SideBarLeft/index.vue'
import TabMenu from './TabMenu.vue'
import Fullscreen from './Fullscreen.vue'
import ChatAndNotice from './ChatAndNotice/index.vue'
import GitLink from './GitLink.vue'
import UserMenu from './UserMenu.vue'
import Setting from './Setting.vue'
import PageFooter from './PageFooter.vue'
import GqaAvatar from 'src/components/GqaAvatar/index.vue'
import UserProfile from './UserProfile/index.vue'
import AddNoteTodo from './AddNoteTodo.vue';
import { useRoute } from 'vue-router';
import useDocument from 'src/composables/useDocument'
import { useQuasar, LocalStorage } from 'quasar';
import XEUtils from 'xe-utils'
import { useI18n } from 'vue-i18n';
import { postAction } from 'src/api/manage';
import AchievementDialog from 'src/plugins/Achievement/AchievementDialog.vue';
// 动态更改网站标题和favicon
useDocument()

const $q = useQuasar();
const { t } = useI18n();
const { darkTheme } = useDarkTheme()
const route = useRoute();
const userStore = useUserStore();
const storageStore = useStorageStore();
const permissionStore = usePermissionStore();
const toggleLeftDrawer = ref(false);
const topMenuChildren = ref({});
const currentTopMenu = ref('');
const fabPos = ref([3, 80]);
const userProfile = ref(null);

const gqaFrontend = computed(() => {
    return storageStore.GetGqaFrontend()
})

const changeTop = (childrenMenu) => {
    topMenuChildren.value = childrenMenu.children
}

onMounted(() => {
    const localDark = LocalStorage.getItem('gqa-dark-theme') || false
    $q.dark.set(localDark)
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
        menuSpan.style.fontWeight = "bold"
        menuSpan.style.animation = 'heartBeat'
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
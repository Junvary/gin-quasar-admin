<template>
    <q-layout view="hHh LpR lFr">
        <q-header reveal elevated :class="darkTheme">
            <q-toolbar>
                <q-btn dense round flat icon="menu" aria-label="Menu" @click="toggleLeftDrawer = !toggleLeftDrawer" />

                <GqaAvatar class="gin-quasar-admin-logo" :src="gqaFrontend.logo" style="margin-left: 5px;" />

                <q-toolbar-title shrink class="text-bold text-italic cursor-pointer" style="padding: 0 5px;">
                    {{ gqaFrontend.subTitle }}
                </q-toolbar-title>

                <q-select v-model="currentItemMenu" :options="topMenu" map-options option-value="name"
                    option-label="title" @update:model-value="changeTop" />

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

            <!-- <q-separator /> -->
            <!-- header下面的标签页 -->
            <div class="row bg-white">
                <TabMenu />
            </div>
        </q-header>

        <q-drawer elevated v-model="toggleLeftDrawer" show-if-above bordered content-class="bg-grey-1">
            <SideBarLeft :topMenuChildren="topMenuChildren" />
        </q-drawer>

        <q-page-container>
            <router-view />
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

    </q-layout>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue';
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
import { useRoute, useRouter } from 'vue-router';
import useDocument from 'src/composables/useDocument'
import { useQuasar, LocalStorage } from 'quasar';
// 动态更改网站标题和favicon
useDocument()

const $q = useQuasar();
const { darkTheme } = useDarkTheme()
const route = useRoute();
const router = useRouter();
const storageStore = useStorageStore();
const permissionStore = usePermissionStore();
const toggleLeftDrawer = ref(false);
const topMenuChildren = ref({});
const currentItemMenu = ref('');
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

    if (findTopItemMenu.value.is_plugin === 'yes') {
        topMenuChildrenPlugin.value = findTopItemMenu.value
    }
    topMenuChildren.value = findTopItemMenu.value
})
watch(route, () => {
    if (findTopItemMenu.value.is_plugin === 'yes') {
        topMenuChildrenPlugin.value = findTopItemMenu.value
    }
    topMenuChildren.value = findTopItemMenu.value
})

const topMenu = computed(() => {
    return permissionStore.topMenu
})

const findTopItemMenu = computed(() => {
    const name = route.name
    let item = {}
    for (let m of topMenu.value) {
        if (m.name === name || (m.children && m.children.find((item) => item.name === name))) {
            item = m
            break
        }
    }
    currentItemMenu.value = item.name
    console.log(item)
    return item
})

const changeTopMenu = (item) => {
    if (item && item.top) {
        currentItemMenu.value = item.top ? item.top.name : ''
        if (item.top.name === 'dashboard') {
            router.push('/dashboard')
        }
        topMenuChildren.value = item
    }
}

const topMenuChildrenPlugin = ref({})
const changeTopMenuPlugin = (item) => {
    currentItemMenu.value = item.top ? item.top.name : ''
    topMenuChildrenPlugin.value = item
    changeTopMenu(item)
}

const moveFab = (ev) => {
    fabPos.value = [fabPos.value[0] - ev.delta.x, fabPos.value[1] - ev.delta.y]
}
</script>

<style lang="scss" scoped>
</style>
<template>
    <q-layout view="hHh LpR lFr">
        <q-header reveal elevated>
            <q-toolbar class="bg-primary glossy ">
                <q-btn dense round glossy push icon="sync_alt" aria-label="Menu"
                    @click="toggleLeftDrawer = !toggleLeftDrawer" />

                <GqaAvatar class="gin-quasar-admin-logo" :src="gqaFrontend.logo" />

                <q-toolbar-title shrink class="text-bold text-italic">
                    {{ gqaFrontend.subTitle }}
                </q-toolbar-title>

                <q-tabs dense inline-label outside-arrows mobile-arrows shrink stretch v-model="currentItemMenu"
                    style="max-width: 100%;" class="text-white">
                    <q-tab v-for="item in topMenu" :key="item.top.name" :name="item.top.name"
                        :label="$t(item.top.title)" @click="changeTopMenu(item)" />
                </q-tabs>

                <q-space />

                <Fullscreen style="margin: 0 5px" />

                <ChatAndNotice />

                <AddNoteTodo />

                <UserMenu style="margin: 0 5px" @showProfile="$refs.userProfile.show()" />
                <!-- <q-language-switcher/> -->
                <Setting style="margin: 0 5px" />
                <GitLink style="margin: 0 5px" v-if="gqaFrontend.showGit === 'yes'" />

            </q-toolbar>

            <!-- <q-separator /> -->
            <!-- header下面的标签页 -->
            <div class="row bg-white">
                <TabMenu v-show="!pageDashboard" />
            </div>

        </q-header>

        <q-drawer elevated v-if="!pageDashboard" v-model="toggleLeftDrawer" show-if-above bordered
            content-class="bg-grey-1">
            <SideBarLeft :topMenuItem="topMenuItem" />
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
// 动态更改网站标题和favicon
useDocument()

const route = useRoute();
const router = useRouter();
const storageStore = useStorageStore();
const permissionStore = usePermissionStore();
const toggleLeftDrawer = ref(false);
const topMenuItem = ref({});
const currentItemMenu = ref('dashboard');
const fabPos = ref([3, 80]);
const userProfile = ref(null);

const gqaFrontend = computed(() => {
    return storageStore.GetGqaFrontend()
})


onMounted(() => {
    topMenuItem.value = findTopItemMenu.value
})
watch(route, () => {
    topMenuItem.value = findTopItemMenu.value
})

const topMenu = computed(() => {
    return permissionStore.topMenu
})
const pageDashboard = computed(() => {
    const menu = topMenuItem.value?.top?.name === 'dashboard' || JSON.stringify(topMenuItem.value) === '{}'
    if (route.name === 'dashboard' && menu) {
        return true
    }
    return false
})
const findTopItemMenu = computed(() => {
    const name = route.name
    let item = {}
    for (let m of topMenu.value) {
        if (m.top.name === name || (m.arrayChildren && m.arrayChildren.find((item) => item.name === name))) {
            item = m
            break
        }
    }
    currentItemMenu.value = item.top ? item.top.name : ''
    return item
})

const changeTopMenu = (item) => {
    if (item.top.name === 'dashboard') {
        router.push('/dashboard')
    }
    topMenuItem.value = item
}

const moveFab = (ev) => {
    fabPos.value = [fabPos.value[0] - ev.delta.x, fabPos.value[1] - ev.delta.y]
}
</script>

<style lang="scss" scoped>
</style>
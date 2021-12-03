<template>
    <q-layout view="hHh LpR lFr">
        <q-header reveal elevated>
            <!-- <div class="row no-wrap shadow-1"> -->
            <q-toolbar class="bg-primary glossy ">
                <q-btn dense round glossy push icon="sync_alt" aria-label="Menu"
                    @click="toggleLeftDrawer = !toggleLeftDrawer" />

                <GqaAvatar class="gin-quasar-admin-logo" :src="gqaFrontend.gqaWebLogo" />

                <q-toolbar-title shrink class="text-bold text-italic">
                    {{ gqaFrontend.gqaSubTitle }}
                </q-toolbar-title>

                <q-tabs dense inline-label outside-arrows mobile-arrows shrink stretch v-model="currentItemMenu"
                    style="max-width: 50%;" class="text-white">
                    <q-tab v-for="item in topMenu" :key="item.top.name" :icon="item.top.icon" :name="item.top.name"
                        :label="item.top.title" @click="changeTopMenu(item)" />
                </q-tabs>

                <q-space />

                <Fullscreen style="margin: 0 5px" />
                <Notice style="margin: 0 5px" />
                <UserMenu style="margin: 0 5px" @showProfile="$refs.userProfile.show()" />
                <!-- <q-language-switcher/> -->
                <Setting style="margin: 0 5px" />
                <GitLink style="margin: 0 5px" v-if="gqaFrontend.gqaShowGit === 'yes'" />

            </q-toolbar>
            <!-- </div> -->
            <q-separator />
            <!-- header下面的标签页 -->
            <div class="row bg-primary">
                <!-- <ChipMenu /> -->
                <TabMenu v-show="!pageDashboard" />
            </div>

        </q-header>

        <q-drawer elevated v-if="!pageDashboard" v-model="toggleLeftDrawer" show-if-above bordered
            content-class="bg-grey-1">

            <SideBarLeft :topMenuItem="topMenuItem" />

        </q-drawer>

        <q-page-container>
            <router-view />
            <q-page-scroller position="bottom-right" :scroll-offset="150" :offset="[18, 18]">
                <q-btn dense fab push icon="keyboard_arrow_up" color="primary" />
            </q-page-scroller>
        </q-page-container>

        <q-footer reveal elevated>
            <PageFooter />
        </q-footer>

        <UserProfile ref="userProfile" />

    </q-layout>
</template>

<script>
import { mapGetters } from 'vuex'
import { gqaFrontendMixin } from 'src/mixins/gqaFrontendMixin'
import SideBarLeft from './SideBarLeft'
import TabMenu from './TabMenu'
import Fullscreen from './Fullscreen'
import Notice from './Notice'
import GitLink from './GitLink'
import UserMenu from './UserMenu'
import Setting from './Setting'
import PageFooter from './PageFooter'
import GqaAvatar from 'src/components/GqaAvatar'
import UserProfile from 'src/pages/UserProfile'

export default {
    name: 'MainLayout',
    mixins: [gqaFrontendMixin],
    components: {
        SideBarLeft,
        TabMenu,
        Fullscreen,
        Notice,
        GitLink,
        UserMenu,
        Setting,
        PageFooter,
        GqaAvatar,
        UserProfile,
    },
    computed: {
        ...mapGetters({
            topMenu: 'permission/topMenu',
        }),
        pageDashboard() {
            const menu = this.topMenuItem.top.name === 'dashboard' || JSON.stringify(this.topMenuItem) === '{}'
            if (this.$route.name === 'dashboard' && menu) {
                return true
            }
            return false
        },
        findTopItemMenu() {
            const name = this.$route.name
            let item = {}
            for (let m of this.topMenu) {
                if (m.top.name === name || (m.arrayChildren && m.arrayChildren.find((item) => item.name === name))) {
                    item = m
                    break
                }
            }
            this.currentItemMenu = item.top ? item.top.name : ''
            return item
        },
    },
    watch: {
        $route() {
            this.topMenuItem = this.findTopItemMenu
        },
    },
    data() {
        return {
            toggleLeftDrawer: false,
            topMenuItem: {},
            currentItemMenu: 'dashboard',
        }
    },
    created() {
        this.topMenuItem = this.findTopItemMenu
    },
    methods: {
        changeTopMenu(item) {
            if (item.top.name === 'dashboard') {
                this.$router.push('/dashboard')
            }
            this.topMenuItem = item
        },
    },
}
</script>

<style lang="scss" scoped>
</style>

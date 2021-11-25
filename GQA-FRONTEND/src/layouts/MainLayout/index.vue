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
                    <q-tab v-for="item in asideMenu" :key="item.name" :icon="item.icon" :name="item.name"
                        :label="item.title" @click="changeItemMenu(item)" />
                </q-tabs>

                <!-- <div class="q-gutter-sm">
                    <q-btn dense glossy push v-for="item in asideMenu" :key="item.name" :icon="item.icon"
                        @click="changeItemMenu(item)">
                        {{item.title}}
                    </q-btn>
                </div> -->

                <!-- <TabMenu /> -->
                <!-- </q-toolbar>

                <q-toolbar class="bg-primary glossy col-4"> -->
                <q-space />
                <!-- 简易面包屑 -->
                <!-- <q-breadcrumbs active-color="white" style="font-size: 13px; margin-left: 20px">
                    <q-breadcrumbs-el :label="item.title" :icon="item.icon"
                        v-for="item in matched.slice(1, matched.length)" :key="item.path" />
                    </q-breadcrumbs> -->

                <Fullscreen style="margin: 0 5px" />

                <Notice style="margin: 0 5px" />

                <UserMenu style="margin: 0 5px" />
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

            <SideBarLeft :itemMenu="itemMenu" />

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
    },
    computed: {
        matched() {
            return this.$route.matched
        },
        ...mapGetters({
            asideMenu: 'permission/aside',
            searchMenu: 'permission/search',
        }),
        pageDashboard() {
            const menu = this.itemMenu.name === 'dashboard' || JSON.stringify(this.itemMenu) === '{}'
            if (this.$route.name === 'dashboard' && menu) {
                return true
            }
            return false
        },
    },
    watch: {
        $route() {
            if (this.searchMenu.length) {
                const parentCode = this.searchMenu.filter((item) => item.name === this.$route.name)[0].parentCode
                if (parentCode) {
                    this.changeItemMenu(this.findInAsideMenu(parentCode, this.asideMenu)[0])
                    this.currentItemMenu = this.findInAsideMenu(parentCode, this.asideMenu)[0].name
                } else {
                    this.changeItemMenu(this.asideMenu[0])
                    this.currentItemMenu = this.asideMenu[0].name
                }
            }
        },
    },
    data() {
        return {
            toggleLeftDrawer: false,
            itemMenu: {},
            currentItemMenu: 'dashboard',
        }
    },
    created() {
        const name = this.$route.name
        const parentCode = this.searchMenu.filter((item) => item.name === name)[0].parentCode
        if (parentCode) {
            this.currentItemMenu = parentCode
            this.changeItemMenu(this.findInAsideMenu(parentCode, this.asideMenu)[0])
        }
    },
    methods: {
        changeItemMenu(item) {
            if (item.name === 'dashboard') {
                this.$router.push('/dashboard')
            }
            this.itemMenu = item
        },
        findInAsideMenu(parentCode, menu) {
            if (menu.filter((item) => item.name === parentCode).length === 0) {
                for (let subMenu of menu) {
                    if (subMenu.children && subMenu.children.filter((subItem) => subItem.name === parentCode).length !== 0) {
                        return [subMenu]
                    }
                }
                for (let subMenu of menu) {
                    return this.findInAsideMenu(parentCode, subMenu)
                }
            } else {
                return menu.filter((item) => item.name === parentCode)
            }
        },
    },
}
</script>

<style lang="scss" scoped>
</style>

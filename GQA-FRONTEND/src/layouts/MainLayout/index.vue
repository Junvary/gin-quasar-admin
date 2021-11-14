<template>
    <q-layout view="hHh LpR lFr">
        <q-header reveal elevated>
            <!-- <div class="row no-wrap shadow-1"> -->
            <q-toolbar class="bg-primary glossy ">
                <q-btn flat dense round icon=" sync_alt" aria-label="Menu" @click="leftDrawerOpen = !leftDrawerOpen" />

                <q-avatar class="gin-quasar-admin-logo">
                    <img src="gqa128.png" />
                </q-avatar>

                <q-toolbar-title shrink class="text-bold text-italic">
                    Gin-Quasar-Admin
                </q-toolbar-title>

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
                <GitLink style="margin: 0 5px" />

            </q-toolbar>
            <!-- </div> -->
            <q-separator />
            <!-- header下面的标签页 -->
            <div class="row bg-primary">
                <!-- <ChipMenu /> -->
                <TabMenu />
            </div>

        </q-header>

        <q-drawer elevated v-model="leftDrawerOpen" show-if-above bordered content-class="bg-grey-1">
            <q-list>
                <q-item clickable v-ripple>
                    <q-item-section class="text-primary text-bold text-center">
                         {{ $t('LayoutMainWelcome') }}
                    </q-item-section>
                </q-item>
                <SideBarLeft />
            </q-list>
        </q-drawer>

        <q-page-container>
            <router-view />
        </q-page-container>

        <q-footer reveal elevated>
            <PageFooter />
        </q-footer>
    </q-layout>
</template>

<script>
import SideBarLeft from './SideBarLeft'
import TabMenu from './TabMenu'
import Fullscreen from './Fullscreen'
import Notice from './Notice'
import GitLink from './GitLink'
import UserMenu from './UserMenu'
import Setting from './Setting'
import PageFooter from './PageFooter'

export default {
    name: 'MainLayout',
    components: {
        SideBarLeft,
        TabMenu,
        Fullscreen,
        Notice,
        GitLink,
        UserMenu,
        Setting,
        PageFooter,
    },
    computed: {
        matched() {
            return this.$route.matched
        },
    },
    data() {
        return {
            leftDrawerOpen: false,
        }
    },
}
</script>

<style lang="scss" scoped>
</style>

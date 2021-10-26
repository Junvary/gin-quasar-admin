<template>
    <q-tabs dense inline-label outside-arrows mobile-arrows class="bg-primary text-white shadow-2" style="width: 100%">
        <q-route-tab exact replace v-for="tab in tabMenus" :to="tab.path" :key="tab.path" :name="tab.path"
            :ripple="{ color: 'primary' }">
            <template v-slot>
                <q-icon size="1.3rem" v-if="tab.meta.icon" :name="tab.meta.icon" />
                <span class="tab-label">{{ tab.meta.title || '未命名' }}</span>
                <q-icon v-if="tab.path !== '/dashboard'" class="tab-close" name="close"
                    @click.prevent.stop="removeTab(tab)" />
                <q-menu touch-position context-menu>
                    <q-list dense bordered separator class="text-primary">
                        <q-item clickable v-close-popup>
                            <q-item-section avatar style="min-width: 0px">
                                <q-icon name="code" />
                            </q-item-section>
                            <q-item-section @click="removeOtherTab(tab)">
                                关闭其他
                            </q-item-section>
                        </q-item>
                        <q-item clickable v-close-popup>
                            <q-item-section avatar style="min-width: 0px">
                                <q-icon name="keyboard_arrow_right" />
                            </q-item-section>
                            <q-item-section @click="removeRightTab(tab)">
                                关闭右侧
                            </q-item-section>
                        </q-item>
                        <q-item clickable v-close-popup>
                            <q-item-section avatar style="min-width: 0px">
                                <q-icon name="keyboard_arrow_left" />
                            </q-item-section>
                            <q-item-section @click="removeLeftTab(tab)">
                                关闭左侧
                            </q-item-section>
                        </q-item>
                        <q-item clickable v-close-popup>
                            <q-item-section avatar style="min-width: 0px">
                                <q-icon name="close" />
                            </q-item-section>
                            <q-item-section @click="removeAllTab">
                                关闭所有
                            </q-item-section>
                        </q-item>
                    </q-list>
                </q-menu>
            </template>
        </q-route-tab>
    </q-tabs>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
    name: 'TabMenu',
    data() {
        return {}
    },
    computed: {
        ...mapState('tabMenu', ['tabMenus', 'currentTab']),
    },
    watch: {
        $route() {
            this.addTabMenu()
        },
    },
    beforeDestroy() {
        this.DestroyTabMenu()
    },
    created() {
        this.addTabMenu()
    },
    methods: {
        ...mapActions('tabMenu', ['AddTabMenu', 'ChangeCurrentTab', 'RemoveTab', 'RemoveRightTab', 'RemoveLeftTab', 'DestroyTabMenu']),
        addTabMenu() {
            this.AddTabMenu(this.$route)
        },
        toTab(tab) {
            this.ChangeCurrentTab(tab)
            this.$nextTick(() => {
                this.$router.push({ path: tab.path })
            })
        },
        removeTab(tab) {
            this.RemoveTab(tab)
            this.$nextTick(() => {
                this.$router.push({ path: this.currentTab.path })
            })
        },
        removeOtherTab(tab) {
            this.DestroyTabMenu()
            this.AddTabMenu(tab)
            // 没有点击在当前激活菜单上，那么跳转到这个菜单
            if (tab.path !== this.$route.path) {
                this.$nextTick(() => {
                    this.$router.push({ path: tab.path })
                })
            }
        },
        removeRightTab(tab) {
            this.RemoveRightTab(tab)
            // 没有点击在当前激活菜单上，那么跳转到这个菜单
            if (tab.path !== this.$route.path) {
                this.$nextTick(() => {
                    this.$router.push({ path: tab.path })
                })
            }
        },
        removeLeftTab(tab) {
            this.RemoveLeftTab(tab)
            // 没有点击在当前激活菜单上，那么跳转到这个菜单
            if (tab.path !== this.$route.path) {
                this.$nextTick(() => {
                    this.$router.push({ path: tab.path })
                })
            }
        },
        removeAllTab() {
            this.DestroyTabMenu()
            this.AddTabMenu()
            this.$nextTick(() => {
                this.$router.push({ path: '/' })
            })
        },
    },
}
</script>

<style lang="scss" scoped>
.tab-label {
    margin: 0 7px;
    white-space: nowrap;
    max-width: 150px;
    overflow: hidden;
    text-overflow: ellipsis;
}
.tab-close {
    display: inline-flex;
    font-size: 1rem;
    border-radius: 0.2rem;
    opacity: 0.5;
    transition: all 0.3s;
    &:hover {
        opacity: 1;
    }
}
</style>
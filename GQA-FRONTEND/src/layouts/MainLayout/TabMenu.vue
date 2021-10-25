<template>
    <q-tabs dense inline-label outside-arrows mobile-arrows class="bg-primary text-white shadow-2" style="width: 100%">
        <q-route-tab exact replace v-for="tab in tabMenus" :to="tab.fullPath" :key="tab.fullPath" :name="tab.fullPath"
            :ripple="{ color: 'primary' }" content-class="activeColor">
            <template v-slot>
                <q-icon size="1.3rem" v-if="tab.meta.icon" :name="tab.meta.icon" />
                <span class="tab-label">{{ tab.meta.title || '未命名' }}</span>
                <q-icon v-if="tab.fullPath !== '/index'" class="tab-close" name="close" @click="removeTab(tab)" />
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
        activeColor() {
            return (tab) => {
                if (tab.fullPath === this.$route.fullPath) {
                    return 'green'
                } else {
                    return 'primary'
                }
            }
        },
    },
    watch: {
        $route() {
            this.addTabMenu()
        },
    },
    beforeDestroy() {
        this.RemoveTabMenu()
    },
    created() {
        this.addTabMenu()
    },
    methods: {
        ...mapActions('tabMenu', ['AddTabMenu', 'ChangeCurrentTab', 'RemoveTab', 'RemoveTabMenu']),

        addTabMenu() {
            this.AddTabMenu(this.$route)
        },
        toTab(tab) {
            this.ChangeCurrentTab(tab)
            this.$nextTick(() => {
                this.$router.push({ path: tab.fullPath })
            })
        },
        removeTab(tab) {
            // actions里返回的是个Promose。。。
            this.RemoveTab(tab).then((res) => {
                if (res) {
                    this.$router.push({ path: this.currentTab.fullPath })
                } else {
                    this.$q.notify({
                        message: '首页不能被关闭！',
                        color: 'negative',
                    })
                }
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
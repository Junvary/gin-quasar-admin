<template>
    <q-tabs dense inline-label outside-arrows mobile-arrows class="text-white bg-primary shadow-2" style="width: 100%"
        v-if="!loginPage">
        <q-route-tab exact replace v-for="tab in tabMenus" :to="tab.path" :key="tab.path" :name="tab.path"
            :ripple="{ color: 'primary' }">
            <template v-slot>
                <q-icon size="1.3rem" v-if="tab.meta.icon" :name="tab.meta.icon" />
                <span class="tab-label">{{ $t(tab.meta.title) || $t('Unknown') }}</span>
                <q-icon v-if="tab.path !== '/dashboard'" class="tab-close" name="close"
                    @click.prevent.stop="removeTab(tab)" />
                <q-menu touch-position context-menu>
                    <q-list dense bordered separator class="text-primary">
                        <q-item clickable v-close-popup>
                            <q-item-section avatar style="min-width: 0px">
                                <q-icon name="code" />
                            </q-item-section>
                            <q-item-section @click="removeOtherTab(tab)">
                                {{ $t('CloseOther') }}
                            </q-item-section>
                        </q-item>
                        <q-item clickable v-close-popup>
                            <q-item-section avatar style="min-width: 0px">
                                <q-icon name="keyboard_arrow_right" />
                            </q-item-section>
                            <q-item-section @click="removeRightTab(tab)">
                                {{ $t('CloseRight') }}
                            </q-item-section>
                        </q-item>
                        <q-item clickable v-close-popup>
                            <q-item-section avatar style="min-width: 0px">
                                <q-icon name="keyboard_arrow_left" />
                            </q-item-section>
                            <q-item-section @click="removeLeftTab(tab)">
                                {{ $t('CloseLeft') }}
                            </q-item-section>
                        </q-item>
                        <q-item clickable v-close-popup>
                            <q-item-section avatar style="min-width: 0px">
                                <q-icon name="close" />
                            </q-item-section>
                            <q-item-section @click="removeAllTab">
                                {{ $t('CloseAll') }}
                            </q-item-section>
                        </q-item>
                    </q-list>
                </q-menu>
            </template>
        </q-route-tab>
    </q-tabs>
</template>

<script setup>
import { computed, ref, watch, onMounted, onUnmounted, nextTick } from 'vue';
import { useTabMenuStore } from 'src/stores/tabMenu'
import { useRoute, useRouter } from 'vue-router';
const tabMenuStore = useTabMenuStore()

const router = useRouter()
const route = useRoute()
const tabMenus = computed(() => tabMenuStore.tabMenus)
const currentTab = computed(() => tabMenuStore.currentTab)
watch(route, () => {
    tabMenuStore.AddTabMenu(Object.assign({}, route))
})
onMounted(() => {
    tabMenuStore.AddTabMenu(Object.assign({}, route))
})
onUnmounted(() => {
    tabMenuStore.DestroyTabMenu()
})
const loginPage = ref(false)

const removeTab = (tab) => {
    tabMenuStore.RemoveTab(tab)
    nextTick(() => {
        router.push({ path: currentTab.value.path })
    })
}
const removeOtherTab = (tab) => {
    tabMenuStore.DestroyTabMenu()
    tabMenuStore.AddTabMenu(tab)
    // 没有点击在当前激活菜单上，那么跳转到这个菜单
    if (tab.path !== route.path) {
        nextTick(() => {
            router.push({ path: tab.path })
        })
    }
}
const removeRightTab = (tab) => {
    tabMenuStore.RemoveRightTab(tab)
    // 没有点击在当前激活菜单上，那么跳转到这个菜单
    if (tab.path !== route.path) {
        nextTick(() => {
            router.push({ path: tab.path })
        })
    }
}
const removeLeftTab = (tab) => {
    tabMenuStore.RemoveLeftTab(tab)
    // 没有点击在当前激活菜单上，那么跳转到这个菜单
    if (tab.path !== route.path) {
        nextTick(() => {
            router.push({ path: tab.path })
        })
    }
}
const removeAllTab = () => {
    tabMenuStore.DestroyTabMenu()
    tabMenuStore.AddTabMenu()
    nextTick(() => {
        router.push({ path: '/' })
    })
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

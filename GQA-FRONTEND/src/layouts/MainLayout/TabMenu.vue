<template>
    <q-tabs dense inline-label outside-arrows mobile-arrows :class="darkThemeTab" style="width: 100%;" v-if="!loginPage"
        align="left">
        <q-route-tab exact replace v-for="tab in tabMenus" :to="tab.path" :key="tab.path" :name="tab.path">
            <template v-slot>
                <q-icon size="1.1rem" v-if="tab.meta.icon" :name="tab.meta.icon" />
                <span class="tab-label">{{ selectRouteLabel(tab) }}</span>
                <q-icon v-if="tab.name !== defaultPage" class="tab-close" name="close"
                    @click.prevent.stop="removeTab(tab)" />
                <q-menu touch-position context-menu>
                    <q-list dense bordered separator class="bg-white text-dark">
                        <q-item clickable v-close-popup v-ripple>
                            <q-item-section @click="removeOtherTab(tab)">
                                {{ $t('CloseOther') }}
                            </q-item-section>
                        </q-item>
                        <q-item clickable v-close-popup v-ripple>
                            <q-item-section @click="removeRightTab(tab)">
                                {{ $t('CloseRight') }}
                            </q-item-section>
                        </q-item>
                        <q-item clickable v-close-popup v-ripple>
                            <q-item-section @click="removeLeftTab(tab)">
                                {{ $t('CloseLeft') }}
                            </q-item-section>
                        </q-item>
                        <q-item clickable v-close-popup v-ripple>
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
import useTheme from 'src/composables/useTheme';
import { usePermissionStore } from 'src/stores/permission';
import useCommon from 'src/composables/useCommon'

const { selectRouteLabel } = useCommon()
const tabMenuStore = useTabMenuStore()
const { darkThemeTab } = useTheme()
const router = useRouter()
const route = useRoute()
const permissionStore = usePermissionStore()

const tabMenus = computed(() => tabMenuStore.tabMenus)
const currentTab = computed(() => tabMenuStore.currentTab)
const defaultPage = computed(() => permissionStore.defaultPage[0])

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
    if (tab.path === currentTab.value.path) {
        tabMenuStore.RemoveTab(tab)
        nextTick(() => {
            router.push({ path: currentTab.value.path })
        })
    } else {
        tabMenuStore.RemoveTab(tab)
    }
}
const removeOtherTab = (tab) => {
    tabMenuStore.DestroyTabMenu()
    tabMenuStore.AddTabMenu(tab)
    // If the current active menu is not clicked, then jump to this menu
    if (tab.path !== route.path) {
        nextTick(() => {
            router.push({ path: tab.path })
        })
    }
}
const removeRightTab = (tab) => {
    tabMenuStore.RemoveRightTab(tab)
    // If the current active menu is not clicked, then jump to this menu
    if (tab.path !== route.path) {
        nextTick(() => {
            router.push({ path: tab.path })
        })
    }
}
const removeLeftTab = (tab) => {
    tabMenuStore.RemoveLeftTab(tab)
    // If the current active menu is not clicked, then jump to this menu
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
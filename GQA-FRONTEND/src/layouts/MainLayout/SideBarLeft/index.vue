<template>
    <div class="q-mini-drawer-hide absolute" style="top: 0px; left: 0px">
        <TabMenu />
    </div>
    <q-card clickable class="absolute-top text-primary text-bold text-center text-h6 cursor-pointer"
        style="height: 60px; width: calc(100% - 48px); margin-left: 48px;">
        <q-item-section class="row items-center justify-center absolute-bottom">
            <q-icon :name="topMenuItem.top?.icon" size="sm" v-if="topMenuItem.top?.icon" />
            {{ topMenuItem.top ? $t(topMenuItem.top.title) : "" }}
        </q-item-section>
    </q-card>
    <q-scroll-area style="height: calc(100% - 60px); margin-top: 60px; margin-left: 48px;">
        <q-list>
            <template v-for="(item, index) in topMenuItem.treeChildren ? topMenuItem.treeChildren : []" :key="index">
                <SideBarLeftItem :addRoutesItem="item" :initLevel="0" />
            </template>
        </q-list>
    </q-scroll-area>
</template>

<script setup>
import TabMenu from '../TabMenu.vue'
import { toRefs } from 'vue';
import SideBarLeftItem from './SideBarLeftItem'
const props = defineProps({
    topMenuItem: {
        type: Object,
        required: false,
        default: () => {
            return {}
        },
    },
})
const { topMenuItem } = toRefs(props)
</script>

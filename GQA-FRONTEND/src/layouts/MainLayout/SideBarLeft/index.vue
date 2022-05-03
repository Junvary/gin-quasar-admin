<template>
    <q-list>
        <q-item clickable class="text-primary text-bold text-center text-h6">
            <q-item-section class="row items-center justify-center">
                <q-icon :name="topMenuItem.top?.icon" size="sm" v-if="topMenuItem.top?.icon" />
                {{ topMenuItem.top ? $t(topMenuItem.top.title) : "" }}
            </q-item-section>
        </q-item>
        <template v-for="(item, index) in topMenuItem.treeChildren ? topMenuItem.treeChildren : []" :key="index">
            <SideBarLeftItem :addRoutesItem="item" :initLevel="0" />
        </template>
    </q-list>
</template>

<script setup>
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

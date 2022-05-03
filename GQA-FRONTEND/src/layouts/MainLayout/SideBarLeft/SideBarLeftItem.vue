<template>
    <component :is="chooseComponent" :addRoutesItem="addRoutesItem" :initLevel="initLevel">
        <template v-if="addRoutesItem.children && addRoutesItem.children.length">
            <SideBarLeftItem :key="item.id" :addRoutesItem="item" v-for="item in addRoutesItem.children"
                :initLevel="initLevel + 0.3" />
        </template>
    </component>
</template>

<script setup>
import ItemSingle from './ItemSingle'
import ItemMultiple from './ItemMultiple'
import { computed, toRefs } from 'vue';

const props = defineProps({
    addRoutesItem: {
        default: function () {
            return null
        },
        type: Object,
    },
    initLevel: {
        type: Number,
        default: 0,
    },
})
const { addRoutesItem, initLevel } = toRefs(props)
const chooseComponent = computed(() => {
    if (addRoutesItem.value?.children?.length) {
        return ItemMultiple
    } else {
        return ItemSingle
    }
})
</script>


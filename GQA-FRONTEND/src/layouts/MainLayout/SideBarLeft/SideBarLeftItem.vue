<template>
    <component :is="chooseComponent" :trueItem="childrenItem" :initLevel="initLevel">
        <template v-if="childrenItem.children && childrenItem.children.length !== 0">
            <SideBarLeftItem v-for="item in childrenItem.children" :childrenItem="item" :initLevel="initLevel + 0.3" />
        </template>
    </component>
</template>

<script setup>
import ItemSingle from './ItemSingle.vue'
import ItemMultiple from './ItemMultiple.vue'
import { computed, toRefs } from 'vue';

const props = defineProps({
    childrenItem: {
        default: function () {
            return {}
        },
        type: Object,
    },
    initLevel: {
        type: Number,
        default: 0,
    },
})
const { childrenItem, initLevel } = toRefs(props)

const chooseComponent = computed(() => {
    if (childrenItem.value.children && childrenItem.value.children.length !== 0) {
        return ItemMultiple
    } else {
        return ItemSingle
    }
})
</script>
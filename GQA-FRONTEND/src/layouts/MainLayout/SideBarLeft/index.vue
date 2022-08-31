<template>
    <q-scroll-area style="height: calc(100%);">
        <q-list>
            <template v-for="(childrenItem, index) in topMenuChildren" :key="index">
                <SideBarLeftItem :childrenItem="childrenItem" :initLevel="0" />
            </template>
        </q-list>
    </q-scroll-area>
</template>

<script setup>
import { toRefs } from 'vue';
import SideBarLeftItem from './SideBarLeftItem'

const props = defineProps({
    topMenuChildren: {
        type: Array,
        required: false,
        default: () => {
            return []
        },
    },
})
const { topMenuChildren } = toRefs(props)
</script>

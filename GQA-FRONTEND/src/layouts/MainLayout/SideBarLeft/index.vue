<template>
    <q-list>
        <q-item clickable v-ripple class="text-primary text-bold text-center text-h6">
            <q-item-section>
                {{ topMenuItem.top ? $t(topMenuItem.top.title) : "" }}
            </q-item-section>
        </q-item>
        <template v-for="(item, index) in topMenuItem.treeChildren ? topMenuItem.treeChildren : []" :key="index">
            <SideBarLeftItem :addRoutesItem="item" :initLevel="0" />
        </template>
    </q-list>
</template>

<script>
import SideBarLeftItem from './SideBarLeftItem'

export default {
    name: 'SideBarLeft',
    components: {
        SideBarLeftItem,
    },
    props: {
        topMenuItem: {
            type: Object,
            required: false,
            default: () => {
                return {}
            },
        },
    },
}
</script>
<template>
    <component :is="chooseComponent" :addRoutesItem="addRoutesItem" :initLevel="initLevel">
        <template v-if="addRoutesItem.children && addRoutesItem.children.length">
            <SideBarLeftItem :key="item.id" :addRoutesItem="item" v-for="item in addRoutesItem.children"
                :initLevel="initLevel + 0.3" />
        </template>
    </component>
</template>

<script>
import ItemSingle from './ItemSingle'
import ItemMultiple from './ItemMultiple'

export default {
    name: 'SideBarLeftItem',
    props: {
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
    },
    components: {
        ItemSingle,
        ItemMultiple,
    },
    computed: {
        chooseComponent() {
            if (this.addRoutesItem.children) {
                return 'ItemMultiple'
            } else {
                return 'ItemSingle'
            }
        },
    },
}
</script>

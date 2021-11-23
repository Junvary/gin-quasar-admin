<template>
    <q-list>
        <q-item clickable v-ripple class="text-primary text-bold text-center text-h6">
            <q-item-section>
                {{ itemMenu.title }}
            </q-item-section>
        </q-item>
        <template v-for="(item, index) in itemMenu.children ? itemMenu.children : []" :key="index">
            <SideBarLeftItem :addRoutesItem="item" :initLevel="0" />
        </template>
    </q-list>
</template>

<script>
import { mapGetters } from 'vuex'
import SideBarLeftItem from './SideBarLeftItem'

export default {
    name: 'SideBarLeft',
    components: {
        SideBarLeftItem,
    },
    props: {
        itemMenu: {
            type: Object,
            required: false,
            default: () => {
                return this.addRoutes[0]
            },
        },
    },
    computed: {
        ...mapGetters({
            addRoutes: 'permission/aside',
        }),
    },
}
</script>
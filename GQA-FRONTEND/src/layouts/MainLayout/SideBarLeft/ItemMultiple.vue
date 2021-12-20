<template>
    <q-expansion-item :group="addRoutesItem.name" v-model="itemOpen" :header-inset-level="initLevel">
        <template v-slot:header>
            <q-item-section avatar>
                <q-icon :name="addRoutesItem.icon" />
            </q-item-section>
            <q-item-section>{{ $t(addRoutesItem.title) }}</q-item-section>
        </template>
        <slot></slot>
    </q-expansion-item>
</template>

<script>
export default {
    name: 'AsyncSubmenu',
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
    watch: {
        $route() {
            this.changeOpen()
        },
    },
    created() {
        this.changeOpen()
    },
    data() {
        return {
            itemOpen: false,
        }
    },
    methods: {
        changeOpen() {
            for (let item of this.addRoutesItem.children) {
                if (item.path === this.$route.path || item.parentCode === this.$route.name) {
                    this.itemOpen = true
                    return
                }
            }
            this.itemOpen = false
        },
    },
}
</script>

<style lang="scss">
</style>
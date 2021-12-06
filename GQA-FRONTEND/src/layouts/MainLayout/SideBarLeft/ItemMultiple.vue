<template>
    <q-expansion-item :group="addRoutesItem.name" :default-opened="itemOpen" :header-inset-level="initLevel">
        <template v-slot:header>
            <q-item-section avatar>
                <q-icon :name="addRoutesItem.icon" color="primary" />
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
    computed: {
        itemOpen() {
            for (let item of this.addRoutesItem.children) {
                if (item.path === this.$route.path) {
                    return true
                }
            }
            return false
        },
    },
}
</script>
<style lang="scss">
</style>
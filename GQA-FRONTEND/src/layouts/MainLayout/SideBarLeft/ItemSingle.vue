<template>
    <q-item clickable v-ripple exact :to="{ path: addRoutesItem.path }" :inset-level="initLevel"
        active-class="bg-primary text-white text-bold text-italic">
        <q-item-section avatar>
            <q-icon :name="addRoutesItem.icon" />
        </q-item-section>
        <q-item-section>{{ $t(addRoutesItem.title) }}</q-item-section>
    </q-item>
</template>

<script>
export default {
    name: 'MenuItem',
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
}
</script>

<style lang="scss" scoped>
.item-active-class {
    color: $primary;
    background: $primary;
    // background: lighten($color: $primary, $amount: 30%);
    // background: scale-color($primary, $lightness: 5%);
}
</style>
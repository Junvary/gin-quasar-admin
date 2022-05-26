<template>
    <q-item clickable exact @click="toPath(addRoutesItem)" :inset-level="initLevel"
        active-class="bg-primary text-white text-bold text-italic">
        <q-item-section avatar>
            <q-icon :name="addRoutesItem.icon" />
        </q-item-section>
        <q-item-section>{{ $t(addRoutesItem.title) }}</q-item-section>
    </q-item>
</template>

<script setup>
import { toRefs } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()
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

const toPath = (addRoutesItem) => {
    if (addRoutesItem.is_link === 'yes') {
        window.open(addRoutesItem.path)
    } else {
        router.push(addRoutesItem.path)
    }
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
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

<script setup>
import { watch, onMounted, ref, toRefs } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()
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
watch(route, () => {
    changeOpen()
})
onMounted(() => {
    changeOpen()
})
const itemOpen = ref(false)
const changeOpen = () => {
    for (let item of addRoutesItem.value.children) {
        if (item.path === route.path || item.parent_code === route.name) {
            itemOpen.value = true
            return
        }
    }
    itemOpen.value = false
}
</script>

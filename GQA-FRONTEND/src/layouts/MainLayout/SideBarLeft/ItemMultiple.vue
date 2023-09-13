<template>
    <q-expansion-item v-model="itemOpen" :group="trueItem.name" :header-inset-level="initLevel" :dense="sidebarDense">
        <template v-slot:header>
            <q-item-section avatar>
                <q-icon :name="trueItem.icon" />
            </q-item-section>
            <q-item-section>
                {{ selectOptionLabel(trueItem) }}
            </q-item-section>
        </template>
        <slot></slot>
    </q-expansion-item>
</template>

<script setup>
import { watch, onMounted, ref, toRefs, computed } from 'vue';
import { useRoute } from 'vue-router';
import useCommon from 'src/composables/useCommon';
import { useSettingStore } from 'src/stores/setting';

const settingStore = useSettingStore();
const { selectOptionLabel } = useCommon()
const route = useRoute()
const props = defineProps({
    trueItem: {
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
const { trueItem, initLevel } = toRefs(props)

const sidebarDense = computed(() => settingStore.GetSidebarDense())

watch(route, () => {
    changeOpen()
})

onMounted(() => {
    changeOpen()
})

const itemOpen = ref(false)
const changeOpen = () => {
    for (let item of trueItem.value.children) {
        if (item.path === route.path || item.parent_code === route.name) {
            itemOpen.value = true
            return
        }
    }
    itemOpen.value = false
}
</script>
<template>
    <q-item clickable v-ripple exact @click="toPath(trueItem)" :inset-level="initLevel" :active="checkActive"
        :active-class="darkThemeSelect" :dense="sidebarDense">
        <q-item-section avatar>
            <q-icon :name="trueItem.icon" />
        </q-item-section>
        <q-item-section>
            {{ selectOptionLabel(trueItem) }}
        </q-item-section>
    </q-item>
</template>

<script setup>
import { computed, inject, toRefs } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import useTheme from 'src/composables/useTheme';
import useCommon from 'src/composables/useCommon';
import { useSettingStore } from 'src/stores/setting';

const settingStore = useSettingStore();
const { selectOptionLabel } = useCommon()
const { darkThemeSelect } = useTheme()
const router = useRouter()
const route = useRoute()
const props = defineProps({
    trueItem: {
        default: function () {
            return {}
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

const bus = inject('bus')
const toPath = (item) => {
    if (item.is_link === 'yesNo_yes') {
        window.open(item.path)
    } else {
        bus.emit('changeRoute', true)
        router.push(item.path).then(() => {
            bus.emit('changeRoute', false)
        })
    }
}

const checkActive = computed(() => {
    if (route.path === trueItem.value.path) {
        return true
    } else {
        return false
    }
})
</script>

<style lang="scss" scoped>
.item-active-class {
    color: $primary;
    background: $primary;
}

.q-item {
    border-radius: 0 32px 32px 0
}
</style>
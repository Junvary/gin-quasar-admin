<template>
    <q-card :style="{backgroundImage: gradualColor, borderRadius: borderRadius}" :bordered="bordered">
        <slot />
    </q-card>
</template>

<script setup>
import { computed, toRefs } from 'vue';

const props = defineProps({
    startColor: {
        type: String,
        required: false,
    },
    endColor: {
        type: String,
        required: false
    },
    borderRadius: {
        type: String,
        required: false,
        default: '15px'
    },
    bordered: {
        type: Boolean,
        required: false,
        default: false
    }
})
const { startColor, endColor, borderRadius, bordered } = toRefs(props)
const gradualColor = computed(() => {
    return `linear-gradient(to bottom right, ${startColor.value}, ${endColor.value}`
})
</script>
<template>
    <td data-th="Name" class="text-left">
        <div v-bind:style="treeTd.setPadding(treeTd.item)"
            :class="treeTd.iconName(treeTd.item) != 'done' ? 'q-pl-lg' : ''">
            <q-btn @click="treeTd.toggle(treeTd.item)" v-if="treeTd.iconName(treeTd.item) != 'done'"
                :icon="treeTd.iconName(treeTd.item)" flat dense>
            </q-btn>
            <span class="q-ml-sm">{{ treeTd.item[firstTd] }}</span>
        </div>
    </td>
</template>

<script setup>
import { toRefs } from 'vue';

const props = defineProps({
    treeTd: {
        type: Object,
        required: true,
    },
    firstTd: {
        type: String,
        required: true,
    }
})
const { treeTd, firstTd } = toRefs(props)
</script>
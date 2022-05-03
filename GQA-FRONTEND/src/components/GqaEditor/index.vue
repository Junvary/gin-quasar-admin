<template>
    <q-editor :toolbar="toolbar" style="width: 100%" v-model="content" />
</template>

<script setup>
import { toRefs } from 'vue';

const props = defineProps({
    content: {
        type: String,
        required: false,
        default: '',
    },
})
const { content } = toRefs(props)
const toolbar = [
    ['unordered', 'ordered'],
    ['outdent', 'indent'],
    [
        {
            label: this.$q.lang.editor.align,
            icon: this.$q.iconSet.editor.align,
            fixedLabel: true,
            options: ['left', 'center', 'right', 'justify'],
        },
    ],
    ['bold', 'italic', 'strike', 'underline', 'subscript', 'superscript', 'hr'],
    [
        {
            label: this.$q.lang.editor.formatting,
            icon: this.$q.iconSet.editor.formatting,
            list: 'icons',
            options: ['h1', 'h2', 'h3', 'h4', 'h5', 'h6', 'p', 'code'],
        },
    ],
    [
        {
            label: this.$q.lang.editor.fontSize,
            icon: this.$q.iconSet.editor.fontSize,
            list: 'no-icons',
            options: ['size-1', 'size-2', 'size-3', 'size-4', 'size-5', 'size-6', 'size-7'],
        },
    ],

    ['link', 'quote', 'removeFormat'],
    ['undo', 'redo', 'print', 'fullscreen', 'viewsource'],
]
</script>

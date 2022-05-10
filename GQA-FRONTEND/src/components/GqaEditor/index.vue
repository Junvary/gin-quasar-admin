<template>
    <q-editor :toolbar="toolbar" style="width: 100%" v-model="editorContent" @update:model-value="contentChange" />
</template>

<script setup>
import { ref, watch } from 'vue';
import { useQuasar } from 'quasar';

const $q = useQuasar();

const props = defineProps({
    modelValue: {
        type: String,
        required: false,
        default: "",
    },
})
const editorContent = ref(props.modelValue)
watch(props, () => {
    editorContent.value = props.modelValue
})

const emit = defineEmits(['update:modelValue'])
const contentChange = (vlaue) => {
    emit('update:modelValue', vlaue)
}
const toolbar = [
    ['unordered', 'ordered'],
    ['outdent', 'indent'],
    [
        {
            label: $q.lang.editor.align,
            icon: $q.iconSet.editor.align,
            fixedLabel: true,
            options: ['left', 'center', 'right', 'justify'],
        },
    ],
    ['bold', 'italic', 'strike', 'underline', 'subscript', 'superscript', 'hr'],
    [
        {
            label: $q.lang.editor.formatting,
            icon: $q.iconSet.editor.formatting,
            list: 'icons',
            options: ['h1', 'h2', 'h3', 'h4', 'h5', 'h6', 'p', 'code'],
        },
    ],
    [
        {
            label: $q.lang.editor.fontSize,
            icon: $q.iconSet.editor.fontSize,
            list: 'no-icons',
            options: ['size-1', 'size-2', 'size-3', 'size-4', 'size-5', 'size-6', 'size-7'],
        },
    ],

    ['link', 'quote', 'removeFormat'],
    ['undo', 'redo', 'print', 'fullscreen', 'viewsource'],
]
</script>

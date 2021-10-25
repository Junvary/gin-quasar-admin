<template>
    <q-chip :color="show.color" text-color="white">
        {{show.text}}
    </q-chip>
</template>

<script>
export default {
    name: 'GqaStatus',
    props: {
        status: {
            type: String,
            required: false,
            default: '',
        },
    },
    computed: {
        show() {
            if (this.status === 'on') {
                return { color: 'positive', text: '启用' }
            } else if (this.status === 'off') {
                return { color: 'negative', text: '禁用' }
            } else {
                return { color: 'warning', text: '未知' }
            }
        },
    },
}
</script>
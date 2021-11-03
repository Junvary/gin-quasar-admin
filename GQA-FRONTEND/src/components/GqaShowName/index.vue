<template>
    <span>
        {{trueName}}
    </span>
</template>

<script>
import { GqaTrueName } from 'src/settings'

export default {
    name: 'GqaShowName',
    props: {
        customName: {
            type: String,
            required: false,
            default: '',
        },
    },
    computed: {
        trueName() {
            const nickname = this.$q.cookies.get('gqa-nickname')
            const realName = this.$q.cookies.get('gqa-realName')
            if (this.customName) {
                return this.customName
            } else if (nickname) {
                return nickname
            } else if (realName) {
                return realName
            } else {
                return GqaTrueName
            }
        },
    },
}
</script>
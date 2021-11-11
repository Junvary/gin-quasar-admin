<template>
    <span>
        {{trueName}}
    </span>
</template>

<script>
import { GqaUsername } from 'src/settings'

export default {
    name: 'GqaShowName',
    props: {
        customNameString: {
            type: String,
            required: false,
            default: '',
        },
        customNameObject: {
            type: Object,
            required: false,
            default: () => {
                return {}
            },
        },
        showMyName: {
            type: Boolean,
            required: false,
            default: false,
        },
    },
    computed: {
        trueName() {
            if (this.customNameString) {
                // 自定义名字
                return this.customNameString
            } else if (JSON.stringify(this.customNameObject) !== '{}') {
                // 其他用户的名字
                if (this.customNameObject.nickname) {
                    return this.customNameObject.nickname
                } else {
                    return this.customNameObject.realName
                }
            } else if (this.showMyName) {
                const nickname = this.$q.cookies.get('gqa-nickname')
                const realName = this.$q.cookies.get('gqa-realName')
                if (nickname) {
                    return nickname
                } else {
                    return realName
                }
            } else {
                return GqaUsername
            }
        },
    },
}
</script>
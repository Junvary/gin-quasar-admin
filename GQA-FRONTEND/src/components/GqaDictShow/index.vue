<template>
    <div>
        <span v-if="dictName === 'statusOnOff'">
            <q-chip dense text-color="white" color="positive" v-if="dictCode === 'on'">{{dictLabel}}</q-chip>
            <q-chip dense text-color="white" color="negative" v-else-if="dictCode === 'of'">{{dictLabel}}</q-chip>
            <q-chip dense v-else>{{dictLabel}}</q-chip>
        </span>
        <span v-else>
            {{dictLabel}}
        </span>
        <span v-if="withExt1">
            {{dictExt1 + ext1}}
        </span>
    </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
    name: 'GqaDictShow',
    props: {
        dictName: {
            type: String,
            required: true,
        },
        dictCode: {
            type: String,
            default: '',
        },
        withExt1: {
            type: Boolean,
            required: false,
            default: false,
        },
        ext1: {
            type: String,
            default: '',
        },
    },
    computed: {
        dictLabel() {
            if (this.dictCode !== '') {
                const codeList = this.dictCode.split(',')
                const dict = this.dictList[this.dictName]
                let label = ''
                for (let d of codeList) {
                    const l = dict.filter((item) => item.dictCode === d)[0].dictLabel
                    label += l + ' '
                }
                return label
            } else {
                return ''
            }
        },
        dictExt1() {
            if (this.dictCode !== '') {
                const codeList = this.dictCode.split(',')
                const dict = this.dictList[this.dictName]
                let ext1 = ''
                for (let d of codeList) {
                    const l = dict.filter((item) => item.dictCode === d)[0].dictExt1
                    ext1 += l + ' '
                }
                return ext1
            } else {
                return ''
            }
        },
    },
    data() {
        return {
            dictList: {},
        }
    },
    async created() {
        const detailLocal = this.$q.localStorage.getItem('gqa-dict')
        if (detailLocal) {
            this.dictList = detailLocal
        } else {
            await this.GetGqaDict()
            this.dictList = this.$q.localStorage.getItem('gqa-dict')
        }
    },
    methods: {
        ...mapActions('storage', ['GetGqaDict']),
    },
}
</script>
<template>
    <span v-if="dictName === 'statusOnOff'">
        <q-chip dense text-color="white" color="positive" v-if="dictCode === 'on'">{{dictLabel}}</q-chip>
        <q-chip dense text-color="white" color="negative" v-else-if="dictCode === 'of'">{{dictLabel}}</q-chip>
        <q-chip dense v-else>{{dictLabel}}</q-chip>
    </span>
    <span v-else>
        {{dictLabel}}
    </span>
</template>

<script>
import { postAction } from 'src/api/manage'
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
    },
    computed: {
        dictLabel() {
            const codeList = this.dictCode.split(',')
            const dict = this.dictList[this.dictName]
            let label = ''
            for (let d of codeList) {
                const l = dict.filter((item) => item.dictCode === d)[0].dictLabel
                label += l + ' '
            }
            return label
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
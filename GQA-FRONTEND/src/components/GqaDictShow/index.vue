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
            const dict = this.dictList[this.dictName]
            const label = dict.filter((item) => item.value === this.dictCode)[0].label
            return label
        },
    },
    data() {
        return {
            dictList: {},
        }
    },
    created() {
        this.dictList = this.$q.localStorage.getItem('gqa-dict')
    },
}
</script>
<template>
    <q-separator />
    <q-table dense hide-bottom hide-header separator="cell" :rows="rows" :columns="columns" row-key="name">
        <template v-slot:body-cell-first="props">
            <q-td :props="props">
                <q-chip dense color="accent" text-color="white">已装插件</q-chip>
            </q-td>
        </template>
        <template v-slot:body-cell-actions="props">
            <q-td :props="props">
                <span>
                    描述
                    <q-tooltip>
                        {{props.row.PluginRemark}}
                    </q-tooltip>
                </span>
            </q-td>
        </template>

    </q-table>
    <q-separator />
</template>

<script>
export default {
    name: 'GqaPluginList',
    computed: {
        rows() {
            return this.$q.localStorage.getItem('gqa-pluginList')
        },
    },
    data() {
        return {
            columns: [
                { name: 'first', align: 'center', label: '已装插件', field: 'first' },
                { name: 'PluginName', align: 'center', label: '插件名称', field: 'PluginName' },
                { name: 'PluginCode', align: 'center', label: '插件编码', field: 'PluginCode' },
                { name: 'PluginVersion', align: 'center', label: '插件版本', field: 'PluginVersion' },
                { name: 'actions', align: 'center', label: '操作', field: 'actions' },
            ],
        }
    },
}
</script>
<template>
    <q-expansion-item dense expand-separator class="bg-accent text-white" expand-icon-class="text-white">
        <template v-slot:header>
            <q-item-section avatar>
                <q-avatar size="md" icon="install_desktop" color="white" text-color="accent" />
            </q-item-section>
            <q-item-section>
                <q-btn dense flat :label="'已装插件（' + rows.length + '）'" />
            </q-item-section>
        </template>
        <q-table dense hide-bottom hide-header separator="cell" :rows="rows" :columns="columns" row-key="name">
            <template v-slot:body-cell-first="props">
                <q-td :props="props">
                    <q-chip dense color="accent" text-color="white">已装插件</q-chip>
                </q-td>
            </template>
            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <q-chip dense color="accent" text-color="white">
                        描述
                        <q-tooltip>
                            {{props.row.PluginRemark}}
                        </q-tooltip>
                    </q-chip>
                </q-td>
            </template>
        </q-table>
    </q-expansion-item>
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
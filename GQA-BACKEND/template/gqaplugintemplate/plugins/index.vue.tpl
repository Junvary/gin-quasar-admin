<template>
    <q-page padding>
        {{.PluginName}}
        {{.PluginCode}}
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'

export default {
    name: '{{.PluginCode}}',
    mixins: [tableDataMixin],
    data() {
        return {
        }
    },
    created() {
        // this.getTableData()
    },
}
</script>
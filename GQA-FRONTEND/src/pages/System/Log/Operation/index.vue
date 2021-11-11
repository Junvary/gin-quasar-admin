<template>
    <q-page padding>
        <q-table row-key="id" separator="cell" :data="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="handleDelete(props.row)" label="删除" />
                    </div>
                </q-td>
            </template>
        </q-table>
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'

export default {
    name: 'Operation',
    mixins: [tableDataMixin],
    data() {
        return {
            url: {
                list: 'api/system/operation_log/',
            },
            columns: [
                { name: 'request_modular', align: 'center', label: '请求模块', field: 'request_modular' },
                { name: 'request_path', align: 'center', label: '请求地址', field: 'request_path' },
                { name: 'request_method', align: 'center', label: '请求方法', field: 'request_method' },
                { name: 'request_ip', align: 'center', label: 'IP地址', field: 'request_ip' },
                { name: 'request_browser', align: 'center', label: '浏览器', field: 'request_browser' },
                { name: 'request_os', align: 'center', label: '操作系统', field: 'request_os' },
                { name: 'response_code', align: 'center', label: '响应码', field: 'response_code' },
                { name: 'json_result', align: 'center', label: '返回信息', field: 'json_result' },
                { name: 'create_datetime', align: 'center', label: '创建时间', field: 'create_datetime' },
                { name: 'actions', align: 'center', label: '操作', field: 'actions' },
            ],
        }
    },
    created() {
        this.getTableData()
    },
}
</script>
<template>
    <q-page padding>

        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.name" label="姓名" />
            <q-input style="width: 20%" v-model="queryParams.username" label="账号" />
            <q-btn color="primary" @click="handleSearch" label="搜索" />
            <q-btn color="primary" @click="resetSearch" label="重置" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" label="新增配置" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <gqa-dict-show dictName="statusOnOff" :dictCode="props.row.status" />
                </q-td>
            </template>

            <template v-slot:body-cell-stable="props">
                <q-td :props="props">
                    <gqa-dict-show dictName="statusYesNo" :dictCode="props.row.stable" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn dense color="primary" @click="handleSave(props.row)" label="保存编辑" />
                        <q-btn color="negative" @click="handleDelete(props.row)" label="删除" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <add-or-edit-dialog ref="addOrEditDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import addOrEditDialog from './modules/addOrEditDialog'
import { putAction } from 'src/api/manage'
import GqaDictShow from 'src/components/GqaDictShow'

export default {
    name: 'Config',
    mixins: [tableDataMixin],
    components: {
        addOrEditDialog,
        GqaDictShow,
    },
    data() {
        return {
            url: {
                list: 'config/config-list',
                edit: 'config/config-edit',
                delete: 'config/config-delete',
            },
            columns: [
                { name: 'sort', align: 'center', label: '排序', field: 'sort' },
                { name: 'gqaOption', align: 'center', label: '配置名', field: 'gqaOption' },
                { name: 'remark', align: 'center', label: '描述', field: 'remark' },
                { name: 'default', align: 'center', label: '默认配置', field: 'default' },
                { name: 'custom', align: 'center', label: '自定义配置', field: 'custom' },
                { name: 'status', align: 'center', label: '状态', field: 'status' },
                { name: 'stable', align: 'center', label: '系统内置', field: 'stable' },
                { name: 'actions', align: 'center', label: '操作', field: 'actions' },
            ],
        }
    },
    created() {
        this.getTableData()
    },
    methods: {
        async handleSave(row) {
            const res = await putAction(this.url.edit, row)
            if (res.code === 1) {
                this.$q.notify({
                    type: 'positive',
                    message: res.message,
                })
                this.getTableData()
            }
        },
    },
}
</script>
<template>
    <q-page style="margin: 0 16px">

        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="searchParams.value" label="字典编码" />
            <q-input style="width: 20%" v-model="searchParams.label" label="字典名" />
            <q-btn color="primary" @click="handleSearch" label="搜索" />
            <q-btn color="primary" @click="resetSearch" label="重置" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" label="新增字典" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <gqa-status :status="props.row.status" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="showEditForm(props.row)" label="编辑" />
                        <q-btn color="warning" @click="handleDetail(props.row)" label="维护" />
                        <q-btn color="negative" @click="handleDelete(props.row)" label="删除" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <add-or-edit-dialog ref="addOrEditDialog" @emitAddOrEdit="emitAddOrEdit" @handleFinish="handleFinish" />
        <dict-detail-dialog ref="dictDetailDialog" @emitAddOrEdit="emitAddOrEdit" @handleFinish="handleFinish" />
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import addOrEditDialog from './modules/addOrEditDialog'
import dictDetailDialog from './modules/dictDetailDialog'
import GqaStatus from 'src/components/GqaStatus'

export default {
    name: 'Dict',
    mixins: [tableDataMixin],
    components: {
        addOrEditDialog,
        dictDetailDialog,
        GqaStatus,
    },
    data() {
        return {
            url: {
                list: 'dict/dict-list',
                delete: 'dict/dict-delete',
            },
            columns: [
                { name: 'sort', align: 'center', label: '排序', field: 'sort' },
                { name: 'value', align: 'center', label: '字典编码', field: 'value' },
                { name: 'label', align: 'center', label: '字典名称', field: 'label' },
                { name: 'status', align: 'center', label: '状态', field: 'status' },
                { name: 'actions', align: 'center', label: '操作', field: 'actions' },
            ],
        }
    },
    methods: {
        handleDetail(row) {
            this.$refs.dictDetailDialog.show(row)
        },
    },
}
</script>
<template>
    <q-page padding>

        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.projectName" label="项目名称" />
            <q-btn color="primary" @click="handleSearch" label="搜索" />
            <q-btn color="primary" @click="resetSearch" label="重置" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" label="新增项目" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-leader="props">
                <q-td :props="props">
                    <GqaShowName v-if="props.row.leader" :customNameObject="props.row.leader" />
                </q-td>
            </template>

            <template v-slot:body-cell-language="props">
                <q-td :props="props">
                    <GqaDictShow dictName="codeLanguage" :dictCode="props.row.language"
                        v-if="props.row.language !== ''" />
                </q-td>
            </template>

            <template v-slot:body-cell-node="props">
                <q-td :props="props">
                    <GqaDictShow dictName="projectNode" :dictCode="props.row.node" v-if="props.row.node !== ''" />
                </q-td>
            </template>

            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusOnOff" :dictCode="props.row.status" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="showEditForm(props.row)" label="编辑" />
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
import GqaDictShow from 'src/components/GqaDictShow'
import GqaShowName from 'src/components/GqaShowName'
import addOrEditDialog from './modules/addOrEditDialog'

export default {
    name: 'News',
    mixins: [tableDataMixin],
    components: {
        GqaDictShow,
        GqaShowName,
        addOrEditDialog,
    },
    data() {
        return {
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 10,
            },
            url: {
                list: 'plugin-xk/project-list',
                delete: 'plugin-xk/project-delete',
            },
            columns: [
                { name: 'sort', align: 'center', label: '排序', field: 'sort' },
                { name: 'projectName', align: 'center', label: '项目名称', field: 'projectName' },
                { name: 'demand', align: 'center', label: '需求单位', field: 'demand' },
                { name: 'leader', align: 'center', label: '牵头人', field: 'leader' },
                { name: 'player', align: 'center', label: '参与人', field: 'player' },
                { name: 'language', align: 'center', label: '项目语言', field: 'language' },
                { name: 'node', align: 'center', label: '项目节点', field: 'node' },
                { name: 'status', align: 'center', label: '状态', field: 'status' },
                { name: 'actions', align: 'center', label: '操作', field: 'actions' },
            ],
        }
    },
    created() {
        this.getTableData()
    },
}
</script>
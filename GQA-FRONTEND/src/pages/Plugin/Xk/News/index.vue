<template>
    <q-page padding>

        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.title" label="标题" />
            <q-btn color="primary" @click="handleSearch" label="搜索" />
            <q-btn color="primary" @click="resetSearch" label="重置" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" label="新增最新要闻" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusOnOff" :dictCode="props.row.status" />
                </q-td>
            </template>

            <template v-slot:body-cell-createdBy="props">
                <q-td :props="props">
                    <GqaShowName v-if="props.row.createdByUser" :customNameObject="props.row.createdByUser" />
                </q-td>
            </template>

            <template v-slot:body-cell-createdAt="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.createdAt) }}
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
import { FormatDateTime } from 'src/utils/date'

export default {
    name: 'News',
    mixins: [tableDataMixin],
    components: {
        GqaDictShow,
        GqaShowName,
        addOrEditDialog,
    },
    computed: {
        showDateTime() {
            return (datetime) => {
                return FormatDateTime(datetime)
            }
        },
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
                list: 'plugin-xk/news-list',
                delete: 'plugin-xk/news-delete',
            },
            columns: [
                { name: 'sort', align: 'center', label: '排序', field: 'sort' },
                { name: 'title', align: 'center', label: '新闻标题', field: 'title' },
                { name: 'createdBy', align: 'center', label: '发布人', field: 'createdBy' },
                { name: 'createdAt', align: 'center', label: '发布时间', field: 'createdAt' },
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
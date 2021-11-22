<template>
    <q-page padding>
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusOnOff" :dictCode="props.row.status" />
                </q-td>
            </template>

            <template v-slot:body-cell-stable="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusYesNo" :dictCode="props.row.stable" />
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
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import GqaDictShow from 'src/components/GqaDictShow'

export default {
    name: 'Example',
    mixins: [tableDataMixin],
    components: {
        GqaDictShow,
    },
    data() {
        return {
            url: {
                list: 'plugin-example/news-list',
            },
            columns: [
                { name: 'sort', align: 'center', label: '排序', field: 'sort' },
                { name: 'title', align: 'center', label: '新闻标题', field: 'title' },
                { name: 'content', align: 'center', label: '新闻内容', field: 'content' },
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
        showEditForm() {
            alert('编辑！')
        },
        handleDelete() {
            alert('删除！')
        },
    },
}
</script>
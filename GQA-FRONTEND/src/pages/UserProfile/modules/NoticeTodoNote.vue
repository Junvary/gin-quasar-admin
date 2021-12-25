<template>
    <div>
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
            <template v-slot:top>
                <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('TodoNote')" />
            </template>
            <template v-slot:body-cell-todoStatus="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusYesNo" :dictCode="props.row.todoStatus" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="showEditForm(props.row)" :label="$t('Edit')" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <NoticeTodoNoteDetail ref="addOrEditDialog" @handleFinish="handleFinish" />
    </div>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import GqaDictShow from 'src/components/GqaDictShow'
import NoticeTodoNoteDetail from './NoticeTodoNoteDetail'

export default {
    name: 'NoticeTodoNote',
    mixins: [tableDataMixin],
    components: {
        GqaDictShow,
        NoticeTodoNoteDetail,
    },
    computed: {
        columns() {
            return [
                { name: 'id', align: 'center', label: 'ID', field: 'id' },
                { name: 'todoDetail', align: 'center', label: this.$t('Detail'), field: 'todoDetail' },
                { name: 'todoStatus', align: 'center', label: this.$t('Done'), field: 'todoStatus' },
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ]
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
                list: 'todo-note/todo-note-list',
            },
        }
    },
    mounted() {
        this.getTableData()
    },
}
</script>
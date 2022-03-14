<template>
    <q-page padding>
        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.dictCode" :label="$t('Dict') + $t('Code')" />
            <q-input style="width: 20%" v-model="queryParams.dictLabel" :label="$t('Dict')+ $t('Name')" />
            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Dict')" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

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
                        <q-btn color="primary" @click="showEditForm(props.row)" :label="$t('Edit')" />
                        <q-btn color="warning" @click="handleDetail(props.row)" :label="$t('Detail')" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <add-or-edit-dialog ref="addOrEditDialog" @handleFinish="handleFinish" />
        <dict-detail-dialog ref="dictDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import addOrEditDialog from './modules/addOrEditDialog'
import dictDetailDialog from './modules/dictDetailDialog'
import GqaDictShow from 'src/components/GqaDictShow'

export default {
    name: 'Dict',
    mixins: [tableDataMixin],
    components: {
        addOrEditDialog,
        dictDetailDialog,
        GqaDictShow,
    },
    computed: {
        columns() {
            return [
                { name: 'sort', align: 'center', label: this.$t('Sort'), field: 'sort' },
                { name: 'dictCode', align: 'center', label: this.$t('Dict') + this.$t('Code'), field: 'dictCode' },
                { name: 'dictLabel', align: 'center', label: this.$t('Dict') + this.$t('Name'), field: 'dictLabel' },
                { name: 'dictExt1', align: 'center', label: 'Ext1', field: 'dictExt1' },
                { name: 'dictExt2', align: 'center', label: 'Ext2', field: 'dictExt2' },
                { name: 'status', align: 'center', label: this.$t('Status'), field: 'status' },
                { name: 'stable', align: 'center', label: this.$t('Stable'), field: 'stable' },
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ]
        },
    },
    data() {
        return {
            url: {
                list: 'dict/dict-list',
                delete: 'dict/dict-delete',
            },
        }
    },
    created() {
        this.getTableData()
    },
    methods: {
        handleDetail(row) {
            this.$refs.dictDetailDialog.show(row)
        },
    },
}
</script>

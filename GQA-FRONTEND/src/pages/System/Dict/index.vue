<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.value" :label="$t('PageSystemDictFilterValue')" />
            <q-input style="width: 20%" v-model="queryParams.label" :label="$t('PageSystemDictFilterLabel')" />
            <q-btn color="primary" @click="handleSearch" :label="$t('PageSystemDictFilterBtnSearch')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('PageSystemDictFilterBtnResetSearch')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" :label="$t('PageSystemDictBtnAdd')" />
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
                        <q-btn color="primary" @click="showEditForm(props.row)" :label="$t('PageSystemDictBtnEdit')" />
                        <q-btn color="warning" @click="handleDetail(props.row)" :label="$t('PageSystemDictBtnDetails')" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('PageSystemDictBtnDelete')" />
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
    data() {
        return {
            url: {
                list: 'dict/dict-list',
                delete: 'dict/dict-delete',
            },
            columns: [
                { name: 'sort', align: 'center', label: this.$t('PageSystemDictTableColumnSort'), field: 'sort' },
                { name: 'value', align: 'center', label: this.$t('PageSystemDictTableColumnValue'), field: 'value' },
                { name: 'label', align: 'center', label: this.$t('PageSystemDictTableColumnLabel'), field: 'label' },
                { name: 'status', align: 'center', label: this.$t('PageSystemDictTableColumnStatus'), field: 'status' },
                { name: 'stable', align: 'center', label: this.$t('PageSystemDictTableColumnStable'), field: 'stable' },
                { name: 'actions', align: 'center', label: this.$t('PageSystemDictTableColumnActions'), field: 'actions' },
            ],
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

<template>
    <q-dialog v-model="dictDetailVisible" position="right">
        <q-card style="width: 800px; max-width: 50vw; height: 100%">
            <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

                <template v-slot:top="props">
                    <q-btn dense color="primary" @click="showAddForm({parentId: parentDict.id})">
                        {{ $t('PageSystemDictDetailDialogTitle', { label: parentDict.label }) }}
                    </q-btn>
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
                            <q-btn dense color="primary" @click="showEditForm(props.row)" :label="$t('PageSystemDictDetailDialogBtnEdit')" />
                            <q-btn dense color="negative" @click="handleDelete(props.row)" :label="$t('PageSystemDictDetailDialogBtnDelete')" />
                        </div>
                    </q-td>
                </template>
            </q-table>
            <q-inner-loading :showing="loading">
                <q-spinner-gears size="50px" color="primary" />
            </q-inner-loading>
        </q-card>

        <add-or-edit-dialog ref="addOrEditDialog" @handleFinish="handleFinish" />
    </q-dialog>
</template>

<script>
import addOrEditDialog from './addOrEditDialog'
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { postAction } from 'src/api/manage'
import GqaDictShow from 'src/components/GqaDictShow'

export default {
    name: 'dictDetailDialog',
    mixins: [tableDataMixin],
    components: {
        addOrEditDialog,
        GqaDictShow,
    },
    data() {
        return {
            dictDetailVisible: false,
            loading: false,
            parentDict: {},
            url: {
                list: 'dict/dict-list',
                add: 'dict/dict-add',
                edit: 'dict/dict-edit',
                delete: 'dict/dict-delete',
                queryById: 'dict/dict-id',
            },
            columns: [
                { name: 'sort', align: 'center', label: this.$t('PageSystemDictDetailDialogTableColumnSort'), field: 'sort' },
                { name: 'value', align: 'center', label: this.$t('PageSystemDictDetailDialogTableColumnValue'), field: 'value' },
                { name: 'label', align: 'center', label: this.$t('PageSystemDictDetailDialogTableColumnLabel'), field: 'label' },
                { name: 'status', align: 'center', label: this.$t('PageSystemDictDetailDialogTableColumnStatus'), field: 'status' },
                { name: 'stable', align: 'center', label: this.$t('PageSystemDictDetailDialogTableColumnStable'), field: 'stable' },
                { name: 'actions', align: 'center', label: this.$t('PageSystemDictDetailDialogTableColumnActions'), field: 'actions' },
            ],
        }
    },
    methods: {
        show(row) {
            this.dictDetailVisible = true
            this.parentDict = row
            this.queryParams.parentId = this.parentDict.id
            this.getTableData()
        },
    },
}
</script>

<template>
    <q-dialog v-model="dictDetailVisible" position="right">
        <q-card style="width: 800px; max-width: 50vw; height: 100%">
            <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

                <template v-slot:top="props">
                    <q-btn dense color="primary" @click="showAddForm({parentCode: parentDict.dictCode})">
                        {{ $t('Add') }} {{ parentDict.dictLabel}} {{ $t('Item') }}:
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
                            <q-btn dense color="primary" @click="showEditForm(props.row)" :label="$t('Edit')" />
                            <q-btn dense color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
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
                { name: 'sort', align: 'center', label: this.$t('Sort'), field: 'sort' },
                { name: 'dictCode', align: 'center', label: this.$t('Dict') + this.$t('Code'), field: 'dictCode' },
                { name: 'dictLabel', align: 'center', label: this.$t('Dict') + this.$t('Name'), field: 'dictLabel' },
                { name: 'dictExt1', align: 'center', label: 'Ext1', field: 'dictExt1' },
                { name: 'dictExt2', align: 'center', label: 'Ext2', field: 'dictExt2' },
                { name: 'status', align: 'center', label: this.$t('Status'), field: 'status' },
                { name: 'stable', align: 'center', label: this.$t('Stable'), field: 'stable' },
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ],
        }
    },
    methods: {
        show(row) {
            this.dictDetailVisible = true
            this.parentDict = row
            this.queryParams.parentCode = this.parentDict.dictCode
            this.getTableData()
        },
    },
}
</script>

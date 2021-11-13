<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.gqaOption" :label="$t('PageSystemConfigFilterOption')" />
            <q-input style="width: 20%" v-model="queryParams.remark" :label="$t('PageSystemConfigFilterRemark')" />
            <q-btn color="primary" @click="handleSearch" :label="$t('PageSystemConfigTableBtnSearch')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('PageSystemConfigTableBtnResetSearch')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" :label="$t('PageSystemConfigBtnAdd')" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-custom="props">
                <q-td :props="props">
                    {{props.row.custom}}
                    <q-popup-edit v-model="props.row.custom" :title="$t('PageSystemConfigOption') + props.row.gqaOption">
                        <q-input v-model="props.row.custom" dense autofocus />
                    </q-popup-edit>
                </q-td>
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
                        <q-btn dense color="primary" @click="handleSave(props.row)" :label="$t('PageSystemConfigBtnSave')" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('PageSystemConfigBtnDelete')" />
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
    name: 'ConfigFrontend',
    mixins: [tableDataMixin],
    components: {
        addOrEditDialog,
        GqaDictShow,
    },
    data() {
        return {
            url: {
                list: 'config-frontend/config-frontend-list',
                edit: 'config-frontend/config-frontend-edit',
                delete: 'config-frontend/config-frontend-delete',
            },
            columns: [
                { name: 'sort', align: 'center', label: this.$t('PageSystemConfigTableColumnSort'), field: 'sort' },
                { name: 'gqaOption', align: 'center', label: this.$t('PageSystemConfigTableColumnOption'), field: 'gqaOption' },
                { name: 'remark', align: 'center', label: this.$t('PageSystemConfigTableColumnRemark'), field: 'remark' },
                { name: 'default', align: 'center', label: this.$t('PageSystemConfigTableColumnDefault'), field: 'default' },
                { name: 'custom', align: 'center', label: this.$t('PageSystemConfigTableColumnCustom'), field: 'custom' },
                { name: 'status', align: 'center', label: this.$t('PageSystemConfigTableColumnStatus'), field: 'status' },
                { name: 'stable', align: 'center', label: this.$t('PageSystemConfigTableColumnStable'), field: 'stable' },
                { name: 'actions', align: 'center', label: this.$t('PageSystemConfigTableColumnActions'), field: 'actions' },
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

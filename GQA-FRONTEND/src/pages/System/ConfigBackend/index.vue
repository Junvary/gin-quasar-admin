<template>
    <q-page padding>

        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.gqaOption" :label="$t('Config') + $t('Name')" />
            <q-input style="width: 20%" v-model="queryParams.remark" :label="$t('Config') + $t('Remark')" />
            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Config')" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-custom="props">
                <q-td :props="props" class="bg-green-1">
                    {{props.row.custom}}
                    <q-popup-edit v-model="props.row.custom" class="bg-green-13">
                        <template v-slot="scope">
                            {{ $t('Custom') + ' ' + props.row.gqaOption }}
                            <q-input v-model="props.row.custom" dense autofocus clearable @keyup.enter="scope.set" />
                        </template>
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
                        <q-btn color="primary" @click="handleSave(props.row)" :label="$t('Save')" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <add-or-edit-dialog ref="addOrEditDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { mapActions } from 'vuex'
import addOrEditDialog from './modules/addOrEditDialog'
import { putAction } from 'src/api/manage'
import GqaDictShow from 'src/components/GqaDictShow'

export default {
    name: 'ConfigBackend',
    mixins: [tableDataMixin],
    components: {
        addOrEditDialog,
        GqaDictShow,
    },
    computed: {
        columns() {
            return [
                { name: 'sort', align: 'center', label: this.$t('Sort'), field: 'sort' },
                { name: 'gqaOption', align: 'center', label: this.$t('Option'), field: 'gqaOption' },
                { name: 'remark', align: 'center', label: this.$t('Remark'), field: 'remark' },
                { name: 'default', align: 'center', label: this.$t('Default'), field: 'default' },
                { name: 'custom', align: 'center', label: this.$t('Custom'), field: 'custom' },
                { name: 'status', align: 'center', label: this.$t('Status'), field: 'status' },
                { name: 'stable', align: 'center', label: this.$t('Stable'), field: 'stable' },
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ]
        },
    },
    data() {
        return {
            url: {
                list: 'config-backend/config-backend-list',
                edit: 'config-backend/config-backend-edit',
                delete: 'config-backend/config-backend-delete',
            },
        }
    },
    created() {
        this.getTableData()
    },
    methods: {
        ...mapActions('storage', ['GetGqaBackend']),
        async handleSave(row) {
            const res = await putAction(this.url.edit, row)
            if (res.code === 1) {
                this.$q.notify({
                    type: 'positive',
                    message: res.message,
                })
                this.GetGqaBackend()
            }
        },
    },
}
</script>

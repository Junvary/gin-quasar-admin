<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.roleCode" :label="$t('Role') + $t('Code')" />
            <q-input style="width: 20%" v-model="queryParams.roleName" :label="$t('Role') + $t('Name')" />

            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Role')" />
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
                        <q-btn color="primary" @click="showEditForm(props.row)"
                            :label="$t('Edit') + ' ' + $t('Role')" />
                        <q-btn color="warning" @click="showRoleUser(props.row)" :label="$t('User')" />
                        <q-btn color="warning" @click="showRolePermission(props.row)" :label="$t('Permission')" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <add-or-edit-dialog ref="addOrEditDialog" @handleFinish="handleFinish" />
        <role-permission-dialog ref="rolePermissionDialog" />
        <role-user-dialog ref="roleUserDialog" />
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import addOrEditDialog from './modules/addOrEditDialog'
import RolePermissionDialog from './modules/RolePermissionDialog'
import RoleUserDialog from './modules/RoleUserDialog'
import GqaDictShow from 'src/components/GqaDictShow'

export default {
    name: 'Role',
    mixins: [tableDataMixin],
    components: {
        addOrEditDialog,
        RolePermissionDialog,
        RoleUserDialog,
        GqaDictShow,
    },
    created() {
        this.getTableData()
    },
    data() {
        return {
            url: {
                list: 'role/role-list',
                delete: 'role/role-delete',
            },
            columns: [
                { name: 'sort', align: 'center', label: this.$t('Sort'), field: 'sort' },
                { name: 'roleCode', align: 'center', label: this.$t('Role') + this.$t('Code'), field: 'roleCode' },
                { name: 'roleName', align: 'center', label: this.$t('Role') + this.$t('Name'), field: 'roleName' },
                { name: 'status', align: 'center', label: this.$t('Status'), field: 'status' },
                { name: 'stable', align: 'center', label: this.$t('Stable'), field: 'stable' },
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ],
        }
    },

    methods: {
        showRolePermission(row) {
            this.$refs.rolePermissionDialog.show(row)
        },
        showRoleUser(row) {
            this.$refs.roleUserDialog.show(row)
        },
    },
}
</script>

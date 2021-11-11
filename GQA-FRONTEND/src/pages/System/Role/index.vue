<template>
    <q-page padding>

        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.roleCode" label="角色编码" />
            <q-input style="width: 20%" v-model="queryParams.roleName" label="角色名称" />

            <q-btn color="primary" @click="handleSearch" label="搜索" />
            <q-btn color="primary" @click="resetSearch" label="重置" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" label="新增角色" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <gqa-dict-show dictName="statusOnOff" :dictCode="props.row.status" />
                </q-td>
            </template>

            <template v-slot:body-cell-stable="props">
                <q-td :props="props">
                    <gqa-dict-show dictName="statusYesNo" :dictCode="props.row.stable" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="showEditForm(props.row)" label="编辑" />
                        <q-btn color="warning" @click="showRoleUser(props.row)" label="用户" />
                        <q-btn color="warning" @click="showRolePermission(props.row)" label="权限" />
                        <q-btn color="negative" @click="handleDelete(props.row)" label="删除" />
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
                { name: 'sort', align: 'center', label: '排序', field: 'sort' },
                { name: 'roleCode', align: 'center', label: '角色编码', field: 'roleCode' },
                { name: 'roleName', align: 'center', label: '角色名称', field: 'roleName' },
                { name: 'status', align: 'center', label: '状态', field: 'status' },
                { name: 'stable', align: 'center', label: '系统内置', field: 'stable' },
                { name: 'actions', align: 'center', label: '操作', field: 'actions' },
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
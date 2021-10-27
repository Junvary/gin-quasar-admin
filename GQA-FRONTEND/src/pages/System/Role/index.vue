<template>
    <q-page padding>

        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="searchParams.roleCode" label="角色编码" />
            <q-input style="width: 20%" v-model="searchParams.roleName" label="角色名" />

            <q-btn color="primary" @click="handleSearch" label="搜索" />
            <q-btn color="primary" @click="resetSearch" label="重置" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" label="新增角色" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <gqa-status :status="props.row.status" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="showEditForm(props.row)" label="编辑"
                            v-if="props.row.roleCode !== 'super-admin'" />
                        <q-btn color="warning" @click="showRolePermission(props.row)" label="用户" />
                        <q-btn color="warning" @click="showRolePermission(props.row)" label="权限"
                            v-if="props.row.roleCode !== 'super-admin'" />
                        <q-btn color="negative" @click="handleDelete(props.row)" label="删除"
                            v-if="props.row.roleCode !== 'super-admin'" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <add-or-edit-dialog ref="addOrEditDialog" @emitAddOrEdit="emitAddOrEdit" @handleFinish="handleFinish" />
        <role-permission-dialog ref="rolePermissionDialog" />
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import addOrEditDialog from './modules/addOrEditDialog'
import GqaStatus from 'src/components/GqaStatus'
import RolePermissionDialog from './modules/RolePermissionDialog'

export default {
    name: 'Role',
    mixins: [tableDataMixin],
    components: {
        addOrEditDialog,
        GqaStatus,
        RolePermissionDialog,
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
                { name: 'roleName', align: 'center', label: '角色名', field: 'roleName' },
                { name: 'status', align: 'center', label: '状态', field: 'status' },
                { name: 'actions', align: 'center', label: '操作', field: 'actions' },
            ],
        }
    },

    methods: {
        showRolePermission(row) {
            this.$refs.rolePermissionDialog.show(row)
        },
    },
}
</script>
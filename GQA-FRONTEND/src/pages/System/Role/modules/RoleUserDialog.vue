<template>
    <q-dialog v-model="roleUserVisible" v-if="roleUserVisible" position="right">
        <q-card style="min-width: 500px; max-width: 45vw">
            <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                :loading="loading" @request="onRequest">

                <template v-slot:top="props">
                    <q-btn dense color="primary" @click="showAddUserForm()" label="添加已有用户" />
                    <q-space />
                    <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                        @click="props.toggleFullscreen" class="q-ml-md" />
                </template>

                <template v-slot:body-cell-actions="props">
                    <q-td :props="props">
                        <div class="q-gutter-xs">
                            <q-btn dense color="negative" @click="handleRemove(props.row)" label="移除"
                                v-if="props.row.id !== 1 && record.roleCode !== 'super-admin'" />
                        </div>
                    </q-td>
                </template>
            </q-table>
        </q-card>
    </q-dialog>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { postAction } from 'src/api/manage'

export default {
    name: 'RoleUserDialog',
    mixins: [tableDataMixin],
    data() {
        return {
            roleUserVisible: false,
            record: {},
            url: {
                list: 'role/role-user',
                removeUser: 'role/role-user-remove',
            },
            columns: [
                { name: 'sort', align: 'center', label: '排序', field: 'sort' },
                { name: 'username', align: 'center', label: '账号', field: 'username' },
                { name: 'nickname', align: 'center', label: '昵称', field: 'nickname' },
                { name: 'realName', align: 'center', label: '真实姓名', field: 'realName' },
                { name: 'actions', align: 'center', label: '操作', field: 'actions' },
            ],
        }
    },
    methods: {
        show(row) {
            this.record = row
            this.queryParams.roleCode = this.record.roleCode
            this.roleUserVisible = true
            this.getTableData()
        },
        showAddUserForm() {},
        handleRemove(row) {
            this.$q
                .dialog({
                    title: '确定移除？',
                    message: `你确定要把此用户移除该角色吗？`,
                    cancel: true,
                    persistent: true,
                })
                .onOk(async () => {
                    if (this.record.roleCode === 'super-admin' && row.id === 1) {
                        this.$q.notify({
                            type: 'negative',
                            message: '抱歉，你不能把超级管理员从超级管理员组中移除！',
                        })
                        return false
                    }
                    const res = await postAction(this.url.removeUser, {
                        roleCode: this.record.roleCode,
                        userId: row.id,
                    })
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                    }
                    this.getTableData()
                })
        },
    },
}
</script>
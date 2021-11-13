<template>
    <q-dialog v-model="roleUserVisible" position="right">
        <q-card style="min-width: 500px; max-width: 45vw">
            <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

                <template v-slot:top="props">
                    <q-btn dense color="primary" @click="showAddUserForm()" :label="$t('PageSystemRoleUserDialogAddUser')" />
                    <q-space />
                    <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                        @click="props.toggleFullscreen" class="q-ml-md" />
                </template>

                <template v-slot:body-cell-actions="props">
                    <q-td :props="props">
                        <div class="q-gutter-xs">
                            <q-btn dense color="negative" @click="handleRemove(props.row)" :label="$t('PageSystemRoleUserDialogDelete')" />
                        </div>
                    </q-td>
                </template>
            </q-table>
        </q-card>
        <SelectUserDialog ref="selectUserDialog" @handleSelectUser="handleAddUser" selection="multiple" />
    </q-dialog>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { postAction } from 'src/api/manage'
import SelectUserDialog from 'src/components/GqaSeleteUser/SelectUserDialog'

export default {
    name: 'RoleUserDialog',
    mixins: [tableDataMixin],
    components: {
        SelectUserDialog,
    },
    data() {
        return {
            roleUserVisible: false,
            record: {},
            url: {
                list: 'role/role-user',
                removeUser: 'role/role-user-remove',
                addUser: 'role/role-user-add',
            },
            columns: [
                { name: 'sort', align: 'center', label: $t('PageSystemRoleUserDialogTableColumnSort'), field: 'sort' },
                { name: 'username', align: 'center', label: $t('PageSystemRoleUserDialogUsername'), field: 'username' },
                { name: 'nickname', align: 'center', label: $t('PageSystemRoleUserDialogNickname'), field: 'nickname' },
                { name: 'realName', align: 'center', label: $t('PageSystemRoleUserDialogRealName'), field: 'realName' },
                { name: 'actions', align: 'center', label: $t('PageSystemRoleUserDialogActions'), field: 'actions' },
            ],
        }
    },
    methods: {
        show(row) {
            this.tableData = []
            this.record = row
            this.queryParams.roleCode = this.record.roleCode
            this.roleUserVisible = true
            this.getTableData()
        },
        showAddUserForm() {
            this.$refs.selectUserDialog.show(this.tableData)
        },
        handleRemove(row) {
            this.$q
                .dialog({
                    title: $t('PageSystemRoleUserDialogDeleteQuestion'),
                    message: $t('PageSystemRoleUserDialogDeleteMessage'),
                    cancel: true,
                    persistent: true,
                })
                .onOk(async () => {
                    if (this.record.roleCode === 'super-admin' && row.id === 1) {
                        this.$q.notify({
                            type: 'negative',
                            message: $t('PageSystemRoleUserDialogDeleteNotAllowed'),
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
        handleAddUser(event) {
            const idList = []
            for (let i of event) {
                idList.push(i.id)
            }
            postAction(this.url.addUser, {
                roleCode: this.record.roleCode,
                userId: idList,
            }).then((res) => {
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

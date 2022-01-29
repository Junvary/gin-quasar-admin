<template>
    <q-dialog v-model="roleUserVisible" position="right">
        <q-card style="min-width: 500px; max-width: 45vw">
            <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

                <template v-slot:top="props">
                    <q-btn dense color="primary" @click="showAddUserForm()" :label="$t('Add') + $t('User')" />
                    <q-space />
                    <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                        @click="props.toggleFullscreen" class="q-ml-md" />
                </template>

                <template v-slot:body-cell-actions="props">
                    <q-td :props="props">
                        <div class="q-gutter-xs">
                            <q-btn dense color="negative" @click="handleRemove(props.row)" :label="$t('Delete')" />
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
                // { name: 'sort', align: 'center', label: this.$t('Sort'), field: 'sort' },
                { name: 'username', align: 'center', label: this.$t('Username'), field: 'username' },
                { name: 'nickname', align: 'center', label: this.$t('Nickname'), field: 'nickname' },
                { name: 'realName', align: 'center', label: this.$t('RealName'), field: 'realName' },
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ],
            pagination: {
                sortBy: 'username',
                descending: false,
                page: 1,
                rowsPerPage: 10,
            },
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
                    title: this.$t('ConfirmDelete'),
                    message: this.$t('ConfirmDeleteMessage'),
                    cancel: true,
                    persistent: true,
                })
                .onOk(async () => {
                    if (this.record.roleCode === 'super-admin' && row.username === 'admin') {
                        this.$q.notify({
                            type: 'negative',
                            message: this.$t('CanNotDeleteThis'),
                        })
                        return false
                    }
                    const res = await postAction(this.url.removeUser, {
                        roleCode: this.record.roleCode,
                        username: row.username,
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
            const usernameList = []
            for (let i of event) {
                usernameList.push(i.username)
            }
            postAction(this.url.addUser, {
                roleCode: this.record.roleCode,
                username: usernameList,
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

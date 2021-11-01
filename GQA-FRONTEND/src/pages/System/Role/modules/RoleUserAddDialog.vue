<template>
    <q-dialog v-model="roleUserAddVisible">
        <q-card style="min-width: 500px; max-width: 45vw">
            <q-card-section>
                <div class="text-h6">选择用户</div>
            </q-card-section>
            <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                :rows-per-page-options="pagination.options" :loading="loading" @request="onRequest" selection="multiple"
                v-model:selected="selected">
                <template v-slot:top="">
                    <q-btn dense color="primary" @click="handleAddUser()" label="确定添加" />
                </template>
            </q-table>
        </q-card>
    </q-dialog>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'

export default {
    name: 'RoleUserAddDialog',
    mixins: [tableDataMixin],
    data() {
        return {
            roleUserAddVisible: false,
            url: {
                list: 'user/user-list',
            },
            columns: [
                { name: 'username', align: 'center', label: '账号', field: 'username' },
                { name: 'nickname', align: 'center', label: '昵称', field: 'nickname' },
                { name: 'realName', align: 'center', label: '真实姓名', field: 'realName' },
            ],
            selected: [],
        }
    },
    methods: {
        show() {
            this.roleUserAddVisible = true
            this.getTableData()
        },
        handleAddUser() {
            this.$emit('handleAddUser', this.selected)
            this.roleUserAddVisible = false
        },
    },
}
</script>
            
       
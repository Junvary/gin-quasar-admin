<template>
    <q-dialog v-model="selectUserVisible">
        <q-card style="min-width: 700px; max-width: 45vw">
            <q-card-section class="row justify-between items-center">
                <div class="text-subtitle1">
                    用户选择器
                    （{{selection === "multiple" ? "多选" : "单选" }}）
                </div>
                <q-btn dense color="primary" @click="handleSelectUser()" label="确定添加" />
            </q-card-section>

            <q-separator />

            <q-card-section class="row q-gutter-md items-center">
                <q-input dense style="width: 20%" v-model="queryParams.username" label="账号" />
                <q-input dense style="width: 20%" v-model="queryParams.realName" label="真实姓名" />
                <q-btn dense color="primary" @click="handleSearch" label="搜索" />
                <q-btn dense color="primary" @click="resetSearch" label="重置" />
            </q-card-section>

            <q-table row-key="id" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest" :selection="selection"
                v-model:selected="selected">
            </q-table>
        </q-card>
    </q-dialog>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'

export default {
    name: 'SelectUserDialog',
    mixins: [tableDataMixin],
    props: {
        // 必须传递单选多选： multiple, single
        selection: {
            type: String,
            required: true,
        },
    },
    data() {
        return {
            selectUserVisible: false,
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
        show(selectUser) {
            this.selected = []
            this.selectUserVisible = true
            this.getTableData()
            if (this.selection === 'multiple') {
                if (typeof selectUser === Array) {
                    this.selected = selectUser
                } else {
                    this.selected = []
                }
            } else {
                this.selected.push(selectUser)
            }
        },
        handleSelectUser() {
            this.$emit('handleSelectUser', this.selected)
            this.selectUserVisible = false
        },
    },
}
</script>
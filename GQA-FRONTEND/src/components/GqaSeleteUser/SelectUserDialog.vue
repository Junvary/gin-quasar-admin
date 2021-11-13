<template>
    <q-dialog v-model="selectUserVisible">
        <q-card style="min-width: 700px; max-width: 45vw">
            <q-card-section class="items-center justify-between row">
                <div class="text-h7">
                    {{ $t('ComponentSelectUserTitle', {oneOrMultiple: selection === "multiple" ? $t('ComponentSelectUserTitleMultiple') : $t('ComponentSelectUserTitleOne')}) }}
                </div>
                <q-btn dense color="primary" @click="handleSelectUser()" :label="$t('ComponentSelectUserBtnSelectUser')" />
            </q-card-section>

            <q-separator />

            <q-card-section class="items-center row q-gutter-md">
                <q-input dense style="width: 20%" v-model="queryParams.username" :label="$t('ComponentSelectUserFilterUsername')" />
                <q-input dense style="width: 20%" v-model="queryParams.realName" :label="$t('ComponentSelectUserFilterRealname')" />
                <q-btn dense color="primary" @click="handleSearch" :label="$t('ComponentSelectUserBtnSearch')" />
                <q-btn dense color="primary" @click="resetSearch" :label="$t('ComponentSelectUserBtnResetSearch')" />
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
                { name: 'username', align: 'center', label: this.$parent.$t('ComponentSelectUserTableColumnUsername'), field: 'username' },
                { name: 'nickname', align: 'center', label: this.$parent.$t('ComponentSelectUserTableColumnNickname'), field: 'nickname' },
                { name: 'realName', align: 'center', label: this.$parent.$t('ComponentSelectUserTableColumnRealname'), field: 'realName' },
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
                this.selected = selectUser
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

<template>
    <q-dialog v-model="selectUserVisible">
        <q-card style="min-width: 700px; max-width: 45vw">
            <q-card-section class="row justify-between items-center">
                <div class="text-subtitle1">
                    {{ $t('ComponentSelectUserTitle', {oneOrMultiple: selection === "multiple" ? $t('SelectMultiple') : $t('SelectOne')}) }}
                </div>
                <q-btn dense color="primary" @click="handleSelectUser()" :label="$t('Select')" />
            </q-card-section>

            <q-separator />

            <q-card-section class="items-center row q-gutter-md">
                <q-input dense style="width: 20%" v-model="queryParams.username" :label="$t('Username')" />
                <q-input dense style="width: 20%" v-model="queryParams.realName" :label="$t('RealName')" />
                <q-btn dense color="primary" @click="handleSearch" :label="$t('Search')" />
                <q-btn dense color="primary" @click="resetSearch" :label="$t('Reset')" />
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
                { name: 'username', align: 'center', label: this.$parent.$t('Username'), field: 'username' },
                { name: 'nickname', align: 'center', label: this.$parent.$t('Nickname'), field: 'nickname' },
                { name: 'realName', align: 'center', label: this.$parent.$t('RealName'), field: 'realName' },
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

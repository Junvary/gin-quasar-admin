<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.username" :label="$t('Username')" />
            <q-input style="width: 20%" v-model="queryParams.realName" :label="$t('RealName')" />
            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('User')" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-avatar="props">
                <q-td :props="props">
                    <GqaAvatar :src="props.row.avatar" />
                </q-td>
            </template>

            <template v-slot:body-cell-gender="props">
                <q-td :props="props">
                    <GqaDictShow dictName="gender" :dictCode="props.row.gender" />
                </q-td>
            </template>

            <template v-slot:body-cell-role="props">
                <q-td :props="props">
                    <div class="column items-center q-gutter-xs">
                        <q-badge class="col" color="primary" v-for="(item, index) in props.row.role" :key="index">
                            {{ item.roleName }}
                        </q-badge>
                    </div>
                </q-td>
            </template>

            <template v-slot:body-cell-dept="props">
                <q-td :props="props">
                    <div class="column items-center q-gutter-xs">
                        <q-badge class="col" color="primary" v-for="(item, index) in props.row.dept" :key="index">
                            {{ item.deptName }}
                        </q-badge>
                    </div>
                </q-td>
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
                        <q-btn color="warning" @click="resetPassword(props.row)"
                            :label="$t('Reset') + $t('Password')" />
                        <q-btn color="primary" @click="showEditForm(props.row)" :label="$t('Edit')" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <add-or-edit-dialog ref="addOrEditDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import addOrEditDialog from './modules/addOrEditDialog'
import GqaAvatar from 'src/components/GqaAvatar'
import GqaDictShow from 'src/components/GqaDictShow'
import { postAction } from 'src/api/manage'

export default {
    name: 'User',
    mixins: [tableDataMixin],
    components: {
        addOrEditDialog,
        GqaAvatar,
        GqaDictShow,
    },
    computed: {
        columns() {
            return [
                // { name: 'sort', align: 'center', label: this.$t('Sort'), field: 'sort' },
                { name: 'avatar', align: 'center', label: this.$t('Avatar'), field: 'avatar' },
                { name: 'username', align: 'center', label: this.$t('Username'), field: 'username' },
                { name: 'nickname', align: 'center', label: this.$t('Nickname'), field: 'nickname' },
                { name: 'realName', align: 'center', label: this.$t('RealName'), field: 'realName' },
                { name: 'gender', align: 'center', label: this.$t('Gender'), field: 'gender' },
                { name: 'role', align: 'center', label: this.$t('Role'), field: 'role' },
                // { name: 'mobile', align: 'center', label: this.$t('Mobile'), field: 'mobile' },
                // { name: 'email', align: 'center', label: '邮箱', field: 'email' },
                { name: 'dept', align: 'center', label: this.$t('Dept'), field: 'dept' },
                { name: 'status', align: 'center', label: this.$t('Status'), field: 'status' },
                { name: 'stable', align: 'center', label: this.$t('Stable'), field: 'stable' },
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ]
        },
    },
    data() {
        return {
            queryParams: {
                withAdmin: true,
            },
            url: {
                list: 'user/user-list',
                delete: 'user/user-delete',
                resetPassword: 'user/user-reset-password',
            },
            pagination: {
                sortBy: 'username',
                descending: false,
                page: 1,
                rowsPerPage: 10,
            },
        }
    },
    created() {
        this.getTableData()
    },
    methods: {
        resetPassword(row) {
            this.$q
                .dialog({
                    title: this.$t('Reset') + this.$t('Password'),
                    message: this.$t('Confirm') + this.$t('Reset') + this.$t('Password') + '?' + '(' + row.username + ')',
                    cancel: true,
                    persistent: true,
                })
                .onOk(async () => {
                    const res = await postAction(this.url.resetPassword, {
                        id: row.id,
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

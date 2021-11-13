<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.username" :label="$t('PageSystemUserFilterUserName')" />
            <q-input style="width: 20%" v-model="queryParams.realName" :label="$t('PageSystemUserFilterRealName')" />
            <q-btn color="primary" @click="handleSearch" :label="$t('PageSystemUserTableBtnSearch')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('PageSystemUserTableBtnResetSearch')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" :label="$t('PageSystemUserTableColumnActionsBtnAdd')" />
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

            <!-- <template v-slot:body-cell-dept="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        {{ getDeptName(props.row.dept) }}
                    </div>
                </q-td>
            </template>

            <template v-slot:body-cell-role="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-chip color="primary" text-color="white" v-for="(item, index) in props.row.role" :key="index">
                            {{ getRoleName(item) }}
                        </q-chip>

                    </div>
                </q-td>
            </template> -->

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="showEditForm(props.row)" :label="$t('PageSystemUserTableColumnActionsBtnEdit')" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('PageSystemUserTableColumnActionsBtnDelete')" />
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

export default {
    name: 'User',
    mixins: [tableDataMixin],
    components: {
        addOrEditDialog,
        GqaAvatar,
        GqaDictShow,
    },
    data() {
        return {
            url: {
                list: 'user/user-list',
                delete: 'user/user-delete',
            },
            columns: [
                { name: 'sort', align: 'center', label: this.$t('PageSystemUserTableActions'), field: 'sort' },
                { name: 'avatar', align: 'center', label: this.$t('PageSystemUserTableActions'), field: 'avatar' },
                { name: 'username', align: 'center', label: this.$t('PageSystemUserTableActions'), field: 'username' },
                { name: 'nickname', align: 'center', label: this.$t('PageSystemUserTableActions'), field: 'nickname' },
                { name: 'realName', align: 'center', label: this.$t('PageSystemUserTableActions'), field: 'realName' },
                { name: 'gender', align: 'center', label: this.$t('PageSystemUserTableActions'), field: 'gender' },
                { name: 'mobile', align: 'center', label: this.$t('PageSystemUserTableActions'), field: 'mobile' },
                // { name: 'email', align: 'center', label: '邮箱', field: 'email' },
                { name: 'dept', align: 'center', label: this.$t('PageSystemUserTableActions'), field: 'dept' },
                { name: 'status', align: 'center', label: this.$t('PageSystemUserTableActions'), field: 'status' },
                { name: 'stable', align: 'center', label: this.$t('PageSystemUserTableActions'), field: 'stable' },
                { name: 'actions', align: 'center', label: this.$t('PageSystemUserTableActions'), field: 'actions' },
            ],
        }
    },
    created() {
        this.getTableData()
    },
}
</script>

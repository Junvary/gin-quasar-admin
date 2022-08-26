<template>
    <q-page padding>
        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.role_code" :label="$t('Role') + $t('Code')" />
            <q-input style="width: 20%" v-model="queryParams.role_name" :label="$t('Role') + $t('Name')" />
            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Role')" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>
            <!-- <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusOnOff" :dictCode="props.row.status" />
                </q-td>
            </template> -->
            <template v-slot:body-cell-stable="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusYesNo" :dictCode="props.row.stable" />
                </q-td>
            </template>
            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="showEditForm(props.row)" :label="$t('Edit')" />
                        <q-btn color="warning" @click="showRoleUser(props.row)" :label="$t('Role') + $t('User')" />
                        <q-btn color="warning" @click="showRolePermission(props.row)"
                            :label="$t('Role') + $t('Permission')" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
        <RolePermissionDialog ref="rolePermissionDialog" @handleFinish="handleFinish" />
        <RoleUserDialog ref="roleUserDialog" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import recordDetail from './modules/recordDetail.vue'
import RolePermissionDialog from './modules/RolePermissionDialog.vue'
import RoleUserDialog from './modules/RoleUserDialog.vue'

const { t } = useI18n()
const url = {
    list: 'role/get-role-list',
    delete: 'role/delete-role-by-id',
}
const columns = computed(() => {
    return [
        { name: 'sort', align: 'center', label: t('Sort'), field: 'sort' },
        { name: 'role_code', align: 'center', label: t('Role') + t('Code'), field: 'role_code' },
        { name: 'role_name', align: 'center', label: t('Role') + t('Name'), field: 'role_name' },
        { name: 'stable', align: 'center', label: t('Stable'), field: 'stable' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    onRequest,
    getTableData,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    getTableData()
})

const rolePermissionDialog = ref(null)
const showRolePermission = (row) => {
    rolePermissionDialog.value.show(row)
}
const roleUserDialog = ref(null)
const showRoleUser = (row) => {
    roleUserDialog.value.show(row)
}

</script>

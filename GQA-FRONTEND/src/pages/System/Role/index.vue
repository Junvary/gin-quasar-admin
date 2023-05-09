<template>
    <q-page padding>
        <q-card flat>
            <q-card-section class="row q-gutter-x-md items-center">
                <q-input outlined dense style="width: 20%" v-model="queryParams.role_code"
                    :label="$t('Role') + $t('Code')" />
                <q-input outlined dense style="width: 20%" v-model="queryParams.role_name"
                    :label="$t('Role') + $t('Name')" />
                <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
                <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
            </q-card-section>
            <q-card-section>
                <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                    :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
                    <template v-slot:top="props">
                        <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Role')"
                            v-has="'role:add'" />
                        <q-space />
                        <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                            @click="props.toggleFullscreen" class="q-ml-md" />
                    </template>
                    <template v-slot:body-cell-default_page="props">
                        <q-td :props="props">
                            {{ $t(props.row.default_page_menu.title) }}
                        </q-td>
                    </template>
                    <template v-slot:body-cell-stable="props">
                        <q-td :props="props">
                            <GqaDictShow :dictCode="props.row.stable" />
                        </q-td>
                    </template>
                    <template v-slot:body-cell-actions="props">
                        <q-td :props="props" class="q-gutter-x-xs">
                            <q-btn flat dense rounded icon="eva-edit-2-outline" color="primary"
                                @click="showEditForm(props.row)" v-has="'role:edit'">
                                <q-tooltip>
                                    {{ $t('Edit') }}
                                </q-tooltip>
                            </q-btn>
                            <q-btn flat dense rounded icon="add" color="warning" @click="showRolePermission(props.row)"
                                v-has="'role:permission'">
                                <q-tooltip>
                                    {{ $t('Role') + $t('Permission') }}
                                </q-tooltip>
                            </q-btn>
                            <q-btn flat dense rounded icon="eva-people-outline" color="positive"
                                @click="showRoleUser(props.row)" v-has="'role:user'">
                                <q-tooltip>
                                    {{ $t('Role') + $t('User') }}
                                </q-tooltip>
                            </q-btn>
                            <q-btn flat dense rounded icon="delete_outline" color="negative"
                                @click="handleDelete(props.row)" v-has="'role:delete'">
                                <q-tooltip>
                                    {{ $t('Delete') }}
                                </q-tooltip>
                            </q-btn>
                        </q-td>
                    </template>
                </q-table>
            </q-card-section>
        </q-card>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
        <RolePermissionDialog ref="rolePermissionDialog" @handleFinish="handleFinish" />
        <RoleUserDialog ref="roleUserDialog" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, onMounted, ref } from 'vue'
import recordDetail from './modules/recordDetail.vue'
import RolePermissionDialog from './modules/RolePermissionDialog.vue'
import RoleUserDialog from './modules/RoleUserDialog.vue'

const url = {
    list: 'role/get-role-list',
    delete: 'role/delete-role-by-id',
}
const columns = computed(() => {
    return [
        { name: 'sort', align: 'center', label: t('Sort'), field: 'sort' },
        { name: 'role_code', align: 'center', label: t('Role') + t('Code'), field: 'role_code' },
        { name: 'role_name', align: 'center', label: t('Role') + t('Name'), field: 'role_name' },
        { name: 'default_page', align: 'center', label: t('Default') + t('Page'), field: 'default_page' },
        { name: 'stable', align: 'center', label: t('Stable'), field: 'stable' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    t,
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

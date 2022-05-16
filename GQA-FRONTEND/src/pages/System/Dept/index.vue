<template>
    <q-page padding>
        <q-btn color="primary" @click="showAddParentForm()" :label="$t('Add') + $t('Parent') + $t('Dept')" />
        <q-hierarchy separator="cell" dense :columns="columns" :data="deptTree">
            <template v-slot:body="props">
                <gqa-tree-td :treeTd="props" firstTd="sort"></gqa-tree-td>
                <td class="text-center">{{ props.item.dept_name }}</td>
                <td class="text-center">{{ props.item.dept_code }}</td>
                <td class="text-center">{{ props.item.leader }}</td>
                <td class="text-center">
                    <div class="q-gutter-xs">
                        <q-btn dense color="primary" @click="showEditForm(props.item)" :label="$t('Edit')" />
                        <q-btn dense color="warning" @click="showDeptUser(props.item.dept_code)"
                            :label="$t('Dept') + $t('User')" />
                        <q-btn dense color="warning" @click="showAddChildrenForm(props.item.dept_code)"
                            :label="$t('Add') + $t('Children') + $t('Dept')" />
                        <q-btn dense color="negative" @click="handleDelete(props.item)" :label="$t('Delete')" />
                    </div>
                </td>
            </template>
        </q-hierarchy>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
        <DeptUser ref="deptUserDialog" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import recordDetail from './modules/recordDetail'
import { ArrayToTree } from 'src/utils/arrayAndTree'
import DeptUser from './modules/DeptUser'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'dept/get-dept-list',
    delete: 'dept/delete-dept-by-id',
}
const columns = computed(() => {
    return [
        { name: 'sort', align: 'center', label: t('Sort'), field: 'sort' },
        { name: 'dept_name', align: 'center', label: t('Dept') + t('Name'), field: 'dept_name' },
        { name: 'dept_code', align: 'center', label: t('Dept') + t('Code'), field: 'dept_code' },
        { name: 'leader', align: 'center', label: t('Status'), field: 'leader' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    GqaAvatar,
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
    pagination.value.rowsPerPage = 99999
    getTableData()
})

const deptTree = computed(() => {
    if (tableData.value.length !== 0) {
        return ArrayToTree(tableData.value, 'dept_code', 'parent_code')
    }
    return []
})

const deptUserDialog = ref(null)
const showDeptUser = (deptCode) => {
    deptUserDialog.value.show(deptCode)
}
const showAddParentForm = () => {
    showAddForm()
}
const showAddChildrenForm = (deptCode) => {
    recordDetailDialog.value.formType = 'add'
    recordDetailDialog.value.show()
    recordDetailDialog.value.recordDetail.value.parent_code = deptCode
}
</script>

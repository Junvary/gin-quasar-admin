<template>
    <q-page padding>
        <q-card flat>
            <q-card-section class="row q-gutter-x-md items-center">
                <q-btn color="primary" @click="showAddParentForm()" :label="$t('Add') + $t('Parent') + $t('Dept')"
                    v-has="'dept:addParent'" />
            </q-card-section>
            <q-card-section>
                <q-hierarchy separator="cell" dense :columns="columns" :data="deptTree">
                    <template v-slot:body="props">
                        <gqa-tree-td :treeTd="props" firstTd="sort"></gqa-tree-td>
                        <td class="text-center">{{ props.item.dept_name }}</td>
                        <td class="text-center">{{ props.item.dept_code }}</td>
                        <td class="text-center">
                            <GqaShowName :customNameObject="props.item.leader_user" />
                        </td>
                        <td class="text-center">
                            <GqaDictShow :dictCode="props.item.status" />
                        </td>
                        <td class="text-center q-gutter-x-xs">
                            <q-btn flat dense rounded icon="eva-edit-2-outline" color="primary"
                                @click="showEditForm(props.item)" v-has="'dept:edit'">
                                <q-tooltip>
                                    {{ $t('Edit') }}
                                </q-tooltip>
                            </q-btn>
                            <q-btn flat dense rounded icon="add" color="warning"
                                @click="showAddChildrenForm(props.item.dept_code)" v-has="'dept:addChildren'">
                                <q-tooltip>
                                    {{ $t('Add') + $t('Children') + $t('Dept') }}
                                </q-tooltip>
                            </q-btn>
                            <q-btn flat dense rounded icon="eva-people-outline" color="positive"
                                @click="showDeptUser(props.item)" v-has="'dept:deptUser'">
                                <q-tooltip>
                                    {{ $t('Dept') + $t('User') }}
                                </q-tooltip>
                            </q-btn>
                            <q-btn flat dense rounded icon="delete_outline" color="negative"
                                @click="handleDelete(props.item)" v-has="'dept:delete'">
                                <q-tooltip>
                                    {{ $t('Delete') }}
                                </q-tooltip>
                            </q-btn>
                        </td>
                    </template>
                </q-hierarchy>
            </q-card-section>
        </q-card>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
        <DeptUser ref="deptUserDialog" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, onMounted, ref } from 'vue'
import recordDetail from './modules/recordDetail.vue'
import { ChangeNullChildren2Array } from 'src/utils/arrayAndTree'
import DeptUser from './modules/DeptUser.vue'

const url = {
    list: 'dept/get-dept-list',
    delete: 'dept/delete-dept-by-id',
}
const columns = computed(() => {
    return [
        { name: 'sort', align: 'center', label: t('Sort'), field: 'sort' },
        { name: 'dept_name', align: 'center', label: t('Dept') + t('Name'), field: 'dept_name' },
        { name: 'dept_code', align: 'center', label: t('Dept') + t('Code'), field: 'dept_code' },
        { name: 'leader', align: 'center', label: t('Leader'), field: 'leader' },
        { name: 'status', align: 'center', label: t('Status'), field: 'status' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    t,
    pagination,
    GqaDictShow,
    GqaShowName,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    getTableData,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    pagination.value.rowsPerPage = 99999
    getTableData()
})

const deptTree = computed(() => {
    return ChangeNullChildren2Array(tableData.value) || []

})

const deptUserDialog = ref(null)
const showDeptUser = (dept) => {
    deptUserDialog.value.show(dept)
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

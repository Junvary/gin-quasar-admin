<template>
    <div class="items-center column">
        <div class="justify-center row" style="width: 100%">
            <q-btn color="primary" @click="handleDataPermission">
                {{ $t('Save') }}
            </q-btn>
        </div>

        <q-select v-model="deptDataPermissionType" :options="dictOptions.deptDataPermissionType" emit-value map-options
            :label="$t('DeptDataPermissionType')" style="width: 100%" @update:model-value="checkCustom"
            v-if="dictOptions.deptDataPermissionType?.length">
            <template v-slot:option="scope">
                <q-item v-ripple v-bind="scope.itemProps">
                    <q-item-section>
                        <q-item-label>
                            {{ scope.opt.label }}
                        </q-item-label>
                        <q-item-label caption>
                            {{ scope.opt.value }}
                        </q-item-label>
                    </q-item-section>
                </q-item>
            </template>
        </q-select>

        <q-card-section style="width: 100%; max-height: 70vh" class="scroll"
            v-if="tableData.length !== 0 && deptDataPermissionType === 'deptDataPermissionType_custom'">
            <q-tree style="width: 100%" :nodes="tableData" default-expand-all node-key="dept_code" label-key="name"
                selected-color="primary" tick-strategy="strict" v-model:ticked="ticked" v-if="tableData.length">
                <template v-slot:default-header="prop">
                    <div class="items-center row">
                        <div class="text-weight-bold">{{ prop.node.dept_name }}</div>
                    </div>
                </template>
            </q-tree>
        </q-card-section>
        <q-inner-loading :showing="loading">
            <q-spinner-gears size="50px" color="primary" />
        </q-inner-loading>
    </div>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { postAction } from 'src/api/manage'
import { onMounted, ref, toRefs } from 'vue'

const url = {
    list: 'dept/get-dept-list',
    roleDeptDataEdit: 'role/edit-role-dept-data-permission',
}
const props = defineProps({
    row: {
        type: Object,
        required: true,
    }
})
const { row } = toRefs(props)
const {
    $q,
    dictOptions,
    pagination,
    loading,
    tableData,
    getTableData,
} = useTableData(url)

const ticked = ref([])
const deptDataPermissionType = ref('')

onMounted(async () => {
    pagination.value.rowsPerPage = 99999
    deptDataPermissionType.value = row.value.dept_data_permission_type
    if (row.value.dept_data_permission_custom !== '') {
        ticked.value = row.value.dept_data_permission_custom.split(',')
    }
    if (row.value.dept_data_permission_type === 'deptDataPermissionType_custom') {
        getTableData()
    }
})
const checkCustom = (value) => {
    if (value === 'deptDataPermissionType_custom') {
        getTableData()
    }
}
const handleDataPermission = () => {
    let custom = ''
    if (deptDataPermissionType.value === 'deptDataPermissionType_custom') {
        custom = ticked.value.join(',')
    } else {
        ticked.value = []
    }
    postAction(url.roleDeptDataEdit, {
        role_code: row.value.role_code,
        dept_data_permission_type: deptDataPermissionType.value,
        dept_data_permission_custom: custom,
    }).then((res) => {
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: res.message,
            })
        }
    })
}
</script>

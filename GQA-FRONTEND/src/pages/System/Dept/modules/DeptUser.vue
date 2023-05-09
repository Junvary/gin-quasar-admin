<template>
    <q-dialog v-model="deptUserVisible" position="right">
        <q-card style="min-width: 500px; max-width: 45vw; height: 100%;">
            <q-card-section>
                {{ record.dept_name }}
            </q-card-section>
            <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
                <template v-slot:top="props">
                    <q-btn dense color="primary" @click="showAddUserForm()" :label="$t('Add') + ' ' + $t('User')" />
                    <q-space />
                    <!-- <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                        @click="props.toggleFullscreen" class="q-ml-md" /> -->
                </template>

                <template v-slot:body-cell-actions="props">
                    <q-td :props="props">
                        <div class="q-gutter-xs">
                            <q-btn dense color="negative" @click="handleRemove(props.row)" :label="$t('Delete')" />
                        </div>
                    </q-td>
                </template>
            </q-table>
        </q-card>
        <SelectUserDialog ref="selectUserDialog" @handleSelectUser="handleAddUser" selection="multiple" />
    </q-dialog>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { postAction } from 'src/api/manage'
import { computed, ref } from 'vue'
import SelectUserDialog from 'src/components/GqaSeleteUser/SelectUserDialog.vue'

const url = {
    list: 'dept/query-user-by-dept',
    removeUser: 'dept/remove-dept-user',
    addUser: 'dept/add-dept-user',
}
const columns = computed(() => {
    return [
        { name: 'username', align: 'center', label: t('Username'), field: 'username' },
        { name: 'nickname', align: 'center', label: t('Nickname'), field: 'nickname' },
        { name: 'real_name', align: 'center', label: t('RealName'), field: 'real_name' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    $q,
    t,
    pagination,
    queryParams,
    pageOptions,
    loading,
    tableData,
    onRequest,
    getTableData,
} = useTableData(url)

const deptUserVisible = ref(false)
const deptCode = ref('')
const record = ref({})

const show = (dept) => {
    pagination.value.sortBy = 'username'
    tableData.value = []
    record.value = dept
    deptCode.value = dept.dept_code
    queryParams.value.dept_code = deptCode.value
    deptUserVisible.value = true
    getTableData()
}
defineExpose({
    show
})
const selectUserDialog = ref(null)
const showAddUserForm = () => {
    selectUserDialog.value.show(tableData.value)
}
const handleRemove = (row) => {
    $q.dialog({
        title: t('Confirm'),
        message: t('Confirm') + t('Delete') + '?',
        cancel: true,
        persistent: true,
    }).onOk(async () => {
        const res = await postAction(url.removeUser, {
            dept_code: deptCode.value,
            username: row.username,
        })
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: res.message,
            })
        }
        getTableData()
    })
}
const handleAddUser = (event) => {
    const usernameList = []
    for (let i of event) {
        usernameList.push(i.username)
    }
    postAction(url.addUser, {
        dept_code: deptCode.value,
        username: usernameList,
    }).then((res) => {
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: res.message,
            })
        }
        getTableData()
    })
}
</script>

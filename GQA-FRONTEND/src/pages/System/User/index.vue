<template>
    <q-page padding>
        <q-card flat class="row">
            <q-card-section class="col-2">
                <q-card-section>
                    {{ $t('Select') + $t('Dept') }}
                </q-card-section>
                <q-separator />
                <q-card-section v-if="deptList.length">
                    <q-tree :nodes="deptList" v-model:selected="selectDept" @update:selected="handleSelectDept"
                        label-key="dept_name" node-key="dept_code" selected-color="primary" default-expand-all />
                </q-card-section>
            </q-card-section>
            <q-card-section class="col q-gutter-y-md">
                <q-card-section class="row q-gutter-x-md items-center">
                    <q-input outlined dense style="width: 20%" v-model="queryParams.username" :label="$t('Username')" />
                    <q-input outlined dense style="width: 20%" v-model="queryParams.real_name" :label="$t('RealName')" />
                    <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
                    <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
                </q-card-section>
                <q-card-section>
                    <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns"
                        v-model:pagination="pagination" :rows-per-page-options="pageOptions" :loading="loading"
                        @request="onRequest">
                        <template v-slot:top="props">
                            <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('User')"
                                v-has="'user:add'" />
                            <q-space />
                            <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                                @click="props.toggleFullscreen" class="q-ml-md" />
                        </template>
                        <template v-slot:body-cell-avatar="props">
                            <q-td :props="props">
                                <gqa-avatar :src="props.row.avatar" />
                            </q-td>
                        </template>
                        <template v-slot:body-cell-gender="props">
                            <q-td :props="props">
                                <GqaDictShow :dictCode="props.row.gender" />
                            </q-td>
                        </template>
                        <template v-slot:body-cell-role="props">
                            <q-td :props="props">
                                <div class="column items-center q-gutter-xs">
                                    <q-badge class="col" color="primary" v-for="(item, index) in props.row.role"
                                        :key="index">
                                        {{ item.role_name }}
                                    </q-badge>
                                </div>
                            </q-td>
                        </template>
                        <template v-slot:body-cell-dept="props">
                            <q-td :props="props">
                                <div class="column items-center q-gutter-xs">
                                    <q-badge class="col" color="primary" v-for="(item, index) in props.row.dept"
                                        :key="index">
                                        {{ item.dept_name }}
                                    </q-badge>
                                </div>
                            </q-td>
                        </template>
                        <template v-slot:body-cell-status="props">
                            <q-td :props="props">
                                <GqaDictShow :dictCode="props.row.status" />
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
                                    @click="showEditForm(props.row)" v-has="'user:edit'">
                                    <q-tooltip>
                                        {{ $t('Edit') }}
                                    </q-tooltip>
                                </q-btn>
                                <q-btn flat dense rounded icon="mdi-lock-reset" color="warning"
                                    @click="resetPassword(props.row)" v-has="'user:password'">
                                    <q-tooltip>
                                        {{ $t('Reset') + $t('Password') }}
                                    </q-tooltip>
                                </q-btn>
                                <q-btn flat dense rounded icon="delete_outline" color="negative"
                                    @click="handleDelete(props.row)" v-has="'user:delete'">
                                    <q-tooltip>
                                        {{ $t('Delete') }}
                                    </q-tooltip>
                                </q-btn>
                            </q-td>
                        </template>
                    </q-table>
                </q-card-section>
            </q-card-section>
        </q-card>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref } from 'vue'
import recordDetail from './modules/recordDetail.vue'

const url = {
    list: 'user/get-user-list',
    delete: 'user/delete-user-by-id',
    resetPassword: 'user/reset-password',
    deptListUrl: 'dept/get-dept-list'
}
const columns = computed(() => {
    return [
        { name: 'avatar', align: 'center', label: t('Avatar'), field: 'avatar' },
        { name: 'username', align: 'center', label: t('Username'), field: 'username' },
        { name: 'nickname', align: 'center', label: t('Nickname'), field: 'nickname' },
        { name: 'real_name', align: 'center', label: t('RealName'), field: 'real_name' },
        { name: 'gender', align: 'center', label: t('Gender'), field: 'gender' },
        { name: 'role', align: 'center', label: t('Role'), field: 'role' },
        { name: 'dept', align: 'center', label: t('Dept'), field: 'dept' },
        { name: 'status', align: 'center', label: t('Status'), field: 'status' },
        { name: 'stable', align: 'center', label: t('Stable'), field: 'stable' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    $q,
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
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    pagination.value.sortBy = 'username'
    queryParams.value = {
        with_admin: true,
    }
    getTableData()
    getDeptList()
})

const selectDept = ref(null)
const handleSelectDept = (target) => {
    queryParams.value.dept_code = target
    getTableData()
}
const deptList = ref([])
const getDeptList = () => {
    deptList.value = []
    postAction(url.deptListUrl, {
        sort_by: 'sort',
        desc: false,
        page: 1,
        page_size: 99999
    }).then(res => {
        deptList.value = res.data.records
    })
}

const resetPassword = (row) => {
    $q.dialog({
        title: t('Reset') + t('Password'),
        message: t('Confirm') + t('Reset') + t('Password') + '?' + '(' + row.username + ')',
        cancel: true,
        persistent: true,
    }).onOk(async () => {
        const res = await postAction(url.resetPassword, {
            id: row.id,
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

const resetSearch = () => {
    selectDept.value = null
    queryParams.value = {
        with_admin: true,
    }
    pagination.value.sortBy = 'username'
    getTableData()
}

</script>

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
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import recordDetail from './modules/recordDetail'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'user/get-user-list',
    delete: 'user/delete-user-by-id',
    resetPassword: 'user/reset-password',
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
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    pagination.value.sortBy = 'username'
    queryParams.value = {
        withAdmin: true,
    }
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    })
})

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
        onRequest({
            pagination: pagination.value,
            queryParams: queryParams.value
        })
    })
}

</script>

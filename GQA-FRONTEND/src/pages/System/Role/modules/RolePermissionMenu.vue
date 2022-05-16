<template>
    <div class="items-center column">
        <div class="justify-between row" style="width: 100%">
            <q-btn color="negative" :disable="row.role_code === 'super-admin'" @click="handleClear">
                {{ $t('Clear') + $t('All') }}
            </q-btn>
            <q-btn color="negative" :disable="row.role_code === 'super-admin'" @click="handleAll">
                {{ $t('Select') + $t('All') }}
            </q-btn>
            <q-btn color="primary" :disable="row.role_code === 'super-admin'" @click="handleRoleMenu">
                {{ $t('Save') }}
            </q-btn>
        </div>
        <q-card-section style="width: 100%; max-height: 70vh" class="scroll">
            <q-tree dense style="width: 100%" :nodes="menuTree" default-expand-all node-key="name" label-key="name"
                selected-color="primary" v-if="menuTree.length !== 0" tick-strategy="strict" v-model:ticked="ticked">
                <template v-slot:default-header="prop">
                    <div class="items-center row">
                        <q-icon :name="prop.node.icon || 'share'" size="28px" class="q-mr-sm" />
                        <div class="text-weight-bold">{{ $t(prop.node.title) }}</div>
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
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref, toRefs } from 'vue'
import { useI18n } from 'vue-i18n'
import { DictOptions } from 'src/utils/dict'
import { FormatDateTime } from 'src/utils/date'
import { ArrayToTree } from 'src/utils/arrayAndTree'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'menu/get-menu-list',
    roleMenuList: 'role/get-role-menu-list',
    roleMenuEdit: 'role/edit-role-menu',
}
const columns = computed(() => {
    return [
        { name: 'id', align: 'center', label: 'ID', field: 'id' },
        { name: 'login_username', align: 'center', label: t('User'), field: 'login_username' },
        { name: 'login_ip', align: 'center', label: 'IP', field: 'login_ip' },
        { name: 'login_browser', align: 'center', label: t('Browser'), field: 'login_browser' },
        { name: 'login_os', align: 'center', label: t('Os'), field: 'login_os' },
        { name: 'login_platform', align: 'center', label: t('Platform'), field: 'login_platform' },
        { name: 'created_at', align: 'center', label: t('CreatedAt'), field: 'created_at' },
        { name: 'login_success', align: 'center', label: t('LoginSuccess'), field: 'login_success' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const props = defineProps({
    row: {
        type: Object,
        required: true,
    }
})
const { row } = toRefs(props)
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

const menuTree = computed(() => {
    if (tableData.value.length !== 0) {
        return ArrayToTree(tableData.value, 'name', 'parent_code')
    }
    return []
})
onMounted(() => {
    pagination.value.rowsPerPage = 99999
    getTableData()
    getRoleMenuList()
})
const ticked = ref([])
const getRoleMenuList = () => {
    // 每次获取前，清空ticked
    ticked.value = []
    postAction(url.roleMenuList, {
        role_code: row.value.role_code,
    }).then((res) => {
        if (res.code === 1) {
            res.data.records.forEach((item) => {
                ticked.value.push(item.sys_menu_name)
            })
        }
    })
}
const handleRoleMenu = () => {
    const roleMenu = []
    for (let i of ticked.value) {
        roleMenu.push({
            role_code: row.value.role_code,
            menu_name: i,
        })
    }
    postAction(url.roleMenuEdit, {
        role_code: row.value.role_code,
        role_menu: roleMenu,
    }).then((res) => {
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: res.message,
            })
            getRoleMenuList()
        }
    })
}
const handleClear = () => {
    ticked.value = []
}
const handleAll = () => {
    tableData.value.forEach((item) => {
        ticked.value.push(item.name)
    })
}
</script>

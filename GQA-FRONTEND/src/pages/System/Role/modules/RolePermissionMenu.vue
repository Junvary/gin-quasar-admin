<template>
    <div class="items-center column">
        <div class="justify-between row" style="width: 100%">
            <q-btn color="negative" @click="handleClear">
                {{ $t('Clear') + $t('All') }}
            </q-btn>
            <q-btn color="negative" @click="handleAll">
                {{ $t('Select') + $t('All') }}
            </q-btn>
            <q-btn color="primary" @click="handleRoleMenu">
                {{ $t('Save') }}
            </q-btn>
        </div>
        <q-card-section style="width: 100%; max-height: 70vh" class="scroll">
            <q-tree dense style="width: 100%" :nodes="menuTree" default-expand-all node-key="name" label-key="name"
                selected-color="primary" v-if="menuTree.length !== 0" tick-strategy="strict" v-model:ticked="ticked">
                <template v-slot:default-header="prop">
                    <div class="items-center row">
                        <q-icon :name="prop.node.icon || 'share'" size="20px" class="q-mr-sm" />
                        <div class="text-weight-bold">
                            {{ selectOptionLabel(prop.node) }}
                        </div>
                        <q-radio v-model="defaultPage" dense :val="prop.node.name">
                            <q-tooltip anchor="center end" self="center left">
                                {{ $t('Default') + $t('Page') }}
                            </q-tooltip>
                        </q-radio>
                        <div class="row q-gutter-x-md" style="margin-left: 10px">
                            <q-checkbox v-model="buttonCheckMap[item.button_code]" dense :label="item.button_name"
                                color="primary" v-for="item in prop.node.button" :false-value="null" />
                        </div>
                    </div>
                </template>
                <template v-slot:default-body="prop">
                    <q-separator />
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
import { computed, onMounted, ref, toRefs } from 'vue'
import { TreeToArray } from 'src/utils/arrayAndTree';
import useCommon from 'src/composables/useCommon'

const { selectOptionLabel } = useCommon()
const url = {
    list: 'menu/get-menu-list',
    roleMenuList: 'role/get-role-menu-list',
    roleButtonList: 'role/get-role-button-list',
    roleMenuEdit: 'role/edit-role-menu',
}

const defaultPage = ref('dashboard')

const props = defineProps({
    row: {
        type: Object,
        required: true,
    }
})
const { row } = toRefs(props)

const {
    $q,
    pagination,
    loading,
    tableData,
    getTableData,
} = useTableData(url)

const menuTree = computed(() => {
    return tableData.value
})
onMounted(() => {
    pagination.value.rowsPerPage = 99999
    getTableData()
    getRoleButtonList()
    getRoleMenuList()
})

const ticked = ref([])
const getRoleMenuList = () => {
    // Clear ticked before each getRoleMenuList
    ticked.value = []
    postAction(url.roleMenuList, {
        role_code: row.value.role_code,
    }).then((res) => {
        if (res.code === 1) {
            res.data.records.forEach((item) => {
                ticked.value.push(item.sys_menu_name)
                defaultPage.value = row.value.default_page
            })
        }
    })
}

const emit = defineEmits(['handleRoleMenuOk'])
const handleRoleMenu = () => {
    const roleMenu = []
    for (let i of ticked.value) {
        roleMenu.push({
            sys_role_role_code: row.value.role_code,
            sys_menu_name: i,
        })
    }
    const roleButton = []
    for (let bc in buttonCheckMap.value) {
        if (buttonCheckMap.value[bc]) {
            roleButton.push({
                sys_role_role_code: row.value.role_code,
                sys_button_button_code: bc,
            })
        }
    }
    postAction(url.roleMenuEdit, {
        role_code: row.value.role_code,
        role_menu: roleMenu,
        role_button: roleButton,
        default_page: defaultPage.value
    }).then((res) => {
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: res.message,
            })
            emit('handleRoleMenuOk', { id: row.value.id })
            getRoleButtonList()
            getRoleMenuList()
        }
    })
}
const handleClear = () => {
    ticked.value = []
}
const handleAll = () => {
    const arrayMenu = TreeToArray(tableData.value)
    arrayMenu.forEach((item) => {
        ticked.value.push(item.name)
    })
}

const buttonCheckMap = ref({})
const getRoleButtonList = () => {
    buttonCheckMap.value = {}
    postAction(url.roleButtonList, {
        role_code: row.value.role_code,
    }).then(res => {
        res.data.records.forEach(item => {
            buttonCheckMap.value[item.sys_button_button_code] = true
        })
    })
}
</script>

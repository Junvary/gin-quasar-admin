<template>
    <div class="items-center column">
        <div class="justify-between row" style="width: 100%">
            <q-btn color="negative" @click="handleClear">
                {{ $t('Clear') + $t('All') }}</q-btn>
            <q-btn color="negative" @click="handleAll">
                {{ $t('Select') + $t('All') }}</q-btn>
            <q-btn color="primary" @click="handleRoleApi">
                {{ $t('Save') }}
            </q-btn>
        </div>

        <q-card-section style="width: 100%; max-height: 70vh" class="scroll">
            <q-splitter v-model="splitterModel">
                <template v-slot:before>
                    <q-tabs v-model="apiTab" dense vertical class="text-grey" active-color="primary"
                        indicator-color="primary">
                        <q-tab v-for="(item, index) in apiData" :name="item.api_group"
                            :label="item.api_group + getThisTickedNumber(item)" :key="index" />
                    </q-tabs>
                </template>
                <template v-slot:after>
                    <q-tab-panels v-model="apiTab" animated swipeable vertical transition-prev="jump-up"
                        transition-next="jump-up">
                        <q-tab-panel v-for="(item, index) in apiData" :name="item.api_group" :key="index">
                            <q-tree dense :nodes="item.children" default-expand-all node-key="trueId"
                                selected-color="primary" v-if="item.children.length !== 0" tick-strategy="strict"
                                v-model:ticked="ticked">
                                <template v-slot:default-header="prop">
                                    <div class="row items-center">
                                        <q-chip color="accent" text-color="white" dense
                                            v-if="prop.node.api_group.substring(0, 7) === 'plugin-'">
                                            {{ prop.node.api_group }}
                                        </q-chip>
                                        <q-chip color="primary" text-color="white" dense v-else>
                                            {{ prop.node.api_group }}
                                        </q-chip>
                                        <q-chip color="primary" text-color="white" dense
                                            v-if="prop.node.api_method === 'POST'">
                                            {{ prop.node.api_method }}
                                        </q-chip>
                                        <q-chip color="positive" text-color="white" dense
                                            v-if="prop.node.api_method === 'GET'">
                                            {{ prop.node.api_method }}
                                        </q-chip>
                                        <q-chip color="negative" text-color="white" dense
                                            v-if="prop.node.api_method === 'DELETE'">
                                            {{ prop.node.api_method }}
                                        </q-chip>
                                        <q-chip color="warning" text-color="white" dense
                                            v-if="prop.node.api_method === 'PUT'">
                                            {{ prop.node.api_method }}
                                        </q-chip>
                                        <div class="text-weight-bold">
                                            {{ prop.node.api_path }}
                                        </div>
                                        <span class="text-weight-light text-black">
                                            （{{ prop.node.memo }}）
                                        </span>
                                    </div>
                                </template>
                            </q-tree>
                        </q-tab-panel>
                    </q-tab-panels>
                </template>
            </q-splitter>
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

const splitterModel = ref(20)
const url = {
    list: 'api/get-api-list',
    roleApiList: 'role/get-role-api-list',
    roleApiEdit: 'role/edit-role-api',
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
    pagination,
    loading,
    tableData,
    getTableData,
} = useTableData(url)

onMounted(() => {
    pagination.value.rowsPerPage = 99999
    getTableData()
    getRoleApiList()
})

const apiDataTrue = ref([])
const apiTab = ref('')

const apiData = computed(() => {
    if (tableData.value.length) {
        const data = tableData.value
        for (let item of data) {
            item.trueId = 'g:' + item.api_group + 'p:' + item.api_path + 'm:' + item.api_method
        }
        apiDataTrue.value = data
        const apiTree = []
        for (let d of data) {
            if (apiTree.find((item) => item.api_group === d.api_group) === undefined) {
                apiTree.push({
                    api_group: d.api_group,
                    children: [],
                })
            }
        }
        for (let d of data) {
            for (let a of apiTree) {
                if (a.api_group === d.api_group) {
                    a.children.push(d)
                }
            }
        }
        apiTab.value = apiTree[0].api_group
        // If the role is super admin, disable all API editing to ensure that the API can have all call permissions
        if (row.value.role_code === 'super-admin') {
            for (let a of apiTree) {
                for (let i of a.children) {
                    i.disabled = true
                }
            }
        }
        return apiTree
    }
    return []
})
const ticked = ref([])

const getThisTickedNumber = computed(() => {
    return (api) => {
        const allNumber = api.children.length
        var tickedNumber = 0
        for (let t of ticked.value) {
            if (api.children.find((item) => item.trueId === t) !== undefined) {
                tickedNumber++
            }
        }
        return '(' + tickedNumber + '/' + allNumber + ')'
    }
})

const getRoleApiList = () => {
    // Clear ticked before each getRoleApiList
    ticked.value = []
    postAction(url.roleApiList, {
        role_code: row.value.role_code,
    }).then((res) => {
        if (res.code === 1) {
            res.data.records.forEach((item) => {
                ticked.value.push('g:' + item.api_group + 'p:' + item.api_path + 'm:' + item.api_method)
            })
        }
    })
}
const handleRoleApi = () => {
    const roleApi = []
    tableData.value.forEach((item) => {
        for (let t of ticked.value) {
            if (t === item.trueId) {
                roleApi.push({
                    role_code: row.value.role_code,
                    api_group: item.api_group,
                    api_method: item.api_method,
                    api_path: item.api_path,
                })
            }
        }
    })
    postAction(url.roleApiEdit, {
        role_code: row.value.role_code,
        role_api: roleApi,
    }).then((res) => {
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: res.message,
            })
            getRoleApiList()
        }
    })
}
const handleClear = () => {
    ticked.value = []
}
const handleAll = () => {
    ticked.value = []
    tableData.value.forEach((item) => {
        ticked.value.push('g:' + item.api_group + 'p:' + item.api_path + 'm:' + item.api_method)
    })
}
</script>

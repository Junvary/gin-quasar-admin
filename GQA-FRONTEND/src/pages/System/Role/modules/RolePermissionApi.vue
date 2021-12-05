<template>
    <div class="items-center column">
        <div class="justify-between row" style="width: 100%">
            <q-btn color="negative" :disable="row.roleCode === 'super-admin'" @click="handleClear">
                {{ $t('ClearAll') }}</q-btn>
            <q-btn color="negative" :disable="row.roleCode === 'super-admin'" @click="handleAll">
                {{ $t('SelectAll') }}</q-btn>
            <q-btn color="primary" :disable="row.roleCode === 'super-admin'" @click="handleRoleApi">
                {{ $t('PageSystemRolePermissionAPISave') }}</q-btn>
        </div>

        <q-card-section style="width: 100%; max-height: 70vh" class="scroll">
            <q-tree dense :nodes="apiData" default-expand-all node-key="trueId" selected-color="primary"
                v-if="tableData.length !== 0" tick-strategy="strict" v-model:ticked="ticked">
                <template v-slot:default-header="prop">
                    <div class="row items-center">
                        <q-chip color="accent" text-color="white" dense
                            v-if="prop.node.apiGroup.substring(0, 7)=== 'plugin-'">
                            {{ prop.node.apiGroup }}
                        </q-chip>
                        <q-chip color="primary" text-color="white" dense v-else>
                            {{ prop.node.apiGroup }}
                        </q-chip>
                        <q-chip color="primary" text-color="white" dense v-if="prop.node.apiMethod=== 'POST'">
                            {{ prop.node.apiMethod }}
                        </q-chip>
                        <q-chip color="positive" text-color="white" dense v-if="prop.node.apiMethod=== 'GET'">
                            {{ prop.node.apiMethod }}
                        </q-chip>
                        <q-chip color="negative" text-color="white" dense v-if="prop.node.apiMethod=== 'DELETE'">
                            {{ prop.node.apiMethod }}
                        </q-chip>
                        <q-chip color="warning" text-color="white" dense v-if="prop.node.apiMethod=== 'PUT'">
                            {{ prop.node.apiMethod }}
                        </q-chip>
                        <div class="text-weight-bold">{{ prop.node.apiPath }}</div>
                        <span class="text-weight-light text-black">
                            （{{ prop.node.remark}}）
                        </span>
                    </div>
                </template>
            </q-tree>
        </q-card-section>
        <q-inner-loading :showing="loading">
            <q-spinner-gears size="50px" color="primary" />
        </q-inner-loading>
    </div>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { postAction, putAction } from 'src/api/manage'

export default {
    name: 'RolePermissionApi',
    mixins: [tableDataMixin],
    props: {
        row: {
            type: Object,
            required: true,
        },
    },
    computed: {
        apiData() {
            if (this.tableData.length) {
                // if (this.row.roleCode === 'super-admin') {
                //     for (let i of this.tableData) {
                //         i.disabled = true
                //     }
                // }
                const data = this.tableData
                for (let item of data) {
                    item.trueId = 'p:' + item.apiPath + 'm:' + item.apiMethod
                }
                return data
            }
            return []
        },
    },
    data() {
        return {
            pagination: {
                sortBy: 'sort',
                descending: false,
                page: 1,
                rowsPerPage: 10000,
                rowsNumber: 0,
            },
            url: {
                list: 'api/api-list',
                roleApiList: 'role/role-api',
                roleApiEdit: 'role/role-api-edit',
            },
            ticked: [],
        }
    },
    created() {
        this.getTableData()
        this.getRoleApiList()
    },
    methods: {
        getRoleApiList() {
            // 每次获取前，清空ticked
            this.ticked = []
            postAction(this.url.roleApiList, {
                roleCode: this.row.roleCode,
            }).then((res) => {
                if (res.code === 1) {
                    res.data.records.forEach((item) => {
                        this.ticked.push('p:' + item[1] + 'm:' + item[2])
                    })
                }
            })
        },
        handleRoleApi() {
            const policy = []
            this.tableData.forEach((item) => {
                for (let t of this.ticked) {
                    if (t === item.trueId) {
                        policy.push({
                            V1: item.apiPath,
                            V2: item.apiMethod,
                        })
                    }
                }
            })
            putAction(this.url.roleApiEdit, {
                roleCode: this.row.roleCode,
                policy: policy,
            }).then((res) => {
                if (res.code === 1) {
                    this.$q.notify({
                        type: 'positive',
                        message: res.message,
                    })
                    this.getRoleApiList()
                }
            })
        },
        handleClear() {
            this.ticked = []
        },
        handleAll() {
            this.tableData.forEach((item) => {
                this.ticked.push('p:' + item.apiPath + 'm:' + item.apiMethod)
            })
        },
    },
}
</script>

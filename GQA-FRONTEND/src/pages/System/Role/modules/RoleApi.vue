<template>
    <div class="column items-center">
        <div class="row justify-between" style="width: 100%">
            <q-btn color="negative" :disable="row.roleCode === 'super-admin'" @click="handleClear">全部清空</q-btn>
            <q-btn color="negative" :disable="row.roleCode === 'super-admin'" @click="handleAll">全部选择</q-btn>
            <q-btn color="primary" :disable="row.roleCode === 'super-admin'" @click="handleRoleApi">保存菜单权限</q-btn>
        </div>
        <q-card-section style="width: 100%; max-height: 50vw" class="scroll">
            <q-tree :nodes="apiData" default-expand-all node-key="trueId" selected-color="primary"
                v-if="tableData.length !== 0" tick-strategy="strict" v-model:ticked="ticked">
                <template v-slot:default-header="prop">
                    <div class="row items-center">
                        <q-chip color="primary" dense>
                            {{ prop.node.V3 }}
                        </q-chip>
                        <q-chip color="primary" dense v-if="prop.node.V2=== 'POST'">
                            {{ prop.node.V2 }}
                        </q-chip>
                        <q-chip color="positive" dense v-if="prop.node.V2=== 'GET'">
                            {{ prop.node.V2 }}
                        </q-chip>
                        <q-chip color="negative" dense v-if="prop.node.V2=== 'DELETE'">
                            {{ prop.node.V2 }}
                        </q-chip>
                        <q-chip color="warning" dense v-if="prop.node.V2=== 'PUT'">
                            {{ prop.node.V2 }}
                        </q-chip>
                        <div class="text-weight-bold">{{ prop.node.V1 }}</div>
                        <span class="text-weight-light text-black">
                            （{{ prop.node.V4}}）
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
    name: 'RoleApi',
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
                if (this.row.roleCode === 'super-admin') {
                    for (let i of this.tableData) {
                        i.disabled = true
                    }
                }
                const data = this.tableData
                for (let item of data) {
                    item.trueId = 'p:' + item.V1 + 'm:' + item.V2
                }
                return data
            }
            return []
        },
    },
    data() {
        return {
            pagination: {
                sortBy: 'desc',
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
                    res.data.info.forEach((item) => {
                        this.ticked.push('p:' + item.V1 + 'm:' + item.V2)
                    })
                } else {
                    this.$q.notify({
                        type: 'negative',
                        message: res.message,
                    })
                }
            })
        },
        handleRoleApi() {
            const roleApi = []
            this.tableData.forEach((item) => {
                for (let t of this.ticked) {
                    if (t === item.trueId) {
                        roleApi.push(item)
                    }
                }
            })
            for (let i of roleApi) {
                i.V0 = this.row.roleCode
                delete i.ID
            }
            putAction(this.url.roleApiEdit, {
                roleCode: this.row.roleCode,
                roleApi: roleApi,
            }).then((res) => {
                if (res.code === 1) {
                    this.$q.notify({
                        type: 'positive',
                        message: res.message,
                    })
                    this.getRoleApiList()
                } else {
                    this.$q.notify({
                        type: 'negative',
                        message: res.message,
                    })
                }
            })
        },
        handleClear() {},
        handleAll() {},
    },
}
</script>
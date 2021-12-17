<template>
    <div class="items-center column">
        <div class="justify-center row" style="width: 100%">
            <q-btn color="primary" :disable="row.roleCode === 'super-admin'" @click="handleDataPermission">
                {{ $t('Save') }}</q-btn>
        </div>

        <q-select v-model="deptDataPermissionType" :options="dictOptions.deptDataPermissionType" emit-value map-options
            :label="$t('DeptDataPermissionType')" style="width: 100%" @update:model-value="checkCustom">
            <template v-slot:option="scope">
                <q-item v-bind="scope.itemProps">
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
            v-if="deptTree.length !== 0 && deptDataPermissionType === 'custom'">
            <q-tree style="width: 100%" :nodes="deptTree" default-expand-all node-key="deptCode" label-key="name"
                selected-color="primary" tick-strategy="strict" v-model:ticked="ticked">
                <template v-slot:default-header="prop">
                    <div class="items-center row">
                        <div class="text-weight-bold">{{ prop.node.deptName }}</div>
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
import { ArrayToTree } from 'src/utils/arrayAndTree'
import { postAction, putAction } from 'src/api/manage'
import { DictOptions } from 'src/utils/dict'

export default {
    name: 'RolePermissionData',
    mixins: [tableDataMixin],
    props: {
        row: {
            type: Object,
            required: true,
        },
    },
    computed: {
        deptTree() {
            if (this.tableData.length !== 0) {
                return ArrayToTree(this.tableData, 'deptCode', 'parentCode')
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
                list: 'dept/dept-list',
                roleDeptDataEdit: 'role/role-dept-data-permission-edit',
            },
            ticked: [],
            deptDataPermissionType: '',
            dictOptions: {},
        }
    },
    async created() {
        this.dictOptions = await DictOptions()
        this.deptDataPermissionType = this.row.deptDataPermissionType
        if (this.row.deptDataPermissionCustom !== '') {
            this.ticked = this.row.deptDataPermissionCustom.split(',')
        }
        if (this.row.deptDataPermissionType === 'custom') {
            this.getTableData()
        }
    },
    methods: {
        checkCustom(value) {
            if (value === 'custom') {
                this.getTableData()
            }
        },
        handleDataPermission() {
            let custom = ''
            if (this.deptDataPermissionType === 'custom') {
                custom = this.ticked.join(',')
            } else {
                this.ticked = []
            }
            putAction(this.url.roleDeptDataEdit, {
                roleCode: this.row.roleCode,
                deptDataPermissionType: this.deptDataPermissionType,
                deptDataPermissionCustom: custom,
            }).then((res) => {
                if (res.code === 1) {
                    this.$q.notify({
                        type: 'positive',
                        message: res.message,
                    })
                }
            })
        },
    },
}
</script>

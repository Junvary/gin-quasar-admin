<template>
    <q-card v-if="addOrEditVisible">
        <div class="row justify-between items-center">
            <q-card-section>
                <div class="text-h6">
                    {{formTypeName}}：
                    {{addOrEditDetail.deptName ? addOrEditDetail.deptName : "部门"}}
                </div>
            </q-card-section>
            <q-card-actions>
                <q-btn :label="'保存' + formTypeName " color="primary" @click="handleAddOrEidt" />
                <q-btn label="取消" color="negative" @click="onClose" />
            </q-card-actions>
        </div>

        <q-separator />

        <q-card-section>
            <q-form ref="addOrEditForm">
                <div class="row justify-between">
                    <div class="q-gutter-md col-9">
                        <div class="row">
                            <q-input class="col" label="ID" v-model="addOrEditDetail.id" disable />
                            <GqaSeleteUser className="col" label="负责人" v-model:selectUser="addOrEditDetail.owner"
                                v-model:selectId="addOrEditDetail.ownerId" selection="single" />
                        </div>

                        <div class="row">
                            <q-field class="col" label="创建时间" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.createdAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" label="创建人" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.createdByUser"
                                        :customNameObject="addOrEditDetail.createdByUser" />
                                </template>
                            </q-field>
                            <q-field class="col" label="更新时间" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.updatedAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" label="更新人" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.updatedByUser"
                                        :customNameObject="addOrEditDetail.updatedByUser" />
                                </template>
                            </q-field>
                        </div>

                        <div class="row">
                            <q-input class="col" label="部门编码" v-model="addOrEditDetail.deptCode"
                                :rules="[ val => val && val.length > 0 || '必须输入部门编码']" />
                            <q-input class="col" label="部门名称" v-model="addOrEditDetail.deptName"
                                :rules="[ val => val && val.length > 0 || '必须输入部门名称']" />
                        </div>

                        <div class="row">
                            <q-input class="col" v-model.number="addOrEditDetail.sort" type="number"
                                :rules="[ val => val >= 1 || '排序必须大于0']" label=" 排序" />
                            <q-input class="col" v-model="addOrEditDetail.phone" label="联系电话" />

                        </div>

                        <div class="row">
                            <q-field class="col" label="是否启用" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.status" :options="options.statusOnOff"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>
                        <q-input v-model="addOrEditDetail.remark" type="textarea" label="备注" />
                    </div>
                    <div class="q-gutter-md col-3">
                        <q-field label="父级部门" stack-label>
                            <template v-slot:control>
                                <q-tree :nodes="menuTree" default-expand-all node-key="id" label-key="name"
                                    selected-color="primary" v-model:selected="addOrEditDetail.parentId"
                                    v-if="menuTree.length !== 0" @update:selected="onSelected">
                                    <template v-slot:default-header="prop">
                                        <div class="row items-center">
                                            <div class="text-weight-bold">{{ prop.node.deptName }}</div>
                                        </div>
                                    </template>
                                </q-tree>
                            </template>
                        </q-field>
                    </div>
                </div>
            </q-form>

        </q-card-section>

        <q-separator />

        <q-inner-loading :showing="loading">
            <q-spinner-gears size="50px" color="primary" />
        </q-inner-loading>
    </q-card>
</template>

<script>
import { addOrEditMixin } from 'src/mixins/addOrEditMixin'
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { ArrayToTree } from 'src/utils/arrayAndTree'
import GqaSeleteUser from 'src/components/GqaSeleteUser'
import GqaShowName from 'src/components/GqaShowName'

export default {
    name: 'addOrEditCard',
    mixins: [addOrEditMixin, tableDataMixin],
    components: {
        GqaShowName,
        GqaSeleteUser,
    },
    computed: {
        menuTree() {
            if (this.tableData.length !== 0) {
                return ArrayToTree(this.tableData)
            }
            return []
        },
    },
    data() {
        return {
            detail: {
                id: '',
                createdAt: '',
                createdBy: '',
                updatedAt: '',
                updatedBy: '',
                sort: 1,
                status: 'on',
                remark: '',
                parentId: 0,
                owner: {},
                deptCode: '',
                deptName: '',
                phone: '',
            },
            url: {
                list: 'dept/dept-list',
                add: 'dept/dept-add',
                edit: 'dept/dept-edit',
                queryById: 'dept/dept-id',
            },
            pagination: {
                sortBy: 'sort',
                descending: false,
                page: 1,
                rowsPerPage: 10000,
                rowsNumber: 0,
            },
            selectedKey: '',
        }
    },
    methods: {
        resetDetail() {
            this.detail = {
                id: '',
                createdAt: '',
                createdBy: '',
                updatedAt: '',
                updatedBy: '',
                sort: 1,
                status: 'on',
                remark: '',
                parentId: 0,
                owner: {},
                deptCode: '',
                deptName: '',
                phone: '',
            }
        },
        show(row) {
            this.loading = true
            this.resetDetail()
            this.addOrEditDetail = Object.assign(this.detail, row)
            this.addOrEditVisible = true
            if (!this.addOrEditDetail.id) {
                this.loading = false
            } else {
                this.handleQueryById(this.addOrEditDetail.id)
            }
            this.getTableData()
        },
        onSelected(key) {
            this.addOrEditDetail.parentId = key
        },
        onClose() {
            this.addOrEditVisible = false
            this.$emit('handleFinish')
        },
    },
}
</script>
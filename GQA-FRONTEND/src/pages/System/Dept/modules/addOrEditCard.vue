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
                <q-btn :label="'保存' + formTypeName " color="primary" @click="emitAddOrEdit" />
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
                            <q-input class="col" label="负责人" v-model="addOrEditDetail.owner.id"
                                v-if="addOrEditDetail.owner.id" />
                        </div>

                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.createAt" label="创建时间" disable />
                            <q-input class="col" v-model="addOrEditDetail.createBy" label="创建人" disable />
                            <q-input class="col" v-model=" addOrEditDetail.updateAt" label="更新时间" disable />
                            <q-input class="col" v-model="addOrEditDetail.updateBy" label="更新人" disable />
                        </div>

                        <div class="row">
                            <q-input class="col" label="部门编码" v-model="addOrEditDetail.deptCode"
                                :rules="[ val => val && val.length > 0 || '必须输入部门编码']" />
                            <q-input class="col" label="部门名称" v-model="addOrEditDetail.deptName"
                                :rules="[ val => val && val.length > 0 || '必须输入部门名称']" />
                        </div>

                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.sort" type="number" label="排序" />
                            <q-input class="col" v-model="addOrEditDetail.phone" label="联系电话" />

                        </div>

                        <div class="row">
                            <q-field class="col" label="是否启用" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.status" :options="options.status"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>
                        <q-input v-model="addOrEditDetail.desc" type="textarea" label="备注" />
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
import { postAction, putAction } from 'src/api/manage'
import { ArrayToTree } from 'src/utils/arrayAndTree'

export default {
    name: 'addOrEditCard',
    mixins: [addOrEditMixin],
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
                createAt: '',
                createBy: '',
                updateAt: '',
                updateBy: '',
                sort: 0,
                status: 0,
                desc: '',
                parentId: 0,
                owner: {},
                deptCode: '',
                deptName: '',
                phone: '',
            },
            addOrEditUrl: {
                list: 'dept/dept-list',
                add: 'dept/dept-add',
                edit: 'dept/dept-edit',
                queryById: 'dept/dept-id',
            },
            tableData: [],
            pagination: {
                sortBy: 'desc',
                descending: false,
                page: 1,
                rowsPerPage: 10000,
                rowsNumber: 0,
            },
            selectedKey: '',
        }
    },
    created() {
        this.getTableData()
    },
    methods: {
        resetDetail() {
            this.detail = {
                id: '',
                createAt: '',
                createBy: '',
                updateAt: '',
                updateBy: '',
                sort: 0,
                status: 0,
                desc: '',
                parentId: 0,
                owner: {},
                deptCode: '',
                deptName: '',
                phone: '',
            }
        },
        onSelected(key) {
            this.addOrEditDetail.parentId = key
        },
        onClose() {
            this.addOrEditVisible = false
            this.$emit('handleFinish')
        },
        getTableData() {
            this.loading = true
            postAction(this.addOrEditUrl.list, {
                page: this.pagination.page,
                pageSize: this.pagination.rowsPerPage,
                ...this.queryParams,
            })
                .then((res) => {
                    this.pagination.rowsNumber = res.data.total
                    this.tableData = res.data.records
                })
                .finally(() => {
                    this.loading = false
                })
        },
        async handleAddOrEidt() {
            const success = await this.$refs.addOrEditForm.validate()
            if (success) {
                if (this.formType === 'edit') {
                    const res = await putAction(this.addOrEditUrl.edit, this.addOrEditDetail)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                    }
                } else if (this.formType === 'add') {
                    const res = await postAction(this.addOrEditUrl.add, this.addOrEditDetail)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                        this.addOrEditVisible = false
                    }
                } else {
                    this.$q.notify({
                        type: 'negative',
                        message: '无法新增或编辑！',
                    })
                }
                this.$emit('handleFinish')
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: '请完善表格信息！',
                })
            }
        },
    },
}
</script>
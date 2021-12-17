<template>
    <q-card v-if="addOrEditVisible">
        <div class="items-center justify-between row">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{ $t('Dept') }}:
                    {{ addOrEditDetail.deptName }}
                </div>
            </q-card-section>
            <q-card-actions>
                <q-btn :label="$t('User')" color="warning" @click="showDeptUser(detail.deptCode)" />
                <q-btn :label="$t('Save')" color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('Cancel')" color="negative" @click="onClose" />
            </q-card-actions>
        </div>

        <q-separator />

        <q-card-section>
            <q-form ref="addOrEditForm">
                <div class="justify-between row">
                    <div class="q-gutter-md col-9">
                        <div class="row">
                            <q-input class="col" label="ID" v-model="addOrEditDetail.id" disable />
                            <GqaSeleteUser className="col" :label="$t('Owner')"
                                v-model:selectUser="addOrEditDetail.owner"
                                v-model:selectUsername="addOrEditDetail.ownerUsername" selection="single" />
                        </div>

                        <div class="row">
                            <q-field class="col" :label="$t('CreatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.createdAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('CreatedBy')" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.createdByUser"
                                        :customNameObject="addOrEditDetail.createdByUser" />
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('UpdatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.updatedAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('UpdatedBy')" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.updatedByUser"
                                        :customNameObject="addOrEditDetail.updatedByUser" />
                                </template>
                            </q-field>
                        </div>

                        <div class="row">
                            <q-input class="col" :label="$t('Dept') + $t('Code')" v-model="addOrEditDetail.deptCode"
                                :rules="[ val => val && val.length > 0 || $t('NeedInput')]" />
                            <q-input class="col" :label="$t('Dept') + $t('Name')" v-model="addOrEditDetail.deptName"
                                :rules="[ val => val && val.length > 0 || $t('NeedInput')]" />
                        </div>

                        <div class="row">
                            <q-input class="col" v-model.number="addOrEditDetail.sort" type="number"
                                :rules="[ val => val >= 1 || $t('SortRule')]" :label="$t('Sort')" />
                            <q-input class="col" v-model="addOrEditDetail.phone" :label="$t('Phone')" />

                        </div>

                        <div class="row">
                            <q-field class="col" :label="$t('Status')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.status" :options="options.statusOnOff"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>
                        <q-input v-model="addOrEditDetail.remark" type="textarea" :label="$t('Remark')" />
                    </div>
                    <div class="q-gutter-md col-3">
                        <q-field :label="$t('Parent')" stack-label>
                            <template v-slot:control>
                                <q-tree :nodes="deptTree" default-expand-all node-key="deptCode" label-key="name"
                                    selected-color="primary" v-model:selected="addOrEditDetail.parentCode"
                                    v-if="deptTree.length !== 0" @update:selected="onSelected">
                                    <template v-slot:default-header="prop">
                                        <div class="items-center row">
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

        <dept-user-dialog ref="deptUserDialog" />
    </q-card>
</template>

<script>
import { addOrEditMixin } from 'src/mixins/addOrEditMixin'
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { ArrayToTree } from 'src/utils/arrayAndTree'
import GqaSeleteUser from 'src/components/GqaSeleteUser'
import GqaShowName from 'src/components/GqaShowName'
import DeptUserDialog from './DeptUserDialog'

export default {
    name: 'addOrEditCard',
    mixins: [addOrEditMixin, tableDataMixin],
    components: {
        GqaShowName,
        GqaSeleteUser,
        DeptUserDialog,
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
            detail: {
                id: '',
                createdAt: '',
                createdBy: '',
                updatedAt: '',
                updatedBy: '',
                sort: 1,
                status: 'on',
                remark: '',
                parentCode: '',
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
                parentCode: '',
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
            this.addOrEditDetail.parentCode = key
        },
        onClose() {
            this.addOrEditVisible = false
            this.$emit('handleFinish')
        },
        showDeptUser(deptCode) {
            this.$refs.deptUserDialog.show(deptCode)
        },
    },
}
</script>

<template>
    <q-card v-if="addOrEditVisible">
        <div class="items-center justify-between row">
            <q-card-section>
                <div class="text-h6">
                    {{formTypeName}}：
                    {{addOrEditDetail.title ? addOrEditDetail.title : $t('PageSystemMenuEditDialogTitle')}}
                </div>
            </q-card-section>
            <q-card-actions>
                <q-btn :label="$t('PageSystemMenuEditDialogBtnSave')" color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('PageSystemMenuEditDialogBtnCancel')" color="negative" @click="onClose" />
            </q-card-actions>
        </div>

        <q-separator />

        <q-card-section>
            <q-form ref="addOrEditForm">
                <div class="justify-between row">
                    <div class="q-gutter-md col-9">
                        <div class="row">
                            <q-input class="col" label="ID" v-model="addOrEditDetail.id" disable />
                            <q-input class="col" :label="$t('PageSystemMenuEditDialogSymbol')" v-model="addOrEditDetail.icon">
                                <template v-slot:before>
                                    <q-icon :name="addOrEditDetail.icon" size="35px" class="q-mr-sm" />
                                </template>
                            </q-input>

                        </div>

                        <div class="row">
                            <q-field class="col" :label="$t('PageSystemMenuEditDialogCreatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.createdAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemMenuEditDialogCreatedBy')" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.createdByUser"
                                        :customNameObject="addOrEditDetail.createdByUser" />
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemMenuEditDialogUpdatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.updatedAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemMenuEditDialogUpdatedBy')" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.updatedByUser"
                                        :customNameObject="addOrEditDetail.updatedByUser" />
                                </template>
                            </q-field>
                        </div>

                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.sort" type="number" :label="$t('PageSystemMenuEditDialogSort')" />
                            <q-input class="col" v-model="addOrEditDetail.title" :label="$t('')"
                                :rules="[ val => val && val.length > 0 || $t('PageSystemMenuEditDialogMenuTitleRule')]" />
                            <q-input class="col" v-model="addOrEditDetail.name" :label="$t('PageSystemMenuEditDialogMenuName')"
                                :rules="[ val => val && val.length > 0 || $t('PageSystemMenuEditDialogMenuNameRule')]" />
                        </div>

                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.path" :label="$t('PageSystemMenuEditDialogMenuPath')"
                                :rules="[ val => val && val.length > 0 || $t('PageSystemMenuEditDialogMenuPathRule')]" />
                            <q-input class="col" v-model="addOrEditDetail.component" label="$t('PageSystemMenuEditDialogMenuComponent')"
                                :rules="[ val => val && val.length > 0 || $t('PageSystemMenuEditDialogMenuComponentRule')]" />
                        </div>

                        <div class="row">
                            <q-field class="col" :label="$t('PageSystemMenuEditDialogMenuIsLink')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.isLink" :options="options.statusYesNo"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemMenuEditDialogMenuKeepAlive')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.keepAlive" :options="options.statusYesNo"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>

                        <div class="row">
                            <q-field class="col" :label="$t('PageSystemMenuEditDialogMenuHidden')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.hidden" :options="options.statusYesNo"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                            <q-field class="col" label="$t('PageSystemMenuEditDialogMenuStatus')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.status" :options="options.statusOnOff"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>
                        <q-input v-model="addOrEditDetail.remark" type="textarea" :label="$t('PageSystemMenuEditDialogMenuRemark')" />
                    </div>
                    <div class="q-gutter-md col-3">
                        <q-field :label="$t('PageSystemMenuEditDialogMenuParent')" stack-label>
                            <template v-slot:control>
                                <q-tree :nodes="menuTree" default-expand-all node-key="id" label-key="name"
                                    selected-color="primary" v-model:selected="addOrEditDetail.parentId"
                                    v-if="menuTree.length !== 0" @update:selected="onSelected">
                                    <template v-slot:default-header="prop">
                                        <div class="items-center row">
                                            <q-icon :name="prop.node.icon || 'share'" size="28px" class="q-mr-sm" />
                                            <div class="text-weight-bold">{{ prop.node.title }}</div>
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
import GqaShowName from 'src/components/GqaShowName'
import { ArrayToTree } from 'src/utils/arrayAndTree'

export default {
    name: 'addOrEditCard',
    mixins: [addOrEditMixin, tableDataMixin],
    components: {
        GqaShowName,
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
                name: '',
                path: '',
                component: '',
                title: '',
                icon: '',
                hidden: 'no',
                keepAlive: 'no',
                isLink: 'no',
            },
            url: {
                list: 'menu/menu-list',
                add: 'menu/menu-add',
                edit: 'menu/menu-edit',
                queryById: 'menu/menu-id',
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
                name: '',
                path: '',
                component: '',
                title: '',
                icon: '',
                hidden: 'no',
                keepAlive: 'no',
                isLink: 'no',
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

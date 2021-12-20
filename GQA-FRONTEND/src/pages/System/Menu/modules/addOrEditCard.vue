<template>
    <q-card v-if="addOrEditVisible">
        <div class="row justify-between items-center">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{ $t('Menu') }}:
                    {{ addOrEditDetail.title }}
                </div>
            </q-card-section>
            <q-card-actions>
                <q-btn :label="$t('Save') + ' ' + formTypeName " color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('Cancel')" color="negative" @click="onClose" />
            </q-card-actions>
        </div>

        <q-separator />

        <q-card-section>
            <q-form ref="addOrEditForm">
                <div class="row justify-between">
                    <div class="q-gutter-md col-9">
                        <div class="row">
                            <q-input class="col" label="ID" v-model="addOrEditDetail.id" disable />
                            <q-input class="col" :label="$t('Icon')" v-model="addOrEditDetail.icon">
                                <template v-slot:before>
                                    <q-icon :name="addOrEditDetail.icon" size="35px" class="q-mr-sm" />
                                </template>
                            </q-input>

                        </div>

                        <div class="row">
                            <q-field class="col" :label="$t('CreatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{ showDateTime(addOrEditDetail.createdAt) }}
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
                                    {{ showDateTime(addOrEditDetail.updatedAt) }}
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
                            <q-input class="col" v-model.number="addOrEditDetail.sort" type="number"
                                :rules="[ val => val >= 1 || $t('SortRule')]" :label="$t('Sort')" />
                            <q-input class="col" v-model="addOrEditDetail.title" :label="$t('Menu') + ' ' + $t('Title')"
                                :rules="[ val => val && val.length > 0 || $t('NeetInput')]" />
                            <q-input class="col" v-model="addOrEditDetail.name"
                                :label="$t('Menu') + 'Name' + '(' + $t('English') + ',' + $t('Unique') + ')'"
                                :rules="[ val => val && val.length > 0 || $t('NeetInput')]" />
                        </div>

                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.path" label="URL"
                                :rules="[ val => val && val.length > 0 || $t('NeetInput')]" />
                            <q-input class="col" v-model="addOrEditDetail.component" :label="$t('Component')"
                                :rules="[ val => val && val.length > 0 || $t('NeetInput')]" />
                        </div>

                        <div class="row">
                            <q-field class="col" :label="$t('IsLink')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.isLink" :options="dictOptions.statusYesNo"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('KeepAlive')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.keepAlive"
                                        :options="dictOptions.statusYesNo" color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>

                        <div class="row">
                            <q-field class="col" :label="$t('Hidden')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.hidden" :options="dictOptions.statusYesNo"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('Status')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.status" :options="dictOptions.statusOnOff"
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
                                <q-scroll-area style="height: 65vh; width: 100%">
                                    <q-tree dense :nodes="menuTree" default-expand-all node-key="name" label-key="name"
                                        selected-color="primary" v-model:selected="addOrEditDetail.parentCode"
                                        v-if="menuTree.length !== 0" @update:selected="onSelected">
                                        <template v-slot:default-header="prop">
                                            <div class="row items-center">
                                                <q-icon :name="prop.node.icon || 'share'" size="sm" class="q-mr-sm" />
                                                <div class="text-weight-bold">{{ $t(prop.node.title) }}</div>
                                            </div>
                                        </template>
                                    </q-tree>
                                </q-scroll-area>
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
                return ArrayToTree(this.tableData, 'name', 'parentCode')
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
                parentCode: '',
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
            this.addOrEditDetail.parentCode = key
        },
        onClose() {
            this.addOrEditVisible = false
            this.$emit('handleFinish')
        },
    },
}
</script>
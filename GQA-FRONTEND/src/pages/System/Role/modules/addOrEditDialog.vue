<template>
    <q-dialog v-model="addOrEditVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{formTypeName}}：
                    {{addOrEditDetail.roleName ? addOrEditDetail.roleName : $t('PageSystemRoleEditDialogTitle')}}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section style="max-height: 80vh" class="scroll q-gutter-md">
                <q-form ref="addOrEditForm">
                    <div class="q-gutter-md">
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.id" label="ID" disable />
                        </div>
                        <div class="row">
                            <q-field class="col" :label="$t('PageSystemRoleEditDialogCreatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.createdAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemRoleEditDialogCreatedBy')" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.createdByUser"
                                        :customNameObject="addOrEditDetail.createdByUser" />
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemRoleEditDialogUpdatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.updatedAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemRoleEditDialogUpdatedBy')" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.updatedByUser"
                                        :customNameObject="addOrEditDetail.updatedByUser" />
                                </template>
                            </q-field>
                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.roleCode" :label="$t('PageSystemRoleEditDialogRoleCode')"
                                :rules="[ val => val && val.length > 0 || $t('PageSystemRoleEditDialogRoleCodeRule')]" />
                            <q-input class="col" v-model="addOrEditDetail.roleName" :label="$t('PageSystemRoleEditDialogRoleName')"
                                :rules="[ val => val && val.length > 0 || $t('PageSystemRoleEditDialogRoleNameRule')]" />
                        </div>
                        <div class="justify-between row">
                            <!-- <div class="col-6">
                                <q-field label="数据归属" stack-label>
                                    <template v-slot:control>
                                        <q-tree :nodes="deptTree" node-key="id" label-key="name" default-expand-all
                                            selected-color="primary" v-model:selected="addOrEditDetail.dept_belong_id"
                                            v-if="deptTree.length !== 0" />
                                    </template>
                                </q-field>
                            </div> -->
                            <div class="col-6 q-gutter-md">
                                <q-field :label="$t('PageSystemRoleEditDialogStatus')" stack-label>
                                    <template v-slot:control>
                                        <q-option-group v-model="addOrEditDetail.status" :options="options.statusOnOff"
                                            color="primary" inline>
                                        </q-option-group>
                                    </template>
                                </q-field>
                                <q-input v-model.number="addOrEditDetail.sort" type="number"
                                    :rules="[ val => val >= 1 || $t('PageSystemRoleEditDialogSortRule')]" :label="$t('PageSystemRoleEditDialogSort')" />
                                <q-input v-model="addOrEditDetail.remark" type="textarea" :label="$t('PageSystemRoleEditDialogRemark')" />
                            </div>
                        </div>
                    </div>
                </q-form>
            </q-card-section>

            <q-separator />

            <q-card-actions align="right">
                <q-btn :label="$t('PageSystemRoleEditDialogBtnSave')" color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('PageSystemRoleEditDialogBtnCancel')" color="negative" v-close-popup />
            </q-card-actions>
        </q-card>
    </q-dialog>
</template>

<script>
import { addOrEditMixin } from 'src/mixins/addOrEditMixin'
import GqaShowName from 'src/components/GqaShowName'

export default {
    name: 'addOrEditDialog',
    mixins: [addOrEditMixin],
    components: {
        GqaShowName,
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
                roleCode: '',
                roleName: '',
            },
            url: {
                add: 'role/role-add',
                edit: 'role/role-edit',
                queryById: 'role/role-id',
            },
        }
    },
    methods: {
        resetDetail() {
            this.detail = {
                id: '',
                createdAt: '',
                createdBy: '',
                updatid: '',
                createdAt: '',
                createdBy: '',
                updatedAt: '',
                updatedBy: '',
                sort: 1,
                status: 'on',
                remark: '',
                roleCode: '',
                roleName: '',
            }
        },
    },
}
</script>

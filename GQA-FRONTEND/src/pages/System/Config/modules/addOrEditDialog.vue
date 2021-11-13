<template>
    <q-dialog v-model="addOrEditVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{formTypeName}}：
                    {{addOrEditDetail.roleName ? addOrEditDetail.roleName : $t('PageSystemConfigEditDialogTitle')}}
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
                            <q-field class="col" :label="$t('PageSystemConfigEditDialogCreatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.createdAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemConfigEditDialogCreatedBy')" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.createdByUser"
                                        :customNameObject="addOrEditDetail.createdByUser" />
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemConfigEditDialogUpdatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.updatedAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemConfigEditDialogUpdatedBy')" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.updatedByUser"
                                        :customNameObject="addOrEditDetail.updatedByUser" />
                                </template>
                            </q-field>
                        </div>
                        <div class="row">
                            <q-input class="col" v-model.number="addOrEditDetail.sort" type="number"
                                :rules="[ val => val >= 1 || $t('PageSystemConfigEditDialogSortRule')]" :label="$t('PageSystemConfigEditDialogSort')" />
                            <q-input class="col" v-model="addOrEditDetail.remark" :label="$t('PageSystemConfigEditDialogRemark')" />
                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.gqaOption" :label="$t('PageSystemConfigEditDialogOption')"
                                :rules="[ val => val && val.length > 0 || $t('PageSystemConfigEditDialogOptionRule')]" />

                            <q-input class="col" v-model="addOrEditDetail.default" :label="$t('PageSystemConfigEditDialogDefault')"
                                :rules="[ val => val && val.length > 0 || $t('PageSystemConfigEditDialogDefaultRule')]" />
                        </div>
                        <q-field :label="$t('PageSystemConfigEditDialogStatus')" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="addOrEditDetail.status" :options="options.statusOnOff"
                                    color="primary" inline>
                                </q-option-group>
                            </template>
                        </q-field>
                    </div>
                </q-form>
            </q-card-section>

            <q-separator />

            <q-card-actions align="right">
                <q-btn :label="$t('PageSystemConfigEditDialogBtnSave')" color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('PageSystemConfigEditDialogBtnCancel')" color="negative" v-close-popup />
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
                gqaOption: '',
                default: '',
                custom: '',
            },
            url: {
                add: 'config/config-add',
            },
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
                gqaOption: '',
                default: '',
                custom: '',
            }
        },
    },
}
</script>

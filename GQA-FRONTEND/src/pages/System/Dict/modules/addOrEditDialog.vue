<template>
    <q-dialog v-model="addOrEditVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{formTypeName}}：
                    {{addOrEditDetail.dictName ? addOrEditDetail.dictName : $t('PageSystemDictEditDialogTitle')}}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section style="max-height: 80vh" class="scroll q-gutter-md">
                <q-form ref="addOrEditForm">
                    <div class="q-gutter-md">
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.id" label="ID" disable />
                            <q-input class="col" v-model="addOrEditDetail.parentId" :label="$t('PageSystemDictEditDialogParent')" disable />
                        </div>
                        <div class="row">
                            <q-field class="col" :label="$t('PageSystemDictEditDialogCreatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.createdAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemDictEditDialogCreatedBy')" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.createdByUser"
                                        :customNameObject="addOrEditDetail.createdByUser" />
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemDictEditDialogUpdatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.updatedAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemDictEditDialogUpdatedBy')" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.updatedByUser"
                                        :customNameObject="addOrEditDetail.updatedByUser" />
                                </template>
                            </q-field>
                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.value" :label="$t('PageSystemDictEditDialogValue')" lazy-rules
                                :rules="[ val => val && val.length > 0 || $t('PageSystemDictEditDialogValueRule')]" />
                            <q-input class="col" v-model="addOrEditDetail.label" :label="$t('PageSystemDictEditDialogLabel')" lazy-rules
                                :rules="[ val => val && val.length > 0 || $t('PageSystemDictEditDialogLabelRule')]" />

                        </div>
                        <div class="row">
                            <q-field class="col" :label="$t('PageSystemDictEditDialogStatus')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.status" :options="options.statusOnOff"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>
                        <q-input v-model="addOrEditDetail.remark" type="textarea" :label="$t('PageSystemDictEditDialogRemark')" />
                    </div>
                </q-form>
            </q-card-section>

            <q-separator />

            <q-card-actions align="right">
                <q-btn :label="$t('PageSystemDictEditDialogBtnSave')" color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('PageSystemDictEditDialogBtnCancel')" color="negative" v-close-popup />
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
                parentId: 0,
                value: '',
                label: '',
            },
            url: {
                add: 'dict/dict-add',
                edit: 'dict/dict-edit',
                queryById: 'dict/dict-id',
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
                parentId: 0,
                value: '',
                label: '',
            }
        },
    },
}
</script>

<template>
    <q-dialog v-model="addOrEditVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{ $t('Dict') }}:
                    {{ addOrEditDetail.dictName }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section style="max-height: 80vh" class="scroll q-gutter-md">
                <q-form ref="addOrEditForm">
                    <div class="q-gutter-md">
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.id" label="ID" disable />
                            <q-input class="col" v-model.number="addOrEditDetail.sort" type="number"
                                :rules="[ val => val >= 1 || $t('SortRule')]" :label="$t('Sort')" />
                            <q-input class="col" v-model="addOrEditDetail.parentCode" :label="$t('Parent')" disable />
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
                            <q-input class="col" v-model="addOrEditDetail.dictCode" :label="$t('Dict') + $t('Code')"
                                lazy-rules :rules="[ val => val && val.length > 0 || $t('NeedInput')]" />
                            <q-input class="col" v-model="addOrEditDetail.dictLabel" :label="$t('Dict') + $t('Name')"
                                lazy-rules :rules="[ val => val && val.length > 0 || $t('NeedInput')]" />
                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.dictExt1" label="Ext1" lazy-rules />
                            <q-input class="col" v-model="addOrEditDetail.dictExt2" label="Ext2" lazy-rules />
                            <q-input class="col" v-model="addOrEditDetail.dictExt3" label="Ext3" lazy-rules />
                            <q-input class="col" v-model="addOrEditDetail.dictExt4" label="Ext4" lazy-rules />
                            <q-input class="col" v-model="addOrEditDetail.dictExt5" label="Ext5" lazy-rules />
                        </div>
                        <div class="row">
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
                </q-form>
            </q-card-section>

            <q-separator />

            <q-card-actions align="right">
                <q-btn :label="$t('Save')" color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('Cancel')" color="negative" v-close-popup />
            </q-card-actions>

            <q-inner-loading :showing="loading">
                <q-spinner-gears size="50px" color="primary" />
            </q-inner-loading>
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
                parentCode: '',
                dictCode: '',
                dictLabel: '',
                dictExt1: '',
                dictExt2: '',
                dictExt3: '',
                dictExt4: '',
                dictExt5: '',
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
                parentCode: '',
                dictCode: '',
                dictLabel: '',
                dictExt1: '',
                dictExt2: '',
                dictExt3: '',
                dictExt4: '',
                dictExt5: '',
            }
        },
    },
}
</script>

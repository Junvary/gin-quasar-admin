<template>
    <q-dialog v-model="addOrEditVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{formTypeName}}：
                    {{addOrEditDetail.dictName ? addOrEditDetail.dictName : "字典"}}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section style="max-height: 80vh" class="scroll q-gutter-md">
                <q-form ref="addOrEditForm">
                    <div class="q-gutter-md">
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.id" label="ID" disable />
                            <q-input class="col" v-model="addOrEditDetail.parentId" label="父ID" disable />
                        </div>
                        <div class="row">
                            <q-field class="col" label="创建时间" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.createdAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" label="创建人" stack-label disable>
                                <template v-slot:control>
                                    {{addOrEditDetail.createdBy}}
                                </template>
                            </q-field>
                            <q-field class="col" label="更新时间" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.updatedAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" label="更新人" stack-label disable>
                                <template v-slot:control>
                                    {{addOrEditDetail.updatedBy}}
                                </template>
                            </q-field>
                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.value" label="字典编码" lazy-rules
                                :rules="[ val => val && val.length > 0 || '必须输入字典编码']" />
                            <q-input class="col" v-model="addOrEditDetail.label" label="字典名称" lazy-rules
                                :rules="[ val => val && val.length > 0 || '必须输入字典名称']" />

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
                </q-form>
            </q-card-section>

            <q-separator />

            <q-card-actions align="right">
                <q-btn :label="'保存' + formTypeName " color="primary" @click="handleAddOrEidt" />
                <q-btn label="取消" color="negative" v-close-popup />
            </q-card-actions>
        </q-card>
    </q-dialog>
</template>

<script>
import { addOrEditMixin } from 'src/mixins/addOrEditMixin'

export default {
    name: 'addOrEditDialog',
    mixins: [addOrEditMixin],
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
<template>
    <q-dialog v-model="addOrEditVisible" position="top">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }}文档:
                    {{ addOrEditDetail.title }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section style="max-height: 80vh" class="scroll q-gutter-md">
                <q-form ref="addOrEditForm">
                    <div class="q-gutter-md">
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.id" label="ID" disable />
                            <q-input class="col" v-model.number="addOrEditDetail.sort" type="number"
                                :rules="[ val => val >= 1 || '排序必须大于0']" label="排序" />
                        </div>
                        <div class="row">
                            <q-field class="col" label="创建时间" stack-label disable>
                                <template v-slot:control>
                                    {{ showDateTime(addOrEditDetail.createdAt) }}
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
                                    {{ showDateTime(addOrEditDetail.updatedAt) }}
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
                            <q-input class="col" v-model="addOrEditDetail.title" label="标题"
                                :rules="[ val => val && val.length > 0 || '必须输入标题']" />
                            <q-field dense label="是否启用" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.status" :options="dictOptions.statusOnOff"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>
                        <q-field label="文档内容" stack-label>
                            <template v-slot:control>
                                <GqaEditor v-model="addOrEditDetail.content" />
                            </template>
                        </q-field>
                        <div class="row ">
                            <GqaUpload class="col" title="上传附件" v-model:attachment="addOrEditDetail.attachment" />
                            <q-separator vertical spaced />
                            <q-input class="col" v-model="addOrEditDetail.remark" type="textarea" label="备注" />
                        </div>

                    </div>
                </q-form>
            </q-card-section>

            <q-separator />

            <q-card-actions align="right">
                <q-btn :label="'保存' + formTypeName " color="primary" @click="handleAddOrEidt" />
                <q-btn label="取消" color="negative" v-close-popup />
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
import GqaEditor from 'src/components/GqaEditor'
import GqaUpload from 'src/components/GqaUpload'

export default {
    name: 'addOrEditDialog',
    mixins: [addOrEditMixin],
    components: {
        GqaShowName,
        GqaEditor,
        GqaUpload,
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
                remark: '',
                status: 'on',
                title: '',
                content: '',
                attachment: '',
            },
            url: {
                add: 'plugin-xk/document-add',
                edit: 'plugin-xk/document-edit',
                queryById: 'plugin-xk/document-id',
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
                remark: '',
                status: 'on',
                title: '',
                content: '',
                attachment: '',
            }
        },
    },
}
</script>
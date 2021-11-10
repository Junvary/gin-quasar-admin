<template>
    <q-dialog v-model="addOrEditVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{formTypeName}}：
                    {{addOrEditDetail.roleName ? addOrEditDetail.roleName : "角色"}}
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
                            <q-input class="col" v-model="addOrEditDetail.createAt" label="创建时间" disable />
                            <q-input class="col" v-model="addOrEditDetail.createBy" label="创建人" disable />
                            <q-input class="col" v-model=" addOrEditDetail.updateAt" label="更新时间" disable />
                            <q-input class="col" v-model="addOrEditDetail.updateBy" label="更新人" disable />
                        </div>
                        <div class="row">
                            <q-input class="col" v-model.number="addOrEditDetail.sort" type="number"
                                :rules="[ val => val >= 1 || '排序必须大于0']" label="排序" />
                            <q-input class="col" v-model="addOrEditDetail.remark" label="描述" />
                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.gqaOption" label="字段名（英）"
                                :rules="[ val => val && val.length > 0 || '必须输入字段名']" />

                            <q-input class="col" v-model="addOrEditDetail.default" label="默认值"
                                :rules="[ val => val && val.length > 0 || '必须输入默认值']" />
                        </div>
                        <q-field label="是否启用" stack-label>
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
                createAt: '',
                createBy: '',
                updateAt: '',
                updateBy: '',
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
                createAt: '',
                createBy: '',
                updateAt: '',
                updateBy: '',
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
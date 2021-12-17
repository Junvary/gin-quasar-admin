<template>
    <q-dialog v-model="addOrEditVisible" position="top">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }}最新要闻:
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
                                    {{showDateTime(addOrEditDetail.createdAt)}}
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
                                    {{showDateTime(addOrEditDetail.updatedAt)}}
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
                            <q-input class="col" v-model="addOrEditDetail.projectName" label="项目名称"
                                :rules="[ val => val && val.length > 0 || '必须输入项目名称']" />
                            <q-field dense label="是否启用" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.status" :options="dictOptions.statusOnOff"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.demand" label="需求单位" />
                            <q-select class="col" v-model="addOrEditDetail.language" :options="dictOptions.codeLanguage"
                                multiple clearable emit-value map-options label="项目语言" />
                            <q-select class="col" v-model="addOrEditDetail.node" :options="dictOptions.projectNode"
                                clearable emit-value map-options label="项目节点" />
                            <GqaSeleteUser className="col" label="牵头人" v-model:selectUser="addOrEditDetail.leader"
                                v-model:selectId="addOrEditDetail.leaderId" selection="single" />
                        </div>
                        <q-select bottom-slots label="参与人" v-model="addOrEditDetail.player" use-input use-chips multiple
                            hide-dropdown-icon input-debounce="0" new-value-mode="add-unique">
                            <template v-slot:hint>
                                使用回车确定单个人选
                            </template>
                        </q-select>
                        <q-input class="col" v-model="addOrEditDetail.remark" type="textarea" label="备注" />

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
import GqaSeleteUser from 'src/components/GqaSeleteUser'
import { getAction, postAction, putAction } from 'src/api/manage'

export default {
    name: 'addOrEditDialog',
    mixins: [addOrEditMixin],
    components: {
        GqaShowName,
        GqaSeleteUser,
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
                projectName: '',
                demand: '',
                leader: {},
                player: [],
                language: [],
                node: '',
            },
            url: {
                add: 'plugin-xk/project-add',
                edit: 'plugin-xk/project-edit',
                queryById: 'plugin-xk/project-id',
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
                projectName: '',
                demand: '',
                leader: {},
                player: [],
                language: [],
                node: '',
            }
        },
        async handleQueryById(id) {
            const res = await postAction(this.url.queryById, {
                id,
            })
            if (res.code === 1) {
                this.addOrEditDetail = res.data.records
                this.addOrEditDetail.player = this.addOrEditDetail.player.split(',')
                this.addOrEditDetail.language = this.addOrEditDetail.language.split(',')
            }
            this.loading = false
        },
        async handleAddOrEidt() {
            const success = await this.$refs.addOrEditForm.validate()
            if (success) {
                this.addOrEditDetail.player = this.addOrEditDetail.player.join(',')
                this.addOrEditDetail.language = this.addOrEditDetail.language.join(',')

                if (this.formType === 'edit') {
                    if (this.url === undefined || !this.url.edit) {
                        this.$q.notify({
                            type: 'negative',
                            message: '请先配置url',
                        })
                        return
                    }
                    const res = await putAction(this.url.edit, this.addOrEditDetail)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                        this.addOrEditVisible = false
                    }
                } else if (this.formType === 'add') {
                    if (this.url === undefined || !this.url.add) {
                        this.$q.notify({
                            type: 'negative',
                            message: '请先配置url',
                        })
                        return
                    }
                    const res = await postAction(this.url.add, this.addOrEditDetail)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                        this.addOrEditVisible = false
                    }
                } else {
                    this.$q.notify({
                        type: 'negative',
                        message: '无法新增或编辑！',
                    })
                }
                this.$emit('handleFinish')
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: '请完善表格信息！',
                })
            }
        },
    },
}
</script>
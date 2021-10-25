<template>
    <q-dialog v-model="addOrEditVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{formTypeName}}：
                    {{addOrEditDetail.nickname ? addOrEditDetail.nickname : addOrEditDetail.realName? addOrEditDetail.realName :"用户"}}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section style="max-height: 80vh" class="scroll q-gutter-md">
                <q-form ref="addOrEditForm">
                    <div class="q-gutter-md">
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.id" label="ID" disable />
                            <q-input class="col" v-model="addOrEditDetail.uuid" label="UUID" disable />
                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.createAt" label="创建时间" disable />
                            <q-input class="col" v-model="addOrEditDetail.createBy" label="创建人" disable />
                            <q-input class="col" v-model=" addOrEditDetail.updateAt" label="更新时间" disable />
                            <q-input class="col" v-model="addOrEditDetail.updateBy" label="更新人" disable />
                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.avatar" label="头像" />
                            <q-input class="col" v-model="addOrEditDetail.username" label="账号" lazy-rules
                                :rules="[ val => val && val.length > 0 || '必须输入账号']" :disable="formType === 'edit'" />
                            <q-input class="col" v-model="addOrEditDetail.nickname" label="昵称" />
                            <q-input class="col" v-model="addOrEditDetail.realName" label="真实姓名" lazy-rules
                                :rules="[ val => val && val.length > 0 || '必须输入真实姓名']" :disable="formType === 'edit'" />

                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.mobile" label="手机" />
                            <q-input class="col" v-model="addOrEditDetail.email" label="邮箱" />
                        </div>
                        <div class="row">
                            <q-field class="col" label="性别" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.gender" :options="options.gender"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                            <q-field class="col" label="是否启用" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.status" :options="options.status"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>
                        <q-input v-model="addOrEditDetail.desc" type="textarea" label="备注" />
                    </div>
                </q-form>
            </q-card-section>

            <q-separator />

            <q-card-actions align="right">
                <q-btn :label="'保存' + formTypeName " color="primary" @click="emitAddOrEdit" />
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
                sort: 0,
                status: 0,
                desc: '',
                avatar: '',
                username: '',
                nickname: '',
                realName: '',
                mobile: '',
                email: '',
                gender: 0,
            },
            addOrEditUrl: {
                add: 'user/user-add',
                edit: 'user/user-edit',
                queryById: 'user/user-id',
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
                sort: 0,
                status: 0,
                desc: '',
                avatar: '',
                username: '',
                nickname: '',
                realName: '',
                mobile: '',
                email: '',
                gender: 0,
            }
        },
    },
}
</script>
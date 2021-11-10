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
                            <q-input class="col" v-model.number="addOrEditDetail.sort" type="number"
                                :rules="[ val => val >= 1 || '排序必须大于0']" label="排序" />
                            <q-file class="col" v-model="avatarFile" label="头像" max-files="1">
                                <template v-slot:prepend>
                                    <gqa-avatar :src="addOrEditDetail.avatar" />
                                </template>
                                <template v-slot:after v-if="avatarFile">
                                    <q-btn dense flat color="primary" icon="cloud_upload" @click="handleUpload" />
                                </template>
                            </q-file>
                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.createAt" label="创建时间" disable />
                            <q-input class="col" v-model="addOrEditDetail.createBy" label="创建人" disable />
                            <q-input class="col" v-model=" addOrEditDetail.updateAt" label="更新时间" disable />
                            <q-input class="col" v-model="addOrEditDetail.updateBy" label="更新人" disable />
                        </div>
                        <div class="row">
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
                                    <q-option-group v-model="addOrEditDetail.status" :options="options.statusOnOff"
                                        color="primary" inline :disable="addOrEditDetail.username ==='admin'">
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
import GqaAvatar from 'src/components/GqaAvatar'
import { postAction } from 'src/api/manage'

export default {
    name: 'addOrEditDialog',
    mixins: [addOrEditMixin],
    components: {
        GqaAvatar,
    },
    data() {
        return {
            avatarFile: null,
            detail: {
                id: '',
                createAt: '',
                createBy: '',
                updateAt: '',
                updateBy: '',
                sort: 1,
                status: 'on',
                remark: '',
                avatar: '',
                username: '',
                nickname: '',
                realName: '',
                mobile: '',
                email: '',
                gender: 'u',
            },
            url: {
                add: 'user/user-add',
                edit: 'user/user-edit',
                queryById: 'user/user-id',
                avatarUrl: 'upload/avatar',
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
                avatar: '',
                username: '',
                nickname: '',
                realName: '',
                mobile: '',
                email: '',
                gender: 'u',
            }
        },
        chooseFile() {
            this.$refs.gqaUploader.show('头像', false)
        },
        handleUpload() {
            if (!this.avatarFile) {
                this.$q.notify({
                    type: 'negative',
                    message: '请选择文件！',
                })
                return
            }
            let form = new FormData()
            form.append('file', this.avatarFile)
            postAction(this.url.avatarUrl, form).then((res) => {
                if (res.code === 1) {
                    this.addOrEditDetail.avatar = res.data.records
                    this.avatarFile = null
                    this.$q.notify({
                        type: 'positive',
                        message: '头像上传成功！',
                    })
                }
            })
        },
    },
}
</script>
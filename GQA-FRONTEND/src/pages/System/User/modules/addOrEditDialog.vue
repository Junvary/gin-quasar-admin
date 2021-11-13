<template>
    <q-dialog v-model="addOrEditVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{formTypeName}}：
                    {{addOrEditDetail.nickname ? addOrEditDetail.nickname : addOrEditDetail.realName? addOrEditDetail.realName : $t('PageSystemUserEditDialogUserNameTitle')}}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section style="max-height: 80vh" class="scroll q-gutter-md">
                <q-form ref="addOrEditForm">
                    <div class="q-gutter-md">
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.id" label="ID" disable />
                            <q-input class="col" v-model.number="addOrEditDetail.sort" type="number"
                                :rules="[ val => val >= 1 || $t('PageSystemUserEditDialogSortRule') ]" label="$t('PageSystemUserEditDialogSort')" />
                            <q-file class="col" v-model="avatarFile" :label="$t('PageSystemUserEditDialogUserPicture')" max-files="1">
                                <template v-slot:prepend>
                                    <GqaAvatar :src="addOrEditDetail.avatar" />
                                </template>
                                <template v-slot:after v-if="avatarFile">
                                    <q-btn dense flat color="primary" icon="cloud_upload" @click="handleUpload" />
                                </template>
                            </q-file>
                        </div>
                        <div class="row">
                            <q-field class="col" :label="$t('PageSystemUserEditDialogCreatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.createdAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemUserEditDialogCreatedBy')" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.createdByUser"
                                        :customNameObject="addOrEditDetail.createdByUser" />
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemUserEditDialogUpdatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.updatedAt)}}
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemUserEditDialogUpdatedBy')" stack-label disable>
                                <template v-slot:control>
                                    <GqaShowName v-if="addOrEditDetail.updatedByUser"
                                        :customNameObject="addOrEditDetail.updatedByUser" />
                                </template>
                            </q-field>
                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.username" :label="$t('PageSystemUserEditDialogUsername')" lazy-rules
                                :rules="[ val => val && val.length > 0 || $t('PageSystemUserEditDialogUsernameRule') ]" :disable="formType === 'edit'" />
                            <q-input class="col" v-model="addOrEditDetail.nickname" :label="$t('PageSystemUserEditDialogNickname')" />
                            <q-input class="col" v-model="addOrEditDetail.realName" :label="$t('PageSystemUserEditDialogRealName')" lazy-rules
                                :rules="[ val => val && val.length > 0 || $t('PageSystemUserEditDialogRealNameRule') ]" :disable="formType === 'edit'" />

                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.mobile" :label="$t('PageSystemUserEditDialogMobile')" />
                            <q-input class="col" v-model="addOrEditDetail.email" :label="$t('PageSystemUserEditDialogEmail')" />
                        </div>
                        <div class="row">
                            <q-field class="col" :label="$t('PageSystemUserEditDialogGender')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.gender" :options="options.gender"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('PageSystemUserEditDialogStatus')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.status" :options="options.statusOnOff"
                                        color="primary" inline :disable="addOrEditDetail.username ==='admin'">
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>
                        <q-input v-model="addOrEditDetail.remark" type="textarea" :label="$t('PageSystemUserEditDialogRemark')" />
                    </div>
                </q-form>
            </q-card-section>

            <q-separator />

            <q-card-actions align="right">
                <q-btn :label="$t('PageSystemUserEditDialogBtnSave')" color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('PageSystemUserEditDialogBtnCancel')" color="negative" v-close-popup />
            </q-card-actions>
        </q-card>
    </q-dialog>
</template>

<script>
import { addOrEditMixin } from 'src/mixins/addOrEditMixin'
import { postAction } from 'src/api/manage'
import GqaAvatar from 'src/components/GqaAvatar'
import GqaShowName from 'src/components/GqaShowName'

export default {
    name: 'addOrEditDialog',
    mixins: [addOrEditMixin],
    components: {
        GqaAvatar,
        GqaShowName,
    },
    data() {
        return {
            avatarFile: null,
            detail: {
                id: '',
                createdAt: '',
                createdBy: '',
                updatedAt: '',
                updatedBy: '',
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
                createdAt: '',
                createdBy: '',
                updatedAt: '',
                updatedBy: '',
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
            this.$refs.gqaUploader.show($t('PageSystemUserEditDialogMessagePicture'), false)
        },
        handleUpload() {
            if (!this.avatarFile) {
                this.$q.notify({
                    type: 'negative',
                    message: $t('PageSystemUserEditDialogMessageSelectFile'),
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
                        message: $t('PageSystemUserEditDialogMessageAvatarUpload'),
                    })
                }
            })
        },
    },
}
</script>

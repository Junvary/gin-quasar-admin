<template>
    <q-dialog v-model="addOrEditVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{ $t('User') }}:
                    {{ addOrEditDetail.nickname ? addOrEditDetail.nickname : addOrEditDetail.realName ? addOrEditDetail.realName : "" }}
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
                            <q-file class="col" v-model="avatarFile" :label="$t('Avatar')" max-files="1"
                                @rejected="rejected" :accept="gqaBackend.avatarExt"
                                :max-file-size="gqaBackend.avatarMaxSize*1024*1024">
                                <template v-slot:prepend>
                                    <GqaAvatar :src="addOrEditDetail.avatar" />
                                </template>
                                <template v-slot:after v-if="avatarFile">
                                    <q-btn dense flat color="primary" icon="cloud_upload" @click="handleUpload" />
                                </template>
                            </q-file>
                        </div>
                        <div class="row">
                            <q-field class="col" :label="$t('CreatedAt')" stack-label disable>
                                <template v-slot:control>
                                    {{showDateTime(addOrEditDetail.createdAt)}}
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
                                    {{showDateTime(addOrEditDetail.updatedAt)}}
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
                            <q-input class="col" v-model="addOrEditDetail.username" :label="$t('Username')" lazy-rules
                                :rules="[ val => val && val.length > 0 || $t('NeedInput') ]"
                                :disable="formType === 'edit'" />
                            <q-input class="col" v-model="addOrEditDetail.nickname" :label="$t('Nickname')" />
                            <q-input class="col" v-model="addOrEditDetail.realName" :label="$t('RealName')" lazy-rules
                                :rules="[ val => val && val.length > 0 || $t('NeedInput') ]"
                                :disable="formType === 'edit'" />

                        </div>
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.mobile" :label="$t('Mobile')" />
                            <q-input class="col" v-model="addOrEditDetail.email" :label="$t('Email')" />
                        </div>
                        <div class="row">
                            <q-field class="col" :label="$t('Role')" stack-label>
                                <template v-slot:control>
                                    <span v-for="(item, index) in addOrEditDetail.role" :key="index">
                                        {{ item.roleName }};&nbsp;
                                    </span>
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('Dept')" stack-label>
                                <template v-slot:control>
                                    <span v-for="(item, index) in addOrEditDetail.dept" :key="index">
                                        {{ item.deptName }};&nbsp;
                                    </span>
                                </template>
                            </q-field>
                        </div>
                        <div class="row">
                            <q-field class="col" :label="$t('Gender')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.gender" :options="dictOptions.gender"
                                        color="primary" inline>
                                    </q-option-group>
                                </template>
                            </q-field>
                            <q-field class="col" :label="$t('Status')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="addOrEditDetail.status" :options="dictOptions.statusOnOff"
                                        color="primary" inline :disable="addOrEditDetail.username ==='admin'">
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
import { gqaBackendMixin } from 'src/mixins/gqaBackendMixin'
import { postAction } from 'src/api/manage'
import GqaAvatar from 'src/components/GqaAvatar'
import GqaShowName from 'src/components/GqaShowName'

export default {
    name: 'addOrEditDialog',
    mixins: [addOrEditMixin, gqaBackendMixin],
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
        // chooseFile() {
        //     this.$refs.gqaUploader.show($t('PageSystemUserEditDialogMessagePicture'), false)
        // },
        handleUpload() {
            if (!this.avatarFile) {
                this.$q.notify({
                    type: 'negative',
                    message: this.$t('PleaseSelectFile'),
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
                        message: this.$t('UploadSuccess'),
                    })
                }
            })
        },
        rejected(rejectedEntries) {
            this.$q.notify({
                type: 'negative',
                message: this.$t('FileRejected'),
            })
        },
    },
}
</script>

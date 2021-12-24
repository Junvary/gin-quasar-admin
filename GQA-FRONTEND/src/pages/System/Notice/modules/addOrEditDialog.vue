<template>
    <q-dialog v-model="addOrEditVisible" position="top">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }} {{$t('Notice')}}:
                    {{ addOrEditDetail.noticeTitle }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section style="max-height: 80vh" class="scroll q-gutter-md">
                <q-form ref="addOrEditForm">
                    <div class="q-gutter-md">
                        <div class="row">
                            <q-input class="col" v-model="addOrEditDetail.id" label="ID" disable />
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

                        <q-input v-model="addOrEditDetail.noticeTitle" :label="$t('Title')"
                            :rules="[ val => val && val.length > 0 || $t('NeedInput')]" />

                        <q-select v-model="addOrEditDetail.noticeType" :options="dictOptions.noticeType" emit-value
                            map-options :rules="[ val => val && val.length > 0 || $t('NeedInput')]"
                            :label="$t('Notice') + $t('Type')" />

                        <q-input v-model="addOrEditDetail.noticeContent" type="textarea" :label="$t('Content')" />

                        <q-field :label="$t('SendTo')" stack-label>
                            <template v-slot:control>
                                <q-option-group :options="noticeToUserTypeOption" name="noticeToUserType"
                                    v-model="addOrEditDetail.noticeToUserType" inline
                                    @update:model-value="changenoticeToUserType" />
                            </template>
                        </q-field>

                        <q-select v-if="this.addOrEditDetail.noticeToUserType=== 'some' && this.changeTableData.length"
                            v-model="selectUser" :options="changeTableData" multiple clearable emit-value map-options
                            :rules="[ val => val && val.length > 0 || $t('NeedInput')]" :label="$t('User')" />

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
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { addOrEditMixin } from 'src/mixins/addOrEditMixin'
import GqaShowName from 'src/components/GqaShowName'
import { getAction, postAction, putAction } from 'src/api/manage'

export default {
    name: 'addOrEditDialog',
    mixins: [addOrEditMixin, tableDataMixin],
    components: {
        GqaShowName,
    },
    computed: {
        noticeToUserTypeOption() {
            return [
                { label: this.$t('All') + this.$t('User'), value: 'all' },
                { label: this.$t('Some') + this.$t('User'), value: 'some' },
            ]
        },
        columns() {
            return [
                { name: 'username', align: 'center', label: this.$parent.$t('Username'), field: 'username' },
                { name: 'nickname', align: 'center', label: this.$parent.$t('Nickname'), field: 'nickname' },
                { name: 'realName', align: 'center', label: this.$parent.$t('RealName'), field: 'realName' },
            ]
        },
        changeTableData() {
            if (this.tableData.length) {
                for (let i of this.tableData) {
                    i.label = i.nickname ? i.nickname : i.realName ? i.realName : i.username
                    i.value = i.username
                }
                return this.tableData
            }
            return []
        },
    },
    data() {
        return {
            detail: {
                id: '',
                createdAt: '',
                createdBy: '',
                updatedAt: '',
                updatedBy: '',
                noticeTitle: '',
                noticeContent: '',
                noticeType: '',
                noticeToUserType: '',
                noticeToUser: [],
            },
            url: {
                list: 'user/user-list',
                add: 'notice/notice-add',
                edit: 'notice/notice-edit',
                queryById: 'notice/notice-id',
            },
            selectUser: [],
        }
    },
    methods: {
        show(row) {
            this.loading = true
            this.resetDetail()
            if (row && row.noticeToUserType === 'some') {
                this.changenoticeToUserType('some')
            }
            this.addOrEditDetail = Object.assign(this.detail, row)
            this.addOrEditVisible = true
            if (!this.addOrEditDetail.id) {
                this.loading = false
            } else {
                this.handleQueryById(this.addOrEditDetail.id)
            }
        },
        resetDetail() {
            this.detail = {
                id: '',
                createdAt: '',
                createdBy: '',
                updatedAt: '',
                updatedBy: '',
                noticeTitle: '',
                noticeContent: '',
                noticeType: '',
                noticeToUserType: '',
                noticeToUser: [],
            }
        },
        changenoticeToUserType(val) {
            if (val === 'some') {
                this.selectUser = []
                this.getTableData().then(() => {
                    if (this.addOrEditDetail.noticeToUserType === 'some') {
                        for (let u of this.addOrEditDetail.noticeToUser) {
                            this.selectUser.push(u.toUser)
                        }
                    }
                })
            }
        },
        async handleAddOrEidt() {
            if (!this.addOrEditDetail.noticeToUserType) {
                this.$q.notify({
                    type: 'negative',
                    message: this.$t('FixForm') + ': ' + this.$t('SendTo'),
                })
                return
            }
            const success = await this.$refs.addOrEditForm.validate()
            if (success) {
                if (this.addOrEditDetail.noticeToUserType === 'some') {
                    this.addOrEditDetail.noticeToUser = this.selectUser
                } else {
                    this.addOrEditDetail.noticeToUser = []
                }
                if (this.formType === 'edit') {
                    if (this.url === undefined || !this.url.edit) {
                        this.$q.notify({
                            type: 'negative',
                            message: this.$t('UrlNotConfig'),
                        })
                        return
                    }
                    this.addOrEditDetail.noticeToUser = []
                    this.selectUser.forEach((item) => {
                        this.addOrEditDetail.noticeToUser.push({
                            noticeId: this.addOrEditDetail.noticeId,
                            toUser: item,
                        })
                    })
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
                            message: this.$t('UrlNotConfig'),
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
                        message: this.$t('CanNotAddOrEdit'),
                    })
                }
                this.$emit('handleFinish')
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: this.$t('FixForm'),
                })
            }
        },
    },
}
</script>
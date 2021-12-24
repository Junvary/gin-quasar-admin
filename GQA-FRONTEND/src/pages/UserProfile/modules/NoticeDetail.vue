<template>
    <q-dialog v-model="addOrEditVisible" position="top" @hide="hide">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
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
                            :rules="[ val => val && val.length > 0 || $t('NeedInput')]" readonly />

                        <q-select v-model="addOrEditDetail.noticeType" :options="dictOptions.noticeType" emit-value
                            map-options :rules="[ val => val && val.length > 0 || $t('NeedInput')]"
                            :label="$t('Notice') + $t('Type')" readonly />

                        <q-input v-model="addOrEditDetail.noticeContent" type="textarea" :label="$t('Content')"
                            readonly />

                    </div>
                </q-form>
            </q-card-section>

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
    name: 'NoticeDetail',
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
                noticeTitle: '',
                noticeContent: '',
                noticeType: '',
                noticeToUserType: '',
                noticeToUser: '',
            },
            url: {
                queryById: 'notice/notice-id-read',
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
                noticeTitle: '',
                noticeContent: '',
                noticeType: '',
                noticeToUserType: '',
                noticeToUser: '',
            }
        },
        hide() {
            this.$bus.emit('noticeGetTableData')
            this.$emit('hide')
        },
    },
}
</script>
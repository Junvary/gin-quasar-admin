<template>
    <q-dialog v-model="projectDetailVisible" position="top">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="row text-h6">
                    {{ project.projectName}}
                    <q-space></q-space>
                    <q-btn label="保存" color="primary" @click="handleSave" />
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section style="max-height: 85vh" class="scroll">
                <q-form ref="projectDetailForm">
                    <q-timeline layout="comfortable" side="right" color="primary">
                        <q-timeline-entry class="project-style" :icon="checkIcon(item)" :color="checkColor(item)"
                            v-for="(item, index) in dictOptions.projectNode.filter(item=>item.dictCode.indexOf('p') !== -1)"
                            :key="index" :side="side[index]">
                            <template v-slot:title>
                                {{item.dictLabel}}
                                <span
                                    v-if="projectForm[item.dictCode].timeRange && projectForm[item.dictCode].timeRange.from"
                                    style="float: right">
                                    {{ projectForm[item.dictCode].timeRange.from }}
                                    至
                                    {{ projectForm[item.dictCode].timeRange.to }}
                                </span>

                            </template>
                            <template v-slot:subtitle>
                                <q-date v-model="projectForm[item.dictCode].timeRange" range :color="checkColor(item)"
                                    mask="YYYY-MM-DD" />
                                <q-separator spaced />
                            </template>
                            <div>
                                <span class="text-subtitle5">
                                    {{ item.remark }}
                                </span>
                                <q-input outlined v-model="projectForm[item.dictCode].content" autogrow
                                    :label="item.dictLabel + ' 备注'" />
                                <GqaUpload v-model:attachment="projectForm[item.dictCode].attachment"
                                    :color="checkColor(item)" />
                            </div>
                        </q-timeline-entry>

                    </q-timeline>

                    <q-separator spaced />

                    <q-card-section>
                        <div class="row text-h6 justify-center">
                            其他材料
                        </div>
                    </q-card-section>

                    <q-timeline layout="comfortable" side="right" color="primary">
                        <q-timeline-entry class="project-style" :icon="checkIcon(item)" :color="checkColor(item)"
                            v-for="(item, index) in dictOptions.projectNode.filter(item=>item.dictCode.indexOf('m') !== -1)"
                            :key="index" :title="item.dictLabel" :side="side[index]">
                            <template v-slot:subtitle>
                                <q-date v-model="projectForm[item.dictCode].timeRange" range :color="checkColor(item)"
                                    mask="YYYY-MM-DD" />
                                <q-separator spaced />
                            </template>
                            <div>
                                <span class="text-subtitle5">
                                    {{ item.remark }}
                                </span>
                                <q-input outlined v-model="projectForm[item.dictCode].content" autogrow
                                    :label="item.dictLabel + ' 备注'" />
                                <GqaUpload title="上传附件" v-model:attachment="projectForm[item.dictCode].attachment"
                                    :color="checkColor(item)" />
                            </div>
                        </q-timeline-entry>
                    </q-timeline>
                </q-form>
            </q-card-section>

            <q-inner-loading :showing="loading">
                <q-spinner-gears size="50px" color="primary" />
            </q-inner-loading>
        </q-card>
    </q-dialog>
</template>

<script>
import { DictOptions } from 'src/utils/dict'
import GqaUpload from 'src/components/GqaUpload'
import { putAction, postAction } from 'src/api/manage'

export default {
    name: 'ProjectDetailDialog',
    components: {
        GqaUpload,
    },
    computed: {
        checkColor() {
            return (item) => {
                if (this.projectForm[item.dictCode].content || this.projectForm[item.dictCode].attachment || (this.projectForm[item.dictCode].timeRange && this.projectForm[item.dictCode].timeRange.from && this.projectForm[item.dictCode].timeRange.to)) {
                    return 'primary'
                } else {
                    return 'grey-5'
                }
            }
        },
        checkIcon() {
            return (item) => {
                if (this.projectForm[item.dictCode].content || this.projectForm[item.dictCode].attachment || (this.projectForm[item.dictCode].timeRange && this.projectForm[item.dictCode].timeRange.from && this.projectForm[item.dictCode].timeRange.to)) {
                    return 'check'
                } else {
                    return 'question_mark'
                }
            }
        },
    },
    data() {
        return {
            loading: false,
            dictOptions: {},
            currentSide: 0,
            side: ['left', 'right'],
            projectDetailVisible: false,
            project: {},
            projectForm: {},
            url: {
                add: 'plugin-xk/project-add',
                edit: 'plugin-xk/project-edit-detail',
                queryById: 'plugin-xk/project-id',
            },
        }
    },
    methods: {
        async show(row) {
            this.loading = true
            this.projectForm = {}
            this.side = ['left', 'right']
            this.project = row
            this.dictOptions = await DictOptions()
            const projectNode = this.dictOptions.projectNode
            for (let item of projectNode) {
                this.projectForm[item.dictCode] = {
                    content: '',
                    attachment: '',
                    timeRange: {
                        from: '',
                        to: '',
                    },
                }
                this.side = this.side.concat(this.side)
            }
            this.projectDetailVisible = true
            this.handleQueryById(this.project.id)
        },
        async handleQueryById(id) {
            const res = await postAction(this.url.queryById, {
                id,
            })
            if (res.code === 1) {
                const projectDetail = res.data.records.projectDetail
                for (let item of projectDetail) {
                    this.projectForm[item.projectNode].content = item.content
                    this.projectForm[item.projectNode].attachment = item.attachment
                    this.projectForm[item.projectNode].timeRange = Object.assign({}, this.projectForm[item.projectNode].timeRange, {
                        from: item.startTime,
                        to: item.endTime,
                    })
                }
            }
            this.loading = false
        },
        async handleSave() {
            const success = await this.$refs.projectDetailForm.validate()
            if (success) {
                const projectDetail = []
                for (let key in this.projectForm) {
                    projectDetail.push({
                        projectId: this.project.projectId,
                        projectNode: key,
                        content: this.projectForm[key].content,
                        attachment: this.projectForm[key].attachment,
                        startTime: this.projectForm[key].timeRange.from,
                        endTime: this.projectForm[key].timeRange.to,
                    })
                }
                putAction(this.url.edit, {
                    projectId: this.project.projectId,
                    projectDetail: projectDetail,
                }).then((res) => {
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                        this.projectDetailVisible = false
                    }
                })
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

<style lang="scss" scoped>
::v-deep(.q-timeline__subtitle) {
    opacity: 1;
}
</style>
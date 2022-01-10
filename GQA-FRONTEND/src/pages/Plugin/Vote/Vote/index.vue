<template>
    <q-page padding>
        <div class="row">
            <div class="col-6" style="padding: 4px">
                <q-toolbar class="bg-primary text-white shadow-1">
                    <q-toolbar-title>党员投票</q-toolbar-title>
                    <q-space></q-space>
                    <q-btn push glossy color="negative" @click="handleVoteV1" v-if="canVoteV1">
                        提交本次党员投票结果
                    </q-btn>
                    <span v-else>
                        本轮已投票或没有投票权限!
                    </span>
                </q-toolbar>
                <q-form ref="candidateFormV1" v-if="canVoteV1">
                    <template v-if="candidateListV1.length">
                        <q-card v-for="(item, index) in candidateListV1" :key="index">
                            <q-card-section style="padding: 0 4px">
                                <div class="row justify-between items-center">
                                    <div class="col22">
                                        <GqaShowName :customNameObject="item.candidateByUser" />
                                    </div>
                                    <div class="col-10 row">
                                        <q-input class="col" v-for="(dict, index) in dictOptions.voteTypeDetail"
                                            :key="index" v-model.number="voteResultV1[item.candidate][dict.dictCode]"
                                            :label="dict.dictLabel" type="number"
                                            :rules="[ val => val >= 1 || '必须大于等于1']" />
                                    </div>
                                </div>
                            </q-card-section>
                        </q-card>
                    </template>
                    <span v-else>
                        党员投票暂时没有候选人!
                    </span>
                </q-form>
            </div>
            <div class="col-6" style="padding: 4px">
                <q-toolbar class="bg-primary text-white shadow-1">
                    <q-toolbar-title>管理人员投票</q-toolbar-title>
                    <q-space></q-space>
                    <q-btn push glossy color="negative" @click="handleVoteV2" v-if="canVoteV2">
                        提交本次管理人员投票结果
                    </q-btn>
                    <span v-else>
                        本轮已投票或没有投票权限!
                    </span>
                </q-toolbar>
                <q-form ref="candidateFormV2" v-if="canVoteV2">
                    <template v-if="candidateListV2.length">
                        <q-card v-for="(item, index) in candidateListV2" :key="index">
                            <q-card-section style="padding: 0 4px">
                                <div class="row justify-between items-center">
                                    <div class="col22">
                                        <GqaShowName :customNameObject="item.candidateByUser" />
                                    </div>
                                    <div class="col-10 row">
                                        <q-input class="col" v-for="(dict, index) in dictOptions.voteTypeDetail"
                                            :key="index" v-model.number="voteResultV2[item.candidate][dict.dictCode]"
                                            :label="dict.dictLabel" type="number"
                                            :rules="[ val => val >= 1 || '必须大于等于1']" />
                                    </div>
                                </div>

                            </q-card-section>
                        </q-card>
                    </template>
                    <span v-else>
                        管理人员投票暂时没有候选人!
                    </span>
                </q-form>
            </div>
            <div class="col-6">

            </div>
        </div>
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { DictOptions } from 'src/utils/dict'
import GqaShowName from 'src/components/GqaShowName'
import { postAction } from 'src/api/manage'

export default {
    name: 'Vote',
    mixins: [tableDataMixin],
    components: {
        GqaShowName,
    },
    data() {
        return {
            url: {
                list: 'plugin-vote/candidate-list',
                vote: 'plugin-vote/vote-handle',
                canVote: 'plugin-vote/can-vote',
            },
            dictOptions: {},
            candidateListV1: [],
            candidateListV2: [],
            voteResultV1: {},
            voteResultV2: {},
            canVoteV1: false,
            canVoteV2: false,
        }
    },
    async created() {
        this.dictOptions = await DictOptions()
        this.checkCanVote()
        this.changeTableData()
    },
    methods: {
        checkCanVote() {
            postAction(this.url.canVote).then((res) => {
                if (res.code === 1) {
                    this.canVoteV1 = res.data.records.v1
                    this.canVoteV2 = res.data.records.v2
                }
            })
        },
        changeTableData() {
            this.getTableData().then(() => {
                this.candidateListV1 = this.tableData.filter((item) => item.voteType === 'v1')
                this.candidateListV2 = this.tableData.filter((item) => item.voteType === 'v2')
                for (let cv1 of this.candidateListV1) {
                    for (let dict of this.dictOptions.voteTypeDetail) {
                        this.voteResultV1[cv1.candidate] = Object.assign({}, this.voteResultV1[cv1.candidate], {
                            [dict.dictCode]: 1,
                        })
                    }
                }
                for (let cv2 of this.candidateListV2) {
                    for (let dict of this.dictOptions.voteTypeDetail) {
                        this.voteResultV2[cv2.candidate] = Object.assign({}, this.voteResultV2[cv2.candidate], {
                            [dict.dictCode]: 1,
                        })
                    }
                }
            })
        },
        async handleVoteV1() {
            const success = await this.$refs.candidateFormV1.validate()
            if (success) {
                this.$q
                    .dialog({
                        title: '党员投票',
                        message: '确定提交本次党员投票结果吗？',
                        cancel: true,
                        persistent: true,
                    })
                    .onOk(async () => {
                        const voteList = []
                        for (let c in this.voteResultV1) {
                            for (let r in this.voteResultV1[c]) {
                                voteList.push({
                                    candidate: c,
                                    voteType: 'v1',
                                    voteTypeDetail: r,
                                    voteScore: this.voteResultV1[c][r],
                                })
                            }
                        }
                        this.handleVotePost(voteList, '党员投票')
                    })
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: this.$t('FixForm'),
                })
            }
        },
        async handleVoteV2() {
            const success = await this.$refs.candidateFormV2.validate()
            if (success) {
                this.$q
                    .dialog({
                        title: '管理人员投票',
                        message: '确定提交本次管理人员投票结果吗？',
                        cancel: true,
                        persistent: true,
                    })
                    .onOk(async () => {
                        const voteList = []
                        for (let c in this.voteResultV2) {
                            for (let r in this.voteResultV2[c]) {
                                voteList.push({
                                    candidate: c,
                                    voteType: 'v2',
                                    voteTypeDetail: r,
                                    voteScore: this.voteResultV2[c][r],
                                })
                            }
                        }
                        this.handleVotePost(voteList, '管理人员投票')
                    })
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: this.$t('FixForm'),
                })
            }
        },
        handleVotePost(voteList, type) {
            postAction(this.url.vote, {
                requestAddScoreDetail: voteList,
            }).then((res) => {
                if (res.code === 1) {
                    this.$q.notify({
                        type: 'positive',
                        message: '投票成功: ' + type,
                    })
                    this.checkCanVote()
                }
            })
        },
    },
}
</script>
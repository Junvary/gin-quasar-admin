<template>
    <q-page padding>
        <div class="row">
            <div class="col" style="padding: 4px">
                <q-toolbar class="bg-primary text-white shadow-1">
                    <q-toolbar-title>
                        民主评议党员投票
                        <span v-if="canVoteJiJian">
                            & 纪检委员评价
                        </span>
                        <span v-if="canVoteZhengGong">
                            & 政工干事评价
                        </span>
                    </q-toolbar-title>
                    <q-space></q-space>
                    <span v-if="canVoteDy && candidateListDy.length">
                        共:
                        {{candidateListDy.length}}
                        人&nbsp;&nbsp;&nbsp;&nbsp;
                    </span>
                    <q-btn push glossy color="negative" @click="handleVoteDy"
                        v-if="canVoteDy && candidateListDy.length">
                        提交本次投票结果
                    </q-btn>
                    <span v-else-if="!candidateListDy.length">
                        暂时没有候选人
                    </span>
                    <span v-else>
                        本轮已投票或没有投票权限!
                    </span>
                </q-toolbar>
                <q-form ref="candidateFormDy" v-if="canVoteDy">
                    <template v-if="candidateListDy.length">
                        <q-card v-for="(item, index) in candidateListDy" :key="index" bordered>
                            <q-card-section style="padding: 0 4px">
                                <div class="row justify-between items-center">
                                    <div class="col-3 row justify-center items-center">
                                        <q-chip class="glossy" color="primary" text-color="white">
                                            <GqaShowName :customNameObject="item.candidateByUser" />
                                            ({{ item.candidate }})
                                            : {{personTotalScore(voteResultDy[item.candidate])}}分
                                        </q-chip>
                                    </div>
                                    <div class="col-9 row">
                                        <q-input class="col" v-for="(dict, index) in trueVoteTypeDetailDy" :key="index"
                                            v-model.number="voteResultDy[item.candidate][dict.dictCode]"
                                            :label="dict.dictLabel" type="number"
                                            :rules="[ val => val >= 1 && val <= 100 ||'必须大于等于1，且小于等于100']" />
                                    </div>
                                </div>
                            </q-card-section>
                        </q-card>
                    </template>
                    <span v-else>
                        暂时没有候选人!
                    </span>
                </q-form>
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
    name: 'VoteDy',
    mixins: [tableDataMixin],
    components: {
        GqaShowName,
    },
    computed: {
        personTotalScore() {
            return (scoreList) => {
                let t = 0
                for (let i in scoreList) {
                    t += scoreList[i]
                }
                return t
            }
        },
    },
    data() {
        return {
            url: {
                list: 'plugin-vote/candidate-list',
                vote: 'plugin-vote/vote-handle',
                canVote: 'plugin-vote/can-vote-dy',
            },
            pagination: {
                sortBy: 'candidate',
                descending: false,
                page: 1,
                rowsPerPage: 40000,
            },
            dictOptions: {},
            candidateListDy: [],
            voteResultDy: {},
            canVoteDy: false,
            canVoteJiJian: false,
            canVoteZhengGong: false,
            trueVoteTypeDetailDy: [],
        }
    },
    async created() {
        this.dictOptions = await DictOptions()
        this.checkCanVote()
    },
    methods: {
        checkCanVote() {
            postAction(this.url.canVote).then((res) => {
                if (res.code === 1) {
                    this.canVoteDy = res.data.records.dy
                    this.canVoteJiJian = res.data.records.jiJian
                    this.canVoteZhengGong = res.data.records.zhengGong
                    if (this.canVoteDy && !this.canVoteJiJian && !this.canVoteZhengGong) {
                        this.trueVoteTypeDetailDy = this.dictOptions.voteTypeDetailDy.filter((item) => item.dictCode !== 'dy_p_zhenggong' && item.dictCode !== 'dy_p_jijian')
                    } else if (this.canVoteDy && !this.canVoteJiJian && this.canVoteZhengGong) {
                        this.trueVoteTypeDetailDy = this.dictOptions.voteTypeDetailDy.filter((item) => item.dictCode === 'dy_p_zhenggong')
                    } else if (this.canVoteDy && this.canVoteJiJian && !this.canVoteZhengGong) {
                        this.trueVoteTypeDetailDy = this.dictOptions.voteTypeDetailDy.filter((item) => item.dictCode === 'dy_p_jijian')
                    } else if (this.canVoteDy && this.canVoteJiJian && this.canVoteZhengGong) {
                        this.trueVoteTypeDetailDy = this.dictOptions.voteTypeDetailDy.filter((item) => item.dictCode === 'dy_p_zhenggong' || item.dictCode === 'dy_p_jijian')
                    } else {
                        this.trueVoteTypeDetailDy = []
                    }
                    this.changeTableData()
                }
            })
        },
        changeTableData() {
            this.getTableData().then(() => {
                this.candidateListDy = this.tableData.filter((item) => item.voteType === 'dy')
                for (let dy of this.candidateListDy) {
                    for (let dict of this.trueVoteTypeDetailDy) {
                        this.voteResultDy[dy.candidate] = Object.assign({}, this.voteResultDy[dy.candidate], {
                            [dict.dictCode]: 95,
                        })
                    }
                }
                // console.log(this.voteResultDy)
            })
        },
        async handleVoteDy() {
            const success = await this.$refs.candidateFormDy.validate()
            if (success) {
                this.$q
                    .dialog({
                        title: '民主评议党员投票',
                        message: '确定提交本次民主评议党员投票结果吗？',
                        cancel: true,
                        persistent: true,
                    })
                    .onOk(async () => {
                        const voteList = []
                        for (let c in this.voteResultDy) {
                            for (let r in this.voteResultDy[c]) {
                                voteList.push({
                                    candidate: c,
                                    voteType: 'dy',
                                    voteTypeDetail: r,
                                    voteScore: this.voteResultDy[c][r],
                                })
                            }
                        }
                        this.handleVotePost(voteList, '民主评议党员投票')
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
<template>
    <q-page padding>
        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="searchMonth" label="选择年月" />
            <q-select style="width: 20%" v-model="queryParams.voteType" :options="dictOptions.voteType" emit-value
                map-options label="投票类型" />
            <q-btn color="primary" @click="changeTableData" label="确定" />
        </div>
        <div ref="monthScoreChart" style="width: 100%; height: 500px"></div>
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="paginationList"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <span class="text-h6">
                    {{ searchMonth }}
                    <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" />
                    投票明细
                </span>
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-voteType="props">
                <q-td :props="props">
                    <GqaDictShow dictName="voteType" :dictCode="props.row.voteType" />
                </q-td>
            </template>

            <template v-slot:body-cell-voteTypeDetail="props">
                <q-td :props="props">
                    <GqaDictShow dictName="voteTypeDetail" :dictCode="props.row.voteTypeDetail" />
                </q-td>
            </template>

            <template v-slot:body-cell-candidate="props">
                <q-td :props="props">
                    <GqaShowName :customNameObject="props.row.candidateByUser" />
                </q-td>
            </template>

            <template v-slot:body-cell-voteFrom="props">
                <q-td :props="props">
                    <GqaShowName :customNameObject="props.row.voteFromByUser" />
                </q-td>
            </template>

            <template v-slot:body-cell-createdAt="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.createdAt) }}
                </q-td>
            </template>

        </q-table>
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { DictOptions } from 'src/utils/dict'
import { date } from 'quasar'
import { markRaw } from 'vue'
import { FormatDateTime } from 'src/utils/date'
import GqaDictShow from 'src/components/GqaDictShow'
import GqaShowName from 'src/components/GqaShowName'

export default {
    name: 'ScoreResult',
    mixins: [tableDataMixin],
    components: {
        GqaDictShow,
        GqaShowName,
    },
    data() {
        return {
            queryParams: {
                voteType: 'v1',
            },
            searchMonth: '',
            myChart: {},
            option: {},
            dictOptions: {},
            url: {
                list: 'plugin-vote/vote-result-list',
            },
            monthScore: {
                user: [],
            },
            monthUserScore: [],
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 10000,
            },
            paginationList: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 10,
            },
        }
    },
    computed: {
        showDateTime() {
            return (datetime) => {
                return FormatDateTime(datetime)
            }
        },
        columns() {
            return [
                { name: 'voteType', align: 'center', label: '投票类型', field: 'voteType' },
                { name: 'voteTypeDetail', align: 'center', label: '投票细类', field: 'voteTypeDetail' },
                { name: 'candidate', align: 'center', label: '候选人', field: 'candidate' },
                { name: 'voteScore', align: 'center', label: '分数', field: 'voteScore' },
                { name: 'voteFrom', align: 'center', label: '投票人', field: 'voteFrom' },
                { name: 'voteMonth', align: 'center', label: '投票周期', field: 'voteMonth' },
                { name: 'createdAt', align: 'center', label: '投票时间', field: 'createdAt' },
            ]
        },
    },
    mounted() {
        // 转化为非响应式
        this.myChart = markRaw(this.$echarts.init(this.$refs.monthScoreChart))
    },
    async created() {
        const timeStamp = Date.now()
        this.searchMonth = date.formatDate(timeStamp, 'YYYYMM')
        this.dictOptions = await DictOptions()
        this.changeTableData()
    },
    methods: {
        changeTableData() {
            this.getTableData().then(() => {
                for (let i of this.tableData) {
                    i.candidateMonth = i.candidate + i.voteMonth
                }
                this.monthScore = {
                    user: [],
                }
                this.monthUserScore = []
                this.getMonthScore(this.searchMonth)
            })
        },
        getMonthScore(month) {
            const monthScoreTemp = this.tableData.filter((item) => item.voteMonth === month)
            const onlyIdList = []
            for (let i of monthScoreTemp) {
                if (onlyIdList.indexOf(i.candidateMonth) === -1) {
                    onlyIdList.push(i.candidateMonth)
                }
            }
            for (let id of onlyIdList) {
                const oneCandidate = monthScoreTemp.filter((item) => item.candidateMonth === id)
                if (oneCandidate.length) {
                    this.monthScore.user.push(oneCandidate[0].candidateByUser.realName)
                    const detailList = []
                    for (let v of oneCandidate) {
                        this.monthScore[v.voteTypeDetail] = this.monthScore[v.voteTypeDetail] ? this.monthScore[v.voteTypeDetail] : []
                        if (detailList.indexOf(v.voteTypeDetail) === -1) {
                            detailList.push(v.voteTypeDetail)
                        }
                    }
                    for (let d of detailList) {
                        const thisDetailList = oneCandidate.filter((item) => item.voteTypeDetail === d)
                        let total = 0
                        for (let t of thisDetailList) {
                            total += t.voteScore
                        }
                        this.monthScore[d].push((total / thisDetailList.length).toFixed(2))
                    }
                }
            }

            for (let u = 0; u < this.monthScore.user.length; u++) {
                let userScore = 0
                let userNumber = 1
                for (let i in this.monthScore) {
                    if (i !== 'user') {
                        userNumber += 1
                        userScore += Number(this.monthScore[i][u])
                    }
                }
                this.monthUserScore.push((userScore / (userNumber - 1)).toFixed(2))
            }
            this.updateMonthScoreEcharts()
        },
        updateMonthScoreEcharts() {
            const series = []
            for (let dict of this.dictOptions.voteTypeDetail) {
                series.push({
                    name: dict.dictLabel,
                    type: 'bar',
                    data: this.monthScore[dict.dictCode],
                    label: {
                        show: true,
                        position: 'inside',
                        formatter: '{a}\n{c}',
                    },
                })
            }
            series.push({
                name: '平均得分',
                type: 'line',
                yAxisIndex: 1,
                data: this.monthUserScore,
                label: {
                    show: true,
                    formatter: '{c}',
                },
            })
            this.option = {
                title: {
                    text: this.searchMonth + '得分情况',
                    // left: 'center',
                },
                tooltip: {
                    trigger: 'item',
                },
                grid: {
                    left: '3%',
                    right: '3%',
                    bottom: '3%',
                    containLabel: true,
                },
                toolbox: {
                    feature: {
                        dataView: { show: true, readOnly: false },
                        magicType: { show: true, type: ['line', 'bar'] },
                        restore: { show: true },
                        saveAsImage: { show: true },
                    },
                },
                legend: {},
                xAxis: [
                    {
                        type: 'category',
                        data: this.monthScore.user,
                        axisPointer: {
                            type: 'shadow',
                        },
                    },
                ],
                yAxis: [
                    {
                        type: 'value',
                        name: '分项得分',
                        axisLabel: {
                            formatter: '{value}',
                        },
                    },
                    {
                        type: 'value',
                        name: '平均得分',
                        axisLabel: {
                            formatter: '{value}',
                        },
                    },
                ],
                series: series,
            }
            this.myChart.setOption(this.option)
        },
    },
}
</script>
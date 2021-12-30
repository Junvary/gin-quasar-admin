<template>
    <q-dialog v-model="projectDetailChartVisible">
        <q-card style="width: 1400px; max-width: 80vw;">
            <div ref="projectDetailChart" style="width: 100%; height: 500px"></div>
        </q-card>
    </q-dialog>
</template>

<script>
import { postAction } from 'src/api/manage'
import { date } from 'quasar'
import { DictOptions } from 'src/utils/dict'

export default {
    name: 'ProjectDetailChartDialog',
    data() {
        return {
            projectDetailChartVisible: false,
            myChart: {},
            option: {},
            url: {
                queryById: 'plugin-xk/project-id',
            },
            projectOptions: [],
            project: {},
            projectDetail: [],
            projectNode: [],
        }
    },
    methods: {
        show(row) {
            this.project = row
            this.projectDetailChartVisible = true
            this.handleQueryById(this.project.id)
        },
        async handleQueryById(id) {
            this.projectOptions = []
            this.projectDetail = []
            this.projectNode = []

            const projectNode = await DictOptions()
            this.projectOptions = projectNode.projectNode

            const res = await postAction(this.url.queryById, {
                id,
            })
            if (res.code === 1) {
                const detail = res.data.records.projectDetail
                if (detail && detail.length !== 0) {
                    for (let item = 0; item < detail.length; item++) {
                        if (detail[item].projectNode.indexOf('p') !== -1) {
                            this.projectDetail.push({
                                value: [item, detail[item].startTime, detail[item].endTime, date.getDateDiff(date.formatDate(detail[item].endTime, 'YYYY-MM-DD'), date.formatDate(detail[item].startTime, 'YYYY-MM-DD'), 'days')],
                            })
                            this.projectNode.push(this.projectOptions.filter((d) => d.dictCode === detail[item].projectNode)[0].dictLabel)
                        }
                    }
                }
                this.$nextTick(() => {
                    this.myChart = this.$echarts.init(this.$refs.projectDetailChart)
                    this.updateEcharts()
                })
            }
        },
        renderItem(params, api) {
            var categoryIndex = api.value(0)
            var startDate = api.coord([api.value(1), categoryIndex])
            var endDate = api.coord([api.value(2), categoryIndex])
            var height = api.size([0, 1])[1] * 0.8
            var rectShape = this.$echarts.graphic.clipRectByRect(
                {
                    x: startDate[0],
                    y: startDate[1] - height / 2,
                    //矩形的宽高
                    width: endDate[0] - startDate[0],
                    height: height,
                },
                {
                    x: params.coordSys.x,
                    y: params.coordSys.y,
                    width: params.coordSys.width,
                    height: params.coordSys.height,
                }
            )
            return {
                type: 'rect',
                shape: rectShape,
                style: api.style(),
            }
        },
        updateEcharts() {
            this.option = {
                tooltip: {
                    formatter: function (params) {
                        const node = params.name
                        const arr = params.value
                        return node + ':<br/>' + arr[1] + ' 至 ' + arr[2] + '<br/>' + '共 ' + arr[3] + ' 天'
                    },
                },
                title: {
                    text: this.project.projectName,
                    left: 'center',
                },

                //图表底板
                grid: {
                    height: 350,
                },
                xAxis: {
                    type: 'time',
                    splitLine: {
                        show: true,
                    },
                    axisLabel: {
                        rotate: 40,
                        formatter: (value, index) => {
                            return date.formatDate(value, 'YYYY-MM-DD')
                        },
                    },
                },
                yAxis: {
                    data: this.projectNode,
                    splitLine: {
                        show: true,
                    },
                },

                series: [
                    {
                        type: 'custom',
                        renderItem: this.renderItem,
                        name: '进度',
                        itemStyle: {
                            color: '#1976D2',
                        },
                        encode: {
                            x: [1, 2],
                            y: 0,
                        },
                        data: this.projectDetail,
                    },
                ],
            }
            this.myChart.setOption(this.option)
        },
    },
}
</script>
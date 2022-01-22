<template>
    <div class="gqa-weapon-code" id="gqa-weapon-code">
        <div ref="codeChart" style="width: 100%; height: 500px"></div>
    </div>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'

export default {
    name: 'PageWeaponCode',
    mixins: [tableDataMixin],
    computed: {
        myChart() {
            return this.$echarts.init(this.$refs.codeChart)
        },
    },
    data() {
        return {
            url: {
                list: 'public/plugin-xk/weapon-language-list',
            },
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 999,
            },
        }
    },
    async mounted() {
        await this.getTableData()
        this.updateEcharts()
    },
    methods: {
        updateEcharts() {
            const option = {
                title: {
                    text: '项目语言',
                    subtext: this.tableData && this.tableData.length ? '项目使用语言' : '纯属虚构',
                    left: 'center',
                },
                tooltip: {
                    trigger: 'item',
                },
                series: [
                    {
                        name: '语言',
                        type: 'pie',
                        radius: '70%',
                        data:
                            this.tableData && this.tableData.length
                                ? this.tableData
                                : [
                                      { value: 1048, name: 'Java' },
                                      { value: 735, name: 'C#' },
                                      { value: 580, name: 'C++' },
                                      { value: 484, name: 'Python' },
                                      { value: 300, name: 'Golang' },
                                  ],
                        emphasis: {
                            itemStyle: {
                                shadowBlur: 10,
                                shadowOffsetX: 0,
                                shadowColor: 'rgba(0, 0, 0, 0.5)',
                            },
                        },
                        label: {
                            show: true,
                            // formatter: '{b}:{c}',
                        },
                    },
                ],
            }
            this.myChart.setOption(option)
        },
    },
}
</script>

<style lang="scss" scoped>
.gqa-weapon-code {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 100px 0;
}
</style>
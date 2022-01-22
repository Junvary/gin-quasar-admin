<template>
    <div class="gqa-weapon-node" id="gqa-weapon-node">
        <div ref="nodeChart" style="width: 100%; height: 500px"></div>
    </div>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'

export default {
    name: 'PageWeaponNode',
    mixins: [tableDataMixin],
    computed: {
        myChart() {
            return this.$echarts.init(this.$refs.nodeChart)
        },
    },
    data() {
        return {
            url: {
                list: 'public/plugin-xk/weapon-node-list',
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
                    text: '项目节点',
                    subtext: this.tableData && this.tableData.length ? '项目当前节点' : '纯属虚构',
                    left: 'center',
                },
                tooltip: {
                    trigger: 'item',
                },
                series: [
                    {
                        name: '节点',
                        type: 'pie',
                        radius: '70%',
                        data:
                            this.tableData && this.tableData.length
                                ? this.tableData
                                : [
                                      { value: 200, name: '交付上线' },
                                      { value: 1, name: '产品下线' },
                                      { value: 30, name: '移交运维' },
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
.gqa-weapon-node {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 100px 0;
}
</style>
<template>
    <div class="gqa-weapon" id="gqa-weapon" style="background: #f3f4fe">
        <div id="weapon-chart" style="width: 100%; height: 600px"></div>
    </div>
</template>

<script>
export default {
    name: 'PageWeapon',
    data() {
        return {
            myChart: {},
            option: {},
        }
    },
    mounted() {
        this.myChart = this.$echarts.init(document.getElementById('weapon-chart'))
        this.updateEcharts()
    },
    methods: {
        updateEcharts() {
            this.option = {
                title: {
                    text: '武器库',
                    subtext: '纯属虚构',
                    left: 'center',
                },
                tooltip: {
                    trigger: 'item',
                },
                series: [
                    {
                        name: '武器',
                        type: 'pie',
                        radius: '70%',
                        data: [
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
                    },
                ],
            }
            this.myChart.setOption(this.option)
        },
    },
}
</script>

<style lang="scss" scoped>
.gqa-weapon {
    background: #f3f4fe;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 100px 0;
}
</style>
<template>
    <q-card>
        <div ref="piechart" style="height: 400px;"></div>
        <q-resize-observer @resize="onResize" />
    </q-card>
</template>

<script setup>
import { markRaw, onMounted, ref, watch } from 'vue';
import useTheme from 'src/composables/useTheme'

const { darkThemeChart } = useTheme()
// const echarts = require('echarts')
import * as echarts from "echarts";
const chart = ref(null)
const piechart = ref(null)

onMounted(() => {
    init()
})
watch(() => darkThemeChart.value, () => {
    init()
})
const init = () => {
    let ct = piechart.value;
    echarts.dispose(ct);
    chart.value = markRaw(echarts.init(ct, darkThemeChart.value));
    chart.value.setOption(options.value);
}
const onResize = () => {
    chart.value?.resize()
}

const options = ref({
    tooltip: {
        trigger: 'item'
    },
    legend: {
        top: '5%',
        left: 'center',
    },
    series: [
        {
            name: 'Access From',
            type: 'pie',
            radius: ['40%', '70%'],
            avoidLabelOverlap: false,
            label: {
                show: false,
                position: 'center'
            },
            emphasis: {
                label: {
                    show: true,
                    fontSize: '40',
                    fontWeight: 'bold'
                }
            },
            labelLine: {
                show: false
            },
            data: [
                { value: 1048, name: 'Search Engine' },
                { value: 735, name: 'Direct' },
                { value: 580, name: 'Email' },
                { value: 484, name: 'Union Ads' },
                { value: 300, name: 'Video Ads' }
            ]
        }
    ]
})

</script>

<template>
    <q-card>
        <div ref="barchart" style="height: 400px;"></div>
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
const barchart = ref(null)

onMounted(() => {
    init()
})
watch(() => darkThemeChart.value, () => {
    init()
})
const init = () => {
    let ct = barchart.value;
    echarts.dispose(ct);
    chart.value = markRaw(echarts.init(ct, darkThemeChart.value));
    chart.value.setOption(options.value);
}
const onResize = () => {
    chart.value?.resize()
}
const options = ref({
    xAxis: {
        type: 'category',
        data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },
    yAxis: {
        type: 'value'
    },
    series: [
        {
            data: [120, 200, 150, 80, 70, 110, 130],
            type: 'bar'
        }
    ]
})

</script>

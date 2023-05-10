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
    legend: {
        top: 'bottom'
    },
    toolbox: {
        show: true,
        feature: {
            mark: { show: true },
            dataView: { show: true, readOnly: false },
            restore: { show: true },
            saveAsImage: { show: true }
        }
    },
    series: [
        {
            name: 'Nightingale Chart',
            type: 'pie',
            radius: [50, 150],
            center: ['50%', '50%'],
            roseType: 'area',
            itemStyle: {
                borderRadius: 8
            },
            data: [
                { value: 40, name: 'rose 1' },
                { value: 38, name: 'rose 2' },
                { value: 32, name: 'rose 3' },
                { value: 30, name: 'rose 4' },
                { value: 28, name: 'rose 5' },
                { value: 26, name: 'rose 6' },
            ]
        }
    ]
})

</script>

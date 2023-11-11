<template>
    <q-page padding style="overflow-x: hidden;" class="q-gutter-y-md">
        <div class="row q-gutter-x-xs">
            <QStatistic bordered class="col" :label="timeWelcome()" boldLabel style="border-radius: 15px;">
                <div class="row items-center text-italic" style="margin-top: 7px;">
                    {{ randomWelcome() }}
                    <q-space />
                    <span class="text-bold">
                        <GqaShowName showMyName />
                    </span>
                    <q-inner-loading :showing="loading">
                        <q-spinner-gears size="50px" color="primary" />
                    </q-inner-loading>
                </div>
            </QStatistic>
            <QStatistic bordered class="col" label="License" boldLabel
                style="backgroundImage: linear-gradient(to bottom right, #ec4786,#b955a4);border-radius: 15px;">
                <div class="row items-center">
                    <q-icon name="fab fa-github" size="32px" dark />
                    <q-space />
                    <span>
                        {{ license }}
                    </span>
                    <q-inner-loading :showing="loading">
                        <q-spinner-gears size="50px" color="primary" />
                    </q-inner-loading>
                </div>
            </QStatistic>
            <QStatistic bordered class="col" label="Stars" boldLabel
                style="backgroundImage: linear-gradient(to bottom right, #865ec0,#5144b4);border-radius: 15px;">
                <div class="row items-center">
                    <q-icon name="star_border" size="32px" dark />
                    <q-space></q-space>
                    <img src="https://img.shields.io/github/stars/junvary/gin-quasar-admin?style=social"
                        placeholder-src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAJYAAACWBAMAAADOL2zRAAAAG1BMVEXMzMyWlpaqqqq3t7fFxcW+vr6xsbGjo6OcnJyLKnDGAAAACXBIWXMAAA7EAAAOxAGVKw4bAAABAElEQVRoge3SMW+DMBiE4YsxJqMJtHOTITPeOsLQnaodGImEUMZEkZhRUqn92f0MaTubtfeMh/QGHANEREREREREREREtIJJ0xbH299kp8l8FaGtLdTQ19HjofxZlJ0m1+eBKZcikd9PWtXC5DoDotRO04B9YOvFIXmXLy2jEbiqE6Df7DTleA5socLqvEFVxtJyrpZFWz/pHM2CVte0lS8g2eDe6prOyqPglhzROL+Xye4tmT4WvRcQ2/m81p+/rdguOi8Hc5L/8Qk4vhZzy08DduGt9eVQyP2qoTM1zi0/uf4hvBWf5c77e69Gf798y08L7j0RERERERERERH9P99ZpSVRivB/rgAAAABJRU5ErkJggg==">
                </div>
            </QStatistic>
            <QStatistic bordered class="col" label="Forks" boldLabel
                style="backgroundImage: linear-gradient(to bottom right, #56cdf3,#719de3);border-radius: 15px;">
                <div class="row items-center">
                    <q-icon name="fork_right" size="32px" dark />
                    <q-space></q-space>
                    <img src="https://img.shields.io/github/forks/junvary/gin-quasar-admin?style=social"
                        placeholder-src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAJYAAACWBAMAAADOL2zRAAAAG1BMVEXMzMyWlpaqqqq3t7fFxcW+vr6xsbGjo6OcnJyLKnDGAAAACXBIWXMAAA7EAAAOxAGVKw4bAAABAElEQVRoge3SMW+DMBiE4YsxJqMJtHOTITPeOsLQnaodGImEUMZEkZhRUqn92f0MaTubtfeMh/QGHANEREREREREREREtIJJ0xbH299kp8l8FaGtLdTQ19HjofxZlJ0m1+eBKZcikd9PWtXC5DoDotRO04B9YOvFIXmXLy2jEbiqE6Df7DTleA5socLqvEFVxtJyrpZFWz/pHM2CVte0lS8g2eDe6prOyqPglhzROL+Xye4tmT4WvRcQ2/m81p+/rdguOi8Hc5L/8Qk4vhZzy08DduGt9eVQyP2qoTM1zi0/uf4hvBWf5c77e69Gf798y08L7j0RERERERERERH9P99ZpSVRivB/rgAAAABJRU5ErkJggg==">
                </div>
            </QStatistic>
            <QStatistic bordered class="col" label="Watchers" boldLabel
                style="backgroundImage: linear-gradient(to bottom right, #fcbc25,#f68057);border-radius: 15px;">
                <div class="row items-center">
                    <q-icon name="visibility" size="32px" dark />
                    <q-space></q-space>
                    <img src="https://img.shields.io/github/watchers/junvary/gin-quasar-admin?style=social"
                        placeholder-src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAJYAAACWBAMAAADOL2zRAAAAG1BMVEXMzMyWlpaqqqq3t7fFxcW+vr6xsbGjo6OcnJyLKnDGAAAACXBIWXMAAA7EAAAOxAGVKw4bAAABAElEQVRoge3SMW+DMBiE4YsxJqMJtHOTITPeOsLQnaodGImEUMZEkZhRUqn92f0MaTubtfeMh/QGHANEREREREREREREtIJJ0xbH299kp8l8FaGtLdTQ19HjofxZlJ0m1+eBKZcikd9PWtXC5DoDotRO04B9YOvFIXmXLy2jEbiqE6Df7DTleA5socLqvEFVxtJyrpZFWz/pHM2CVte0lS8g2eDe6prOyqPglhzROL+Xye4tmT4WvRcQ2/m81p+/rdguOi8Hc5L/8Qk4vhZzy08DduGt9eVQyP2qoTM1zi0/uf4hvBWf5c77e69Gf798y08L7j0RERERERERERH9P99ZpSVRivB/rgAAAABJRU5ErkJggg==">
                </div>
            </QStatistic>
        </div>
        <div class="row q-gutter-x-md">
            <q-card class="col">
                <StarHistory />
            </q-card>
            <q-card class="col">
                <PieChart2 class="col" />
            </q-card>
            <q-card class="col">
                <PieChart class="col" />
            </q-card>
        </div>
        <div class="row q-gutter-x-md">
            <q-card class="col">
                <TimeLine />
            </q-card>
        </div>
    </q-page>
</template>

<script setup>
import GqaShowName from 'src/components/GqaShowName/index.vue'
import PieChart from "./charts/PieChart.vue";
import PieChart2 from "./charts/PieChart2.vue";
import StarHistory from './charts/StarHistory.vue'
import { timeWelcome, randomWelcome } from 'src/utils/welcome'
import TimeLine from './charts/TimeLine.vue'
import axios from 'axios'
import { onMounted, ref } from 'vue';
import { runSequentialPromises } from 'quasar'

const githubGet = axios.create()
const loading = ref(false)
const license = ref('MIT')
onMounted(() => {
    loading.value = true
    runSequentialPromises([
        () => githubGet.get("https://api.github.com/repos/junvary/gin-quasar-admin/license"),
        // 更多接口
    ]).then(resultAggregator => {
        license.value = resultAggregator[0].value.data.license.name
    }).finally(() => {
        loading.value = false
    })
})
</script>

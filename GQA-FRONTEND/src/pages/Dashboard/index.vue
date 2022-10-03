<template>
    <q-page padding style="overflow-x: hidden;" class="q-gutter-y-sm">
        <q-card bordered class="dashboard-card">
            <q-card-section>
                <div class="row items-center no-wrap">
                    <q-item class="col">
                        <q-item-section avatar>
                            <GqaAvatar loginUser size="50px" />
                        </q-item-section>
                        <q-item-section>
                            <q-item-label class="text-h6 text-bold">
                                {{ timeWelcome() }}
                                <GqaShowName showMyName />
                            </q-item-label>
                            <q-item-label caption>
                                {{ randomWelcome() }}
                            </q-item-label>
                        </q-item-section>
                    </q-item>
                    <div class="col-auto">
                        <div class="column items-center">
                            <span class="text-h8">项目数量</span>
                            <span class="text-h5">
                                21
                            </span>
                        </div>
                    </div>
                </div>
            </q-card-section>
        </q-card>

        <div class="row q-gutter-x-sm">
            <GqaCard bordered class="col" startColor="#ec4786" endColor="#b955a4">
                <q-card-section>
                    <span class="text-h6">
                        License
                    </span>
                    <div class="row items-center">
                        <q-icon name="fab fa-github" size="32px" dark />
                        <q-space></q-space>
                        <span>MIT</span>
                    </div>
                </q-card-section>
            </GqaCard>
            <GqaCard bordered class="col" startColor="#865ec0" endColor="#5144b4">
                <q-card-section>
                    <span class="text-h6">
                        Stars
                    </span>
                    <div class="row items-center">
                        <q-icon name="star_border" size="32px" dark />
                        <q-space></q-space>
                        <img src="https://img.shields.io/github/stars/junvary/gin-quasar-admin?style=social"
                            placeholder-src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAJYAAACWBAMAAADOL2zRAAAAG1BMVEXMzMyWlpaqqqq3t7fFxcW+vr6xsbGjo6OcnJyLKnDGAAAACXBIWXMAAA7EAAAOxAGVKw4bAAABAElEQVRoge3SMW+DMBiE4YsxJqMJtHOTITPeOsLQnaodGImEUMZEkZhRUqn92f0MaTubtfeMh/QGHANEREREREREREREtIJJ0xbH299kp8l8FaGtLdTQ19HjofxZlJ0m1+eBKZcikd9PWtXC5DoDotRO04B9YOvFIXmXLy2jEbiqE6Df7DTleA5socLqvEFVxtJyrpZFWz/pHM2CVte0lS8g2eDe6prOyqPglhzROL+Xye4tmT4WvRcQ2/m81p+/rdguOi8Hc5L/8Qk4vhZzy08DduGt9eVQyP2qoTM1zi0/uf4hvBWf5c77e69Gf798y08L7j0RERERERERERH9P99ZpSVRivB/rgAAAABJRU5ErkJggg==">
                    </div>
                </q-card-section>
            </GqaCard>
            <GqaCard bordered class="col" startColor="#56cdf3" endColor="#719de3">
                <q-card-section>
                    <span class="text-h6">
                        Forks
                    </span>
                    <div class="row items-center">
                        <q-icon name="fork_right" size="32px" dark />
                        <q-space></q-space>
                        <img src="https://img.shields.io/github/forks/junvary/gin-quasar-admin?style=social"
                            placeholder-src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAJYAAACWBAMAAADOL2zRAAAAG1BMVEXMzMyWlpaqqqq3t7fFxcW+vr6xsbGjo6OcnJyLKnDGAAAACXBIWXMAAA7EAAAOxAGVKw4bAAABAElEQVRoge3SMW+DMBiE4YsxJqMJtHOTITPeOsLQnaodGImEUMZEkZhRUqn92f0MaTubtfeMh/QGHANEREREREREREREtIJJ0xbH299kp8l8FaGtLdTQ19HjofxZlJ0m1+eBKZcikd9PWtXC5DoDotRO04B9YOvFIXmXLy2jEbiqE6Df7DTleA5socLqvEFVxtJyrpZFWz/pHM2CVte0lS8g2eDe6prOyqPglhzROL+Xye4tmT4WvRcQ2/m81p+/rdguOi8Hc5L/8Qk4vhZzy08DduGt9eVQyP2qoTM1zi0/uf4hvBWf5c77e69Gf798y08L7j0RERERERERERH9P99ZpSVRivB/rgAAAABJRU5ErkJggg==">
                    </div>
                </q-card-section>
            </GqaCard>
            <GqaCard bordered class="col" startColor="#fcbc25" endColor="#f68057">
                <q-card-section>
                    <span class="text-h6">
                        Watchers
                    </span>
                    <div class="row items-center">
                        <q-icon name="visibility" size="32px" dark />
                        <q-space></q-space>
                        <img src="https://img.shields.io/github/watchers/junvary/gin-quasar-admin?style=social"
                            placeholder-src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAJYAAACWBAMAAADOL2zRAAAAG1BMVEXMzMyWlpaqqqq3t7fFxcW+vr6xsbGjo6OcnJyLKnDGAAAACXBIWXMAAA7EAAAOxAGVKw4bAAABAElEQVRoge3SMW+DMBiE4YsxJqMJtHOTITPeOsLQnaodGImEUMZEkZhRUqn92f0MaTubtfeMh/QGHANEREREREREREREtIJJ0xbH299kp8l8FaGtLdTQ19HjofxZlJ0m1+eBKZcikd9PWtXC5DoDotRO04B9YOvFIXmXLy2jEbiqE6Df7DTleA5socLqvEFVxtJyrpZFWz/pHM2CVte0lS8g2eDe6prOyqPglhzROL+Xye4tmT4WvRcQ2/m81p+/rdguOi8Hc5L/8Qk4vhZzy08DduGt9eVQyP2qoTM1zi0/uf4hvBWf5c77e69Gf798y08L7j0RERERERERERH9P99ZpSVRivB/rgAAAABJRU5ErkJggg==">
                    </div>
                </q-card-section>
            </GqaCard>
        </div>

        <div class="row q-gutter-x-sm">
            <GqaCard bordered class="col">
                <PieChart />
            </GqaCard>
            <GqaCard bordered class="col">
                <BarChart />
            </GqaCard>
        </div>

        <div class="row q-gutter-x-sm">
            <GqaCard bordered class="col">
                <LineChart />
            </GqaCard>
            <GqaCard bordered class="col">
                <PieChart2 />
            </GqaCard>
        </div>

    </q-page>
</template>

<script setup>
import GqaAvatar from 'src/components/GqaAvatar/index.vue'
import GqaShowName from 'src/components/GqaShowName/index.vue'
import BarChart from "./charts/BarChart";
import PieChart from "./charts/PieChart.vue";
import PieChart2 from "./charts/PieChart2.vue";
import LineChart from "./charts/LineChart.vue";
import { timeWelcome, randomWelcome } from 'src/utils/welcome'
import GqaCard from 'src/components/GqaCard/index.vue'
</script>

<style lang="scss" scoped>
.dashboard-card {
    border-radius: 15px;
}
</style>

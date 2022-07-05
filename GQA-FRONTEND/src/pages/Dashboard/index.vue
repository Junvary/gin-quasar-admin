<template>
    <q-page padding style="overflow-x: hidden;">
        <q-card class="my-card" flat bordered>
            <q-item v-ripple>
                <q-item-section avatar>
                    <GqaAvatar loginUser size="40px" />
                </q-item-section>

                <q-item-section>
                    <q-item-label class="text-h6">
                        {{ timeWelcome() }}
                        <GqaShowName showMyName />
                    </q-item-label>
                    <q-item-label caption>
                        {{ randomWelcome() }}
                    </q-item-label>
                </q-item-section>
            </q-item>

            <q-separator />

            <q-card-section style="padding: 0;">
                <div class="row q-col-gutter-sm q-ml-xs q-mr-sm q-py-sm">
                    <div class="col-lg-6 col-md-6 col-sm-12 col-xs-12">
                        <PieChart />
                    </div>
                    <div class="col-lg-6 col-md-6 col-sm-12 col-xs-12">
                        <BarChart />
                    </div>
                    <div class="col-lg-6 col-md-6 col-sm-12 col-xs-12">
                        <LineChart />
                    </div>
                    <div class="col-lg-6 col-md-6 col-sm-12 col-xs-12">
                        <PieChart2 />
                    </div>
                </div>
            </q-card-section>
        </q-card>
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
</script>

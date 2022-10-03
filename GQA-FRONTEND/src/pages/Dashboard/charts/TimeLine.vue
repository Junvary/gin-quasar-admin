<template>
    <q-timeline layout="comfortable" color="primary">
        <q-timeline-entry v-for="(item, index) in dataList">
            <template v-slot:subtitle>
                {{showDateTime(item.commit.author.date)}}
            </template>
            <div>
                {{item.commit.message}}
            </div>
        </q-timeline-entry>
        <q-inner-loading :showing="loading">
            <q-spinner-gears size="50px" color="primary" />
        </q-inner-loading>
    </q-timeline>
</template>

<script setup>
import { onMounted, ref, computed } from 'vue';
import axios from 'axios'
import { FormatDateTime } from 'src/utils/date'

const loading = ref(false)
const githubGet = axios.create()

const dataList = ref([])
onMounted(() => {
    loading.value = true
    githubGet.get('https://api.github.com/repos/junvary/gin-quasar-admin/commits?page=0&per_page=6').then(res => {
        dataList.value = res.data
    }).finally(() => {
        loading.value = false
    })
})
const showDateTime = computed(() => {
    return (datetime) => {
        return FormatDateTime(datetime)
    }
})
</script>
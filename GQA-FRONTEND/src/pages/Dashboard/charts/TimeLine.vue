<template>
    <q-card-section>
        <q-table title="Commits" :rows="dataList" :columns="columns" row-key="name" :loading="loading" hide-bottom bordered
            separator="cell">
            <template v-slot:body-cell-author="props">
                <q-td :props="props">
                    {{ props.row.commit.author.name }}
                </q-td>
            </template>
            <template v-slot:body-cell-time="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.commit.author.date) }}
                </q-td>
            </template>
            <template v-slot:body-cell-commit="props">
                <q-td :props="props">
                    {{ props.row.commit.message }}
                </q-td>
            </template>
        </q-table>
    </q-card-section>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import axios from 'axios'
import useCommon from 'src/composables/useCommon'

const { showDateTime } = useCommon()

const loading = ref(false)
const githubGet = axios.create()

const columns = [
    { name: 'author', align: 'center', label: 'author', field: 'author' },
    { name: 'time', align: 'center', label: 'time', field: 'time' },
    { name: 'commit', align: 'left', label: 'commits', field: 'commit' },
]

const dataList = ref([])
onMounted(() => {
    loading.value = true
    githubGet.get('https://api.github.com/repos/junvary/gin-quasar-admin/commits?page=0&per_page=5').then(res => {
        dataList.value = res.data
    }).finally(() => {
        loading.value = false
    })
})
</script>
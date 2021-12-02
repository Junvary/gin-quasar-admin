<template>
    <div class="row items-center justify-center gqa-project" id="gqa-project">
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest"
            style="min-width: 50vw; max-width: 90vw">

            <template v-slot:top>
                <span class="text-h6">
                    项目进度
                </span>
                <q-space />
                <!-- <q-input dense v-model="queryParams.title" label="标题" />
                <q-btn dense color="primary" @click="handleSearch" label="搜索" style="margin: 0 10px" /> -->
                <q-btn dense round push icon="cached" @click="getTableData" />
            </template>

            <template v-slot:body-cell-leader="props">
                <q-td :props="props">
                    <GqaShowName v-if="props.row.leader" :customNameObject="props.row.leader" />
                </q-td>
            </template>

            <template v-slot:body-cell-language="props">
                <q-td :props="props">
                    <GqaDictShow dictName="codeLanguage" :dictCode="props.row.language" />
                </q-td>
            </template>

            <template v-slot:body-cell-node="props">
                <q-td :props="props">
                    <GqaDictShow dictName="projectNode" :dictCode="props.row.node" />
                </q-td>
            </template>

        </q-table>
    </div>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import GqaShowName from 'src/components/GqaShowName'
import GqaDictShow from 'src/components/GqaDictShow'

export default {
    name: 'PageNews',
    mixins: [tableDataMixin],
    components: {
        GqaShowName,
        GqaDictShow,
    },
    data() {
        return {
            showDetailVisible: false,
            detail: {},
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 10,
            },
            url: {
                list: 'public/plugin-xk/project-list',
            },
            columns: [
                { name: 'projectName', align: 'center', label: '项目名称', field: 'projectName' },
                { name: 'demand', align: 'center', label: '需求单位', field: 'demand' },
                { name: 'leader', align: 'center', label: '牵头人', field: 'leader' },
                { name: 'player', align: 'center', label: '参与人', field: 'player' },
                { name: 'language', align: 'center', label: '项目语言', field: 'language' },
                { name: 'node', align: 'center', label: '项目节点', field: 'node' },
            ],
        }
    },
    async created() {
        try {
            await this.getTableData()
        } catch (error) {
            this.tableData = [
                {
                    projectName: 'Gin-Quasar-Admin',
                    demand: 'Gin-Quasar-Admin',
                },
            ]
        }
    },
    methods: {},
}
</script>

<style lang="scss" scoped>
.gqa-project {
    background: rgba(170, 142, 245, 1);
    width: 100%;
    padding: 100px 0;
}
</style>
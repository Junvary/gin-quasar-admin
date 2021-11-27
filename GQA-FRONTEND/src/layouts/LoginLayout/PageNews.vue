<template>
    <div class="q-pa-md row items-center justify-center gqa-news" id="gqa-news">
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest" style="width: 50vw">
            <template v-slot:top>
                <span>
                    最新要闻
                </span>
                <q-space />
                <q-btn dense round push icon="cached" @click="getTableData" />
            </template>
            <template v-slot:body-cell-createdBy="props">
                <q-td :props="props">
                    <GqaShowName v-if="props.row.createdByUser" :customNameObject="props.row.createdByUser" />
                </q-td>
            </template>
        </q-table>
    </div>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import GqaShowName from 'src/components/GqaShowName'

export default {
    name: 'PageNews',
    mixins: [tableDataMixin],
    components: {
        GqaShowName,
    },
    data() {
        return {
            url: {
                list: 'public/plugin-xk/news-list',
            },
            columns: [
                { name: 'title', align: 'center', label: '新闻标题', field: 'title' },
                { name: 'createdBy', align: 'center', label: '发布人', field: 'createdBy' },
            ],
        }
    },
    created() {
        try {
            this.getTableData()
        } catch (error) {
            this.tableData = [
                {
                    title: 'Gin-Quasar-Admin',
                },
            ]
        }
    },
}
</script>

<style lang="scss" scoped>
.gqa-news {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 100px 0;
}
</style>
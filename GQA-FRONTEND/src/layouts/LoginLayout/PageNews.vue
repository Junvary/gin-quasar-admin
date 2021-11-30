<template>
    <div class="q-pa-md row items-center justify-center gqa-news" id="gqa-news">
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest"
            style="min-width: 50vw; max-width: 90vw">

            <template v-slot:top>
                <span class="text-h6">
                    最新要闻
                </span>
                <q-space />
                <q-input dense v-model="queryParams.title" label="标题" />
                <q-btn dense color="primary" @click="handleSearch" label="搜索" style="margin: 0 10px" />
                <q-btn dense round push icon="cached" @click="getTableData" />
            </template>

            <template v-slot:body-cell-title="props">
                <q-td :props="props">
                    <q-btn flat dense rounded glossy @click="showDetail(props.row)">
                        {{props.row.title}}
                    </q-btn>
                </q-td>
            </template>

            <template v-slot:body-cell-createdAt="props">
                <q-td :props="props">
                    {{showDateTime(props.row.createdAt)}}
                </q-td>
            </template>

            <template v-slot:body-cell-createdBy="props">
                <q-td :props="props">
                    <GqaShowName v-if="props.row.createdByUser" :customNameObject="props.row.createdByUser" />
                </q-td>
            </template>

        </q-table>

        <q-dialog v-model="showDetailVisible">
            <q-card style="width: 1400px; max-width: 80vw;">
                <q-card-section class="row items-center q-pb-none">
                    <div class="text-h6">
                        {{detail.title}}
                    </div>
                    <q-space />
                    <q-btn icon="close" flat round dense v-close-popup />
                </q-card-section>

                <q-separator />

                <q-card-section v-html="detail.content" style="max-height: 80vh" class="scroll" />
            </q-card>
        </q-dialog>
    </div>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import GqaShowName from 'src/components/GqaShowName'
import { FormatDataTime } from 'src/utils/date'

export default {
    name: 'PageNews',
    mixins: [tableDataMixin],
    components: {
        GqaShowName,
    },
    computed: {
        showDateTime() {
            return (datetime) => {
                return FormatDataTime(datetime)
            }
        },
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
                list: 'public/plugin-xk/news-list',
            },
            columns: [
                { name: 'title', align: 'center', label: '新闻标题', field: 'title' },
                { name: 'createdBy', align: 'center', label: '发布人', field: 'createdBy' },
                { name: 'createdAt', align: 'center', label: '发布时间', field: 'createdAt' },
            ],
        }
    },
    async created() {
        try {
            await this.getTableData()
        } catch (error) {
            this.tableData = [
                {
                    title: 'Gin-Quasar-Admin',
                    content: 'Gin-Quasar-Admin',
                },
            ]
        }
    },
    methods: {
        showDetail(row) {
            this.detail = row
            this.showDetailVisible = true
        },
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
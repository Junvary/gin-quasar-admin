<template>
    <q-page padding>
        <div class="row">
            <div class="col-4 column items-center q-gutter-md">
                <q-date minimal v-model="queryParams.inAreaTime" mask="YYYY-MM-DD" @update:model-value="changeDate" />
                <q-card style="width: 80%; max-width: 400px; padding: 10px" v-if="dayList.length">
                    <q-card-section horizontal>
                        <GqaAvatar size="150px" />
                        <q-card-section class="column items-center justify-center" style="width: 100%">
                            <q-chip class="glossy" color="orange" text-color="white" icon-right="star">
                                今日最先打卡
                            </q-chip>

                            <span>
                                {{showDateTime(dayList[0].inAreaTime)}}

                            </span>
                            <span>
                                {{dayList[0].userName}}(
                                {{dayList[0].workNumber}}
                                )

                            </span>
                        </q-card-section>
                    </q-card-section>
                </q-card>
                <q-card style="width: 80%; max-width: 400px; padding: 10px" v-if="yearFirstList.length">
                    <q-card-section horizontal>
                        <GqaAvatar size="150px" />
                        <q-card-section class="column items-center justify-center" style="width: 100%">
                            <q-chip class="glossy" color="deep-orange" text-color="white" icon-right="star">
                                本年度最先打卡
                            </q-chip>

                            <span>
                                {{showDateTime(yearFirstList[0].inAreaTime)}}

                            </span>
                            <span>
                                {{yearFirstList[0].userName}}(
                                {{yearFirstList[0].workNumber}}
                                )

                            </span>
                        </q-card-section>
                    </q-card-section>
                </q-card>
            </div>
            <div class="col">
                <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
                    <q-input style="width: 20%" v-model="queryParams.workNumber" label="工号" />
                    <q-input style="width: 20%" v-model="queryParams.userName" label="姓名" />
                    <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
                    <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
                </div>
                <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns"
                    v-model:pagination="pagination" :rows-per-page-options="pageOptions" :loading="loading"
                    @request="onRequest">

                    <template v-slot:body-cell-inAreaTime="props">
                        <q-td :props="props">
                            {{showDateTime(props.row.inAreaTime)}}
                        </q-td>
                    </template>

                    <template v-slot:body-cell-actions="props">
                        <q-td :props="props">
                            <div class="q-gutter-xs">
                                <q-btn dense color="primary" @click="personResult(props.row)" label="个人记录" />
                            </div>
                        </q-td>
                    </template>
                </q-table>
            </div>
        </div>

        <q-dialog v-model="personVisible">
            <q-spinner color="primary" size="3em" v-if="personLoading" />
            <q-card style="width: 100%; max-width: 400px; padding: 10px" v-if="personFirstList.length">
                <q-card-section horizontal>
                    <GqaAvatar size="150px" />
                    <q-card-section class="column items-center justify-center" style="width: 100%">
                        <span>
                            {{personFirstList[0].userName}}(
                            {{personFirstList[0].workNumber}}
                            )
                        </span>
                        <q-chip class="glossy" color="deep-orange" text-color="white" icon-right="star">
                            本年度最先打卡
                        </q-chip>
                        <span>
                            创造于：
                        </span>
                        <span>
                            {{showDateTime(personFirstList[0].inAreaTime)}}
                        </span>
                    </q-card-section>
                </q-card-section>
            </q-card>

        </q-dialog>

    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { FormatDateTime } from 'src/utils/date'
import GqaAvatar from 'src/components/GqaAvatar'
import { postAction } from 'src/api/manage'
import { date } from 'quasar'

export default {
    name: 'All',
    mixins: [tableDataMixin],
    components: {
        GqaAvatar,
    },
    computed: {
        showDateTime() {
            return (datetime) => {
                return FormatDateTime(datetime)
            }
        },
    },
    data() {
        return {
            queryParams: {
                inAreaTime: '',
            },
            url: {
                list: 'plugin-attendance/in-area-list',
                year: 'plugin-attendance/in-area-year',
            },
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 10,
            },
            columns: [
                { name: 'workNumber', align: 'center', label: '工号', field: 'workNumber' },
                { name: 'userName', align: 'center', label: '姓名', field: 'userName' },
                { name: 'inAreaTime', align: 'center', label: '打卡时间', field: 'inAreaTime' },
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ],
            dayList: [],
            yearFirstList: [],
            personVisible: false,
            personFirstList: [],
            personLoading: false,
        }
    },
    created() {
        const dateNow = new Date()
        this.queryParams.inAreaTime = date.formatDate(dateNow, 'YYYY-MM-DD')
        this.getDayResult()
        this.getYearResult()
    },
    methods: {
        async getDayResult() {
            this.dayList = []
            await this.getTableData()
            this.dayList = this.tableData
        },
        getYearResult() {
            this.yearFirstList = []
            postAction(this.url.year, {
                workNumber: '',
            }).then((res) => {
                if (res.code === 1) {
                    this.yearFirstList = res.data.records
                }
            })
        },
        changeDate(value) {
            this.pagination = {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 10,
            }
            this.queryParams.inAreaTime = value
            this.getDayResult()
        },
        personResult(row) {
            this.personFirstList = []
            this.personLoading = true
            this.personVisible = true
            postAction(this.url.year, {
                workNumber: row.workNumber,
            })
                .then((res) => {
                    if (res.code === 1) {
                        this.personFirstList = res.data.records
                    }
                })
                .finally(() => {
                    this.personLoading = false
                })
        },
        resetSearch() {
            this.queryParams = {}
            const dateNow = new Date()
            this.queryParams.inAreaTime = date.formatDate(dateNow, 'YYYY-MM-DD')
            this.getTableData()
        },
    },
}
</script>
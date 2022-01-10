<template>
    <q-dialog v-model="randomUserVisible" position="top">
        <q-card style="width: 1000px; max-width: 80vw;">
            <q-card-section align="center">
                <q-table row-key="id" separator="cell" :rows="tableDataBase" :columns="columns"
                    v-model:pagination="paginationBase" :rows-per-page-options="pageOptions" :loading="loadingBase"
                    @request="onRequestBase">
                    <template v-slot:top>
                        <span class="text-h6">
                            <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" />
                            : 固定投票人
                        </span>

                    </template>
                    <template v-slot:body-cell-avatar="props">
                        <q-td :props="props">
                            <GqaAvatar :src="props.row.voterByUser.avatar" />
                        </q-td>
                    </template>

                    <template v-slot:body-cell-username="props">
                        <q-td :props="props">
                            {{ props.row.voter }}
                        </q-td>
                    </template>

                    <template v-slot:body-cell-nickname="props">
                        <q-td :props="props">
                            {{ props.row.voterByUser.nickname }}
                        </q-td>
                    </template>

                    <template v-slot:body-cell-realName="props">
                        <q-td :props="props">
                            {{ props.row.voterByUser.realName }}
                        </q-td>
                    </template>

                    <template v-slot:body-cell-voteType="props">
                        <q-td :props="props">
                            <GqaDictShow dictName="voteType" :dictCode="props.row.voteType" />
                        </q-td>
                    </template>
                </q-table>
                <q-input v-model="memo" label="投票说明/备注/版本" />
                <div class="q-gutter-xs">
                    <q-btn label="仅使用固定投票人" color="primary" @click="handleSaveRandom('base')" />
                    <q-btn label="取消" color="negative" v-close-popup />
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section align="center">
                <span class="text-h6">
                    <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" />
                    : 随机投票人
                </span>
                <q-input v-model.number="queryParams.randomNumber" type="number"
                    :rules="[ val => val >= 0 || '随机投票人必须大于等于0']" label="选择随机投票人数量" />
                <div class="q-gutter-xs">
                    <q-btn label="选取随机投票人" color="primary" @click="handleRandom" />
                </div>
                <div v-if="tableData.length">
                    <q-card-section align="center">
                        <div class="text-h6">
                            系统为你随机选取了
                            {{ queryParams.randomNumber }}
                            个投票人如下:
                        </div>
                    </q-card-section>
                    <q-card-section align="center">
                        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns"
                            v-model:pagination="pagination" :rows-per-page-options="pageOptions" :loading="loading"
                            @request="onRequest">
                            <template v-slot:body-cell-avatar="props">
                                <q-td :props="props">
                                    <GqaAvatar :src="props.row.avatar" />
                                </q-td>
                            </template>

                            <template v-slot:body-cell-username="props">
                                <q-td :props="props">
                                    {{ props.row.username }}
                                </q-td>
                            </template>

                            <template v-slot:body-cell-nickname="props">
                                <q-td :props="props">
                                    {{ props.row.nickname }}
                                </q-td>
                            </template>

                            <template v-slot:body-cell-realName="props">
                                <q-td :props="props">
                                    {{ props.row.realName }}
                                </q-td>
                            </template>

                            <template v-slot:body-cell-voteType="props">
                                <q-td :props="props">
                                    <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" />
                                </q-td>
                            </template>
                        </q-table>
                    </q-card-section>
                    <q-separator spaced />
                    <q-card-actions align="center">
                        <q-btn label="确认选取固定+随机投票人" color="primary" @click="handleSaveRandom('all')" />
                        <q-btn label="取消" color="negative" v-close-popup />
                    </q-card-actions>
                </div>
            </q-card-section>

        </q-card>
    </q-dialog>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { postAction } from 'src/api/manage'
import GqaDictShow from 'src/components/GqaDictShow'
import GqaAvatar from 'src/components/GqaAvatar'

export default {
    name: 'RandomUserDialog',
    mixins: [tableDataMixin],
    components: {
        GqaDictShow,
        GqaAvatar,
    },
    computed: {
        columns() {
            return [
                { name: 'avatar', align: 'center', label: this.$t('Avatar'), field: 'avatar' },
                { name: 'username', align: 'center', label: this.$t('Username'), field: 'username' },
                { name: 'nickname', align: 'center', label: this.$t('Nickname'), field: 'nickname' },
                { name: 'realName', align: 'center', label: this.$t('RealName'), field: 'realName' },
                { name: 'voteType', align: 'center', label: '投票类型', field: 'voteType' },
            ]
        },
    },
    data() {
        return {
            loadingBase: false,
            tableDataBase: [],
            paginationBase: {
                sortBy: 'sort',
                descending: false,
                page: 1,
                rowsPerPage: 10,
            },
            queryParams: {
                randomNumber: 0,
                voteType: '',
            },
            randomUserVisible: false,
            url: {
                baseList: 'plugin-vote/voter-list',
                list: 'plugin-vote/voter-random-get',
                add: 'plugin-vote/voter-random-add',
            },
            memo: '',
        }
    },
    methods: {
        show(voteType) {
            this.queryParams.voteType = voteType
            this.queryParams.randomNumber = 0
            this.tableData = []
            this.randomUserVisible = true
            this.getTableDataBase()
        },
        async getTableDataBase() {
            await this.onRequestBase({ pagination: this.paginationBase })
        },
        async onRequestBase(props) {
            this.tableDataBase = []
            this.loadingBase = true
            // 组装分页和过滤条件
            const params = {}
            params.sortBy = props.pagination.sortBy
            params.desc = props.pagination.descending
            params.page = props.pagination.page
            params.pageSize = props.pagination.rowsPerPage
            const allParams = Object.assign({}, params, this.queryParams)
            // 带参数请求数据
            await postAction(this.url.baseList, allParams)
                .then((res) => {
                    if (res.code === 1) {
                        // 最终要把分页给同步掉
                        this.paginationBase = props.pagination
                        // 并且加入总数字段
                        this.paginationBase.rowsNumber = res.data.total
                        this.tableDataBase = res.data.records
                    }
                })
                .finally(() => {
                    this.loadingBase = false
                })
        },
        handleRandom() {
            this.getTableData()
        },
        handleSaveRandom(type) {
            const voterList = []
            for (let item of this.tableDataBase) {
                voterList.push(item.voter)
            }
            if (type === 'all') {
                for (let r of this.tableData) {
                    voterList.push(r.username)
                }
            }
            postAction(this.url.add, {
                voteType: this.queryParams.voteType,
                memo: this.memo,
                voter: voterList,
            }).then((res) => {
                if (res.code === 1) {
                    this.$q.notify({
                        type: 'positive',
                        message: '创建新版本投票人列表成功!',
                    })
                }
                this.$emit('raomdomOver')
                this.randomUserVisible = false
            })
        },
    },
}
</script>
<template>
    <q-dialog v-model="randomUserVisible" position="top">
        <q-card style="width: 80vw; max-width: 80vw;">
            <div class="row">
                <q-card-section align="center" class="col">
                    <q-table row-key="id" separator="cell" :rows="tableDataBase" :columns="columns"
                        v-model:pagination="paginationBase" :rows-per-page-options="[0]" :loading="loadingBase"
                        class="my-sticky-header-table" @request="onRequestBase">
                        <template v-slot:top>
                            <span class="row text-h6">
                                <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" />
                                : 固定投票人
                            </span>
                            <q-space></q-space>
                            <q-btn label="仅使用固定投票人" color="primary" @click="handleSaveRandom('base')" />
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
                        <template v-slot:body-cell-voteRatio="props">
                            <q-td :props="props" v-if="props.row.voteType === 'dy'">
                                <GqaDictShow dictName="voteDyRatio" :dictCode="props.row.voteRatio" withExt1 ext1="%" />
                            </q-td>
                            <q-td :props="props" v-if="props.row.voteType === 'gl'">
                                <GqaDictShow dictName="voteGlRatio" :dictCode="props.row.voteRatio" withExt1 ext1="%" />
                            </q-td>
                        </template>
                    </q-table>
                </q-card-section>
                <q-card-section align="center" class="col">
                    <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns"
                        v-model:pagination="pagination" :rows-per-page-options="[0]"
                        :class="tableData.length ? 'my-sticky-header-table2' : ''" :loading="loading"
                        @request="onRequest">
                        <template v-slot:top>
                            <span class="row text-h6">
                                <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" />
                                : 随机投票人
                            </span>
                            <q-space></q-space>
                        </template>
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

                        <template v-slot:body-cell-voteRatio="props">
                            <q-td :props="props" v-if="queryParams.voteType === 'dy'">
                                其他评议人员 40 %
                            </q-td>
                            <q-td :props="props" v-if="queryParams.voteType === 'gl'">
                                职工代表/普通职工 30%
                            </q-td>
                        </template>

                    </q-table>

                    <q-form ref="randomUserForm">
                        <q-input v-model.number="queryParams.randomNumber" type="number"
                            :rules="[ val => val > 0 || '随机投票人数量必须大于0才能抽取']" label="选择随机投票人数量" />
                    </q-form>
                    <div class="q-gutter-xs">
                        <q-btn label="选取随机投票人" color="primary" @click="handleRandom" />
                        <q-btn v-if="tableData.length" label="确认选取固定+随机投票人" color="negative"
                            @click="handleSaveRandom('all')" />
                    </div>
                </q-card-section>
            </div>
            <q-form ref="randomMemoForm" style="margin: 0 20px">
                <q-input v-model="memo" label="投票说明/备注/版本" :rules="[ val => val && val.length > 0 || $t('NeedInput') ]"
                    placeholder="xxxx年第xx期xx投票" />
            </q-form>

            <!-- <q-separator /> -->

            <div class="row justify-center items-center q-gutter-xs" style="margin: 10px 0">
                <q-btn label="取消" color="negative" v-close-popup />
            </div>

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
                { name: 'voteRatio', align: 'center', label: '投票权重', field: 'voteRatio' },
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
                rowsPerPage: 20000,
            },
            pagination: {
                sortBy: 'sort',
                descending: false,
                page: 1,
                rowsPerPage: 20000,
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
        async handleRandom() {
            const success = await this.$refs.randomUserForm.validate()
            if (success) {
                this.getTableData()
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: this.$t('FixForm'),
                })
            }
        },
        async handleSaveRandom(type) {
            const success = await this.$refs.randomMemoForm.validate()
            if (success) {
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
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: this.$t('FixForm'),
                })
            }
        },
    },
}
</script>

<style lang="scss" scoped>
.my-sticky-header-table {
    height: 75vh;
    thead tr th {
        position: sticky;
        z-index: 1;
    }

    thead tr:first-child th {
        top: 10;
    }
}
.my-sticky-header-table2 {
    height: 65vh;
    thead tr th {
        position: sticky;
        z-index: 1;
    }

    thead tr:first-child th {
        top: 10;
    }
}
</style>
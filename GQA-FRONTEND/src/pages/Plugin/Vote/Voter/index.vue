<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.voter" label="固定投票人" clearable />
            <q-select style="width: 20%" v-model="queryParams.voteType" :options="dictOptions.voteType" emit-value
                map-options label="投票类型" @update:model-value="handleSearchWithRatio" />
            <q-select style="width: 20%" v-if="queryParams.voteType === 'dy'" v-model="queryParams.voteRatio"
                :options="dictOptions.voteDyRatio" emit-value map-options label="投票权重"
                @update:model-value="handleSearch" />
            <q-select style="width: 20%" v-if="queryParams.voteType === 'gl'" v-model="queryParams.voteRatio"
                :options="dictOptions.voteGlRatio" emit-value map-options label="投票权重"
                @update:model-value="handleSearch" />
            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <div class="q-gutter-xs" v-if="queryParams.voteType === 'dy'">
                    <q-btn v-for="(item, index) in dictOptions.voteDyRatio" :key="index" dense color="primary"
                        @click="showAddUserForm(item.dictCode, item.dictLabel)">
                        新增
                        {{item.dictLabel}}({{item.dictExt1}}%)
                    </q-btn>
                </div>

                <div class="q-gutter-xs" v-if="queryParams.voteType === 'gl'">
                    <q-btn v-for="(item, index) in dictOptions.voteGlRatio" :key="index" dense color="primary"
                        @click="showAddUserForm(item.dictCode, item.dictLabel)">
                        新增
                        {{item.dictLabel}}({{item.dictExt1}}%)
                    </q-btn>
                </div>

                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
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
                <q-td :props="props" v-if="queryParams.voteType === 'dy'">
                    <GqaDictShow dictName="voteDyRatio" :dictCode="props.row.voteRatio" withExt1 ext1="%" />
                </q-td>
                <q-td :props="props" v-if="queryParams.voteType === 'gl'">
                    <GqaDictShow dictName="voteGlRatio" :dictCode="props.row.voteRatio" withExt1 ext1="%" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>

        <SelectUserDialog ref="selectUserDialog" @handleSelectUser="handleAddUser" selection="multiple"
            :title="selectRatioLabel" />
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import SelectUserDialog from 'src/components/GqaSeleteUser/SelectUserDialog'
import GqaAvatar from 'src/components/GqaAvatar'
import GqaDictShow from 'src/components/GqaDictShow'
import { postAction } from 'src/api/manage'
import { DictOptions } from 'src/utils/dict'

export default {
    name: 'Voter',
    mixins: [tableDataMixin],
    components: {
        SelectUserDialog,
        GqaAvatar,
        GqaDictShow,
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
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ]
        },
    },
    data() {
        return {
            queryParams: {
                voteType: 'dy',
                voteRatio: '',
            },
            url: {
                list: 'plugin-vote/voter-list',
                add: 'plugin-vote/voter-add',
                delete: 'plugin-vote/voter-delete',
            },
            dictOptions: {},
            selectRatio: '',
            selectRatioLabel: '',
        }
    },
    async created() {
        this.getTableData()
        this.dictOptions = await DictOptions()
    },
    methods: {
        resetSearch() {
            this.queryParams = {
                voteType: 'dy',
                voteRatio: '',
            }
            this.getTableData()
        },
        handleSearchWithRatio() {
            this.queryParams.voteRatio = ''
            this.handleSearch()
        },
        showAddUserForm(dictCode, dictLabel) {
            this.selectRatio = dictCode
            this.selectRatioLabel = dictLabel
            this.$refs.selectUserDialog.show()
        },
        handleAddUser(event) {
            const usernameList = []
            for (let i of event) {
                usernameList.push(i.username)
            }
            postAction(this.url.add, {
                voter: usernameList,
                voteType: this.queryParams.voteType,
                voteRatio: this.selectRatio,
            }).then((res) => {
                if (res.code === 1) {
                    this.$q.notify({
                        type: 'positive',
                        message: res.message,
                    })
                }
                this.getTableData()
            })
        },
    },
}
</script>

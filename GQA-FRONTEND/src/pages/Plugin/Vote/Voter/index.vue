<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.voter" label="固定投票人" />
            <q-select style="width: 20%" v-model="queryParams.voteType" :options="dictOptions.voteType" emit-value
                map-options label="投票类型" />
            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn dense color="primary" @click="showAddUserForm()">
                    新增投票人(
                    <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" />
                    )
                </q-btn>
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

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>

        <SelectUserDialog ref="selectUserDialog" @handleSelectUser="handleAddUser" selection="multiple" />
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
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ]
        },
    },
    data() {
        return {
            queryParams: {
                voteType: 'v1',
            },
            url: {
                list: 'plugin-vote/voter-list',
                add: 'plugin-vote/voter-add',
                delete: 'plugin-vote/voter-delete',
            },
            dictOptions: {},
        }
    },
    async created() {
        this.getTableData()
        this.dictOptions = await DictOptions()
    },
    methods: {
        showAddUserForm() {
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

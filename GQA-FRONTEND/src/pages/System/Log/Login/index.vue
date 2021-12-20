<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.loginUsername" :label="$t('User')" />
            <q-select style="width: 20%" v-model="queryParams.loginSuccess" :options="dictOptions.statusYesNo"
                emit-value map-options :label="$t('LoginSuccess')" />

            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-loginSuccess="props">
                <q-td :props="props">
                    <q-badge align="middle" :color="props.row.loginSuccess === 'no' ? 'negative' : 'positive'">
                        <GqaDictShow dictName="statusYesNo" :dictCode="props.row.loginSuccess" />
                    </q-badge>
                </q-td>
            </template>

            <template v-slot:body-cell-createdAt="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.createdAt) }}
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="warning" :label="$t('Detail')">
                            <q-tooltip>
                                {{ props.row.remark }}
                            </q-tooltip>
                        </q-btn>
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import GqaDictShow from 'src/components/GqaDictShow'
import { DictOptions } from 'src/utils/dict'
import { FormatDataTime } from 'src/utils/date'

export default {
    name: 'Login',
    mixins: [tableDataMixin],
    components: {
        GqaDictShow,
    },
    computed: {
        columns() {
            return [
                { name: 'id', align: 'center', label: 'ID', field: 'id' },
                { name: 'loginUsername', align: 'center', label: this.$t('User'), field: 'loginUsername' },
                { name: 'loginIp', align: 'center', label: 'IP', field: 'loginIp' },
                { name: 'loginBrowser', align: 'center', label: this.$t('Browser'), field: 'loginBrowser' },
                { name: 'loginOs', align: 'center', label: this.$t('Os'), field: 'loginOs' },
                { name: 'loginPlatform', align: 'center', label: this.$t('Platform'), field: 'loginPlatform' },
                { name: 'createdAt', align: 'center', label: this.$t('CreatedAt'), field: 'createdAt' },
                { name: 'loginSuccess', align: 'center', label: this.$t('LoginSuccess'), field: 'loginSuccess' },
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ]
        },
        showDateTime() {
            return (datetime) => {
                return FormatDataTime(datetime)
            }
        },
    },
    data() {
        return {
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 10,
            },
            url: {
                list: 'log/log-login-list',
                delete: 'log/log-login-delete',
            },
            dictOptions: {},
        }
    },
    async created() {
        this.getTableData()
        this.dictOptions = await DictOptions()
    },
}
</script>

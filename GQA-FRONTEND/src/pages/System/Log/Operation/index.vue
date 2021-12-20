<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.operationUsername" :label="$t('User')" />

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

            <template v-slot:body-cell-createdAt="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.createdAt) }}
                </q-td>
            </template>

        </q-table>
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import GqaDictShow from 'src/components/GqaDictShow'
import { DictOptions } from 'src/utils/dict'
import { FormatDateTime } from 'src/utils/date'

export default {
    name: 'Operation',
    mixins: [tableDataMixin],
    components: {
        GqaDictShow,
    },
    computed: {
        columns() {
            return [
                { name: 'id', align: 'center', label: 'ID', field: 'id' },
                { name: 'operationUsername', align: 'center', label: this.$t('User'), field: 'operationUsername' },
                { name: 'operationIp', align: 'center', label: 'IP', field: 'operationIp' },
                { name: 'operationMethod', align: 'center', label: this.$t('Method'), field: 'operationMethod' },
                { name: 'operationApi', align: 'center', label: this.$t('Api'), field: 'operationApi' },
                { name: 'operationStatus', align: 'center', label: this.$t('Status'), field: 'operationStatus' },
                { name: 'createdAt', align: 'center', label: this.$t('CreatedAt'), field: 'createdAt' },
            ]
        },
        showDateTime() {
            return (datetime) => {
                return FormatDateTime(datetime)
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
                list: 'log/log-operation-list',
                delete: 'log/log-operation-delete',
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

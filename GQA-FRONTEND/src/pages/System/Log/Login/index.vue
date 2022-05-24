<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.login_username" :label="$t('User')" />
            <q-select style="width: 20%" v-model="queryParams.login_success" :options="dictOptions.statusYesNo"
                emit-value map-options :label="$t('Login') + $t('Success')" />

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

            <template v-slot:body-cell-login_success="props">
                <q-td :props="props">
                    <q-badge align="middle" :color="props.row.login_success === 'no' ? 'negative' : 'positive'">
                        <GqaDictShow dictName="statusYesNo" :dictCode="props.row.login_success" />
                    </q-badge>
                </q-td>
            </template>

            <template v-slot:body-cell-created_at="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.created_at) }}
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="warning" :label="$t('Detail')">
                            <q-tooltip>
                                {{ props.row.memo }}
                            </q-tooltip>
                        </q-btn>
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { FormatDateTime } from 'src/utils/date'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'log/get-log-login-list',
    delete: 'log/delete-log-login-by-id',
}
const columns = computed(() => {
    return [
        { name: 'id', align: 'center', label: 'ID', field: 'id' },
        { name: 'login_username', align: 'center', label: t('User'), field: 'login_username' },
        { name: 'login_ip', align: 'center', label: 'IP', field: 'login_ip' },
        { name: 'login_browser', align: 'center', label: t('Browser'), field: 'login_browser' },
        { name: 'login_os', align: 'center', label: t('Os'), field: 'login_os' },
        { name: 'login_platform', align: 'center', label: t('Platform'), field: 'login_platform' },
        { name: 'created_at', align: 'center', label: t('CreatedAt'), field: 'created_at' },
        { name: 'login_success', align: 'center', label: t('Login') + t('Success'), field: 'login_success' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    dictOptions,
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    GqaAvatar,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    onRequest,
    getTableData,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(async () => {
    pagination.value.sortBy = 'created_at'
    pagination.value.descending = true
    getTableData()
})

const showDateTime = computed(() => {
    return (datetime) => {
        return FormatDateTime(datetime)
    }
})
</script>

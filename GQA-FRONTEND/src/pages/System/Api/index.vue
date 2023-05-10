<template>
    <q-page padding>
        <q-card flat>
            <q-card-section class="row q-gutter-x-md items-center">
                <q-input outlined dense style="width: 20%" v-model="queryParams.api_group"
                    :label="$t('Api') + $t('Group')" />
                <q-input outlined dense style="width: 20%" v-model="queryParams.api_method"
                    :label="$t('Api') + $t('Method')" />
                <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
                <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
            </q-card-section>
            <q-card-section>
                <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                    :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
                    <template v-slot:top="props">
                        <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Api')"
                            v-has="'api:add'" />
                        <q-space />
                        <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                            @click="props.toggleFullscreen" class="q-ml-md" />
                    </template>
                    <template v-slot:body-cell-status="props">
                        <q-td :props="props">
                            <GqaDictShow :dictCode="props.row.status" />
                        </q-td>
                    </template>
                    <template v-slot:body-cell-stable="props">
                        <q-td :props="props">
                            <GqaDictShow :dictCode="props.row.stable" />
                        </q-td>
                    </template>
                    <template v-slot:body-cell-actions="props">
                        <q-td :props="props" class="q-gutter-x-xs">
                            <q-btn flat dense rounded icon="eva-edit-2-outline" color="primary"
                                @click="showEditForm(props.row)" v-has="'api:edit'">
                                <q-tooltip>
                                    {{ $t('Edit') }}
                                </q-tooltip>
                            </q-btn>
                            <q-btn flat dense rounded icon="delete_outline" color="negative"
                                @click="handleDelete(props.row)" v-has="'api:delete'">
                                <q-tooltip>
                                    {{ $t('Delete') }}
                                </q-tooltip>
                            </q-btn>
                        </q-td>
                    </template>
                </q-table>
            </q-card-section>
        </q-card>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, onMounted } from 'vue'
import recordDetail from './modules/recordDetail.vue'

const url = {
    list: 'api/get-api-list',
    delete: 'api/delete-api-by-id',
}
const columns = computed(() => {
    return [
        { name: 'sort', align: 'center', label: t('Sort'), field: 'sort' },
        { name: 'api_group', align: 'center', label: t('Api') + t('Group'), field: 'api_group' },
        { name: 'api_method', align: 'center', label: t('Api') + t('Method'), field: 'api_method' },
        { name: 'api_path', align: 'center', label: t('Api') + t('Path'), field: 'api_path' },
        { name: 'memo', align: 'center', label: t('Memo'), field: 'memo' },
        { name: 'status', align: 'center', label: t('Status'), field: 'status' },
        { name: 'stable', align: 'center', label: t('Stable'), field: 'stable' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions', },
    ]
})
const {
    t,
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
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

onMounted(() => {
    getTableData()
})
</script>

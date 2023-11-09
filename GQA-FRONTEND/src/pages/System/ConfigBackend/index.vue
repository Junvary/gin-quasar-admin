<template>
    <q-page padding>
        <q-card flat>
            <q-card-section class="row q-gutter-x-md items-center">
                <q-input outlined dense style="width: 20%" v-model="queryParams.config_item"
                    :label="$t('Config') + $t('Item')" />
                <q-input outlined dense style="width: 20%" v-model="queryParams.memo" :label="$t('Memo')" />
                <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
                <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
            </q-card-section>
            <q-card-section>
                <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                    :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

                    <template v-slot:top="props">
                        <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Config')"
                            v-has="'config-backend:add'" />
                        <q-space />
                        <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                            @click="props.toggleFullscreen" class="q-ml-md" />
                    </template>

                    <template v-slot:body-cell-item_custom="props">
                        <q-td :props="props">
                            <template v-if="props.row.config_item === 'captchaKeyLong'">
                                <q-input hint="2-6" v-model.number="props.row.item_custom" dense outlined clearable
                                    type="number" :rules="[val => val && val >= 2 && val <= 6 || '2-6']" no-error-icon />
                            </template>
                            <template v-else-if="props.row.config_item === 'captchaWidth'">
                                <q-input hint="200-260" v-model.number="props.row.item_custom" dense outlined clearable
                                    type="number" :rules="[val => val && val >= 200 && val <= 260 || '200-260']"
                                    no-error-icon />
                            </template>
                            <template v-else-if="props.row.config_item === 'captchaHeight'">
                                <q-input hint="60-100" v-model.number="props.row.item_custom" dense outlined clearable
                                    type="number" :rules="[val => val && val >= 60 && val <= 100 || '60-100']"
                                    no-error-icon />
                            </template>
                            <template v-else-if="props.row.config_item === 'jwtExpiresAt'">
                                <q-input v-model.number="props.row.item_custom" dense outlined clearable type="number" />
                            </template>
                            <template v-else-if="props.row.config_item === 'jwtRefreshAt'">
                                <q-input v-model.number="props.row.item_custom" dense outlined clearable type="number" />
                            </template>
                            <template v-else-if="props.row.config_item === 'avatarMaxSize'
                                || props.row.config_item === 'logoMaxSize'
                                || props.row.config_item === 'faviconMaxSize'
                                || props.row.config_item === 'bannerImageMaxSize'">
                                <q-input hint="1-5" v-model.number="props.row.item_custom" dense outlined clearable
                                    type="number" :rules="[val => val && val >= 1 && val <= 5 || '1-5']" no-error-icon />
                            </template>
                            <template v-else-if="props.row.config_item === 'fileMaxSize'">
                                <q-input hint="1-10" v-model.number="props.row.item_custom" dense outlined clearable
                                    type="number" :rules="[val => val && val >= 1 && val <= 10 || '1-10']" no-error-icon />
                            </template>
                            <template v-else>
                                <q-input v-model="props.row.item_custom" dense outlined clearable />
                            </template>
                        </q-td>
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
                        <q-td :props="props" class="q-gutter-x-md">
                            <q-btn flat dense color="primary" @click="handleSave(props.row)" :label="$t('Save')"
                                v-has="'config-backend:save'">
                            </q-btn>
                            <q-btn flat dense color="warning" @click="handleReset(props.row)" :label="$t('Reset')"
                                v-has="'config-backend:reset'">
                            </q-btn>
                            <q-btn flat dense color="negative" :label="$t('Delete')" v-if="props.row.stable !== 'yesNo_yes'"
                                @click="handleDelete(props.row)" v-has="'config-backend:delete'">
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
import { postAction } from 'src/api/manage'
import { computed, onMounted } from 'vue'
import { useStorageStore } from 'src/stores/storage'
import recordDetail from './modules/recordDetail.vue'
import useTheme from 'src/composables/useTheme';

const storageStore = useStorageStore()
const { darkThemeSelect } = useTheme()
const url = {
    list: 'config-backend/get-config-backend-list',
    edit: 'config-backend/edit-config-backend',
    delete: 'config-backend/delete-config-backend-by-id',
}
const columns = computed(() => {
    return [
        { name: 'sort', align: 'center', label: t('Sort'), field: 'sort' },
        { name: 'config_item', align: 'center', label: t('Config') + t('Item'), field: 'config_item' },
        { name: 'memo', align: 'center', label: t('Memo'), field: 'memo' },
        { name: 'item_default', align: 'center', label: t('Default'), field: 'item_default' },
        { name: 'item_custom', align: 'center', label: t('Custom'), field: 'item_custom' },
        { name: 'status', align: 'center', label: t('Status'), field: 'status' },
        { name: 'stable', align: 'center', label: t('Stable'), field: 'stable' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    $q,
    t,
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
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

const handleReset = (row) => {
    row.item_custom = ''
}

const handleSave = async (row) => {
    row.item_custom = String(row.item_custom)
    const res = await postAction(url.edit, row)
    if (res.code === 1) {
        $q.notify({
            type: 'positive',
            message: res.message,
        })
        storageStore.SetGqaBackend()
    }
}
</script>

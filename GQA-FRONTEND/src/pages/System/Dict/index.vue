<template>
    <q-page padding>
        <q-card flat>
            <q-card-section class="row q-gutter-x-md items-center">
                <q-btn color="primary" @click="showAddParentForm()" :label="$t('Add') + $t('Parent') + $t('Dict')"
                    v-has="'dict:addParent'" />
            </q-card-section>
            <q-card-section>
                <q-hierarchy separator="cell" dense :columns="columns" :data="dictTree">
                    <template v-slot:body="props">
                        <gqa-tree-td :treeTd="props" firstTd="sort"></gqa-tree-td>
                        <td class="text-center">{{ props.item.dict_code }}</td>
                        <td class="text-center">{{ props.item.dict_label }}</td>
                        <td class="text-center">{{ props.item.dict_ext_1 }}</td>
                        <td class="text-center">{{ props.item.dict_ext_2 }}</td>
                        <td class="text-center">{{ props.item.dict_ext_3 }}</td>
                        <td class="text-center">{{ props.item.dict_ext_4 }}</td>
                        <td class="text-center">{{ props.item.dict_ext_5 }}</td>
                        <td class="text-center">
                            <GqaDictShow :dictCode="props.item.status" />
                        </td>
                        <td class="text-center">
                            <GqaDictShow :dictCode="props.item.stable" />
                        </td>
                        <td class="text-center q-gutter-x-md">
                            <q-btn flat dense color="primary" :label="$t('Edit')" @click="showEditForm(props.item)"
                                v-has="'dict:edit'">
                            </q-btn>
                            <q-btn-dropdown flat dense color="primary" :label="$t('More')" menu-anchor="bottom left"
                                menu-self="top left">
                                <q-list dense>
                                    <q-item clickable v-close-popup @click="showAddChildrenForm(props.item.dict_code)"
                                        v-has="'dict:addChildren'">
                                        <q-item-section>
                                            <q-item-label>{{ $t('Add') + $t('Children') + $t('Dict') }}</q-item-label>
                                        </q-item-section>
                                    </q-item>

                                    <q-item clickable v-close-popup @click="handleDelete(props.item)" v-has="'dict:delete'">
                                        <q-item-section>
                                            <q-item-label>{{ $t('Delete') }}</q-item-label>
                                        </q-item-section>
                                    </q-item>
                                </q-list>
                            </q-btn-dropdown>
                        </td>
                    </template>
                </q-hierarchy>
            </q-card-section>
        </q-card>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, onMounted } from 'vue'
import recordDetail from './modules/recordDetail.vue'
import { ChangeNullChildren2Array } from 'src/utils/arrayAndTree'

const url = {
    list: 'dict/get-dict-list',
    delete: 'dict/delete-dict-by-id',
}
const columns = computed(() => {
    return [
        { name: 'sort', align: 'center', label: t('Sort'), field: 'sort' },
        { name: 'dict_code', align: 'center', label: t('Dict') + t('Code'), field: 'dict_code' },
        { name: 'dict_label', align: 'center', label: t('Dict') + t('Name'), field: 'dict_label' },
        { name: 'dict_ext_1', align: 'center', label: 'Ext1', field: 'dict_ext_1' },
        { name: 'dict_ext_2', align: 'center', label: 'Ext2', field: 'dict_ext_2' },
        { name: 'dict_ext_3', align: 'center', label: 'Ext3', field: 'dict_ext_3' },
        { name: 'dict_ext_4', align: 'center', label: 'Ext4', field: 'dict_ext_4' },
        { name: 'dict_ext_5', align: 'center', label: 'Ext5', field: 'dict_ext_5' },
        { name: 'status', align: 'center', label: t('Status'), field: 'status' },
        { name: 'stable', align: 'center', label: t('Stable'), field: 'stable' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    t,
    pagination,
    GqaDictShow,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    getTableData,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    pagination.value.rowsPerPage = 99999
    getTableData()
})

const dictTree = computed(() => {
    return ChangeNullChildren2Array(tableData.value) || []
})

const showAddParentForm = () => {
    showAddForm()
}
const showAddChildrenForm = (dictCode) => {
    recordDetailDialog.value.formType = 'add'
    recordDetailDialog.value.show()
    recordDetailDialog.value.recordDetail.value.parent_code = dictCode
}
</script>


<template>
    <q-page padding>
        <q-card flat>
            <q-card-section class="row q-gutter-x-md items-center">
                <q-btn color="primary" @click="showAddParentForm()" :label="$t('Add') + $t('Parent') + $t('Menu')"
                    v-has="'menu:addParent'" />
            </q-card-section>
            <q-card-section>
                <q-hierarchy separator="cell" dense :columns="columns" :data="menuTree">
                    <template v-slot:body="props">
                        <gqa-tree-td :treeTd="props" firstTd="sort"></gqa-tree-td>
                        <td class="text-center">
                            <q-icon size="md" :name="props.item.icon" />
                        </td>
                        <td class="text-center">
                            {{ t(props.item.title) ? t(props.item.title) : props.item.title }}
                        </td>
                        <td class="text-center">
                            {{ props.item.path }}
                        </td>
                        <td class="text-center">
                            {{ props.item.redirect }}
                        </td>
                        <td class="text-center">
                            {{ props.item.component }}
                        </td>
                        <td class="text-center">
                            <GqaDictShow :dictCode="props.item.is_link" />
                        </td>
                        <td class="text-center">
                            <GqaDictShow :dictCode="props.item.hidden" />
                        </td>
                        <td class="text-center">
                            <GqaDictShow :dictCode="props.item.stable" />
                        </td>
                        <td class="text-center">
                            <div class="q-gutter-md">
                                <q-btn flat dense icon="eva-edit-2-outline" color="primary"
                                    @click="showEditForm(props.item)" :label="$t('Edit')" v-has="'menu:edit'" />
                                <q-btn flat dense icon="add" color="warning"
                                    @click="showAddChildrenForm(props.item.name)" v-has="'menu:addChildren'"
                                    :label="$t('Add') + $t('Children') + $t('Menu')" />
                                <q-btn flat dense icon="mdi-gesture-tap-button" color="positive"
                                    @click="showButtonForm(props.item)" v-has="'menu:buttonRegister'"
                                    :label="$t('Button') + $t('Register')" />
                                <q-btn flat dense icon="delete_outline" color="negative"
                                    @click="handleDelete(props.item)" :label="$t('Delete')" v-has="'menu:delete'" />
                            </div>
                        </td>
                    </template>
                </q-hierarchy>
            </q-card-section>
        </q-card>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
        <buttonDetail ref="buttonDetailDialog" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import recordDetail from './modules/recordDetail'
import { ChangeNullChildren2Array } from 'src/utils/arrayAndTree'
import buttonDetail from './modules/buttonDetail.vue'

const { t } = useI18n()
const url = {
    list: 'menu/get-menu-list',
    delete: 'menu/delete-menu-by-id',
}
const columns = computed(() => {
    return [
        { name: 'sort', align: 'center', label: t('Sort'), field: 'sort' },
        { name: 'icon', align: 'center', label: t('Icon'), field: 'icon' },
        { name: 'title', align: 'center', label: t('Name'), field: 'title' },
        { name: 'path', align: 'center', label: t('Path'), field: 'path' },
        { name: 'redirect', align: 'center', label: t('Redirect'), field: 'redirect' },
        { name: 'component', align: 'center', label: t('Component'), field: 'component' },
        { name: 'is_link', align: 'center', label: t('IsLink'), field: 'is_link' },
        { name: 'hidden', align: 'center', label: t('Hidden'), field: 'hidden' },
        { name: 'stable', align: 'center', label: t('Stable'), field: 'stable' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
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

const menuTree = computed(() => {
    return ChangeNullChildren2Array(tableData.value) || []
})

const showAddParentForm = () => {
    showAddForm()
}
const showAddChildrenForm = (name) => {
    recordDetailDialog.value.formType = 'add'
    recordDetailDialog.value.show()
    recordDetailDialog.value.recordDetail.value.parent_code = name
}

const buttonDetailDialog = ref(null)
const showButtonForm = (item) => {
    buttonDetailDialog.value.show(item)
}

</script>

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
                            {{ selectOptionLabel(props.item) }}
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
                        <td class="text-center q-gutter-x-md">
                            <q-btn flat dense color="primary" :label="$t('Edit')" @click="showEditForm(props.item)"
                                v-has="'menu:edit'">
                            </q-btn>
                            <q-btn-dropdown flat dense color="primary" :label="$t('More')" menu-anchor="bottom left"
                                menu-self="top left">
                                <q-list dense>
                                    <q-item clickable v-close-popup @click="showAddChildrenForm(props.item.name)"
                                        v-has="'menu:addChildren'">
                                        <q-item-section>
                                            <q-item-label>{{ $t('Add') + $t('Children') + $t('Menu') }}</q-item-label>
                                        </q-item-section>
                                    </q-item>

                                    <q-item clickable v-close-popup @click="showButtonForm(props.item)"
                                        v-has="'menu:buttonRegister'">
                                        <q-item-section>
                                            <q-item-label>{{ $t('Button') + $t('Register') }}</q-item-label>
                                        </q-item-section>
                                    </q-item>

                                    <q-item clickable v-close-popup @click="handleDelete(props.item)" v-has="'menu:delete'">
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
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" :menuTree="menuTree" />
        <buttonDetail ref="buttonDetailDialog" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, onMounted, ref } from 'vue'
import recordDetail from './modules/recordDetail.vue'
import { ChangeNullChildren2Array } from 'src/utils/arrayAndTree'
import buttonDetail from './modules/buttonDetail.vue'
import useCommon from 'src/composables/useCommon'

const { selectOptionLabel } = useCommon()
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

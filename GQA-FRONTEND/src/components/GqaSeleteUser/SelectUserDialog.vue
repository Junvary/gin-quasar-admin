<template>
    <q-dialog v-model="selectUserVisible">
        <q-card style="min-width: 700px; max-width: 45vw">
            <q-card-section class="row justify-between items-center">
                <div class="text-subtitle1">
                    <span v-if="title !== ''">
                        {{ title }}
                        {{ selection === "multiple" ? $t('SelectMultiple') : $t('SelectSingle') }}
                    </span>
                    <span v-else>
                        {{ $t('SelectUserComponent', {
                            oneOrMultiple: selection === "multiple" ? $t('SelectMultiple') :
                                $t('SelectSingle')
                        })
                        }}
                    </span>
                </div>
                <!-- <span class="text-subtitle2 text-negative row justify-center" v-if="selection === 'multiple'">
                    {{ $t('GqaSelectUserHelp') }}
                </span> -->
                <q-btn dense color="primary" @click="handleSelectUser()" :label="$t('Select')" />
            </q-card-section>

            <q-separator />

            <q-card-section class="items-center row q-gutter-md">
                <q-input dense outlined style="width: 20%" v-model="queryParams.username" :label="$t('Username')" />
                <q-input dense outlined style="width: 20%" v-model="queryParams.real_name" :label="$t('RealName')" />
                <q-btn dense color="primary" @click="handleSearch" :label="$t('Search')" />
                <q-btn dense color="primary" @click="resetSearch" :label="$t('Reset')" />
            </q-card-section>

            <q-table row-key="username" bordered separator="cell" :rows="tableData" :columns="columns"
                v-model:pagination="pagination" :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest"
                :selection="selection" v-model:selected="selected">
            </q-table>
        </q-card>
    </q-dialog>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, ref, toRefs } from 'vue'
import { useI18n } from 'vue-i18n'
import { ArrayOrObject } from 'src/utils/arrayOrObject'

const { t } = useI18n()
const props = defineProps({
    // multiple, single
    selection: {
        type: String,
        required: true,
    },
    title: {
        type: String,
        required: false,
        default: '',
    }
})
const { selection, title } = toRefs(props)
const url = {
    list: 'user/get-user-list',
}
const columns = computed(() => {
    return [
        { name: 'username', align: 'center', label: t('Username'), field: 'username' },
        { name: 'nickname', align: 'center', label: t('Nickname'), field: 'nickname' },
        { name: 'real_name', align: 'center', label: t('RealName'), field: 'real_name' },
    ]
})
const {
    pagination,
    queryParams,
    pageOptions,
    loading,
    tableData,
    onRequest,
    getTableData,
    handleSearch,
    resetSearch,
} = useTableData(url)

const selectUserVisible = ref(false)
const selected = ref(null)

const show = (selectUser) => {
    pagination.value.sortBy = 'username'
    pagination.value.descending = false
    selected.value = []
    selectUserVisible.value = true
    getTableData()
    // 如果需要加载已选择的用户，放开下面的注释
    // if (selection.value === 'multiple') {
    //     if (ArrayOrObject(selectUser) === 'Array') {
    //         selected.value = selectUser
    //     } else {
    //         selected.value = []
    //     }
    // } else if (selectUser) {
    //     selected.value.push(selectUser)
    // } else {
    //     selected.value = []
    // }
}

defineExpose({
    show
})
const emit = defineEmits(['handleSelectUser'])
const handleSelectUser = () => {
    emit('handleSelectUser', selected.value)
    selectUserVisible.value = false
}
</script>

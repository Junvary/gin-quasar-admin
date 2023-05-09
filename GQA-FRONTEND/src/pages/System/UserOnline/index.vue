<template>
    <q-page padding>
        <q-card flat>
            <q-card-section class="row q-gutter-x-md items-center">
                <q-input outlined dense style="width: 20%" v-model="queryParams.username" :label="$t('Username')" />
                <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
                <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
            </q-card-section>
            <q-card-section>
                <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
                    :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
                    <template v-slot:top="props">
                        <span class="text-h6">
                            {{ $t('UserOnline') }}
                        </span>
                        <q-space />
                        <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                            @click="props.toggleFullscreen" class="q-ml-md" />
                    </template>
                    <template v-slot:body-cell-avatar="props">
                        <q-td :props="props">
                            <gqa-avatar :src="props.row.avatar" />
                        </q-td>
                    </template>
                    <template v-slot:body-cell-username="props">
                        <q-td :props="props">
                            {{ props.row.user.username }}
                        </q-td>
                    </template>
                    <template v-slot:body-cell-nickname="props">
                        <q-td :props="props">
                            {{ props.row.user.nickname }}
                        </q-td>
                    </template>
                    <template v-slot:body-cell-real_name="props">
                        <q-td :props="props">
                            {{ props.row.user.real_name }}
                        </q-td>
                    </template>
                    <template v-slot:body-cell-actions="props">
                        <q-td :props="props" class="q-gutter-x-xs">
                            <q-btn flat dense rounded icon="mdi-logout-variant" color="negative"
                                @click="handleKickOut(props.row)" v-has="'user-online:kick'">
                                <q-tooltip>
                                    {{ $t('KickOut') }}
                                </q-tooltip>
                            </q-btn>
                        </q-td>
                    </template>
                </q-table>
            </q-card-section>
        </q-card>
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { postAction } from 'src/api/manage'
import { computed, onMounted } from 'vue'

const url = {
    list: 'user-online/get-user-online-list',
    kick: 'user-online/kick-online-user',
}
const columns = computed(() => {
    return [
        { name: 'avatar', align: 'center', label: t('Avatar'), field: 'avatar' },
        { name: 'username', align: 'center', label: t('Username'), field: 'username' },
        { name: 'nickname', align: 'center', label: t('Nickname'), field: 'nickname' },
        { name: 'real_name', align: 'center', label: t('RealName'), field: 'real_name' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    $q,
    t,
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

onMounted(() => {
    pagination.value.sortBy = 'username'
    getTableData()
})

const handleKickOut = (row) => {
    $q.dialog({
        title: t('KickOut') + t('User'),
        message: t('Confirm') + t('KickOut') + t('User') + '?' + '(' + row.username + ')',
        cancel: true,
        persistent: true,
    }).onOk(async () => {
        const res = await postAction(url.kick, {
            username: row.username
        })
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: res.message,
            })
        }
        getTableData()
    })
}
</script>

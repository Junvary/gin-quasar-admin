<template>
    <q-dialog v-model="rolePermissionVisible" position="top" @hide="$emit('handleFinish')">
        <q-card style="min-width: 900px; max-width: 50vw">
            <q-card-section>
                <div class="text-h6">
                    {{ $t('Role') + $t('Permission') }}: {{ row.role_name }}
                </div>
            </q-card-section>

            <q-tabs v-model="tab" dense class="text-grey" active-color="primary" indicator-color="primary" align="justify"
                narrow-indicator>
                <q-tab name="menu" :label="$t('Menu') + '/' + $t('Button') + $t('Permission')" />
                <q-tab name="api" :label="$t('Api') + $t('Permission')" />
                <q-tab name="data" :label="$t('Data') + $t('Permission')" />
            </q-tabs>

            <q-separator />

            <q-tab-panels v-model="tab" animated>
                <q-tab-panel name="menu">
                    <role-permission-menu :row="row" v-if="tab === 'menu'" @handleRoleMenuOk="show" />
                </q-tab-panel>

                <q-tab-panel name="api">
                    <role-permission-api :row="row" v-if="tab === 'api'" />
                </q-tab-panel>

                <q-tab-panel name="data">
                    <role-permission-data :row="row" v-if="tab === 'data'" />
                </q-tab-panel>
            </q-tab-panels>
        </q-card>
    </q-dialog>
</template>

<script setup>
import RolePermissionMenu from './RolePermissionMenu.vue'
import RolePermissionApi from './RolePermissionApi.vue'
import RolePermissionData from './RolePermissionData.vue'
import { ref } from 'vue';
import { postAction } from 'src/api/manage';

const url = {
    queryById: 'role/query-role-by-id'
}

const rolePermissionVisible = ref(false)
const row = ref({})
const tab = ref('menu')

const show = (record) => {
    postAction(url.queryById, {
        id: record.id
    }).then(res => {
        if (res.code === 1) {
            row.value = res.data.records
            rolePermissionVisible.value = true
        }
    })
}
defineExpose({
    show
})
</script>

<template>
    <q-dialog v-model="rolePermissionVisible" position="top">
        <q-card style="min-width: 700px; max-width: 50vw">
            <q-card-section>
                <div class="text-h6">
                    {{ $t('PageSystemRolePermissionDialogRights') }} {{row.roleName}}
                </div>
            </q-card-section>

            <q-tabs v-model="tab" dense class="text-grey" active-color="primary" indicator-color="primary"
                align="justify" narrow-indicator>
                <q-tab name="menu" :label="$t('Menu') + ' ' + $t('Permission')" />
                <q-tab name="api" :label="$t('Api') + ' ' + $t('Permission')" />
                <q-tab name="movies" label="Movies" />
            </q-tabs>

            <q-separator />

            <q-tab-panels v-model="tab" animated>
                <q-tab-panel name="menu">
                    <role-permission-menu :row="row" />
                </q-tab-panel>

                <q-tab-panel name="api">
                    <role-permission-api :row="row" />
                </q-tab-panel>

                <q-tab-panel name="movies">
                    <div class="text-h6">Movies</div>
                    Lorem ipsum dolor sit amet consectetur adipisicing elit.
                </q-tab-panel>
            </q-tab-panels>
        </q-card>
    </q-dialog>
</template>

<script>
import RolePermissionMenu from './RolePermissionMenu'
import RolePermissionApi from './RolePermissionApi'

export default {
    name: 'RolePermissionDialog',
    components: {
        RolePermissionMenu,
        RolePermissionApi,
    },
    data() {
        return {
            rolePermissionVisible: false,
            row: {},
            tab: 'menu',
        }
    },
    methods: {
        show(row) {
            this.row = row
            this.rolePermissionVisible = true
        },
    },
}
</script>

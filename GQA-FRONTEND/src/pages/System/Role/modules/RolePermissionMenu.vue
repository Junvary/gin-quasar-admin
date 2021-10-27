<template>
    <div class="column items-center">
        <div class="row justify-between" style="width: 100%">
            <q-btn color="negative" @click="handleClear">全部清空</q-btn>
            <q-btn color="negative" @click="handleAll">全部选择</q-btn>
            <q-btn color="primary" @click="handleRolePermissionMenu">保存菜单权限</q-btn>
        </div>
        <q-tree style="width: 100%" :nodes="menuTree" default-expand-all node-key="id" label-key="name"
            selected-color="primary" v-if="menuTree.length !== 0" tick-strategy="strict" v-model:ticked="ticked">
            <template v-slot:default-header="prop">
                <div class="row items-center">
                    <q-icon :name="prop.node.icon || 'share'" size="28px" class="q-mr-sm" />
                    <div class="text-weight-bold">{{ prop.node.title }}</div>
                </div>
            </template>
        </q-tree>
    </div>

</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { ArrayToTree } from 'src/utils/arrayAndTree'
import { postAction, putAction } from 'src/api/manage'

export default {
    name: 'RolePermissionMenu',
    mixins: [tableDataMixin],
    computed: {
        menuTree() {
            if (this.tableData.length !== 0) {
                return ArrayToTree(this.tableData)
            }
            return []
        },
    },
    props: {
        row: {
            type: Object,
            required: true,
        },
    },
    created() {
        this.getRoleMenuList()
    },
    data() {
        return {
            url: {
                list: 'menu/menu-list',
                roleMenuList: 'role/role-menu',
                roleMenuEdit: 'role/role-menu-edit',
            },
            ticked: [],
        }
    },
    methods: {
        getRoleMenuList() {
            // 每次获取前，清空ticked
            this.ticked = []
            postAction(this.url.roleMenuList, {
                roleCode: this.row.roleCode,
            }).then((res) => {
                if (res.code === 1) {
                    res.data.info.forEach((item) => {
                        this.ticked.push(item.MenuId)
                    })
                } else {
                    this.$q.notify({
                        type: 'negative',
                        message: res.message,
                    })
                }
            })
        },
        handleRolePermissionMenu() {
            const roleMenu = []
            for (let i of this.ticked) {
                roleMenu.push({
                    roleCode: this.row.roleCode,
                    menuId: i,
                })
            }
            console.log(roleMenu)
            putAction(this.url.roleMenuEdit, {
                roleCode: this.row.roleCode,
                roleMenu: roleMenu,
            }).then((res) => {
                if (res.code === 1) {
                    this.$q.notify({
                        type: 'positive',
                        message: res.message,
                    })
                } else {
                    this.$q.notify({
                        type: 'negative',
                        message: res.message,
                    })
                }
                this.getRoleMenuList()
            })
        },
        handleAll() {},
        handleClear() {},
    },
}
</script>
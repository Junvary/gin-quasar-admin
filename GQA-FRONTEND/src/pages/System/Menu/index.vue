<template>
    <q-page padding>
        <q-splitter v-model="splitterModel">
            <template v-slot:before>
                <div class="q-pa-md column">
                    <div class="col">
                        <q-btn label="新增菜单" color="primary" @click="showAddForm" />
                    </div>
                    <q-separator />
                    <q-tree :nodes="menuTree" default-expand-all node-key="id" label-key="name" selected-color="primary"
                        v-model:selected="selectedKey" v-if="menuTree.length !== 0" @update:selected="onSelected">
                        <template v-slot:default-header="prop">
                            <div class="row items-center">
                                <q-icon :name="prop.node.icon || 'share'" size="28px" class="q-mr-sm" />
                                <div class="text-weight-bold">{{ prop.node.title }}</div>
                            </div>
                            <q-space></q-space>
                            <gqa-status :status="prop.node.status" dense />
                            <q-btn label="删除" style="float-right" color="negative" dense
                                @click="handleDelete(prop.node)" />
                        </template>
                    </q-tree>
                    <q-inner-loading :showing="loading">
                        <q-spinner-gears size="50px" color="primary" />
                    </q-inner-loading>
                </div>
            </template>
            <template v-slot:after>
                <div class="q-pa-md">
                    <add-or-edit-card ref="addOrEditDialog" @emitAddOrEdit="emitAddOrEdit"
                        @handleFinish="handleFinish" />
                </div>
            </template>
        </q-splitter>
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import AddOrEditCard from './modules/addOrEditCard'
import { getAction, postAction, deleteAction, putAction } from 'src/api/manage'
import { menuUrl } from 'src/api/url'
import { ArrayToTree } from 'src/utils/arrayAndTree'
import GqaStatus from 'src/components/GqaStatus'

export default {
    name: 'Menu',
    mixins: [tableDataMixin],
    components: {
        AddOrEditCard,
        GqaStatus,
    },
    computed: {
        menuTree() {
            if (this.tableData.length !== 0) {
                return ArrayToTree(this.tableData)
            }
            return []
        },
    },
    data() {
        return {
            url: {
                list: 'menu/menu-list',
                delete: 'menu/menu-delete',
            },
            pagination: {
                sortBy: 'desc',
                descending: false,
                page: 1,
                rowsPerPage: 10000,
                rowsNumber: 0,
            },
            splitterModel: 30,
            selectedKey: '',
        }
    },
    created() {},
    methods: {
        showAddForm() {
            this.$refs.addOrEditDialog.onClose()
            this.$refs.addOrEditDialog.formType = 'add'
            this.$refs.addOrEditDialog.show({})
        },
        onSelected(key) {
            const row = this.tableData.filter((item) => {
                return item.id === key
            })
            if (key) {
                this.showEditForm(row[0])
            } else {
                this.$refs.addOrEditDialog.onClose()
            }
        },
        handleFinish() {
            this.selectedKey = ''
            this.getTableData()
        },
        handleDelete(row) {
            this.$q
                .dialog({
                    title: '确定删除？',
                    message: `你确定要删除此项吗？`,
                    cancel: true,
                    persistent: true,
                })
                .onOk(async () => {
                    const res = await deleteAction(this.url.delete, {
                        id: row.id,
                    })
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
                    this.$refs.addOrEditDialog.onClose()
                })
        },
    },
}
</script>
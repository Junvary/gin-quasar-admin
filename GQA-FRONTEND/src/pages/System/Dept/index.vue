<template>
    <q-page padding>
        <q-splitter v-model="splitterModel">
            <template v-slot:before>
                <div class="q-pa-md column">
                    <div class="col">
                        <q-btn label="新增部门" color="primary" @click="showAddForm" />
                    </div>
                    <q-separator />
                    <q-tree :nodes="deptTree" default-expand-all node-key="id" label-key="name" selected-color="primary"
                        v-model:selected="selectedKey" v-if="deptTree.length !== 0" @update:selected="onSelected">
                        <template v-slot:default-header="prop">
                            <div class="row items-center">
                                <q-chip dense color="primary" text-color="white">
                                    {{ prop.node.sort}}
                                </q-chip>
                                <div class="text-weight-bold">
                                    {{ prop.node.deptName }}
                                </div>
                            </div>
                            <q-space></q-space>
                            <GqaDictShow dictName="statusOnOff" :dictCode="prop.node.status" />
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
                    <add-or-edit-card ref="addOrEditDialog" @handleFinish="handleFinish" />
                </div>
            </template>
        </q-splitter>
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import AddOrEditCard from './modules/addOrEditCard'
import { getAction, postAction, deleteAction, putAction } from 'src/api/manage'
import { ArrayToTree } from 'src/utils/arrayAndTree'
import GqaDictShow from 'src/components/GqaDictShow'

export default {
    name: 'Dept',
    mixins: [tableDataMixin],
    components: {
        AddOrEditCard,
        GqaDictShow,
    },
    computed: {
        deptTree() {
            if (this.tableData.length !== 0) {
                return ArrayToTree(this.tableData, 'deptCode', 'parentCode')
            }
            return []
        },
    },
    data() {
        return {
            url: {
                list: 'dept/dept-list',
                delete: 'dept/dept-delete',
            },
            pagination: {
                sortBy: 'sort',
                descending: false,
                page: 1,
                rowsPerPage: 10000,
                rowsNumber: 0,
            },
            splitterModel: 30,
            selectedKey: '',
        }
    },
    created() {
        this.getTableData()
    },
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
            if (this.$refs.addOrEditDialog.formType === 'add') {
                this.selectedKey = ''
            }
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
                    }
                    this.$refs.addOrEditDialog.onClose()
                })
        },
    },
}
</script>
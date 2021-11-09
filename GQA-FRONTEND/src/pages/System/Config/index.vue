<template>
    <q-page padding>
        <q-btn color="primary" @click="showAdd" label="新增配置" />
        <q-list bordered separator>
            <q-form ref="addOrEditForm">
                <q-item v-if="addFormVisible" active active-class="bg-teal-1">
                    <q-item-section>
                        <div class="column">
                            <div class="col">
                                <q-input dense v-model.number="form.sort" type="number"
                                    :rules="[ val => val >= 1 || '排序必须大于0']" label="排序" />
                            </div>
                            <div class="col">
                                <q-input v-model="form.remark" label="描述" />
                            </div>
                        </div>
                    </q-item-section>

                    <q-item-section>
                        <div class="column">
                            <div class="col">
                                <q-input v-model="form.gqaOption" label="字段名（英）"
                                    :rules="[ val => val && val.length > 0 || '必须输入字段名']" />
                            </div>
                            <div class="col">
                                <q-input v-model="form.default" label="默认值"
                                    :rules="[ val => val && val.length > 0 || '必须输入默认值']" />
                            </div>
                        </div>
                    </q-item-section>

                    <q-item-section side>
                        <div class="q-gutter-md">
                            <q-btn dense color="negative" @click="handleCancel" label="取消" />
                            <q-btn dense color="primary" @click="handleAdd" label="确定新增" />
                        </div>
                    </q-item-section>
                </q-item>
            </q-form>
            <q-item v-for="(item, index) in tableData" :key="index">
                <q-item-section class="col-4 gt-sm">
                    <q-item-label class="q-mt-sm">
                        {{item.sort}}：
                        {{item.gqaOption}}
                        （{{item.remark}}）
                    </q-item-label>
                </q-item-section>

                <q-item-section top>
                    <q-item-label lines="1">
                        <span class="text-weight-medium">默认值：</span>
                        <span class="text-grey-8">
                            {{item.default}}
                        </span>
                    </q-item-label>
                    <q-item-label lines="1">
                        <q-input v-model="item.custom" label="自定义" />
                    </q-item-label>
                </q-item-section>

                <q-item-section side>
                    <div class="q-gutter-md">
                        <q-btn dense color="primary" @click="handleSave(item)" label="保存编辑" />
                    </div>
                </q-item-section>
            </q-item>

        </q-list>
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { putAction, postAction } from 'src/api/manage'

export default {
    name: 'Config',
    mixins: [tableDataMixin],
    data() {
        return {
            pagination: {
                sortBy: 'sort',
                descending: false,
                page: 1,
                rowsPerPage: 9999,
                rowsNumber: 0,
                options: [10, 30, 50, 100],
            },
            url: {
                list: 'config/config-list',
                edit: 'config/config-edit',
                add: 'config/config-add',
            },
            addFormVisible: false,
            form: {
                sort: 1,
                remark: '',
                gqaOption: '',
                default: '',
            },
        }
    },
    created() {
        this.getTableData()
    },
    methods: {
        showAdd() {
            this.addFormVisible = true
        },
        handleCancel() {
            this.form = {
                sort: 1,
                remark: '',
                gqaOption: '',
                default: '',
            }
            this.addFormVisible = false
        },
        async handleSave(item) {
            const res = await putAction(this.url.edit, item)
            if (res.code === 1) {
                this.$q.notify({
                    type: 'positive',
                    message: res.message,
                })
                this.getTableData()
            }
        },
        async handleAdd() {
            const success = await this.$refs.addOrEditForm.validate()
            if (success) {
                const res = await postAction(this.url.add, this.form)
                if (res.code === 1) {
                    this.$q.notify({
                        type: 'positive',
                        message: res.message,
                    })
                    this.handleCancel()
                    this.getTableData()
                }
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: '请完善表格信息！',
                })
            }
        },
    },
}
</script>
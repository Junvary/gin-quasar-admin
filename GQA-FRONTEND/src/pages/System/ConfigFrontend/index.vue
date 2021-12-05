<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.gqaOption" :label="$t('Config') + $t('Name')" />
            <q-input style="width: 20%" v-model="queryParams.remark" :label="$t('Config') + $t('Remark')" />
            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Config')" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-default="props">
                <q-td :props="props">
                    <template v-if="props.row.gqaOption === 'gqaWebLogo'">
                        <GqaAvatar :src="props.row.default" />
                    </template>
                    <template v-else-if="props.row.gqaOption === 'gqaHeaderLogo'">
                        <GqaAvatar src="favicon.ico" />
                    </template>
                    <template v-else>
                        {{props.row.default}}
                    </template>
                </q-td>
            </template>

            <template v-slot:body-cell-custom="props">
                <q-td :props="props" class="bg-green-1">
                    <template v-if="props.row.gqaOption === 'gqaWebLogo'">
                        <GqaAvatar :src="props.row.custom" />
                    </template>
                    <template v-else-if="props.row.gqaOption === 'gqaHeaderLogo'">
                        <GqaAvatar :src="props.row.custom || 'favicon.ico'" />
                    </template>
                    <template v-else>
                        {{props.row.custom}}
                    </template>

                    <q-popup-edit v-model="props.row.custom" class="bg-green-13">
                        <!-- 使用是否字典显示 -->
                        <template v-slot="scope" v-if="props.row.gqaOption === 'gqaShowGit'">
                            {{ $t('Custom') + ' ' + props.row.gqaOption }}
                            <q-option-group v-model="props.row.custom" :options="options.statusYesNo" color="primary"
                                inline @update:model-value="scope.set">
                            </q-option-group>
                        </template>
                        <!-- 配置插件可选登录页 -->
                        <template v-slot="scope" v-else-if="props.row.gqaOption === 'gqaPluginLoginLayout'">
                            {{ $t('Custom') + ' ' + props.row.gqaOption }}
                            <GqaPluginList showChoose @changeSuccess="handleSetLoginLayout($event, scope)"
                                :choosePlugin="props.row.custom" />
                        </template>
                        <!-- 网站logo -->
                        <template v-slot="scope" v-else-if="props.row.gqaOption === 'gqaWebLogo'">
                            {{ $t('Custom') + ' ' + props.row.gqaOption }}
                            <q-file v-model="webLogoFile" clearable max-files="1" @rejected="rejected"
                                :accept="gqaBackend.webLogoExt" :max-file-size="gqaBackend.webLogoMaxSize*1024*1024">
                                <template v-slot:prepend>
                                    <GqaAvatar :src="props.row.custom" />
                                </template>
                                <template v-slot:after v-if="webLogoFile">
                                    <q-btn dense flat color="primary" icon="cloud_upload"
                                        @click="handleUploadWeb(scope)" />
                                </template>
                            </q-file>
                        </template>
                        <!-- 标签页logo -->
                        <template v-slot="scope" v-else-if="props.row.gqaOption === 'gqaHeaderLogo'">
                            {{ $t('Custom') + ' ' + props.row.gqaOption }}
                            <q-file v-model="headerLogoFile" clearable max-files="1" @rejected="rejected"
                                :accept="gqaBackend.headerLogoExt"
                                :max-file-size="gqaBackend.headerLogoMaxSize*1024*1024">
                                <template v-slot:prepend>
                                    <GqaAvatar :src="props.row.custom" />
                                </template>
                                <template v-slot:after v-if="headerLogoFile">
                                    <q-btn dense flat color="primary" icon="cloud_upload"
                                        @click="handleUploadHeader(scope)" />
                                </template>
                            </q-file>
                        </template>
                        <!-- 默认的输入框 -->
                        <template v-slot="scope" v-else>
                            {{ $t('Custom') + ' ' + props.row.gqaOption }}
                            <q-input v-model="props.row.custom" dense autofocus clearable @keyup.enter="scope.set" />
                        </template>
                    </q-popup-edit>
                </q-td>
            </template>

            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusOnOff" :dictCode="props.row.status" />
                </q-td>
            </template>

            <template v-slot:body-cell-stable="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusYesNo" :dictCode="props.row.stable" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn dense color="primary" @click="handleSave(props.row)" :label="$t('Save')" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <add-or-edit-dialog ref="addOrEditDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { gqaBackendMixin } from 'src/mixins/gqaBackendMixin'
import { mapActions } from 'vuex'
import addOrEditDialog from './modules/addOrEditDialog'
import { putAction, postAction } from 'src/api/manage'
import GqaDictShow from 'src/components/GqaDictShow'
import { addOrEditMixin } from 'src/mixins/addOrEditMixin'
import GqaAvatar from 'src/components/GqaAvatar'
import GqaPluginList from 'src/components/GqaPluginList'

export default {
    name: 'ConfigFrontend',
    mixins: [tableDataMixin, addOrEditMixin, gqaBackendMixin],
    components: {
        addOrEditDialog,
        GqaDictShow,
        GqaAvatar,
        GqaPluginList,
    },
    data() {
        return {
            url: {
                list: 'config-frontend/config-frontend-list',
                edit: 'config-frontend/config-frontend-edit',
                delete: 'config-frontend/config-frontend-delete',
                uploadWebUrl: 'upload/web-logo',
                uploadHeaderUrl: 'upload/header-logo',
            },
            columns: [
                { name: 'sort', align: 'center', label: this.$t('Sort'), field: 'sort' },
                { name: 'gqaOption', align: 'center', label: this.$t('Option'), field: 'gqaOption' },
                { name: 'remark', align: 'center', label: this.$t('Remark'), field: 'remark' },
                { name: 'default', align: 'center', label: this.$t('Default'), field: 'default' },
                { name: 'custom', align: 'center', label: this.$t('Custom'), field: 'custom' },
                { name: 'status', align: 'center', label: this.$t('Status'), field: 'status' },
                { name: 'stable', align: 'center', label: this.$t('Stable'), field: 'stable' },
                { name: 'actions', align: 'center', label: this.$t('Actions'), field: 'actions' },
            ],
            webLogoFile: null,
            headerLogoFile: null,
        }
    },
    created() {
        this.getTableData()
    },
    methods: {
        ...mapActions('storage', ['GetGqaFrontend']),
        async handleSave(row) {
            const res = await putAction(this.url.edit, row)
            if (res.code === 1) {
                this.$q.notify({
                    type: 'positive',
                    message: res.message,
                })
                this.GetGqaFrontend()
            }
        },
        handleUploadWeb(scope) {
            if (!this.webLogoFile) {
                this.$q.notify({
                    type: 'negative',
                    message: '请选择文件！',
                })
                return
            }
            let form = new FormData()
            form.append('file', this.webLogoFile)
            postAction(this.url.uploadWebUrl, form).then((res) => {
                if (res.code === 1) {
                    const gqaWebLogo = this.tableData.filter((item) => {
                        return item.gqaOption === 'gqaWebLogo'
                    })
                    gqaWebLogo[0].custom = res.data.records
                    this.webLogoFile = null
                    this.$q.notify({
                        type: 'positive',
                        message: '网站Logo上传成功！',
                    })
                    scope.set()
                }
            })
        },
        handleUploadHeader(scope) {
            if (!this.headerLogoFile) {
                this.$q.notify({
                    type: 'negative',
                    message: '请选择文件！',
                })
                return
            }
            let form = new FormData()
            form.append('file', this.headerLogoFile)
            postAction(this.url.uploadHeaderUrl, form).then((res) => {
                if (res.code === 1) {
                    const gqaHeaderLogo = this.tableData.filter((item) => {
                        return item.gqaOption === 'gqaHeaderLogo'
                    })
                    gqaHeaderLogo[0].custom = res.data.records
                    this.headerLogoFile = null
                    this.$q.notify({
                        type: 'positive',
                        message: '标签页Logo上传成功！',
                    })
                    scope.set()
                }
            })
        },
        rejected(rejectedEntries) {
            this.$q.notify({
                type: 'negative',
                message: '文件重复或大小/类型不被允许，请联系管理员！',
            })
        },
        handleSetLoginLayout(event, scope) {
            const gqaPluginLoginLayout = this.tableData.filter((item) => {
                return item.gqaOption === 'gqaPluginLoginLayout'
            })
            gqaPluginLoginLayout[0].custom = event
            scope.set(event)
        },
    },
}
</script>

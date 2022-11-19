<template>
    <q-page padding class="q-gutter-y-md">
        <div class="row q-gutter-x-md items-center">
            <q-input style="width: 20%" v-model="queryParams.config_item" :label="$t('Config') + $t('Name')" />
            <q-input style="width: 20%" v-model="queryParams.memo" :label="$t('Config') + $t('Memo')" />
            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" :label="$t('Add') + ' ' + $t('Config')"
                    v-has="'config-frontend:add'" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-item_default="props">
                <q-td :props="props">
                    <template v-if="props.row.config_item === 'logo'">
                        <GqaAvatar :src="props.row.item_default" />
                    </template>
                    <template v-else-if="props.row.config_item === 'favicon'">
                        <GqaAvatar src="favicon.ico" />
                    </template>
                    <template v-else-if="props.row.config_item === 'bannerImage'">
                        默认效果
                    </template>
                    <template v-else-if="props.row.config_item === 'loginLayoutStyle'">
                        <GqaDictShow :dictCode="props.row.item_default" />
                    </template>
                    <template v-else-if="props.row.config_item === 'showGit'">
                        <GqaDictShow :dictCode="props.row.item_default" />
                    </template>
                    <template v-else>
                        {{ props.row.item_default }}
                    </template>
                </q-td>
            </template>

            <template v-slot:header-cell-item_custom="props">
                <q-th :props="props">
                    {{ props.col.label }}
                    <q-icon name="edit" size="1.3em" />
                </q-th>
            </template>

            <template v-slot:body-cell-item_custom="props">
                <q-td :props="props">
                    <template v-if="props.row.config_item === 'bannerImage'">
                        <GqaAvatar :src="props.row.item_custom" v-if="props.row.item_custom !== ''" />
                    </template>
                    <template v-else-if="props.row.config_item === 'logo'">
                        <GqaAvatar :src="props.row.item_custom" />
                    </template>
                    <template v-else-if="props.row.config_item === 'favicon'">
                        <GqaAvatar :src="props.row.item_custom || 'favicon.ico'" />
                    </template>
                    <template v-else-if="props.row.config_item === 'loginLayoutStyle'">
                        <GqaDictShow :dictCode="props.row.item_custom" v-if="props.row.item_custom !== ''" />
                    </template>
                    <template v-else-if="props.row.config_item === 'showGit'">
                        <GqaDictShow :dictCode="props.row.item_custom" v-if="props.row.item_custom !== ''" />
                    </template>
                    <template v-else>
                        {{ props.row.item_custom }}
                    </template>

                    <q-popup-edit v-model="props.row.item_custom" :class="darkThemeSelect">
                        <template v-slot="scope">
                            {{ $t('Customize') + ' ' + props.row.config_item }}
                            <q-option-group v-if="props.row.config_item === 'loginLayoutStyle'"
                                v-model="props.row.item_custom" :options="dictOptions.displayStyle" color="primary"
                                inline @update:model-value="scope.set">
                            </q-option-group>
                            <q-option-group v-else-if="props.row.config_item === 'showGit'"
                                v-model="props.row.item_custom" :options="dictOptions.yesNo" color="primary" inline
                                @update:model-value="scope.set">
                            </q-option-group>
                            <GqaPluginList v-else-if="props.row.config_item === 'pluginLoginLayout'" showChoose
                                @changeSuccess="handleSetLoginLayout($event, scope)"
                                :choosePlugin="props.row.item_custom" />
                            <q-file v-else-if="props.row.config_item === 'bannerImage'" v-model="bannerImage" clearable
                                max-files="1" @rejected="rejected" :accept="gqaBackend.bannerImageExt"
                                :max-file-size="gqaBackend.bannerImageMaxSize * 1024 * 1024">
                                <template v-slot:prepend v-if="props.row.item_custom !== ''">
                                    <GqaAvatar :src="props.row.item_custom" />
                                </template>
                                <template v-slot:after v-if="bannerImage">
                                    <q-btn dense flat color="primary" icon="cloud_upload"
                                        @click="handleUploadBannerImage(scope)" />
                                </template>
                            </q-file>
                            <q-file v-else-if="props.row.config_item === 'logo'" v-model="logoFile" clearable
                                max-files="1" @rejected="rejected" :accept="gqaBackend.logoExt"
                                :max-file-size="gqaBackend.logoMaxSize * 1024 * 1024">
                                <template v-slot:prepend>
                                    <GqaAvatar :src="props.row.item_custom" />
                                </template>
                                <template v-slot:after v-if="logoFile">
                                    <q-btn dense flat color="primary" icon="cloud_upload"
                                        @click="handleUploadLogo(scope)" />
                                </template>
                            </q-file>
                            <q-file v-else-if="props.row.config_item === 'favicon'" v-model="faviconFile" clearable
                                max-files="1" @rejected="rejected" :accept="gqaBackend.faviconExt"
                                :max-file-size="gqaBackend.faviconMaxSize * 1024 * 1024">
                                <template v-slot:prepend>
                                    <GqaAvatar :src="props.row.item_custom" />
                                </template>
                                <template v-slot:after v-if="faviconFile">
                                    <q-btn dense flat color="primary" icon="cloud_upload"
                                        @click="handleUploadFavicon(scope)" />
                                </template>
                            </q-file>
                            <q-input v-else v-model="props.row.item_custom" dense autofocus clearable
                                @keyup.enter="scope.set" />
                        </template>
                    </q-popup-edit>
                </q-td>
            </template>

            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <GqaDictShow :dictCode="props.row.status" />
                </q-td>
            </template>

            <template v-slot:body-cell-stable="props">
                <q-td :props="props">
                    <GqaDictShow :dictCode="props.row.stable" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="handleSave(props.row)" :label="$t('Save')"
                            v-has="'config-frontend:save'" />
                        <q-btn color="warning" @click="handleReset(props.row)" :label="$t('Reset')"
                            v-has="'config-frontend:reset'" />
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')"
                            v-has="'config-frontend:delete'" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import GqaPluginList from 'src/components/GqaPluginList/index.vue'
import { useStorageStore } from 'src/stores/storage'
import recordDetail from './modules/recordDetail'
import useTheme from 'src/composables/useTheme';

const $q = useQuasar()
const { t } = useI18n()
const storageStore = useStorageStore()
const { darkThemeSelect } = useTheme()
const url = {
    list: 'config-frontend/get-config-frontend-list',
    edit: 'config-frontend/edit-config-frontend',
    delete: 'config-frontend/delete-config-frontend-by-id',
    uploadBannerImage: 'upload/upload-banner-image',
    uploadLogo: 'upload/upload-logo',
    uploadFavicon: 'upload/upload-favicon',
}
const columns = computed(() => {
    return [
        { name: 'sort', align: 'center', label: t('Sort'), field: 'sort' },
        { name: 'config_item', align: 'center', label: t('Config') + t('Item'), field: 'config_item' },
        { name: 'memo', align: 'center', label: t('Memo'), field: 'memo' },
        { name: 'item_default', align: 'center', label: t('Default'), field: 'item_default' },
        { name: 'item_custom', align: 'center', label: t('Custom'), field: 'item_custom' },
        { name: 'status', align: 'center', label: t('Status'), field: 'status' },
        { name: 'stable', align: 'center', label: t('Stable'), field: 'stable' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    gqaBackend,
    dictOptions,
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    GqaAvatar,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    onRequest,
    getTableData,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    getTableData()
})

const bannerImage = ref(null)
const logoFile = ref(null)
const faviconFile = ref(null)

const handleReset = (row) => {
    row.item_custom = ''
}
const handleSave = async (row) => {
    const res = await postAction(url.edit, row)
    if (res.code === 1) {
        $q.notify({
            type: 'positive',
            message: res.message,
        })
        storageStore.SetGqaFrontend()
    }
}
const handleUploadBannerImage = (scope) => {
    if (!bannerImage.value) {
        $q.notify({
            type: 'negative',
            message: '请选择图片！',
        })
        return
    }
    let form = new FormData()
    form.append('file', bannerImage.value)
    postAction(url.uploadBannerImage, form).then((res) => {
        if (res.code === 1) {
            const bi = tableData.value.filter((item) => {
                return item.config_item === 'bannerImage'
            })
            bi[0].item_custom = res.data.records
            bannerImage.value = null
            $q.notify({
                type: 'positive',
                message: '首页大图上传成功！',
            })
            scope.set()
        }
    })
}
const handleUploadLogo = (scope) => {
    if (!logoFile.value) {
        $q.notify({
            type: 'negative',
            message: '请选择文件！',
        })
        return
    }
    let form = new FormData()
    form.append('file', logoFile.value)
    postAction(url.uploadLogo, form).then((res) => {
        if (res.code === 1) {
            const logo = tableData.value.filter((item) => {
                return item.config_item === 'logo'
            })
            logo[0].item_custom = res.data.records
            logoFile.value = null
            $q.notify({
                type: 'positive',
                message: '网站Logo上传成功！',
            })
            scope.set()
        }
    })
}
const handleUploadFavicon = (scope) => {
    if (!faviconFile.value) {
        $q.notify({
            type: 'negative',
            message: '请选择文件！',
        })
        return
    }
    let form = new FormData()
    form.append('file', faviconFile.value)
    postAction(url.uploadFavicon, form).then((res) => {
        if (res.code === 1) {
            const favicon = tableData.value.filter((item) => {
                return item.config_item === 'favicon'
            })
            favicon[0].item_custom = res.data.records
            faviconFile.value = null
            $q.notify({
                type: 'positive',
                message: '网站Favicon上传成功！',
            })
            scope.set()
        }
    })
}
const rejected = (rejectedEntries) => {
    $q.notify({
        type: 'negative',
        message: '文件重复或大小/类型不被允许，请联系管理员！',
    })
}
const handleSetLoginLayout = (event, scope) => {
    // console.log(event)
    const pluginLoginLayout = tableData.value.filter((item) => {
        return item.config_item === 'pluginLoginLayout'
    })
    pluginLoginLayout[0].item_custom = event
    scope.set(event)
}
</script>

<template>
    <q-page padding>
        <q-card flat>
            <q-card-section class="row q-gutter-x-md items-center">
                <q-input outlined dense style="width: 20%" v-model="queryParams.config_item"
                    :label="$t('Config') + $t('Item')" />
                <q-input outlined dense style="width: 20%" v-model="queryParams.memo" :label="$t('Memo')" />
                <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
                <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
            </q-card-section>
            <q-card-section>
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
                                <gqa-avatar :src="props.row.item_default" />
                            </template>
                            <template v-else-if="props.row.config_item === 'favicon'">
                                <gqa-avatar src="favicon.ico" />
                            </template>
                            <template v-else-if="props.row.config_item === 'bannerImage'">
                                Default
                            </template>
                            <template v-else-if="props.row.config_item === 'loginLayoutStyle'">
                                <GqaDictShow :dictCode="props.row.item_default" />
                            </template>
                            <template v-else-if="props.row.config_item === 'pluginLoginLayout'">
                                系统默认
                            </template>
                            <template v-else-if="props.row.config_item === 'portalPage'">
                                <GqaDictShow :dictCode="props.row.item_default" />
                            </template>
                            <template v-else-if="props.row.config_item === 'showGit'">
                                <GqaDictShow :dictCode="props.row.item_default" />
                            </template>
                            <template v-else-if="props.row.config_item === 'starColor'">
                                <div class="row justify-center items-center q-gutter-x-xs">
                                    {{ props.row.item_default }}
                                    <div :style="{ background: props.row.item_default }"
                                        style="width: 10px; height: 10px; border-radius: 100%;"></div>
                                </div>
                            </template>
                            <template v-else>
                                {{ props.row.item_default }}
                            </template>
                        </q-td>
                    </template>
                    <template v-slot:body-cell-item_custom="props">
                        <q-td :props="props">
                            <template v-if="props.row.config_item === 'bannerImage'">
                                <q-file dense outlined v-model="bannerImage" clearable max-files="1" @rejected="rejected"
                                    :accept="gqaBackend.bannerImageExt"
                                    :max-file-size="gqaBackend.bannerImageMaxSize * 1024 * 1024">
                                    <template v-slot:prepend v-if="props.row.item_custom !== ''">
                                        <gqa-avatar :src="props.row.item_custom" />
                                    </template>
                                    <template v-slot:after v-if="bannerImage">
                                        <q-btn dense flat color="primary" icon="cloud_upload"
                                            @click="handleUploadBannerImage" />
                                    </template>
                                </q-file>
                            </template>
                            <template v-else-if="props.row.config_item === 'logo'">
                                <q-file dense outlined v-model="logoFile" clearable max-files="1" @rejected="rejected"
                                    :accept="gqaBackend.logoExt" :max-file-size="gqaBackend.logoMaxSize * 1024 * 1024">
                                    <template v-slot:prepend>
                                        <gqa-avatar :src="props.row.item_custom" />
                                    </template>
                                    <template v-slot:after v-if="logoFile">
                                        <q-btn dense flat color="primary" icon="cloud_upload" @click="handleUploadLogo" />
                                    </template>
                                </q-file>
                            </template>
                            <template v-else-if="props.row.config_item === 'favicon'">
                                <q-file dense outlined v-model="faviconFile" clearable max-files="1" @rejected="rejected"
                                    :accept="gqaBackend.faviconExt"
                                    :max-file-size="gqaBackend.faviconMaxSize * 1024 * 1024">
                                    <template v-slot:prepend>
                                        <gqa-avatar :src="props.row.item_custom" />
                                    </template>
                                    <template v-slot:after v-if="faviconFile">
                                        <q-btn dense flat color="primary" icon="cloud_upload"
                                            @click="handleUploadFavicon" />
                                    </template>
                                </q-file>
                            </template>
                            <template v-else-if="props.row.config_item === 'loginLayoutStyle'">
                                <q-option-group v-model="props.row.item_custom" :options="dictOptions.loginLayoutStyle"
                                    color="primary" inline>
                                    <template v-slot:label="opt">
                                        <div class="row items-center">
                                            <span>{{ $t(opt.label) }}</span>
                                        </div>
                                    </template>
                                </q-option-group>
                            </template>
                            <template v-else-if="props.row.config_item === 'pluginLoginLayout'">
                                <q-select dense outlined v-model="props.row.item_custom" :options="pluginList"
                                    :label="$t('Choose') + $t('Plugin') + '(' + pluginList.length + ')'" emit-value
                                    map-options option-value="plugin_code" option-label="plugin_name"
                                    @update:model-value="handelChangePlugin" />
                            </template>
                            <template v-else-if="props.row.config_item === 'showGit'">
                                <q-option-group v-model="props.row.item_custom" :options="dictOptions.yesNo" color="primary"
                                    inline>
                                    <template v-slot:label="opt">
                                        <div class="row items-center">
                                            <span>{{ $t(opt.label) }}</span>
                                        </div>
                                    </template>
                                </q-option-group>
                            </template>
                            <template v-else-if="props.row.config_item === 'starColor'">
                                <q-input dense outlined v-model="props.row.item_custom">
                                    <template v-slot:prepend v-if="props.row.item_custom">
                                        <div :style="{ background: props.row.item_custom }"
                                            style="width: 10px; height: 10px; border-radius: 100%;"></div>
                                    </template>
                                    <template v-slot:append>
                                        <q-icon name="colorize" class="cursor-pointer">
                                            <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                                                <q-color v-model="props.row.item_custom" />
                                            </q-popup-proxy>
                                        </q-icon>
                                    </template>
                                </q-input>
                            </template>
                            <template v-else>
                                <q-input v-model="props.row.item_custom" dense outlined clearable />
                            </template>
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
                        <q-td :props="props" class="q-gutter-x-md">
                            <q-btn flat dense color="primary" @click="handleSave(props.row)" :label="$t('Save')"
                                v-has="'config-frontend:save'">
                            </q-btn>
                            <q-btn flat dense color="warning" @click="handleReset(props.row)" :label="$t('Reset')"
                                v-has="'config-frontend:reset'">
                            </q-btn>
                            <q-btn flat dense color="negative" :label="$t('Delete')" v-if="props.row.stable !== 'yesNo_yes'"
                                @click="handleDelete(props.row)" v-has="'config-frontend:delete'">
                            </q-btn>
                        </q-td>
                    </template>
                </q-table>
            </q-card-section>
        </q-card>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import useTableData from 'src/composables/useTableData'
import { postAction } from 'src/api/manage'
import { useStorageStore } from 'src/stores/storage'
import recordDetail from './modules/recordDetail.vue'
import { CheckComponent } from 'src/utils/check'

const storageStore = useStorageStore()
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
    $q,
    t,
    gqaBackend,
    dictOptions,
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
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

const pluginList = computed(() => {
    if (Array.isArray($q.localStorage.getItem('gqa-pluginList'))) {
        return $q.localStorage.getItem('gqa-pluginList')
    }
    return []
})

const pluginsFile = import.meta.glob('../../../plugins/**/LoginLayout/index.vue')
const handelChangePlugin = (cp) => {
    const cc = CheckComponent(cp, pluginsFile, 4)
    if (cc[2]) {
        $q.notify({
            type: 'negative',
            message: t('PluginNotSupportLogin'),
        })
        const pll = tableData.value.filter((item) => {
            return item.config_item === 'pluginLoginLayout'
        })
        pll[0].item_custom = ''
    }
}

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
const handleUploadBannerImage = () => {
    if (!bannerImage.value) {
        $q.notify({
            type: 'negative',
            message: t('Please') + t('Select') + t('File'),
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
                message: t('Upload') + t('Success'),
            })
        }
    })
}
const handleUploadLogo = () => {
    if (!logoFile.value) {
        $q.notify({
            type: 'negative',
            message: t('Please') + t('Select') + t('File'),
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
                message: t('Upload') + t('Success'),
            })
        }
    })
}
const handleUploadFavicon = () => {
    if (!faviconFile.value) {
        $q.notify({
            type: 'negative',
            message: t('Please') + t('Select') + t('File'),
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
                message: t('Upload') + t('Success'),
            })
        }
    })
}
const rejected = (rejectedEntries) => {
    $q.notify({
        type: 'negative',
        message: t('SizeOrExtError'),
    })
}
</script>

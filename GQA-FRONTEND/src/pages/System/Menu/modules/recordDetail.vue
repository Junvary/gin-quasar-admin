<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 1200px; max-width: 80vw;">
            <q-card-section class="text-h6">
                {{ formTypeName }} {{ $t('Menu') }}:
                {{ recordDetail.value.title ? $t(recordDetail.value.title) : '' }}
            </q-card-section>
            <gqa-form-top :recordDetail="recordDetail" style="margin: 0 16px;"></gqa-form-top>
            <div class="row">
                <q-card-section class="q-gutter-md col-8">
                    <q-form ref="recordDetailForm">
                        <div class="row q-gutter-md">
                            <q-input outlined hint="" class="col" :label="$t('Icon')" v-model="recordDetail.value.icon"
                                clearable>
                                <template v-slot:prepend>
                                    <q-icon :name="recordDetail.value.icon" class="cursor-pointer">
                                        <q-popup-proxy v-model="iconData.showIconPicker">
                                            <q-input v-model="iconData.filter" label="Filter" outlined clearable dense
                                                class="q-ma-md" />
                                            <q-icon-picker v-model="recordDetail.value.icon" icon-set="material-icons"
                                                :filter="iconData.filter" v-model:model-pagination="iconData.pagination"
                                                tooltips style="height: 300px; width: 300px; background-color: white;" />
                                        </q-popup-proxy>
                                    </q-icon>
                                </template>
                            </q-input>
                            <q-input outlined hint="" class="col" v-model.number="recordDetail.value.sort" type="number"
                                :rules="[val => val >= 1 || $t('SortRule')]" :label="$t('Sort')" />
                        </div>
                        <div class="row q-gutter-md">
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.title"
                                :label="$t('Menu') + $t('Name')"
                                :rules="[val => val && val.length > 0 || $t('NeetInput')]" />
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.name"
                                :label="$t('Menu') + $t('Code')" :disable="formType === 'edit'"
                                :rules="[val => val && val.length > 0 || $t('NeetInput')]" />
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.parent_code"
                                :label="$t('Parent') + $t('Code')" readonly />
                        </div>
                        <div class="row q-gutter-md">
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.path" label="URL" />
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.redirect"
                                :label="t('Redirect')" />
                            <q-input outlined hint="" class="col" v-model="recordDetail.value.component"
                                :label="$t('Component')" />
                        </div>

                        <div class="row q-gutter-md">
                            <q-field outlined hint="" class="col" :label="$t('IsLink')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="recordDetail.value.is_link" :options="dictOptions.yesNo"
                                        color="primary" inline>
                                        <template v-slot:label="opt">
                                            <div class="row items-center">
                                                <span>{{ $t(opt.label) }}</span>
                                            </div>
                                        </template>
                                    </q-option-group>
                                </template>
                            </q-field>
                            <q-field outlined hint="" class="col" :label="$t('KeepAlive')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="recordDetail.value.keep_alive" :options="dictOptions.yesNo"
                                        color="primary" inline>
                                        <template v-slot:label="opt">
                                            <div class="row items-center">
                                                <span>{{ $t(opt.label) }}</span>
                                            </div>
                                        </template>
                                    </q-option-group>
                                </template>
                            </q-field>
                            <q-field outlined hint="" class="col" :label="$t('Hidden')" stack-label>
                                <template v-slot:control>
                                    <q-option-group v-model="recordDetail.value.hidden" :options="dictOptions.yesNo"
                                        color="primary" inline>
                                        <template v-slot:label="opt">
                                            <div class="row items-center">
                                                <span>{{ $t(opt.label) }}</span>
                                            </div>
                                        </template>
                                    </q-option-group>
                                </template>
                            </q-field>
                        </div>
                        <q-input outlined hint="" v-model="recordDetail.value.memo" type="textarea" :label="$t('Memo')" />
                    </q-form>
                </q-card-section>
                <q-card-section class="q-gutter-md col">
                    <q-field outlined hint="" dense :label="$t('Parent') + $t('Menu')" stack-label
                        :disable="!recordDetail.value.parent_code">
                        <template v-slot:control>
                            <q-tree :nodes="menuTree" v-model:selected="recordDetail.value.parent_code" style="width: 100%;"
                                selected-color="primary" node-key="name" default-expand-all
                                @update:selected="handleSelectParentMenu">
                                <template v-slot:default-header="prop">
                                    <div class="row items-center">
                                        <div>{{ t(prop.node.title) }}</div>
                                    </div>
                                </template>
                            </q-tree>
                        </template>
                    </q-field>
                </q-card-section>
            </div>
            <q-separator />
            <q-card-actions align="right">
                <q-btn :label="$t('Save')" color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('Cancel')" color="negative" v-close-popup />
            </q-card-actions>
            <q-inner-loading :showing="loading">
                "primary" />
            </q-inner-loading>
        </q-card>
    </q-dialog>
</template>


<script setup>
import useRecordDetail from 'src/composables/useRecordDetail'
import { ref, watch, toRefs } from 'vue'

const emit = defineEmits(['handleFinish'])
const url = {
    list: 'menu/get-menu-list',
    add: 'menu/add-menu',
    edit: 'menu/edit-menu',
    queryById: 'menu/query-menu-by-id',
}

const props = defineProps({
    menuTree: {
        type: Array,
        required: false,
        default: () => [],
    },
})
const { menuTree } = toRefs(props)

const {
    $q,
    t,
    dictOptions,
    formType,
    formTypeName,
    recordDetail,
    recordDetailVisible,
    loading,
    show,
    recordDetailForm,
    handleAddOrEidt
} = useRecordDetail(url, emit)

defineExpose({
    show,
    formType,
    recordDetail
})
const iconData = ref({
    filter: '',
    showIconPicker: false,
    pagination: {
        itemsPerPage: 35,
        page: 0
    }
})

watch(recordDetail, () => {
    iconData.value.showIconPicker = false
})

const handleSelectParentMenu = (target) => {
    if (target === recordDetail.value.name) {
        recordDetail.value.parent_code = ''
        $q.notify({ type: 'warning', message: t('ParentNotThis') })
    }
}
</script>

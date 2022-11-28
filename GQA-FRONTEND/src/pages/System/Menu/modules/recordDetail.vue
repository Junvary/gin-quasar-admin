<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 900px; max-width: 80vw;">
            <q-card-section class="text-h6">
                {{ formTypeName }} {{ $t('Menu') }}:
                {{ recordDetail.value.title ? $t(recordDetail.value.title) : '' }}
            </q-card-section>
            <q-separator />
            <q-card-section>
                <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
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
                            :label="$t('Menu') + $t('Code')"
                            :rules="[val => val && val.length > 0 || $t('NeetInput')]" />
                        <q-input outlined hint="" class="col" v-model="recordDetail.value.parent_code"
                            :label="$t('Parent') + $t('Code')" />
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
            <q-separator />
            <q-card-actions align="right">
                <q-btn :label="$t('Save')" color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('Cancel')" color="negative" v-close-popup />
            </q-card-actions>
            <q-inner-loading :showing="loading">
                <q-spinner-gears size="50px" color="primary" />
            </q-inner-loading>
        </q-card>
    </q-dialog>
</template>


<script setup>
import useRecordDetail from 'src/composables/useRecordDetail'
import { ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const emit = defineEmits(['handleFinish'])
const url = {
    list: 'menu/get-menu-list',
    add: 'menu/add-menu',
    edit: 'menu/edit-menu',
    queryById: 'menu/query-menu-by-id',
}
const {
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
</script>

<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 800px; max-width: 80vw;">
            <div class="row justify-between items-center">
                <q-card-section>
                    <div class="text-h6">
                        {{ formTypeName }} {{ $t('Menu') }}:
                        {{ recordDetail.value.title }}
                    </div>
                </q-card-section>
            </div>

            <q-separator />

            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <div class="row">
                        <q-input class="col" :label="$t('Icon')" v-model="recordDetail.value.icon" clearable>
                            <template v-slot:before>
                                <q-icon :name="recordDetail.value.icon" size="35px" class="q-mr-sm" />
                            </template>
                            <template v-slot:prepend>
                                <q-icon name="insert_emoticon" class="cursor-pointer">
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
                        <q-input class="col" v-model.number="recordDetail.value.sort" type="number"
                            :rules="[val => val >= 1 || $t('SortRule')]" :label="$t('Sort')" />
                    </div>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.title" :label="$t('Menu') + $t('Name')"
                            :rules="[val => val && val.length > 0 || $t('NeetInput')]" />
                        <q-input class="col" v-model="recordDetail.value.name" :label="$t('Menu') + $t('Code')"
                            :rules="[val => val && val.length > 0 || $t('NeetInput')]" />
                        <q-input class="col" v-model="recordDetail.value.parent_code"
                            :label="$t('Parent') + $t('Code')" />
                    </div>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.path" label="URL" />
                        <q-input class="col" v-model="recordDetail.value.component" :label="$t('Component')" />
                    </div>

                    <div class="row">
                        <q-field class="col" :label="$t('IsLink')" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="recordDetail.value.is_link" :options="dictOptions.statusYesNo"
                                    color="primary" inline>
                                </q-option-group>
                            </template>
                        </q-field>
                        <q-field class="col" :label="$t('KeepAlive')" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="recordDetail.value.keep_alive"
                                    :options="dictOptions.statusYesNo" color="primary" inline>
                                </q-option-group>
                            </template>
                        </q-field>
                        <q-field class="col" :label="$t('Hidden')" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="recordDetail.value.hidden" :options="dictOptions.statusYesNo"
                                    color="primary" inline>
                                </q-option-group>
                            </template>
                        </q-field>
                        <q-field class="col" :label="$t('Status')" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="recordDetail.value.status" :options="dictOptions.statusOnOff"
                                    color="primary" inline>
                                </q-option-group>
                            </template>
                        </q-field>
                    </div>
                    <q-input v-model="recordDetail.value.memo" type="textarea" :label="$t('Memo')" />
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
import GqaAvatar from 'src/components/GqaAvatar'
import GqaShowName from 'src/components/GqaShowName'
import { postAction } from 'src/api/manage'
import { useStorageStore } from 'src/stores/storage'
import { ref, computed, watch } from 'vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import GqaSeleteUser from 'src/components/GqaSeleteUser'

const $q = useQuasar()
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
    showDateTime,
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

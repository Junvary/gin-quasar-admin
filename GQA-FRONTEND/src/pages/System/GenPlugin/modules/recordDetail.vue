<template>
    <q-dialog v-model="recordDetailVisible" position="top" persistent>
        <q-card style="width: 1500px; max-width: 90vw;">
            <q-card-section class="text-h6">
                {{ formTypeName }} {{ $t('Plugin') }}
            </q-card-section>
            <q-separator />
            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <div class="row q-gutter-md">
                        <q-input dense outlined hint="" class="col" v-model.number="form.plugin_sort" type="number"
                            :label="$t('Plugin') + $t('Sort')" :rules="[val => val && val > 1000 || '>1000']" />
                        <q-input dense outlined hint="Camelcase" class="col" v-model.trim="form.plugin_code"
                            prefix="plugin-" :label="$t('Plugin') + $t('Code')"
                            :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        <q-input dense outlined hint="" class="col" v-model.trim="form.plugin_name"
                            :label="$t('Plugin') + $t('Name')" :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                        <q-input dense outlined hint="Camelcase" class="col" v-model.trim="inputModel"
                            :label="$t('Add') + 'Model'">
                            <template v-slot:after>
                                <q-btn icon="add" round dense @click="handleInputModel" />
                            </template>
                        </q-input>
                    </div>

                    <div class="row q-gutter-md" v-if="form.plugin_model.length" style="margin-top: 10px;">
                        <q-card style="width: 100%">
                            <q-tabs v-model="modelStruct" align="left" inline-label outside-arrows mobile-arrows>
                                <q-tab v-for="(item, index) in form.plugin_model" :name="item.model_name">
                                    <span style="white-space:nowrap">
                                        {{ item.model_name }}
                                        <q-icon class="tab-close" name="close" style="margin-left: 5px;"
                                            @click.prevent.stop="removeModel(item, index)" />
                                    </span>
                                </q-tab>
                            </q-tabs>
                            <q-separator />

                            <q-tab-panels v-model="modelStruct" animated style="width: 100%">
                                <q-tab-panel v-for="(item, index) in form.plugin_model" :name="item.model_name"
                                    style="padding: 0;">
                                    <q-table :rows="item.column_list" :columns="columns" row-key="column_name"
                                        hide-pagination>
                                        <template v-slot:top class="q-gutter-md">
                                            <q-btn color="primary" @click="addColumn(item)">
                                                {{ $t('Add') }}
                                                {{ $t('Column') }}
                                            </q-btn>
                                            <q-toggle v-model="item.with_gqa_column"
                                                :label="($t('With') + $t('Basic') + $t('Column'))" />
                                            <q-toggle v-model="item.with_public_list"
                                                :label="$t('With') + $t('Public') + 'List'" />
                                            <q-toggle v-model="item.with_data_permission"
                                                :label="($t('With') + ' ' + $t('DataPermission'))" />
                                            <q-toggle v-model="item.with_log_operation"
                                                :label="($t('With') + ' ' + $t('LogOperation'))" />
                                        </template>

                                        <template v-slot:header-cell-column_name="props">
                                            <q-th :props="props">
                                                {{ props.col.label }}
                                                <q-icon name="edit" />
                                            </q-th>
                                        </template>
                                        <template v-slot:body-cell-column_name="props">
                                            <q-td :props="props">
                                                {{ props.row.column_name }}
                                                <q-popup-edit v-model="props.row.column_name" v-slot="scope">
                                                    <q-input v-model.trim="scope.value" dense autofocus
                                                        @keyup.enter="scope.set" />
                                                </q-popup-edit>
                                            </q-td>
                                        </template>

                                        <template v-slot:header-cell-column_type="props">
                                            <q-th :props="props">
                                                {{ props.col.label }}
                                                <q-icon name="edit" />
                                            </q-th>
                                        </template>
                                        <template v-slot:body-cell-column_type="props">
                                            <q-td :props="props">
                                                {{ props.row.column_type }}
                                                <q-popup-edit v-model="props.row.column_type" v-slot="scope">
                                                    <q-select v-model.trim="scope.value" :options="column_typeOptions" dense
                                                        @update:model-value="scope.set" />
                                                </q-popup-edit>
                                            </q-td>
                                        </template>

                                        <template v-slot:header-cell-column_comment="props">
                                            <q-th :props="props">
                                                {{ props.col.label }}
                                                <q-icon name="edit" />
                                            </q-th>
                                        </template>
                                        <template v-slot:body-cell-column_comment="props">
                                            <q-td :props="props">
                                                {{ props.row.column_comment }}
                                                <q-popup-edit v-model="props.row.column_comment" v-slot="scope">
                                                    <q-input v-model.trim="scope.value" dense autofocus
                                                        @keyup.enter="scope.set" />
                                                </q-popup-edit>
                                            </q-td>
                                        </template>

                                        <template v-slot:header-cell-column_default="props">
                                            <q-th :props="props">
                                                {{ props.col.label }}
                                                <q-icon name="edit" />
                                            </q-th>
                                        </template>
                                        <template v-slot:body-cell-column_default="props">
                                            <q-td :props="props">
                                                {{ props.row.column_default }}
                                                <q-popup-edit v-model="props.row.column_default" v-slot="scope">
                                                    <q-input v-model.trim="scope.value" dense autofocus
                                                        @keyup.enter="scope.set" />
                                                </q-popup-edit>
                                            </q-td>
                                        </template>
                                        <template v-slot:body-cell-actions="props">
                                            <q-td :props="props">
                                                <div class="q-gutter-xs">
                                                    <q-btn color="negative" @click="removeColumn(item, props.row)"
                                                        label="移除" />
                                                </div>
                                            </q-td>
                                        </template>
                                    </q-table>
                                </q-tab-panel>
                            </q-tab-panels>
                        </q-card>
                    </div>
                </q-form>
            </q-card-section>
            <q-separator />
            <q-card-actions align="right">
                <q-btn :label="$t('Gen')" color="primary" @click="handleAddOrEidt" v-has="'genPlugin:gen'" />
                <q-btn :label="$t('Cancel')" color="negative" v-close-popup />
            </q-card-actions>
            <q-inner-loading :showing="loading">
                <q-spinner-gears size="50px" color="primary" />
            </q-inner-loading>
        </q-card>
    </q-dialog>
</template>


<script setup>
import { ref, computed } from 'vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import { postAction } from 'src/api/manage'
import useRecordDetail from 'src/composables/useRecordDetail'

const emit = defineEmits(['handleFinish'])
const $q = useQuasar()
const { t } = useI18n()
const inputModel = ref('')
const modelStruct = ref('')
const form = ref({
    plugin_sort: 5000,
    plugin_name: '',
    plugin_code: '',
    plugin_model: [],
})
const url = {
    add: 'gen-plugin/gen-plugin',
}

const column_typeOptions = [
    'string', 'uint', 'time.Time', 'float64'
]

const columns = computed(() => {
    return [
        { name: 'column_name', label: t('Column') + t('Name'), field: 'column_name', align: 'center' },
        { name: 'column_type', label: t('Data') + t('Type'), field: 'column_type', align: 'center' },
        { name: 'column_comment', label: t('Comment'), field: 'column_comment', align: 'center' },
        { name: 'column_default', label: t('Default'), field: 'column_default', align: 'center' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})

const removeModel = (item, index) => {
    form.value.plugin_model = form.value.plugin_model.filter(m => m.model_name != item.model_name)
    if (modelStruct.value === item.model_name) {
        if (index - 1 >= 0) {
            modelStruct.value = form.value.plugin_model[index - 1].model_name
        } else {
            modelStruct.value = form.value.plugin_model.length ? form.value.plugin_model[0].model_name : ''
        }
    }
}

const removeColumn = (item, row) => {
    item.column_list = item.column_list.filter(item => item.column_name !== row.column_name)
}

const handleInputModel = () => {
    const model = inputModel.value
    if (model) {
        form.value.plugin_model.push({
            model_name: model,
            column_list: [],
            with_gqa_column: true,
            with_public_list: false,
            with_data_permission: false,
            with_log_operation: false,
        })
        inputModel.value = ''
        modelStruct.value = model
    } else {
        $q.notify({
            type: 'negative',
            message: '请输入Model名称',
        })
    }
}
const addColumn = (item) => {
    item.column_list.push({
        column_name: '',
        column_type: '',
        column_comment: '',
        column_default: '',
    })
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
} = useRecordDetail(url, emit)

defineExpose({
    show,
    formType,
    recordDetail
})

const handleAddOrEidt = async () => {
    const success = await recordDetailForm.value.validate()
    if (success) {
        const ok = await postAction(url.add, form.value)
        if (ok.code === 1) {
            $q.notify({
                type: 'positive',
                message: t('Gen') + t('Success'),
            })
            emit('handleFinish')
            recordDetailVisible.value = false
        } else {
            $q.notify({
                type: 'negative',
                message: t('Gen') + t('Failed'),
            })
        }
    } else {
        $q.notify({
            type: 'negative',
            message: t('FixForm'),
        })
    }
}
</script>

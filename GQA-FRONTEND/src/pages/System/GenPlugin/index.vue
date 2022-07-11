<template>
    <q-page padding>
        <q-form ref="genPluginForm">
            <div class="q-gutter-x-md">
                <div class="row q-gutter-md">
                    <q-input class="col" v-model="form.plugin_code"
                        :label="$t('Plugin') + $t('Name') + '(' + $t('English') + ')'" lazy-rules
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                    <q-input class="col" v-model="form.plugin_name" :label="$t('Plugin') + $t('Name')" lazy-rules
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                    <q-toggle class="col-2" v-model="form.with_model" label="生成Model" />
                    <div class="col">
                        <q-input v-model="inputModel" v-if="form.with_model" label="新增Model">
                            <template v-slot:after>
                                <q-btn icon="add" round dense @click="handleInputModel" />
                            </template>
                        </q-input>
                    </div>
                </div>
                <div class="row q-gutter-md" v-if="form.with_model">
                    <q-card style="width: 100%">
                        <q-tabs v-model="modelStruct" align="left" inline-label outside-arrows mobile-arrows>
                            <q-tab v-for="(item, index) in form.plugin_model" :name="item.model_name">
                                <span style="white-space:nowrap">
                                    维护 {{ item.model_name }}模型
                                    <q-icon class="tab-close" name="close"
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
                                        <q-btn color="primary" @click="addColumn(item)">添加字段</q-btn>
                                        <q-toggle v-model="item.with_gqa_column" label="携带基础字段" />
                                        <q-toggle v-model="item.with_public_list" label="携带非鉴权List" />
                                    </template>
                                    <template v-slot:body-cell-column_name="props">
                                        <q-td :props="props">
                                            {{ props.row.column_name }}
                                            <q-popup-edit v-model="props.row.column_name" v-slot="scope">
                                                <q-input v-model="scope.value" dense autofocus
                                                    @keyup.enter="scope.set" />
                                            </q-popup-edit>
                                        </q-td>
                                    </template>
                                    <template v-slot:body-cell-column_type="props">
                                        <q-td :props="props">
                                            {{ props.row.column_type }}
                                            <q-popup-edit v-model="props.row.column_type" v-slot="scope">
                                                <q-select v-model="scope.value" :options="column_typeOptions"
                                                    @update:model-value="scope.set" />
                                            </q-popup-edit>
                                        </q-td>
                                    </template>
                                    <template v-slot:body-cell-column_comment="props">
                                        <q-td :props="props">
                                            {{ props.row.column_comment }}
                                            <q-popup-edit v-model="props.row.column_comment" v-slot="scope">
                                                <q-input v-model="scope.value" dense autofocus
                                                    @keyup.enter="scope.set" />
                                            </q-popup-edit>
                                        </q-td>
                                    </template>
                                    <template v-slot:body-cell-column_default="props">
                                        <q-td :props="props">
                                            {{ props.row.column_default }}
                                            <q-popup-edit v-model="props.row.column_default" v-slot="scope">
                                                <q-input v-model="scope.value" dense autofocus
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

                <div class="row justify-center" style="margin-top: 10px">
                    <q-btn color="primary" @click="handleGen">
                        {{ $t('Gen') + $t('Plugin') }}
                    </q-btn>
                </div>
            </div>
        </q-form>
    </q-page>
</template>

<script setup>
import { postBlobAction } from 'src/api/manage'
import { computed, ref } from 'vue';
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import { fromByteArray } from 'ipaddr.js';

const $q = useQuasar()
const { t } = useI18n()
const url = 'gen-plugin/gen-plugin'

const inputModel = ref('')
const modelStruct = ref('')
const form = ref({
    plugin_name: '',
    plugin_code: '',
    with_model: false,
    plugin_model: [],
})
const columns = [
    { name: 'column_name', label: '字段名', field: 'column_name', align: 'center' },
    { name: 'column_type', label: '数据类型', field: 'column_type', align: 'center' },
    { name: 'column_comment', label: '字段描述', field: 'column_comment', align: 'center' },
    { name: 'column_default', label: '默认值', field: 'column_default', align: 'center' },
    { name: 'actions', label: '操作', field: 'actions', align: 'center' },
]
const column_typeOptions = [
    'string', 'uint', 'time.Time', 'float64'
]

const addColumn = (item) => {
    item.column_list.push({
        column_name: '',
        column_type: '',
        column_comment: '',
        column_default: '',
    })
}

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
        })
        inputModel.value = ''
        modelStruct.value = model
    }
}

const genPluginForm = ref(null)
const handleGen = async () => {
    const success = await genPluginForm.value.validate()
    if (success) {
        console.log(form.value)
        const data = await postBlobAction(url, form.value)
        if (data.headers?.success === 'false') {
            return
        } else {
            $q.notify({
                type: 'positive',
                message: '插件生成成功，开始下载！',
            })
        }
        const blob = new Blob([data])
        const fileName = 'gqa-gen-plugin.zip'
        if ('download' in document.createElement('a')) {
            // 不是IE浏览器
            const url = window.URL.createObjectURL(blob)
            const link = document.createElement('a')
            link.style.display = 'none'
            link.href = url
            link.setAttribute('download', fileName)
            document.body.appendChild(link)
            link.click()
            document.body.removeChild(link)
            window.URL.revokeObjectURL(url)
        } else {
            // IE 10+
            window.navigator.msSaveBlob(blob, fileName)
        }
    } else {
        $q.notify({
            type: 'negative',
            message: t('FixForm'),
        })
    }
}
</script>

<style lang="scss" scoped>
.tab-close {
    display: inline-flex;
    font-size: 1rem;
    border-radius: 0.2rem;
    opacity: 0.5;
    transition: all 0.3s;

    &:hover {
        opacity: 1;
    }
}
</style>

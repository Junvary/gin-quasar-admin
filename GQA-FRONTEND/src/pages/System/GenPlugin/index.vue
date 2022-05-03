<template>
    <q-page padding>
        <q-form ref="genPluginForm">
            <div class="q-gutter-md">
                <div class="row q-gutter-md">
                    <q-input class="col" v-model="form.pluginName" :label="$t('Plugin') + $t('Name')" lazy-rules
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                    <q-input class="col" v-model="form.pluginCode" :label="$t('Plugin') + $t('Code')" lazy-rules
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]" />
                </div>
                <div class="row justify-center">
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

const $q = useQuasar()
const { t } = useI18n()
const url = 'gen-plugin/gen-plugin'
const form = ref({
    plugin_name: '',
    plugin_code: '',
})

const genPluginForm = ref(null)
const handleGen = async () => {
    const success = await genPluginForm.value.validate()
    if (success) {
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

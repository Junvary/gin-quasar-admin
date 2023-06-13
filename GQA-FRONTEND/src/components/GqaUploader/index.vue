<template>
    <q-btn push color="primary" label="导入数据">
        <q-popup-proxy v-model="showUploader">
            <q-banner>
                <q-file v-model="dataFile" label="选择文件" max-files="1" @rejected="rejected">
                    <template v-slot:after v-if="dataFile">
                        <q-btn dense flat color="primary" icon="cloud_upload" @click="handleUpload" />
                    </template>
                </q-file>
            </q-banner>
        </q-popup-proxy>
    </q-btn>
</template>

<script setup>
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import { ref, toRefs } from 'vue';
import { postAction } from 'src/api/manage'

const $q = useQuasar()
const { t } = useI18n()
const dataFile = ref(null)
const showUploader = ref(false)

const props = defineProps({
    url: {
        type: String,
        required: true,
    },

})
const { url } = toRefs(props)

const rejected = (rejectedEntries) => {
    $q.notify({
        type: 'negative',
        message: t('FileRejected'),
    })
}

const emit = defineEmits(['importFinish'])

const handleUpload = () => {
    if (!dataFile.value) {
        $q.notify({
            type: 'negative',
            message: t('PleaseSelectFile'),
        })
        return
    }
    let form = new FormData()
    form.append('file', dataFile.value)
    postAction(url.value, form).then((res) => {
        if (res.code === 1) {
            dataFile.value = null
            showUploader.value = false
            $q.notify({
                type: 'positive',
                message: t('Upload') + t('Success'),
            })
            emit('importFinish')
        }
    })
}
</script>
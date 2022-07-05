<template>
    <div>
        <q-uploader style="width: 100%" :multiple="multiple" :factory="factoryFn"
            :label="title === '' ? $t('Upload') + $t('File') : title" @uploaded="uploaded" :accept="gqaBackend.fileExt"
            :max-file-size="gqaBackend.fileMaxSize * 1024 * 1024" @failed="failed" @finish="finish" @removed="removed"
            @rejected="rejected" :color="color">
            <template v-slot:list="scope">
                <q-list separator>
                    <q-item v-for="file in scope.files" :key="file.name" v-ripple>
                        <q-item-section>
                            <q-item-label class="full-width ellipsis">
                                {{ file.name }}
                            </q-item-label>

                            <q-item-label caption :class="file.__status === 'uploaded' ? 'text-positive' : ''">
                                <q-badge>
                                    {{ $t('Size') }}
                                </q-badge>:{{ file.__sizeLabel }}
                                <q-badge>
                                    {{ $t('Progress') }}
                                </q-badge>:{{ file.__progressLabel }}
                                <q-badge>
                                    {{ $t('Status') }}
                                </q-badge>: {{ file.__status }}
                            </q-item-label>
                        </q-item-section>

                        <q-item-section v-if="file.__img" thumbnail class="gt-xs">
                            <img :src="file.__img.src">
                        </q-item-section>

                        <q-item-section top side>
                            <q-btn class="gt-xs" size="12px" flat dense round icon="delete"
                                @click="scope.removeFile(file)" />
                        </q-item-section>
                    </q-item>

                </q-list>
            </template>
        </q-uploader>
        <q-list bordered separator v-if="fileList.length">
            <q-item clickable active v-ripple v-for="(item, index) in fileList" :key="index" @click="handleDownload(item)"
                active-class="text-primary">
                <q-item-section>
                    {{ item.filename }}
                </q-item-section>
                <q-item-section avatar>
                    <q-btn round push dense color="primary" icon="delete" @click.stop="deleteOne(item)" />
                </q-item-section>
            </q-item>
        </q-list>
    </div>
</template>

<script setup>
import useCommon from 'src/composables/useCommon';
import { downloadAction } from 'src/api/manage'
import { useUserStore } from 'src/stores/user'
import { ref, onMounted, computed, watch, toRefs } from 'vue';
import { useQuasar } from 'quasar';
import { useI18n } from 'vue-i18n';

const userStore = useUserStore()
const $q = useQuasar()
const { t } = useI18n()
const { gqaBackend } = useCommon()
const token = computed(() => userStore.GetToken())

const props = defineProps({
    title: {
        type: String,
        required: false,
        default: '',
    },
    attachment: {
        type: String,
        required: false,
        default: '',
    },
    multiple: {
        type: Boolean,
        required: false,
        default: true,
    },
    color: {
        type: String,
        required: false,
        default: 'primary',
    },
})
const { title, attachment, multiple, color } = toRefs(props)

const uploadUrl = process.env.API + 'upload/upload-file'

const fileList = ref([])
watch(props, () => {
    if (attachment.value && attachment.value !== '') {
        fileList.value = JSON.parse(attachment.value)
    }
})
onMounted(() => {
    if (attachment.value && attachment.value !== '') {
        fileList.value = JSON.parse(attachment.value)
    }
})

const factoryFn = (files) => {
    return {
        url: uploadUrl,
        headers: [{ name: 'Gqa-Token', value: token.value }],
        fieldName: 'file',
        method: 'POST',
    }
}
const uploaded = (info) => {
    $q.notify({
        type: 'positive',
        message: t('Upload') + t('Success'),
    })
    const res = JSON.parse(info.xhr.response)
    fileList.value.push({
        filename: info.files[0].name,
        fileUrl: res.data.records,
    })
}
const failed = (info) => {
    $q.notify({
        type: 'negative',
        message: t('Upload') + t('Failed'),
    })
}
const removed = (files) => {
    files.forEach((item) => {
        fileList.value.splice(fileList.value.map((i) => i.filename).indexOf(item.name), 1)
    })
}
const rejected = (rejectedEntries) => {
    $q.notify({
        type: 'negative',
        message: t('SizeOrExtError'),
    })
}
const handleDownload = (item) => {
    downloadAction(item.fileUrl.substring(11), item.filename)
}
const emit = defineEmits(['update:attachment'])
const finish = () => {
    if (fileList.value.length) {
        emit('update:attachment', JSON.stringify(fileList.value))
    } else {
        emit('update:attachment', '')
    }
}
const deleteOne = (item) => {
    fileList.value = fileList.value.filter((file) => file.fileUrl !== item.fileUrl)
    if (fileList.value.length) {
        emit('update:attachment', JSON.stringify(fileList.value))
    } else {
        emit('update:attachment', '')
    }
}
</script>

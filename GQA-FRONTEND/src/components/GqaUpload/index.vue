<template>
    <div>
        <q-list bordered separator v-if="showAttachmentList">
            <q-item-label header class="row items-center justify-between">
                <span>附件列表:</span>
                <q-btn dense color="negative" @click="newUpload">重新上传</q-btn>
            </q-item-label>
            <q-separator />
            <q-item clickable v-ripple v-for="(item, index) in fileList" :key="index" @click="handleDownload(item)">
                <q-item-section>
                    {{item.filename}}
                </q-item-section>
            </q-item>
        </q-list>
        <q-uploader v-else style="width: 100%" :multiple="multiple" :factory="factoryFn" :label="title"
            @uploaded="uploaded" :accept="gqaBackend.fileExt" :max-file-size="gqaBackend.fileMaxSize*1024*1024"
            @failed="failed" @start="start" @finish="finish" @removed="removed" @rejected="rejected">
            <template v-slot:list="scope">
                <q-list separator>
                    <q-item v-for="file in scope.files" :key="file.name">
                        <q-item-section>
                            <q-item-label class="full-width ellipsis">
                                {{ file.name }}
                            </q-item-label>

                            <q-item-label caption :class="file.__status === 'uploaded' ? 'text-positive': ''">
                                <q-badge>大小</q-badge>:{{ file.__sizeLabel }}
                                <q-badge>进度</q-badge>:{{ file.__progressLabel }}
                                <q-badge>状态</q-badge>: {{ file.__status  }}
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
    </div>

</template>

<script>
import { mapGetters } from 'vuex'
import { gqaBackendMixin } from 'src/mixins/gqaBackendMixin'
import { downloadAction } from 'src/api/manage'

export default {
    name: 'GqaUpload',
    mixins: [gqaBackendMixin],
    props: {
        title: {
            type: String,
            required: false,
            default: '上传附件',
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
    },
    computed: {
        ...mapGetters({
            token: 'user/token',
        }),
    },
    data() {
        return {
            uploadUrl: '/gqa-api/upload/file',
            fileList: [],
            showAttachmentList: false,
        }
    },
    created() {
        if (this.attachment && this.attachment !== '') {
            this.fileList = JSON.parse(this.attachment)
            this.showAttachmentList = true
        } else {
            this.newUpload()
        }
    },
    methods: {
        factoryFn(files) {
            return {
                url: this.uploadUrl,
                headers: [{ name: 'Gqa-Token', value: this.token }],
                fieldName: 'file',
                method: 'POST',
            }
        },
        uploaded(info) {
            this.$q.notify({
                type: 'positive',
                message: info.files[0].name + '上传成功！',
            })
            const res = JSON.parse(info.xhr.response)
            this.fileList.push({
                filename: info.files[0].name,
                fileUrl: res.data.records,
            })
        },
        failed(info) {
            this.$q.notify({
                type: 'negative',
                message: info.files[0].name + '上传失败！',
            })
        },
        start() {},
        finish() {
            this.$emit('update:attachment', JSON.stringify(this.fileList))
        },
        removed(files) {
            files.forEach((item) => {
                this.fileList.splice(this.fileList.map((i) => i.filename).indexOf(item.name), 1)
            })
        },
        rejected(rejectedEntries) {
            this.$q.notify({
                type: 'negative',
                message: '文件重复或大小/类型不被允许，请联系管理员！',
            })
        },
        newUpload() {
            this.fileList = []
            this.showAttachmentList = false
        },
        handleDownload(item) {
            downloadAction(item.fileUrl.substring(11), item.filename)
        },
    },
}
</script>
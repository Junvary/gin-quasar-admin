<template>
    <div>
        <q-uploader style="width: 100%" :multiple="multiple" :factory="factoryFn"
            :label="title === '' ? $t('UploadFiles') : title " @uploaded="uploaded" :accept="gqaBackend.fileExt"
            :max-file-size="gqaBackend.fileMaxSize*1024*1024" @failed="failed" @finish="finish" @removed="removed"
            @rejected="rejected" :color="color">
            <template v-slot:list="scope">
                <q-list separator>
                    <q-item v-for="file in scope.files" :key="file.name">
                        <q-item-section>
                            <q-item-label class="full-width ellipsis">
                                {{ file.name }}
                            </q-item-label>

                            <q-item-label caption :class="file.__status === 'uploaded' ? 'text-positive': ''">
                                <q-badge>
                                    {{ $t('Size') }}
                                </q-badge>:{{ file.__sizeLabel }}
                                <q-badge>
                                    {{ $t('Progress')  }}
                                </q-badge>:{{ file.__progressLabel }}
                                <q-badge>
                                    {{ $t('Status') }}
                                </q-badge>: {{ file.__status  }}
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
            <q-item clickable v-ripple active v-for="(item, index) in fileList" :key="index"
                @click="handleDownload(item)" active-class="text-primary">
                <q-item-section>
                    {{item.filename}}
                </q-item-section>
                <q-item-section avatar>
                    <q-btn round push dense color="primary" icon="delete" @click.stop="deleteOne(item)" />
                </q-item-section>
            </q-item>
        </q-list>
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
    },
    computed: {
        ...mapGetters({
            token: 'user/token',
        }),
    },
    watch: {
        attachment() {
            if (this.attachment && this.attachment !== '') {
                this.fileList = JSON.parse(this.attachment)
            }
        },
    },
    data() {
        return {
            uploadUrl: process.env.API + 'upload/file',
            fileList: [],
        }
    },
    created() {
        if (this.attachment && this.attachment !== '') {
            this.fileList = JSON.parse(this.attachment)
        } else {
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
                message: info.files[0].name + this.$t('UploadSuccess'),
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
                message: info.files[0].name + this.$t('UploadFailed'),
            })
        },
        removed(files) {
            files.forEach((item) => {
                this.fileList.splice(this.fileList.map((i) => i.filename).indexOf(item.name), 1)
            })
        },
        rejected(rejectedEntries) {
            this.$q.notify({
                type: 'negative',
                message: this.$t('SizeOrExtError'),
            })
        },
        handleDownload(item) {
            downloadAction(item.fileUrl.substring(11), item.filename)
        },
        finish() {
            if (this.fileList.length) {
                this.$emit('update:attachment', JSON.stringify(this.fileList))
            } else {
                this.$emit('update:attachment', '')
            }
        },
        deleteOne(item) {
            this.fileList = this.fileList.filter((file) => file.fileUrl !== item.fileUrl)
            if (this.fileList.length) {
                this.$emit('update:attachment', JSON.stringify(this.fileList))
            } else {
                this.$emit('update:attachment', '')
            }
        },
    },
}
</script>
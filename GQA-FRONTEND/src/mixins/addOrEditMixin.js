import { getAction, postAction, putAction } from 'src/api/manage'
import { FormatDateTime } from 'src/utils/date'
import { DictOptions } from 'src/utils/dict'

export const addOrEditMixin = {
    computed: {
        formTypeName() {
            if (this.formType === 'edit') {
                return this.$t('Edit')
            } else if (this.formType === 'add') {
                return this.$t('Add')
            } else {
                return this.$t('Error')
            }
        },
        showDateTime() {
            return (datetime) => {
                return FormatDateTime(datetime)
            }
        }
    },
    data() {
        return {
            addOrEditDetail: {},
            addOrEditVisible: false,
            formType: '',
            loading: false,
            dictOptions: {},
        }
    },
    async created() {
        this.dictOptions = await DictOptions()
    },
    methods: {
        show(row) {
            this.loading = true
            this.resetDetail()
            this.addOrEditDetail = Object.assign(this.detail, row)
            this.addOrEditVisible = true
            if (!this.addOrEditDetail.id) {
                this.loading = false
            } else {
                this.handleQueryById(this.addOrEditDetail.id)
            }
        },
        async handleQueryById(id) {
            const res = await postAction(this.url.queryById, {
                id,
            })
            if (res.code === 1) {
                this.addOrEditDetail = res.data.records
            }
            this.loading = false
        },
        async handleAddOrEidt() {
            const success = await this.$refs.addOrEditForm.validate()
            if (success) {
                // if (this.addOrEditDetail.avatar) {
                //     this.addOrEditDetail.avatar = JSON.stringify(this.addOrEditDetail.avatar)
                // }
                if (this.formType === 'edit') {
                    if (this.url === undefined || !this.url.edit) {
                        this.$q.notify({
                            type: 'negative',
                            message: "请先配置url",
                        })
                        return
                    }
                    const res = await putAction(this.url.edit, this.addOrEditDetail)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                        this.addOrEditVisible = false
                    }
                } else if (this.formType === 'add') {
                    if (this.url === undefined || !this.url.add) {
                        this.$q.notify({
                            type: 'negative',
                            message: "请先配置url",
                        })
                        return
                    }
                    const res = await postAction(this.url.add, this.addOrEditDetail)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                        this.addOrEditVisible = false
                    }
                } else {
                    this.$q.notify({
                        type: 'negative',
                        message: this.$t('CanNotAddOrEdit'),
                    })
                }
                this.$emit('handleFinish')
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: this.$t('FixForm'),
                })
            }
        },
    }
}
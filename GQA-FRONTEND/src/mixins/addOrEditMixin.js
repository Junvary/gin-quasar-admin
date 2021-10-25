import { getAction, postAction, putAction } from 'src/api/manage'

export const addOrEditMixin = {
    computed: {
        formTypeName() {
            if (this.formType === 'edit') {
                return '编辑'
            } else if (this.formType === 'add') {
                return '新增'
            } else {
                return '错误'
            }
        },
    },
    data() {
        return {
            addOrEditDetail: {},
            addOrEditVisible: false,
            formType: '',
            loading: false,
            options: {},
            dictUrl: {
                list: "dict/dict--detail-list"
            }
        }
    },
    created() {
        this.options = this.$q.cookies.get("gqa-dict")
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
            const res = await postAction(this.addOrEditUrl.queryById, {
                id,
            })
            if (res.code === 1) {
                this.addOrEditDetail = res.data.info
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: res.message,
                })
            }
            this.loading = false
        },
        emitAddOrEdit() {
            this.$emit('emitAddOrEdit')
        },
        async handleAddOrEidt() {
            const success = await this.$refs.addOrEditForm.validate()
            if (success) {
                if (this.formType === 'edit') {
                    const res = await putAction(this.addOrEditUrl.edit, this.addOrEditDetail)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                        this.addOrEditVisible = false
                    } else {
                        this.$q.notify({
                            type: 'negative',
                            message: res.message,
                        })
                    }
                } else if (this.formType === 'add') {
                    const res = await postAction(this.addOrEditUrl.add, this.addOrEditDetail)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                        this.addOrEditVisible = false
                    } else {
                        this.$q.notify({
                            type: 'negative',
                            message: res.message,
                        })
                    }
                } else {
                    this.$q.notify({
                        type: 'negative',
                        message: '无法新增或编辑！',
                    })
                }
                this.$emit('handleFinish')
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: '请完善表格信息！',
                })
            }
        },
    }
}
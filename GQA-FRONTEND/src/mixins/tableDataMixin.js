import { postAction, deleteAction } from "src/api/manage"

export const tableDataMixin = {
    data() {
        return {
            loading: false,
            queryParams: {},
            tableData: [],
            pagination: {
                sortBy: 'sort',
                descending: false,
                page: 1,
                rowsPerPage: 10,
            },
            pageOptions: [10, 30, 50, 100]
        }
    },
    methods: {
        async getTableData() {
            await this.onRequest({ pagination: this.pagination })
        },
        async onRequest(props) {
            if (this.url === undefined || !this.url.list) {
                this.$q.notify({
                    type: 'negative',
                    message: "请先配置url",
                })
                return
            }
            this.tableData = []
            this.loading = true
            // 组装分页和过滤条件
            const params = {}
            params.sortBy = props.pagination.sortBy
            params.desc = props.pagination.descending
            params.page = props.pagination.page
            params.pageSize = props.pagination.rowsPerPage
            const allParams = Object.assign({}, params, this.queryParams)
            // 带参数请求数据
            await postAction(this.url.list, allParams).then(res => {
                if (res.code === 1) {
                    // 最终要把分页给同步掉
                    this.pagination = props.pagination
                    // 并且加入总数字段
                    this.pagination.rowsNumber = res.data.total
                    this.tableData = res.data.records
                }
            }).finally(() => {
                this.loading = false
            })
        },
        handleSearch() {
            this.getTableData()
        },
        resetSearch() {
            this.queryParams = {}
            this.getTableData()
        },
        handleDelete(row) {
            if (this.url === undefined || !this.url.delete) {
                this.$q.notify({
                    type: 'negative',
                    message: "请先配置url",
                })
                return
            }
            this.$q
                .dialog({
                    title: '确定删除？',
                    message: `你确定要删除此项吗？`,
                    cancel: true,
                    persistent: true,
                })
                .onOk(async () => {
                    const res = await deleteAction(this.url.delete, {
                        id: row.id,
                    })
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                    }
                    this.getTableData()
                })
        },
        showAddForm(row) {
            this.$refs.addOrEditDialog.formType = 'add'
            this.$refs.addOrEditDialog.show(row)
        },
        showEditForm(row) {
            this.$refs.addOrEditDialog.formType = 'edit'
            this.$refs.addOrEditDialog.show(row)
        },
        handleFinish() {
            this.getTableData()
        }

    }
}
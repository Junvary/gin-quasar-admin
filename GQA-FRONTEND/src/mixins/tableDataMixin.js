import { getAction, postAction, deleteAction, putAction } from "src/api/manage"

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
                rowsNumber: 0,
                options: [10, 30, 50, 100]
            },
        }
    },
    computed: {
        queryAllParams() {
            const params = { ...this.queryParams }
            params.sortBy = this.pagination.sortBy
            params.desc = this.pagination.descending
            params.page = this.pagination.page
            params.pageSize = this.pagination.rowsPerPage
            return params
        }
    },
    methods: {
        getTableData() {
            if (this.url === undefined || !this.url.list) {
                this.$q.notify({
                    type: 'negative',
                    message: "请先配置url",
                })
                return
            }
            this.tableData = []
            this.pagination.rowsNumber = 0
            this.loading = true
            postAction(this.url.list, {
                ...this.queryAllParams
            }).then(res => {
                this.pagination.rowsNumber = res.data.total
                this.tableData = res.data.records
            }).finally(() => {
                this.loading = false
            })
        },
        handleSearch() {
            this.getTableData()
        },
        resetSearch() {
            this.getTableData()
        },
        onRequest(props) {
            const { page, rowsPerPage, sortBy, descending } = props.pagination
            this.pagination.page = page
            this.pagination.rowsPerPage = rowsPerPage
            this.pagination.sortBy = sortBy
            this.pagination.descending = descending
            this.getTableData()
        },
        handleDelete(row) {
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
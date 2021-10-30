import { getAction, postAction, deleteAction, putAction } from "src/api/manage"

export const tableDataMixin = {
    data() {
        return {
            loading: false,
            queryParams: {},
            tableData: [],
            pagination: {
                sortBy: 'desc',
                descending: false,
                page: 1,
                rowsPerPage: 10,
                rowsNumber: 0
            },
        }
    },
    methods: {
        getTableData() {
            this.loading = true
            const params = Object.assign({}, this.queryParams)
            postAction(this.url.list, {
                page: this.pagination.page,
                pageSize: this.pagination.rowsPerPage,
                ...params,
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
        showAddForm() {
            this.$refs.addOrEditDialog.formType = 'add'
            this.$refs.addOrEditDialog.show({})
        },
        showEditForm(row) {
            this.$refs.addOrEditDialog.formType = 'edit'
            this.$refs.addOrEditDialog.show(row)
        },
        emitAddOrEdit() {
            this.$refs.addOrEditDialog.handleAddOrEidt()
        },
        handleFinish() {
            this.getTableData()
        }

    }
}
import { ref, onMounted } from 'vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import { postAction } from 'src/api/manage'
import { DictOptions } from 'src/utils/dict'
import useCommon from './useCommon'

export default function useTableData(url) {
    const { t } = useI18n()
    const $q = useQuasar()
    const dictOptions = ref({})
    onMounted(async () => {
        dictOptions.value = await DictOptions()
    })
    const loading = ref(false)
    const tableData = ref([])
    const recordDetailDialog = ref(null)
    const pagination = ref({
        sortBy: 'sort',
        descending: false,
        page: 1,
        rowsPerPage: 10,
    })
    const queryParams = ref({})
    const pageOptions = ref([10, 30, 50, 100])

    const showAddForm = () => {
        recordDetailDialog.value.formType = 'add'
        recordDetailDialog.value.show()
    }
    const showEditForm = (row) => {
        recordDetailDialog.value.formType = 'edit'
        recordDetailDialog.value.show(row)
    }
    const onRequest = async (props) => {
        if (url === undefined || !url.list) {
            $q.notify({
                type: 'negative',
                message: t('UrlNotConfig'),
            })
            return
        }
        tableData.value = []
        loading.value = true
        // 组装分页和过滤条件
        const params = {}
        params.sort_by = props.pagination.sortBy
        params.desc = props.pagination.descending
        params.page = props.pagination.page
        params.page_size = props.pagination.rowsPerPage
        const allParams = Object.assign({}, params, queryParams.value)
        // 带参数请求数据
        await postAction(url.list, allParams).then(res => {
            if (res.code === 1) {
                // 最终要把分页给同步掉
                pagination.value = props.pagination
                // 并且加入总数字段
                pagination.value.rowsNumber = res.data.total
                tableData.value = res.data.records
            }
        }).finally(() => {
            loading.value = false
        })
    }
    const getTableData = () => onRequest({ pagination: pagination.value, queryParams: queryParams.value })
    const handleSearch = () => {
        getTableData()
    }
    const resetSearch = () => {
        queryParams.value = {}
        getTableData()
    }
    const handleFinish = () => {
        getTableData()
    }
    const handleDelete = (row) => {
        if (!url || !url.delete) {
            $q.notify({
                type: 'negative',
                message: "请先配置url",
            })
            return
        }
        $q.dialog({
            title: t('Confirm'),
            message: t('Confirm') + t('Delete') + '?',
            cancel: true,
            persistent: true,
        }).onOk(async () => {
            const res = await postAction(url.delete, {
                id: row.id,
            })
            if (res.code === 1) {
                $q.notify({
                    type: 'positive',
                    message: res.message,
                })
            }
            getTableData()
        })
    }
    // 引入useCommon中的方法
    const { showDateTime, gqaFrontend, gqaBackend, GqaDictShow, GqaShowName, GqaAvatar, } = useCommon()
    return {
        showDateTime,
        gqaBackend,
        gqaFrontend,
        dictOptions,
        pagination,
        queryParams,
        pageOptions,
        GqaDictShow,
        GqaShowName,
        GqaAvatar,
        loading,
        tableData,
        recordDetailDialog,
        showAddForm,
        showEditForm,
        onRequest,
        handleSearch,
        resetSearch,
        handleFinish,
        handleDelete,
    }
}
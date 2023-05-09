import { DictOptions } from 'src/utils/dict'
import { onMounted, ref, computed, reactive } from 'vue'
import { postAction } from 'src/api/manage'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import useCommon from './useCommon'
import { inject } from 'vue'

export default function useRecordDetail(url, emit) {
    const { t } = useI18n()
    const $q = useQuasar()
    const dictOptions = ref({})
    const bus = inject('bus')
    onMounted(async () => {
        dictOptions.value = await DictOptions()
    })
    const formType = ref('')
    const formTypeName = computed(() => {
        if (formType.value === 'edit') {
            return t('Edit')
        } else if (formType.value === 'add') {
            return t('Add')
        } else {
            return t('Error')
        }
    })
    const recordDetail = reactive({})
    const recordDetailVisible = ref(false)
    const loading = ref(false)
    const show = (row) => {
        loading.value = true
        recordDetail.value = {}
        recordDetailVisible.value = true
        if (row && row.id) {
            handleQueryById(row.id)
        } else {
            recordDetail.value = {}
            recordDetailVisible.value = true
            loading.value = false
        }
    }
    const handleQueryById = async (id) => {
        postAction(url.queryById, {
            id,
        }).then(res => {
            if (res.code === 1) {
                recordDetail.value = res.data.records
            }
        }).finally(() => {
            loading.value = false
        })
    }

    const recordDetailForm = ref(null)

    const handleAddOrEidt = async () => {
        const success = await recordDetailForm.value.validate()
        if (success) {
            if (formType.value === 'edit') {
                if (url === undefined || !url.edit) {
                    $q.notify({
                        type: 'negative',
                        message: t('UrlNotConfig'),
                    })
                    return
                }
                const res = await postAction(url.edit, recordDetail.value)
                if (res.code === 1) {
                    $q.notify({
                        type: 'positive',
                        message: res.message,
                    })
                    recordDetailVisible.value = false
                }
            } else if (formType.value === 'add') {
                if (url === undefined || !url.add) {
                    $q.notify({
                        type: 'negative',
                        message: t('UrlNotConfig'),
                    })
                    return
                }
                const res = await postAction(url.add, recordDetail.value)
                if (res.code === 1) {
                    $q.notify({
                        type: 'positive',
                        message: res.message,
                    })
                    recordDetailVisible.value = false
                }
            } else {
                $q.notify({
                    type: 'negative',
                    message: t('CanNotAddOrEdit'),
                })
            }
            emit('handleFinish')
        } else {
            $q.notify({
                type: 'negative',
                message: t('FixForm'),
            })
        }
    }
    const { showDateTime, showDate, gqaFrontend, gqaBackend, GqaDictShow, GqaShowName, } = useCommon()
    return {
        $q,
        t,
        bus,
        GqaShowName,
        GqaDictShow,
        gqaFrontend,
        gqaBackend,
        dictOptions,
        showDateTime,
        showDate,
        loading,
        show,
        formType,
        formTypeName,
        recordDetail,
        recordDetailVisible,
        handleQueryById,
        recordDetailForm,
        handleAddOrEidt,
    }
}
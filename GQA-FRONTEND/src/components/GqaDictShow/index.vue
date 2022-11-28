<template>
    <div>
        <q-chip dense text-color="white" color="positive" v-if="dictCode === 'onOff_on'">
            {{ $t(dictLabel) }}
        </q-chip>
        <q-chip dense text-color="white" color="negative" v-else-if="dictCode === 'onOff_off'">
            {{ $t(dictLabel) }}
        </q-chip>
        <span v-else-if="dictCode === 'yesNo_yes'
        || dictCode === 'yesNo_no'
        || dictCode === 'gender_male'
        || dictCode === 'gender_female'
        || dictCode === 'gender_unknown'">
            {{ $t(dictLabel) }}
        </span>
        <span v-else>
            {{ dictLabel }}
        </span>
    </div>
</template>

<script setup>
import { ref, computed, onBeforeMount, toRefs } from 'vue';
import { useQuasar } from 'quasar';
import { useStorageStore } from 'src/stores/storage';
import XEUtils from 'xe-utils'

const storageStore = useStorageStore();
const $q = useQuasar()
const props = defineProps({
    dictCode: {
        type: String,
        default: '',
    },
    withExt1: {
        type: Boolean,
        required: false,
        default: false,
    },
    ext1: {
        type: String,
        default: '',
    },
})
const { dictCode, withExt1, ext1 } = toRefs(props)

const dictList = ref([])

onBeforeMount(async () => {
    const detailLocal = $q.localStorage.getItem('gqa-dict')
    if (detailLocal) {
        dictList.value = detailLocal
    } else {
        await storageStore.SetGqaDict()
        dictList.value = $q.localStorage.getItem('gqa-dict')
    }
})

const dictLabel = computed(() => {
    if (dictCode.value !== '') {
        for (let dict in dictList.value) {
            const dictDetail = XEUtils.findTree(dictList.value[dict], item => item.dict_code === dictCode.value)
            if (dictDetail) {
                return dictDetail.item.dict_label
            }
        }
        return ""
    } else {
        return ""
    }
})
// const dictExt1 = computed(() => {
//     if (dictCode.value !== '') {
//         const codeList = dictCode.value.split(',')
//         const dict = dictList.value[dictName.value]
//         let ext1 = ''
//         for (let d of codeList) {
//             const l = dict.filter((item) => item.dict_code === d)[0].dict_ext_1
//             ext1 += l + ' '
//         }
//         return ext1
//     } else {
//         return ''
//     }
// })
</script>
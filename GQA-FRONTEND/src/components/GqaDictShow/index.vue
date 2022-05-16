<template>
    <div>
        <span v-if="dictName === 'statusOnOff'">
            <q-chip dense text-color="white" color="positive" v-if="dictCode === 'on'">{{ dictLabel }}</q-chip>
            <q-chip dense text-color="white" color="negative" v-else-if="dictCode === 'of'">{{ dictLabel }}</q-chip>
            <q-chip dense v-else>{{ dictLabel }}</q-chip>
        </span>
        <span v-else>
            {{ dictLabel }}
        </span>
        <span v-if="withExt1">
            {{ dictExt1 + ext1 }}
        </span>
    </div>
</template>

<script setup>
import { ref, computed, onBeforeMount, toRefs } from 'vue';
import { useQuasar } from 'quasar';
import { useStorageStore } from 'src/stores/storage';

const storageStore = useStorageStore();
const $q = useQuasar()
const props = defineProps({
    dictName: {
        type: String,
        required: true,
    },
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
const { dictName, dictCode, withExt1, ext1 } = toRefs(props)

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
        const codeList = dictCode.value.split(',')
        const dict = dictList.value[dictName.value]
        let label = ''
        for (let d of codeList) {
            const l = dict.filter((item) => item.dict_code === d)[0].dict_label
            label += l + ' '
        }
        return label
    } else {
        return ''
    }
})
const dictExt1 = computed(() => {
    if (dictCode.value !== '') {
        const codeList = dictCode.value.split(',')
        const dict = dictList.value[dictName.value]
        let ext1 = ''
        for (let d of codeList) {
            const l = dict.filter((item) => item.dict_code === d)[0].dict_ext_1
            ext1 += l + ' '
        }
        return ext1
    } else {
        return ''
    }
})
</script>
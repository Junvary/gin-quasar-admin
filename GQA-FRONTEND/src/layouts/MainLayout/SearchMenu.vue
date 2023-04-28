<template>
    <q-select style="width: 200px;" outlined v-if="searchVisible" dense borderless :options="searchMenu" map-options
        v-model="currentMenu" option-value="name" @update:model-value="changeRouter" use-input input-debounce="0"
        @filter="filterFn" :option-label="opt => selectOptionLabel(opt)">
    </q-select>
    <q-btn dense round flat icon="ion-ios-search" @click="showSearch">
        <q-tooltip>
            {{ $t('Search') + ' ' + $t('Menu') }}
        </q-tooltip>
    </q-btn>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { usePermissionStore } from 'src/stores/permission';
import useCommon from 'src/composables/useCommon'
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';

const permissionStore = usePermissionStore();
const router = useRouter()
const { t } = useI18n()

const { selectOptionLabel } = useCommon()

const searchVisible = ref(false)
const showSearch = () => {
    searchVisible.value = !searchVisible.value
}

const searchMenu = ref([])
onMounted(() => {
    searchMenu.value = permissionStore.searchMenu
})

const currentMenu = ref(null)
const changeRouter = (item) => {
    if (item.is_link === 'yesNo_yes') {
        window.open(item.path, '_blank')
    } else {
        router.push({ path: item.path })
    }
    searchVisible.value = false
}

const filterFn = (val, update) => {
    if (val === '') {
        update(() => {
            searchMenu.value = permissionStore.searchMenu
        })
        return
    }
    update(() => {
        searchMenu.value = permissionStore.searchMenu.filter(v => {
            return t(v.title).indexOf(val) > -1
        })
    })
}
</script>
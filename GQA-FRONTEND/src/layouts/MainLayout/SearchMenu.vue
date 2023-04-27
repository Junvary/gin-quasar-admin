<template>
    <q-select style="width: 200px;" outlined v-if="searchVisible" dense borderless :options="searchMenu" map-options
        v-model="currentMenu" option-value="name" @update:model-value="changeRouter"
        :option-label="opt => selectOptionLabel(opt)">
    </q-select>
    <q-btn dense round flat icon="ion-ios-search" @click="showSearch">
        <q-tooltip>
            {{ $t('Search') + ' ' + $t('Menu') }}
        </q-tooltip>
    </q-btn>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { usePermissionStore } from 'src/stores/permission';
import useCommon from 'src/composables/useCommon'
import { useRoute, useRouter } from 'vue-router';

const permissionStore = usePermissionStore();
const router = useRouter()
const route = useRoute()

const searchMenu = computed(() => {
    return permissionStore.searchMenu
})

const { selectOptionLabel } = useCommon()

const searchVisible = ref(false)
const showSearch = () => {
    searchVisible.value = !searchVisible.value
}

onMounted(() => {
    currentMenu.value = route.name
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
</script>
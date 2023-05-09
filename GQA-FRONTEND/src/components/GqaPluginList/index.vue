<template>
    <q-expansion-item dense expand-separator class="bg-primary text-white" expand-icon-class="text-white"
        v-if="rows && rows.length">
        <template v-slot:header>
            <q-item-section avatar>
                <q-avatar size="md" icon="install_desktop" color="white" text-color="primary" />
            </q-item-section>
            <q-item-section>
                <q-btn dense flat :label="$t('Installed') + $t('Plugin') + '（' + rows.length + '）'" />
            </q-item-section>
        </template>
        <q-table dense hide-bottom separator="cell" :rows="rows" :columns="columns" row-key="name">

            <template v-slot:body-cell-first="props">
                <q-td :props="props">
                    <q-chip dense color="primary" text-color="white">
                        {{ $t('Installed') + $t('Plugin') }}
                    </q-chip>
                </q-td>
            </template>

            <template v-slot:body-cell-memo="props">
                <q-td :props="props">
                    <q-chip dense color="primary" text-color="white">
                        {{ $t('Plugin') + $t('Describe') }}
                        <q-tooltip>
                            {{ props.row.plugin_memo }}
                        </q-tooltip>
                    </q-chip>
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props" v-if="showChoose">
                <q-td :props="props">
                    <q-radio dense name="plugin-login-layout" v-model="choosePlugin" :val="props.row.plugin_code"
                        @update:model-value="choosePluginLoginLayout(props.row.plugin_code)" />
                </q-td>
            </template>
        </q-table>
    </q-expansion-item>
</template>

<script setup>
import { computed, toRefs } from 'vue';
import { useQuasar } from 'quasar';
import { useI18n } from 'vue-i18n';

const { t } = useI18n()

const $q = useQuasar()
const props = defineProps({
    showChoose: {
        type: Boolean,
        required: false,
        default: false,
    },
    choosePlugin: {
        type: String,
        required: false,
        default: '',
    },
})
const { showChoose, choosePlugin } = toRefs(props)
const rows = computed(() => {
    if (Array.isArray($q.localStorage.getItem('gqa-pluginList'))) {
        return $q.localStorage.getItem('gqa-pluginList')
    }
    return []
})
const columns = computed(() => {
    if (showChoose.value) {
        return [
            { name: 'first', align: 'center', label: t('Installed') + t('Plugin'), field: 'first' },
            { name: 'plugin_name', align: 'center', label: t('Plugin') + t('Name'), field: 'plugin_name' },
            { name: 'plugin_code', align: 'center', label: t('Plugin') + t('Code'), field: 'plugin_code' },
            { name: 'plugin_version', align: 'center', label: t('Plugin') + t('Version'), field: 'plugin_version' },
            { name: 'memo', align: 'center', label: t('Plugin') + t('Describe'), field: 'memo' },
            { name: 'actions', align: 'center', label: t('Choose') + t('Plugin'), field: 'actions' },
        ]
    } else {
        return [
            { name: 'first', align: 'center', label: t('Installed') + t('Plugin'), field: 'first' },
            { name: 'plugin_name', align: 'center', label: t('Plugin') + t('Name'), field: 'plugin_name' },
            { name: 'plugin_code', align: 'center', label: t('Plugin') + t('Code'), field: 'plugin_code' },
            { name: 'plugin_version', align: 'center', label: t('Plugin') + t('Version'), field: 'plugin_version' },
            { name: 'memo', align: 'center', label: t('Plugin') + t('Describe'), field: 'memo' },
        ]
    }
})
const emit = defineEmits(['changeSuccess'])
const choosePluginLoginLayout = (code) => {
    const pluginCode = code.slice(7)
    const tryImport = new Promise((resolve, reject) => {
        require(`src/plugins/${pluginCode}/LoginLayout/index.vue`).default
        resolve()
    })
    tryImport.then(() => {
        $q.notify({
            type: 'positive',
            message: 'Switching succeeded',
        })
        emit('changeSuccess', code)
    }).catch(() => {
        $q.notify({
            type: 'negative',
            message: 'This plugin does not support the login page yet',
        })
        emit('changeSuccess', '')
    })
}
</script>

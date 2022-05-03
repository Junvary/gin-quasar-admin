<template>
    <q-expansion-item dense expand-separator class="bg-accent text-white" expand-icon-class="text-white"
        v-if="rows && rows.length">
        <template v-slot:header>
            <q-item-section avatar>
                <q-avatar size="md" icon="install_desktop" color="white" text-color="accent" />
            </q-item-section>
            <q-item-section>
                <q-btn dense flat :label="$t('InstalledPlugins') + '（' + rows.length + '）'" />
            </q-item-section>
        </template>
        <q-table dense hide-bottom separator="cell" :rows="rows" :columns="columns" row-key="name">

            <template v-slot:body-cell-first="props">
                <q-td :props="props">
                    <q-chip dense color="accent" text-color="white">
                        {{ $t('InstalledPlugins') }}
                    </q-chip>
                </q-td>
            </template>

            <template v-slot:body-cell-remark="props">
                <q-td :props="props">
                    <q-chip dense color="accent" text-color="white">
                        {{ $t('PluginDescription') }}
                        <q-tooltip>
                            {{ props.row.PluginRemark }}
                        </q-tooltip>
                    </q-chip>
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props" v-if="showChoose.value">
                <q-td :props="props">
                    <q-radio dense name="plugin-login-layout" v-model="choosePlugin.value" :val="props.row.PluginCode"
                        @update:model-value="choosePluginLoginLayout(props.row.PluginCode)" />
                </q-td>
            </template>
        </q-table>
    </q-expansion-item>
</template>

<script setup>
import { computed, toRefs } from 'vue';
import { useQuasar } from 'quasar';

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
            { name: 'first', align: 'center', label: this.$t('InstalledPlugins'), field: 'first' },
            { name: 'PluginName', align: 'center', label: this.$t('PluginName'), field: 'PluginName' },
            { name: 'PluginCode', align: 'center', label: this.$t('PluginCode'), field: 'PluginCode' },
            { name: 'PluginVersion', align: 'center', label: this.$t('PluginVersion'), field: 'PluginVersion' },
            { name: 'remark', align: 'center', label: this.$t('PluginDescription'), field: 'remark' },
            { name: 'actions', align: 'center', label: this.$t('PluginChoose'), field: 'actions' },
        ]
    } else {
        return [
            { name: 'first', align: 'center', label: this.$t('InstalledPlugins'), field: 'first' },
            { name: 'PluginName', align: 'center', label: this.$t('PluginName'), field: 'PluginName' },
            { name: 'PluginCode', align: 'center', label: this.$t('PluginCode'), field: 'PluginCode' },
            { name: 'PluginVersion', align: 'center', label: this.$t('PluginVersion'), field: 'PluginVersion' },
            { name: 'remark', align: 'center', label: this.$t('PluginDescription'), field: 'remark' },
        ]
    }
})
const emit = defineEmits(['changeSuccess'])
const choosePluginLoginLayout = () => {
    const tryImport = new Promise((resolve, reject) => {
        require(`src/layouts/LoginLayout/${code}/index.vue`).default
        resolve()
    })
    tryImport.then(() => {
        $q.notify({
            type: 'positive',
            message: '切换成功！',
        })
        emit('changeSuccess', code)
    }).catch(() => {
        $q.notify({
            type: 'negative',
            message: '此插件还未支持登录页面！',
        })
        emit('changeSuccess', '')
    })
}
</script>

<template>
    <q-page padding class="q-gutter-y-md">
        <q-input dense v-model="filter" label="Filter" outlined clearable style="width: 100%" />
        <q-tabs v-model="tab" dense class="text-grey" active-color="primary" indicator-color="primary" align="justify"
            narrow-indicator>
            <q-tab name="materia" label="material-icons" />
            <q-tab name="fontawesome" label="fontawesome-v6" />
            <q-tab name="eva" label="eva-icons" />
            <q-tab name="mdi" label="mdi-v6" />
            <q-tab name="ionicons" label="ionicons-v4" />
        </q-tabs>
        <q-tab-panels v-model="tab" animated>
            <q-tab-panel name="materia">
                <div class="row" style="width: 95%">
                    <q-icon-picker v-model="materiaData.value" v-model:model-pagination="materiaData.pagination"
                        icon-set="material-icons" :filter="filter" style="height: 60vh" size="30px" tooltips
                        @click="copyIcon(materiaData.value)" />
                </div>
            </q-tab-panel>
            <q-tab-panel name="fontawesome">
                <div class="row" style="width: 95%">
                    <q-icon-picker v-model="fontawesomeData.value" v-model:model-pagination="fontawesomeData.pagination"
                        icon-set="fontawesome-v5" :filter="filter" style="height: 60vh" size="30px" tooltips
                        @click="copyIcon(fontawesomeData.value)" />
                </div>
            </q-tab-panel>
            <q-tab-panel name="eva">
                <div class="row" style="width: 95%">
                    <q-icon-picker v-model="evaData.value" v-model:model-pagination="evaData.pagination"
                        icon-set="eva-icons" :filter="filter" style="height: 60vh" size="30px" tooltips
                        @click="copyIcon(evaData.value)" />
                </div>
            </q-tab-panel>
            <q-tab-panel name="mdi">
                <div class="row" style="width: 95%">
                    <q-icon-picker v-model="mdiData.value" v-model:model-pagination="mdiData.pagination"
                        icon-set="mdi-v6" :filter="filter" style="height: 60vh" size="30px" tooltips
                        @click="copyIcon(mdiData.value)" />
                </div>
            </q-tab-panel>
            <q-tab-panel name="ionicons">
                <div class="row" style="width: 95%">
                    <q-icon-picker v-model="ioniconsData.value" v-model:model-pagination="ioniconsData.pagination"
                        icon-set="ionicons-v4" :filter="filter" style="height: 60vh" size="30px" tooltips
                        @click="copyIcon(ioniconsData.value)" />
                </div>
            </q-tab-panel>
        </q-tab-panels>
    </q-page>
</template>

<script setup>
import { ref } from 'vue';
import { copyToClipboard, useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'

const $q = useQuasar();
const { t } = useI18n();
const tab = ref('materia');
const filter = ref('')
const materiaData = ref({
    value: '',
    pagination: {
        itemsPerPage: 100,
        page: 0
    }
});
const fontawesomeData = ref({
    value: '',
    pagination: {
        itemsPerPage: 100,
        page: 0
    }
});
const evaData = ref({
    value: '',
    pagination: {
        itemsPerPage: 100,
        page: 0
    }
})
const mdiData = ref({
    value: '',
    pagination: {
        itemsPerPage: 100,
        page: 0
    }
})
const ioniconsData = ref({
    value: '',
    pagination: {
        itemsPerPage: 100,
        page: 0
    }
})
const copyIcon = (item) => {
    if (item) {
        copyToClipboard(item).then(() => {
            $q.notify({
                type: 'positive',
                message: t('CopyToClipboard') + ' ' + t('Success') + ': ' + item,
            })
            materiaData.value.value = '';
            fontawesomeData.value.value = '';
            evaData.value.value = '';
            mdiData.value.value = '';
            ioniconsData.value.value = ''
        }).catch(() => {
            $q.notify({
                type: 'negative',
                message: t('CopyToClipboard') + ' ' + t('Failed'),
            })
        })
    }

}
</script>
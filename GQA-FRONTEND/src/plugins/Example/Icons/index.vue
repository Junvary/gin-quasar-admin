<template>
    <q-page padding>
        <q-tabs v-model="tab" dense class="text-grey" active-color="primary" indicator-color="primary" align="justify"
            narrow-indicator>
            <q-tab name="materia" :label="$t('Material')" />
            <q-tab name="fontawesome" :label="$t('FontAwesome')" />
        </q-tabs>

        <q-separator />

        <q-tab-panels v-model="tab" animated>
            <q-tab-panel name="materia">
                <div class="row" style="width: 95%">
                    <q-input v-model="materiaData.filter" label="Filter" outlined clearable style="width: 100%" />
                    <q-icon-picker v-model="materiaData.value" v-model:model-pagination="materiaData.pagination"
                        icon-set="material-icons" :filter="materiaData.filter" style="height: 60vh" size="30px" tooltips
                        @click="copyIcon(materiaData.value)" />
                </div>
            </q-tab-panel>

            <q-tab-panel name="fontawesome">
                <div class="row" style="width: 95%">
                    <q-input v-model="fontawesomeData.filter" label="Filter" outlined clearable style="width: 100%" />
                    <q-icon-picker v-model="fontawesomeData.value" v-model:model-pagination="fontawesomeData.pagination"
                        icon-set="fontawesome-v5" :filter="fontawesomeData.filter" style="height: 60vh" size="30px"
                        tooltips @click="copyIcon(fontawesomeData.value)" />
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
const materiaData = ref({
    value: '',
    filter: '',
    pagination: {
        itemsPerPage: 100,
        page: 0
    }
});
const fontawesomeData = ref({
    value: '',
    filter: '',
    pagination: {
        itemsPerPage: 100,
        page: 0
    }
});
const copyIcon = (item) => {
    if (item) {
        copyToClipboard(item).then(() => {
            $q.notify({
                type: 'positive',
                message: t('CopyToClipboard') + ' ' + t('Success') + ': ' + item,
            })
            materiaData.value.value = '';
            fontawesomeData.value.value = '';
        }).catch(() => {
            $q.notify({
                type: 'negative',
                message: t('CopyToClipboard') + ' ' + t('Failed'),
            })
        })
    }

}
</script>

<style lang="scss" scoped>

</style>

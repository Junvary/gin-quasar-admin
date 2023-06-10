<template>
    <q-page padding>
        <q-card style="width: 100%;">
            <q-card-section>
                <q-input dense v-model="filterKey" label="Filter" outlined clearable style="width: 100%" />
            </q-card-section>
            <q-card-section class="row q-gutter-xs">
                <q-btn-group>
                    <q-btn dense outline color="primary" label="material-icons-outlined"
                        @click="setIcon(materialIconsOutlined)" />
                    <q-btn dense outline color="primary" label="material-icons-round"
                        @click="setIcon(materialIconsRound)" />
                    <q-btn dense outline color="primary" label="material-icons-sharp"
                        @click="setIcon(materialIconsSharp)" />
                    <q-btn dense outline color="primary" label="material-icons" @click="setIcon(materialIcons)" />
                    <q-btn dense outline color="primary" label="material-symbols-outlined"
                        @click="setIcon(materialSymbolsOutlined)" />
                    <q-btn dense outline color="primary" label="material-symbols-rounded"
                        @click="setIcon(materialSymbolsRounded)" />
                    <q-btn dense outline color="primary" label="mmaterial-symbols-sharp"
                        @click="setIcon(materialSymbolsSharp)" />
                </q-btn-group>
                <q-btn-group>
                    <q-btn dense outline color="primary" label="eva-icons" @click="setIcon(evaIcons)" />
                    <q-btn dense outline color="primary" label="fontawesome-v5" @click="setIcon(fontawesomeV5)" />
                    <q-btn dense outline color="primary" label="fontawesome-v6" @click="setIcon(fontawesomeV6)" />
                    <q-btn dense outline color="primary" label="ionicons-v4" @click="setIcon(ioniconsV4)" />
                    <q-btn dense outline color="primary" label="ionicons-v5" @click="setIcon(ioniconsV5)" />
                    <q-btn dense outline color="primary" label="ionicons-v6" @click="setIcon(ioniconsV6)" />
                    <q-btn dense outline color="primary" label="ionicons-v7" @click="setIcon(ioniconsV7)" />
                    <q-btn dense outline color="primary" label="line-awesome" @click="setIcon(lineAwesome)" />
                    <q-btn dense outline color="primary" label="mdi-v4" @click="setIcon(mdiV4)" />
                    <q-btn dense outline color="primary" label="mdi-v5" @click="setIcon(mdiV5)" />
                    <q-btn dense outline color="primary" label="mdi-v6" @click="setIcon(mdiV6)" />
                    <q-btn dense outline color="primary" label="mdi-v7" @click="setIcon(mdiV7)" />
                </q-btn-group>

            </q-card-section>
            <q-card-section class="row q-gutter-sm">
                <q-btn class="col-auto" v-for="(item, key, index) in dataObj" :icon="item" @click="copyIcon(key)">
                    <q-tooltip anchor="top middle" self="center middle">
                        {{ key }}
                    </q-tooltip>
                </q-btn>
            </q-card-section>
            <q-card-actions align="center">
                <q-pagination v-model="pagination.page" :max="pagination.max" />
            </q-card-actions>
        </q-card>
    </q-page>
</template>

<script setup>
import { computed, ref } from 'vue';
import { copyToClipboard, useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import { filterObj } from 'src/utils/object'
import * as evaIcons from '@quasar/extras/eva-icons'
import * as fontawesomeV5 from '@quasar/extras/fontawesome-v5'
import * as fontawesomeV6 from '@quasar/extras/fontawesome-v6'
import * as ioniconsV4 from '@quasar/extras/ionicons-v4'
import * as ioniconsV5 from '@quasar/extras/ionicons-v5'
import * as ioniconsV6 from '@quasar/extras/ionicons-v6'
import * as ioniconsV7 from '@quasar/extras/ionicons-v7'
import * as lineAwesome from '@quasar/extras/line-awesome'
import * as materialIconsOutlined from '@quasar/extras/material-icons-outlined'
import * as materialIconsRound from '@quasar/extras/material-icons-round'
import * as materialIconsSharp from '@quasar/extras/material-icons-sharp'
import * as materialIcons from '@quasar/extras/material-icons'
import * as materialSymbolsOutlined from '@quasar/extras/material-symbols-outlined'
import * as materialSymbolsRounded from '@quasar/extras/material-symbols-rounded'
import * as materialSymbolsSharp from '@quasar/extras/material-symbols-sharp'
import * as mdiV4 from '@quasar/extras/mdi-v4'
import * as mdiV5 from '@quasar/extras/mdi-v5'
import * as mdiV6 from '@quasar/extras/mdi-v6'
import * as mdiV7 from '@quasar/extras/mdi-v7'

const $q = useQuasar();
const { t } = useI18n();
const itemsPerPage = ref(200)
const filterKey = ref('')
const iconSet = ref(materialIcons)
const iconSetKeyList = ref(Object.keys(materialIcons))

const setIcon = (quasarIcon) => {
    iconSet.value = quasarIcon
    iconSetKeyList.value = Object.keys(quasarIcon)
}

const pagination = ref({
    page: 1,
    max: Math.ceil(iconSetKeyList.value.length / itemsPerPage.value)
})

const dataObj = computed(() => {
    if (filterKey.value !== "") {
        const filterArray = iconSetKeyList.value.filter(item => {
            return item.toLowerCase().includes(filterKey.value.toLowerCase())
        })
        pagination.value = {
            page: pagination.value.page,
            max: Math.ceil(filterArray.length / itemsPerPage.value)
        }
        const keyArray = filterArray.slice((pagination.value.page - 1) * itemsPerPage.value, (pagination.value.page - 1) * itemsPerPage.value + itemsPerPage.value - 1)
        return filterObj(iconSet.value, keyArray)
    }
    pagination.value = {
        page: pagination.value.page,
        max: Math.ceil(iconSetKeyList.value.length / itemsPerPage.value)
    }
    const keyArray = iconSetKeyList.value.slice((pagination.value.page - 1) * itemsPerPage.value, (pagination.value.page - 1) * itemsPerPage.value + itemsPerPage.value - 1)
    return filterObj(iconSet.value, keyArray)
})

const copyIcon = (key) => {
    if (key) {
        copyToClipboard(key).then(() => {
            $q.notify({
                type: 'positive',
                message: t('CopyToClipboard') + ' ' + t('Success') + ': ' + key,
            })
        }).catch(() => {
            $q.notify({
                type: 'negative',
                message: t('CopyToClipboard') + ' ' + t('Failed'),
            })
        })
    }
}

</script>
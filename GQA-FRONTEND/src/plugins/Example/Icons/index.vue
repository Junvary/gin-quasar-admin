<template>
    <q-page padding>
        <q-card style="width: 100%;">
            <q-card-section>
                <q-input dense v-model="filterKey" label="Filter" outlined clearable style="width: 100%" />
            </q-card-section>
            <q-card-section class="row q-gutter-md">
                <q-radio dense v-model="iconSet" :val="materialIconsOutlined" label="material-icons-outlined"
                    @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="materialIconsRound" label="material-icons-round" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="materialIconsSharp" label="material-icons-sharp" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="materialIcons" label="material-icons" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="materialSymbolsOutlined" label="material-symbols-outlined"
                    @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="materialSymbolsRounded" label="material-symbols-rounded"
                    @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="materialSymbolsSharp" label="mmaterial-symbols-sharp"
                    @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="evaIcons" label="eva-icons" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="fontawesomeV5" label="fontawesome-v5" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="fontawesomeV6" label="fontawesome-v6" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="ioniconsV4" label="ionicons-v4" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="ioniconsV5" label="ionicons-v5" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="ioniconsV6" label="ionicons-v6" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="ioniconsV7" label="ionicons-v7" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="lineAwesome" label="line-awesome" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="mdiV4" label="mdi-v4" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="mdiV5" label="mdi-v5" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="mdiV6" label="mdi-v6" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="mdiV7" label="mdi-v7" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="themify" label="themify" @click="setIcon" />
                <q-radio dense v-model="iconSet" :val="bootstrapIcons" label="bootstrap-icons" @click="setIcon" />

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
import * as themify from '@quasar/extras/themify'
import * as bootstrapIcons from '@quasar/extras/bootstrap-icons'

const $q = useQuasar();
const { t } = useI18n();
const itemsPerPage = ref(200)
const filterKey = ref('')
const iconSet = ref(materialIcons)
const iconSetKeyList = ref(Object.keys(materialIcons))

const setIcon = () => {
    iconSetKeyList.value = Object.keys(iconSet.value)
}

const pagination = ref({
    page: 1,
    max: Math.ceil(iconSetKeyList.value.length / itemsPerPage.value)
})

const dataObj = computed(() => {
    let dataList = iconSetKeyList.value
    if (filterKey.value !== "" && filterKey.value) {
        dataList = dataList.filter(item => item.toLowerCase().includes(filterKey.value.toLowerCase()))
    }
    pagination.value = {
        page: pagination.value.page,
        max: Math.ceil(dataList.length / itemsPerPage.value)
    }
    const keyArray = dataList.slice((pagination.value.page - 1) * itemsPerPage.value, (pagination.value.page - 1) * itemsPerPage.value + itemsPerPage.value - 1)
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
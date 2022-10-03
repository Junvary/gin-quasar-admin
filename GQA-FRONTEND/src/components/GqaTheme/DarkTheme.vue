<template>
    <q-field :label="$t('Theme')" stack-label>
        <template v-slot:control>
            <q-btn-toggle v-model="darkTheme" toggle-color="primary" :options="[
                { slot: 'sun', value: false },
                { slot: 'moon', value: true },
            ]">
                <template v-slot:sun>
                    <div class="row items-center no-wrap">
                        <q-icon name="fas fa-sun" />
                    </div>
                </template>
                <template v-slot:moon>
                    <div class="row items-center no-wrap">
                        <q-icon name="fas fa-moon" />
                    </div>
                </template>
            </q-btn-toggle>
        </template>
    </q-field>
</template>

<script setup>
import { useQuasar } from 'quasar'
import { onMounted, ref, watch } from 'vue';
import { useSettingStore } from 'src/stores/setting';

const darkTheme = ref(false)
const $q = useQuasar()
const settingStore = useSettingStore()

onMounted(() => {
    darkTheme.value = settingStore.GetDarkTheme()
})
watch(darkTheme, (val) => {
    settingStore.SetDarkTheme(val)
    $q.dark.set(darkTheme.value)
})
</script>
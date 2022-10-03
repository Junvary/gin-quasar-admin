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
import { LocalStorage, useQuasar } from 'quasar'
import { onMounted, ref, watch } from 'vue';

const darkTheme = ref(false)
const $q = useQuasar()

onMounted(() => {
    const localDark = LocalStorage.getItem('gqa-dark-theme')
    darkTheme.value = localDark
})
watch(darkTheme, (val) => {
    LocalStorage.set('gqa-dark-theme', val)
    $q.dark.set(darkTheme.value)
})
</script>
<template>
    <div>
        <q-btn dense round flat :icon="`${!right ? 'settings' : 'close'}`" @click="right = !right" />
        <q-dialog v-model="right" position="right">
            <q-card style="height: 100%; width: 400px" class="text-center">
                <q-card-section>
                    <div class="text-h6">{{ $t('System') + $t('Setting') }}</div>
                </q-card-section>
                <q-card-section>
                    <q-tabs v-model="tab" dense class="text-grey" active-color="primary" indicator-color="primary"
                        align="justify" narrow-indicator>
                        <q-tab name="settings1" :label="$t('Basic') + $t('Setting')" />
                        <q-tab name="settings2" :label="$t('Theme') + $t('Color')" />
                    </q-tabs>

                    <q-separator />

                    <q-tab-panels v-model="tab" animated>
                        <q-tab-panel name="settings1" class="text-left">
                            <GqaLanguage />
                            <DarkTheme />
                            <SideDrawer />
                        </q-tab-panel>

                        <q-tab-panel name="settings2">
                            <GqaTheme />
                        </q-tab-panel>
                    </q-tab-panels>
                </q-card-section>
            </q-card>
        </q-dialog>
    </div>
</template>

<script setup>
import GqaLanguage from 'src/components/GqaLanguage'
import GqaTheme from 'src/components/GqaTheme'
import DarkTheme from 'src/components/GqaTheme/DarkTheme.vue';
import SideDrawer from './SideDrawer.vue'
import { ref } from 'vue';

const right = ref(false)
const tab = ref('settings1')
</script>

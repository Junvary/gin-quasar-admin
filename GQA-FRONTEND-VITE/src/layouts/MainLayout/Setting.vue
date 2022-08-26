<template>
    <div>
        <q-btn dense round flat :icon="`${!right ? 'settings' : 'close'}`" @click="right = !right" />
        <q-dialog v-model="right" position="right">
            <q-card style="height: 100%; width: 400px" class="text-center">
                <q-card-section>
                    <div class="text-h6">{{ $t('System') + $t('Setting') }}</div>
                </q-card-section>
                <q-card-section>
                    <DarkTheme />
                    <GqaLanguage />

                    <q-separator />
                    <GqaTheme />
                </q-card-section>

            </q-card>
        </q-dialog>
    </div>
</template>

<script setup>
import GqaLanguage from 'src/components/GqaLanguage/index.vue'
import GqaTheme from 'src/components/GqaTheme/index.vue'
import DarkTheme from 'src/components/GqaTheme/DarkTheme.vue'
import { ref } from 'vue'

const right = ref(false)
</script>

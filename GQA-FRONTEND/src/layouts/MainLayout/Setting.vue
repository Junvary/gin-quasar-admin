<template>
    <div>
        <q-btn dense round glossy push color="primary" :icon="`${!right ? 'settings' : 'close'}`"
            @click="right = !right" />
        <q-dialog v-model="right" position="right">
            <q-card style="height: 100%; width: 400px">
                <q-card-section>
                    <div class="text-h6">{{ $t('System') + $t('Setting') }}</div>
                </q-card-section>
                <q-card-section>
                    <GqaLanguage />
                    <q-separator />
                    <GqaTheme />
                </q-card-section>

            </q-card>
        </q-dialog>
    </div>
</template>

<script setup>
import GqaLanguage from 'src/components/GqaLanguage'
import GqaTheme from 'src/components/GqaTheme'
import { ref } from 'vue';
const right = ref(false)
</script>

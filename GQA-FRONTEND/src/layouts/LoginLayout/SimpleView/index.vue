<template>
    <q-card-section horizontal style="width: 50%;">
        <q-card-section class="row justify-center" style="width: 100%;">
            <SimpleViewHelp />
        </q-card-section>
    </q-card-section>
    <q-separator vertical dark />
    <q-card-section horizontal :class="darkThemeLoginHelp"
        style="width: 50%; height: 100%; background: #f3fbff; border-radius: 0 20px 20px 0;">
        <q-card-section class="row justify-center" style="width: 100%;">
            <SimpleViewForm />
        </q-card-section>
    </q-card-section>
</template>

<script setup>
import SimpleViewForm from './SimpleViewForm.vue'
import SimpleViewHelp from './SimpleViewHelp.vue'
import useTheme from 'src/composables/useTheme';

const { darkThemeLoginHelp } = useTheme()
</script>

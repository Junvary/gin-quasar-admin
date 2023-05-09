<template>
    <q-btn dense round flat :icon="$q.fullscreen.isActive ? 'las la-compress' : 'ion-ios-qr-scanner'"
        @click="$q.fullscreen.toggle()">
        <q-tooltip>
            {{ $t('FullScreen') }}
        </q-tooltip>
    </q-btn>
</template>

<script setup></script>

<template>
    <q-btn dense round flat :icon="$q.fullscreen.isActive ? 'ion-md-contract' : 'ion-md-expand'"
        @click="$q.fullscreen.toggle()" />
</template>

<script setup></script>

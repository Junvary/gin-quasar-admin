<template>
    <q-btn dense round flat :icon="$q.fullscreen.isActive ? 'fullscreen_exit' : 'fullscreen'"
        @click="$q.fullscreen.toggle()" />
</template>

<script setup></script>

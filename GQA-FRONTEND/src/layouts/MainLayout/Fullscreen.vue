<template>
    <q-btn dense round glossy push color="primary" :icon="$q.fullscreen.isActive ? 'fullscreen_exit' : 'fullscreen'"
        @click="$q.fullscreen.toggle()" />
</template>

<script>
export default {
    name: 'Fullscreen',
}
</script>
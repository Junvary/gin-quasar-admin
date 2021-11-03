<template>
    <div>
        <q-btn round glossy push color="primary" :icon="`${!right ? 'settings' : 'close'}`" @click="right = !right" />
        <q-drawer v-model="right" side="right" overlay bordered>
            abc
        </q-drawer>
    </div>
</template>

<script>
export default {
    name: 'Setting',
    data() {
        return {
            right: false,
        }
    },
}
</script>
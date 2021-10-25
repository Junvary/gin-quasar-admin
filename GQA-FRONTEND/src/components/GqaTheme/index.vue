<template>
    <div class="q-gutter-sm q-ma-none">
        <q-badge color="blue" text-color="black" class="skin-color q-ml-none" @click="setBrand('primary', '#1976D2')" />
        <q-badge color="brown" text-color="black" class="skin-color" @click="setBrand('primary', 'brown')" />
        <q-badge color="teal" text-color="black" class="skin-color" @click="setBrand('primary', 'teal')" />
        <q-badge color="dark" text-color="black" class="skin-color" @click="setBrand('primary', '#263238')" />
        <q-badge color="indigo" text-color="black" class="skin-color" @click="setBrand('primary', 'indigo')" />
        <q-badge color="orange" text-color="black" class="skin-color" @click="setBrand('primary', '#ff6a00')" />
    </div>
</template>

<script>
import { setCssVar } from 'quasar'

export default {
    name: 'GqaTheme',
    methods: {
        setBrand: setCssVar,
    },
}
</script>

<style lang="scss" scoped>
.skin-color {
    padding: 8px;
    border-radius: 50%;
}
</style>

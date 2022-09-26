<template>
    <q-dialog v-model="visible" position="bottom" persistent transition-show="rotate" transition-hide="rotate">
        <q-card style="width: 350px; padding: 5px;box-shadow: 0 0 60px 15px orange;" class="bg-orange">
            <div style="width: 100%; font-size: 0.8rem;" class="text-center text-bold">
                新成就达成
            </div>
            <q-separator dark="" />
            <q-item v-ripple class="bg-orange">
                <q-item-section avatar>
                    <q-chip class="glossy" color="orange-8" text-color="white" icon="star">
                        {{category}}
                    </q-chip>
                </q-item-section>
                <q-item-section style="font-size: 1rem;">
                    {{name}}
                </q-item-section>
            </q-item>
        </q-card>
    </q-dialog>
</template>

<script setup>
import { ref } from 'vue';

const visible = ref(false)

const category = ref("")
const name = ref("")
const show = (detail) => {
    category.value = detail?.category
    name.value = detail?.name
    visible.value = true
    setTimeout(() => {
        visible.value = false
    }, 4000)
}

defineExpose({
    show
})
</script>
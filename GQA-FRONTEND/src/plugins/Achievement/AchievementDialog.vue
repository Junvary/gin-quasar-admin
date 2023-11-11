<template>
    <q-dialog v-model="visible" position="bottom" persistent transition-show="rotate" transition-hide="rotate">
        <GqaCard startColor="#fcbc25" endColor="#f68057" style="width: 500px; box-shadow: 0 0 50px 10px #fcbc25;">
            <q-card-section>
                <span class="row justify-center text-h6 text-bold text-white">
                    新成就解锁
                </span>
                <q-separator spaced dark />
                <div class="row items-center">
                    <q-chip class="glossy col-2" color="orange-8" text-color="white" icon="emoji_events">
                        {{ category }}
                    </q-chip>
                    <div class="col text-dark row justify-center">
                        {{ name }}
                    </div>
                    <q-chip class="glossy col-2" color="orange-8" text-color="white" icon="star">
                        80
                    </q-chip>
                </div>
            </q-card-section>
        </GqaCard>
    </q-dialog>
</template>

<script setup>
import GqaCard from 'src/components/GqaCard/index.vue'
import { onBeforeUnmount, ref } from 'vue';

const visible = ref(false)

const category = ref("")
const name = ref("")
const timer = ref(null)
const show = (detail) => {
    category.value = detail?.category
    name.value = detail?.name
    visible.value = true
    timer.value = setTimeout(() => {
        visible.value = false
    }, 4000)
}

onBeforeUnmount(() => {
    clearTimeout(timer.value)
})

defineExpose({
    show
})
</script>
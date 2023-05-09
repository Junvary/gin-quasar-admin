<template>
    <q-dialog v-model="loginVisible">
        <q-card bordered style="width: 1000px; max-width: 50vw;">
            <q-card-section horizontal>
                <q-img class="col-6" :src="bannerImage" fit="cover" />
                <LoginForm class="col-6" />
            </q-card-section>
        </q-card>
    </q-dialog>
</template>

<script setup>
import LoginForm from './LoginForm.vue';
import useCommon from 'src/composables/useCommon'
import { ref, computed } from 'vue';

const { gqaFrontend } = useCommon()
const loginVisible = ref(false);

const bannerImage = computed(() => {
    if (gqaFrontend.value.bannerImage && gqaFrontend.value.bannerImage.substring(0, 11) === 'gqa-upload:') {
        return process.env.API + gqaFrontend.value.bannerImage.substring(11)
    }
    return "/img/sky.jpg"
})

const show = () => {
    loginVisible.value = true
}

defineExpose({
    show
})
</script>

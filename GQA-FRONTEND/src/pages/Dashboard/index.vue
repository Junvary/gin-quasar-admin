<template>
    <q-page>
        <q-img :src="bannerImage" class="absolute-full">
            <div class="absolute-full text-subtitle2 flex flex-center">
                <span style="font-size: 40px">
                    {{ $t('WelcomeTo') }}
                    {{ gqaFrontend.subTitle }}
                </span>
            </div>
        </q-img>
    </q-page>
</template>

<script setup>
import { useStorageStore } from 'src/stores/storage'
import { computed } from 'vue';

const storageStore = useStorageStore()
const gqaFrontend = computed(() => storageStore.GetGqaFrontend())
const randomImg = 'https://api.ixiaowai.cn/api/api.php'
const bannerImage = computed(() => {
    if (gqaFrontend.value.bannerImage && gqaFrontend.value.bannerImage.substring(0, 11) === 'gqa-upload:') {
        return process.env.API + gqaFrontend.value.bannerImage.substring(11)
    }
    return randomImg
})
</script>

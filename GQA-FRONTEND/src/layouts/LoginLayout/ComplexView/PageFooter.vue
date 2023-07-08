<template>
    <footer>
        <q-img :src="bannerImage" fit="cover" style="width: 100%; height: 5vh;">
            <div class="row items-center justify-center" style="width: 100%; height: 100%">
                {{ gqaFrontend.subTitle }}
                is powered by&nbsp;
                <a href="https://github.com/Junvary/gin-quasar-admin" target="_blank"
                    style="color: white; text-decoration: none;">
                    Gin-Quasar-Admin
                    {{ gqaVersion }}
                </a>
            </div>
        </q-img>
    </footer>
</template>

<script setup>
import { computed } from 'vue';
import { useStorageStore } from 'src/stores/storage'
import config from '../../../../package.json'

const gqaVersion = config.version
const storageStore = useStorageStore()
const gqaFrontend = computed(() => storageStore.GetGqaFrontend())
const bannerImage = computed(() => {
    if (gqaFrontend.value.bannerImage && gqaFrontend.value.bannerImage.substring(0, 11) === 'gqa-upload:') {
        return process.env.API + gqaFrontend.value.bannerImage.substring(11)
    }
    return "planet.png"
})

</script>

<style lang="scss" scoped>
footer {
    position: relative;
    z-index: 0;
    overflow: hidden;
    color: #ffffff;
}
</style>
<template>
    <footer>
        <div v-if="bannerImage !== ''">
            <q-img :src="bannerImage" fit="cover" style="width: 100%; height: 25vh;">
                <div class="column items-center" style="width: 100%; height: 100%">
                    <div class="row justify-center" style="width: 100%">
                        <div class="column items-center q-gutter-y-md">
                            <div class="col">
                                <gqa-avatar size="xl" :src="gqaFrontend.logo" />
                            </div>
                            <div class="col" style="font-size: 35px; font-weight: bold;letter-spacing: 5px;">
                                {{ gqaFrontend.subTitle }}
                            </div>
                            <span class="col" style="text-transform: capitalize;">
                                {{ gqaFrontend.webDescribe }}
                            </span>
                        </div>
                    </div>

                    <q-separator />

                    <div class="column items-center" style="width: 100%">
                        <span class="col">
                            {{ $t('TechnicalSupport') }}: {{ thisYear }}
                        </span>
                        <q-separator />
                        <span class="col">
                            Powered by
                            <a href="https://github.com/Junvary/gin-quasar-admin" target="_blank" style="color: white">
                                Gin-Quasar-Admin
                            </a>
                        </span>
                    </div>
                </div>
            </q-img>
        </div>
        <div v-else class="gin-quasar-admin-banner" style="padding-top: 30px; height: 25vh;">
            <div class="row justify-center">
                <div class="column items-center q-gutter-y-md">
                    <div class="col">
                        <gqa-avatar size="xl" :src="gqaFrontend.logo" />
                    </div>
                    <div class="col" style="font-size: 35px; font-weight: bold;letter-spacing: 5px;">
                        {{ gqaFrontend.subTitle }}
                    </div>
                    <span class="col" style="text-transform: capitalize;">
                        {{ gqaFrontend.webDescribe }}
                    </span>
                </div>
            </div>

            <q-separator />

            <div class="column items-center" style="width: 100%">
                <span class="col">
                    {{ $t('TechnicalSupport') }}: {{ thisYear }}
                </span>
                <q-separator />
                <span class="col">
                    Powered by
                    <a href="https://github.com/Junvary/gin-quasar-admin" target="_blank" style="color: white">
                        Gin-Quasar-Admin
                    </a>
                </span>
            </div>
        </div>
    </footer>
</template>

<script setup>
import { computed, ref } from 'vue';
import { useStorageStore } from 'src/stores/storage'

const storageStore = useStorageStore()
const thisYear = ref(new Date().getFullYear());
const gqaFrontend = computed(() => storageStore.GetGqaFrontend())
const bannerImage = computed(() => {
    if (gqaFrontend.value.bannerImage && gqaFrontend.value.bannerImage.substring(0, 11) === 'gqa-upload:') {
        return process.env.API + gqaFrontend.value.bannerImage.substring(11)
    }
    return ''
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
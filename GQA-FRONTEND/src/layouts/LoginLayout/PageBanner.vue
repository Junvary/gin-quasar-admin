<template>
    <div>
        <section id="gqa-banner" v-if="bannerImage !== ''">
            <q-img :src="bannerImage" fit="cover" style="width: 100%; max-height: 95vh">
                <div class="container-custom">
                    <div class="container-title">
                        <h1>
                            {{ gqaFrontend.subTitle }}
                        </h1>
                        <p class="small-title">
                            {{ gqaFrontend.webDescribe }}
                        </p>
                        <div class="buttons">
                            <q-btn push color="primary" @click="openLink('https://gitee.com/junvary/gin-quasar-admin')"
                                v-if="gqaFrontend.showGit === 'yes'">
                                Gitee
                            </q-btn>

                            <q-btn push color="primary" @click="showLoginForm" :disable="checkDbStatus">
                                {{ checkDbStatus ? $t('Need') + $t('Init') : $t('Login') }}
                            </q-btn>

                            <q-btn push color="primary" @click="openLink('https://github.com/Junvary/gin-quasar-admin')"
                                v-if="gqaFrontend.showGit === 'yes'">
                                Github
                            </q-btn>
                        </div>
                    </div>
                </div>
            </q-img>
        </section>

        <section class="gin-quasar-admin-banner" id="gqa-banner" v-else>
            <div class="container-default">
                <div class="container-title">
                    <h1>
                        {{ gqaFrontend.subTitle }}
                    </h1>
                    <p class="small-title">
                        {{ gqaFrontend.webDescribe }}
                    </p>
                    <div class="buttons">
                        <q-btn push color="primary" @click="openLink('https://gitee.com/junvary/gin-quasar-admin')"
                            v-if="gqaFrontend.showGit === 'yes'">
                            Gitee
                        </q-btn>

                        <q-btn push color="primary" @click="showLoginForm" :disable="checkDbStatus">
                            {{ checkDbStatus ? $t('Need') + $t('Init') : $t('Login') }}
                        </q-btn>

                        <q-btn push color="primary" @click="openLink('https://github.com/Junvary/gin-quasar-admin')"
                            v-if="gqaFrontend.showGit === 'yes'">
                            Github
                        </q-btn>
                    </div>
                </div>
                <div class="container-image">
                    <img src="~src/assets/code.png" alt="container-image">
                </div>
                <div class="container-team1"></div>
                <div class="container-team2"></div>
            </div>
        </section>

        <login-dialog ref="loginDialog" />

    </div>
</template>

<script setup>
import { ref, computed, toRefs } from 'vue';
import LoginDialog from './LoginDialog.vue'
import { useStorageStore } from 'src/stores/storage'

const loginDialog = ref(null);
const props = defineProps({
    checkDbStatus: {
        type: Boolean,
        required: true,
    },
})
const { checkDbStatus } = toRefs(props);
const gqaFrontend = computed(() => {
    const storageStore = useStorageStore()
    return storageStore.GetGqaFrontend()
})
const bannerImage = computed(() => {
    if (gqaFrontend.value.gqaBannerImage && gqaFrontend.value.gqaBannerImage.substring(0, 11) === 'gqa-upload:') {
        return process.env.API + gqaFrontend.value.gqaBannerImage.substring(11)
    }
    return ''
})

const showLoginForm = () => {
    loginDialog.value.show()
}
defineExpose({
    showLoginForm
})
const openLink = (url) => {
    window.open(url)
}
</script>

<style lang="scss" scoped>
.container-custom {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;

    .container-title {
        width: 60%;
        display: flex;
        justify-content: center;
        align-items: center;
        flex-wrap: wrap;
        flex-direction: column;
        z-index: 99;
        margin-bottom: 10vh;

        h1 {
            color: #ffffff;
            font-weight: 700;
            font-size: 50px;
            line-height: 70px;
            text-align: center;
            margin-bottom: 30px;
            letter-spacing: 15px;
            user-select: none;
        }

        .small-title {
            font-weight: 400;
            font-size: 20px;
            letter-spacing: 2px;
            line-height: 40px;
            text-align: center;
            color: #ffffff;
            opacity: 0.8;
            max-width: 750px;
            margin: auto;
            margin-bottom: 30px;
            user-select: none;
            text-transform: capitalize;
        }

        .buttons {
            width: 40%;
            display: flex;
            justify-content: space-around;
        }
    }
}

.container-default {
    width: 100%;
    padding-top: 20vh;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;

    .container-title {
        width: 60%;
        display: flex;
        justify-content: center;
        align-items: center;
        flex-wrap: wrap;
        flex-direction: column;
        z-index: 99;

        h1 {
            color: #ffffff;
            font-weight: 700;
            font-size: 50px;
            line-height: 70px;
            text-align: center;
            margin-bottom: 30px;
            letter-spacing: 20px;
            user-select: none;
        }

        .small-title {
            font-weight: 400;
            font-size: 20px;
            letter-spacing: 2px;
            line-height: 40px;
            text-align: center;
            color: #ffffff;
            opacity: 0.8;
            max-width: 750px;
            margin: auto;
            margin-bottom: 30px;
            user-select: none;
            text-transform: capitalize;
        }

        .buttons {
            width: 40%;
            display: flex;
            justify-content: space-around;
        }
    }

    .container-image {
        margin-top: 150px;
        max-width: 845px;
        // margin: auto;
        text-align: center;
        // position: relative;
        z-index: 1;
        display: flex;
        flex: 1;

        // align-items: flex-end;
        img {
            max-width: 100%;
            text-align: center;
            opacity: 0.5;
            border-radius: 20px;
        }
    }

    .container-team1 {
        position: absolute;
        background-image: url('src/assets/team1.svg');
        background-size: cover;
        width: 55vh;
        height: 35vh;
        left: 30px;
        top: 240px;
    }

    .container-team2 {
        position: absolute;
        background-image: url('src/assets/team2.svg');
        background-size: cover;
        width: 55vh;
        height: 35vh;
        right: 0;
        top: 500px;
    }
}
</style>
<template>
    <div>
        <section id="gqa-banner" class="bg-black">
            <div style="position: absolute; z-index: 10;">
                <div class="doc-stars doc-stars--sm"></div>
                <div class="doc-stars doc-stars--md"></div>
                <div class="doc-stars doc-stars--lg"></div>
            </div>

            <div class="main-title row justify-center items-center">
                <gqa-avatar size="xl" :src="gqaFrontend.logo" style="margin-right: 10px;" />
                {{ gqaFrontend.mainTitle }}
            </div>
            <q-img :src="bannerImage" fit="cover" style="width: 100%; height: 95vh">
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
                                v-if="gqaFrontend.showGit === 'yesNo_yes'">
                                Gitee
                            </q-btn>

                            <q-btn push color="primary" @click="showLoginForm">
                                {{ $t('Login') }}
                            </q-btn>

                            <q-btn push color="primary" @click="openLink('https://github.com/Junvary/gin-quasar-admin')"
                                v-if="gqaFrontend.showGit === 'yesNo_yes'">
                                Github
                            </q-btn>
                        </div>
                    </div>
                </div>
            </q-img>
        </section>
        <login-dialog ref="loginDialog" />
    </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import LoginDialog from '../LoginDialog.vue'
import { useStorageStore } from 'src/stores/storage'

const loginDialog = ref(null);
const storageStore = useStorageStore()
const gqaFrontend = computed(() => {
    return storageStore.GetGqaFrontend()
})
const bannerImage = computed(() => {
    if (gqaFrontend.value.bannerImage && gqaFrontend.value.bannerImage.substring(0, 11) === 'gqa-upload:') {
        return process.env.API + gqaFrontend.value.bannerImage.substring(11)
    }
    return "planet.png"
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

const starColor = computed(() => {
    if (gqaFrontend.value.starColor) {
        return gqaFrontend.value.starColor
    }
    return "#00b4ff"
})
</script>

<style lang="scss" scoped>
$starColor : v-bind(starColor);

.main-title {
    position: absolute;
    top: 30px;
    left: 30px;
    color: #ffffff;
    font-size: 30px;
    z-index: 10;
    user-select: none;
    font-weight: bold;
    letter-spacing: 2px;
}

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
}


.doc-stars {
    position: absolute
}

@function boxShadow($n) {
    $shadows: '#{random(10000)}px #{random(10000)}px #{$starColor}';

    @for $i from 2 through $n {
        $shadows: '#{$shadows}, #{random(10000)}px #{random(10000)}px #{$starColor}'
    }

    ;
    @return unquote($shadows);
}

.doc-stars--sm {
    animation: animateStar 70s linear infinite;
    background: transparent;
    box-shadow: boxShadow(600);
    height: 1px;
    width: 1px
}

.doc-stars--md {
    animation: animateStar 100s linear infinite;
    background: transparent;
    box-shadow: boxShadow(600);
    height: 2px;
    width: 2px
}

.doc-stars--lg {
    animation: animateStar 150s linear infinite;
    background: transparent;
    box-shadow: boxShadow(600);
    height: 3px;
    width: 3px
}

@keyframes animateStar {
    0% {
        transform: translateY(0)
    }

    to {
        transform: translateY(-3500px)
    }
}
</style>
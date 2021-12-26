<template>
    <section class="gin-quasar-admin-banner" id="gqa-banner">
        <div class="container">
            <div class="container-title">
                <h1>
                    {{ gqaFrontend.gqaSubTitle }}
                </h1>
                <p class="small-title">
                    {{ gqaFrontend.gqaDescribe }}
                </p>
                <div class="buttons">
                    <q-btn push color="primary" @click="openLink('https://gitee.com/junvary/gin-quasar-admin')"
                        v-if="gqaFrontend.gqaShowGit === 'yes'">
                        Gitee
                    </q-btn>

                    <q-btn push color="primary" @click="showLoginForm" :disable="checkDbStatus">
                        {{checkDbStatus ? $t('LoginLayoutPageBannerCheckDB') : $t('Login')}}
                    </q-btn>

                    <q-btn push color="primary" @click="openLink('https://github.com/Junvary/gin-quasar-admin')"
                        v-if="gqaFrontend.gqaShowGit === 'yes'">
                        Github
                    </q-btn>
                </div>
            </div>
            <div class="container-image">
                <img src="~src/assets/login/code.png" alt="container-image">
            </div>
            <div class="container-team1"></div>
            <div class="container-team2"></div>
        </div>
        <login-dialog ref="loginDialog" />
    </section>
</template>

<script>
import { gqaFrontendMixin } from 'src/mixins/gqaFrontendMixin'
import LoginDialog from 'src/pages/Login'
export default {
    name: 'PageBanner',
    mixins: [gqaFrontendMixin],
    props: {
        checkDbStatus: {
            type: Boolean,
            required: true,
        },
    },
    components: {
        LoginDialog,
    },
    methods: {
        showLoginForm() {
            this.$refs.loginDialog.show()
        },
        openLink(url) {
            window.open(url)
        },
    },
}
</script>

<style lang="scss" scoped>
.container {
    width: 100%;
    padding-top: 180px;
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
        align-items: flex-end;
        img {
            max-width: 100%;
            text-align: center;
            opacity: 0.5;
            border-radius: 20px;
        }
    }
    .container-team1 {
        position: absolute;
        background-image: url('~src/assets/login/team1.svg');
        background-size: cover;
        width: 55vh;
        height: 35vh;
        left: 30px;
        top: 240px;
    }
    .container-team2 {
        position: absolute;
        background-image: url('~src/assets/login/team2.svg');
        background-size: cover;
        width: 55vh;
        height: 35vh;
        right: 0;
        top: 500px;
    }
}
</style>
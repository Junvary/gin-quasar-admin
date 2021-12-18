<template>
    <q-dialog persistent v-model="loginVisible" transition-hide="slide-down">
        <q-card bordered style="width: 700px; max-width: 45vw;">
            <q-card-section horizontal>
                <q-img class="col-6" :src="randomImg" />
                <q-card-section>
                    <div class="text-center">
                        <GqaAvatar size="xl" :src="gqaFrontend.gqaWebLogo" />
                    </div>
                    <div class="text-h4 text-center text-primary ">
                        {{ gqaFrontend.gqaMainTitle }}
                    </div>
                    <div class="text-h6 text-center text-primary q-mt-md q-mb-xs">
                        {{ $t('LoginTitle') }}
                    </div>
                    <q-form @submit="onSubmit" class="q-mt-lg">
                        <q-input :disable="loading" outlined dense no-error-icon v-model.trim="form.username"
                            :placeholder="$t('Username')" :rules="[(val) =>(val && val.length > 0) || $t('NeedInput'),]"
                            autocomplete="username" />
                        <q-input :disable="loading" outlined dense no-error-icon :type="isPwd ? 'password' : 'text'"
                            v-model.trim="form.password" :placeholder="$t('Password')"
                            :rules="[(val) =>(val && val.length > 0) || $t('NeedInput'),]"
                            autocomplete="current-password">
                            <template v-slot:append>
                                <q-icon :name="isPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer"
                                    @click="isPwd = !isPwd" />
                            </template>
                        </q-input>
                        <q-input :disable="loading" outlined dense no-error-icon v-model.trim="form.captcha"
                            :placeholder="$t('Captcha')"
                            :rules="[(val) => (val && val.length > 0) || $t('NeedInput'),]">
                            <template v-slot:after>
                                <q-btn padding="none" style="width: 120px; height: 100%" @click="getCaptcha">
                                    <q-img :src="captchaImage" />
                                </q-btn>
                            </template>
                        </q-input>
                        <div class="column q-gutter-y-md q-mt-none">
                            <q-checkbox :disable="loading" v-model="rememberMe" :label="$t('LoginRememberMe')" dense
                                @update:model-value="changeRememberMe" />
                        </div>
                        <div class="items-center justify-around q-mt-md row">
                            <q-btn :label="$t('Login')" type="submit" color="primary" :loading="loading"
                                style="width: 48%" />
                            <q-btn :label="$t('Cancel')" @click="loginVisible = false" color="negative"
                                :loading="loading" style="width: 48%" />
                        </div>
                    </q-form>
                    <q-inner-loading :showing="loading">
                        <q-spinner-hourglass color="primary" size="3em" />
                    </q-inner-loading>
                </q-card-section>
            </q-card-section>
            <q-separator />

            <q-card-actions v-if="$q.screen.gt.sm">
                <GqaLanguage />
            </q-card-actions>

        </q-card>
    </q-dialog>
</template>

<script>
import { mapActions } from 'vuex'
import { getAction } from 'src/api/manage'
import { captchaUrl } from 'src/api/url'
import GqaLanguage from 'src/components/GqaLanguage'
import GqaAvatar from 'src/components/GqaAvatar'
import { gqaFrontendMixin } from 'src/mixins/gqaFrontendMixin'

export default {
    name: 'Login',
    mixins: [gqaFrontendMixin],
    components: {
        GqaLanguage,
        GqaAvatar,
    },
    data() {
        return {
            loginVisible: false,
            isPwd: true,
            randomImg: 'https://api.ixiaowai.cn/api/api.php',
            form: {
                username: '',
                password: '',
                captcha: '',
                captchaKey: '',
            },
            rememberMe: true,
            captchaImage: '',
            loading: false,
        }
    },

    methods: {
        show() {
            this.form = {
                username: 'admin',
                password: '123456',
                captcha: '',
                captchaKey: '',
            }
            this.getCaptcha()
            this.loginVisible = true
        },
        ...mapActions('user', ['HandleLogin', 'ChangeRememberMe']),
        getCaptcha() {
            getAction(captchaUrl).then((res) => {
                this.captchaImage = res.data.picPath
                this.form.captchaKey = res.data.captchaId
            })
        },
        async onSubmit() {
            this.loading = true
            const Username = this.form.username
            // const password = this.$md5(this.form.password)
            const Password = this.form.password
            const Captcha = this.form.captcha
            const CaptchaId = this.form.captchaKey
            const res = await this.HandleLogin({
                Username,
                Password,
                Captcha,
                CaptchaId,
            })
            this.getCaptcha()
            if (res) {
                this.loading = false
                this.$router.push(this.$route.query.redirect || '/')
            } else {
                this.form.captcha = ''
                this.loading = false
            }
        },
        changeRememberMe(value) {
            this.ChangeRememberMe(value)
        },
    },
}
</script>

<style lang="scss" scoped>
</style>

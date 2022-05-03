<template>
    <q-dialog persistent v-model="loginVisible" transition-hide="slide-down">
        <q-card bordered style="width: 700px; max-width: 45vw;">
            <q-card-section horizontal>
                <q-img class="col-6" :src="bannerImage" fit="cover" />
                <q-card-section>
                    <div class="text-center">
                        <GqaAvatar size="xl" :src="gqaFrontend.logo" />
                    </div>
                    <div class="text-h5 text-center text-primary ">
                        {{ gqaFrontend.subTitle }}
                    </div>
                    <div class="text-h7 text-center text-primary q-mt-md q-mb-xs">
                        {{ $t('WelcomeBack') }}
                    </div>
                    <q-form @submit="onSubmit" class="q-mt-lg gqa-form">
                        <q-input :disable="loading" outlined dense no-error-icon v-model.trim="form.username"
                            :placeholder="$t('Username')" :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]"
                            autocomplete="username" />
                        <q-input :disable="loading" outlined dense no-error-icon :type="isPwd ? 'password' : 'text'"
                            v-model.trim="form.password" :placeholder="$t('Password')"
                            :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]"
                            autocomplete="current-password">
                            <template v-slot:append>
                                <q-icon :name="isPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer"
                                    @click="isPwd = !isPwd" />
                            </template>
                        </q-input>
                        <q-input :disable="loading" outlined dense no-error-icon v-model.trim="form.captcha"
                            :placeholder="$t('Captcha')" :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                            <template v-slot:after>
                                <q-btn padding="none" style="width: 120px; height: 100%" @click="getCaptcha">
                                    <q-img :src="captchaImage" />
                                </q-btn>
                            </template>
                        </q-input>
                        <div class="column q-gutter-y-md q-mt-none">
                            <q-checkbox :disable="loading" v-model="rememberMe" :label="$t('RememberMe')" dense
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

<script setup>
import { ref, computed } from 'vue';
import GqaLanguage from 'src/components/GqaLanguage/index.vue'
import GqaAvatar from 'src/components/GqaAvatar/index.vue'
import { useUserStore } from 'src/stores/user'
import { useStorageStore } from 'src/stores/storage';
import { postAction } from 'src/api/manage'
import { useRouter, useRoute } from 'vue-router';

const loginVisible = ref(false);
const isPwd = ref(true)
const randomImg = 'https://api.ixiaowai.cn/api/api.php'
const form = ref({
    username: '',
    password: '',
    captcha: '',
    captcha_id: '',
})
const rememberMe = ref(true)
const captchaImage = ref('')
const loading = ref(false)
const storageStore = useStorageStore()
const userStore = useUserStore()
const router = useRouter()
const route = useRoute()

const gqaFrontend = computed(() => {
    return storageStore.GetGqaFrontend()
})
const bannerImage = computed(() => {
    if (gqaFrontend.bannerImage && gqaFrontend.bannerImage.substring(0, 11) === 'gqa-upload:') {
        return process.env.API + gqaFrontend.bannerImage.substring(11)
    }
    return randomImg
})

const show = () => {
    form.value = {
        username: '',
        password: '',
        captcha: '',
        captcha_id: '',
    }
    getCaptcha()
    loginVisible.value = true
}

defineExpose({
    show
})

const getCaptcha = () => {
    postAction('public/get-captcha').then((res) => {
        captchaImage.value = res.data.captcha_image
        form.value.captcha_id = res.data.captcha_id
    })
}
const changeRememberMe = (value) => {
    userStore.ChangeRememberMe(value)
}

const onSubmit = async () => {
    loading.value = true
    const res = await userStore.HandleLogin({
        username: form.value.username,
        password: form.value.password,
        captcha: form.value.captcha,
        captcha_id: form.value.captcha_id
    })
    getCaptcha()
    if (res) {
        loading.value = false
        router.push(route.query.redirect || '/')
        form.value.captcha = ''
        loading.value = false
    } else {
        form.value.captcha = ''
        loading.value = false
    }
}
</script>


<style lang="scss" scoped>
</style>

<template>
    <q-card-section>
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
                    :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]" autocomplete="current-password">
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
                    <q-btn :label="$t('Login')" type="submit" color="primary" :loading="loading" style="width: 100%" />
                </div>
            </q-form>
            <q-inner-loading :showing="loading">
                <q-spinner-hourglass color="primary" size="3em" />
            </q-inner-loading>
        </q-card-section>
        <q-separator />

        <q-card-actions v-if="$q.screen.gt.sm">
            <GqaLanguage />
        </q-card-actions>
    </q-card-section>
</template>

<script setup>
import useCommon from 'src/composables/useCommon'
import { ref, onMounted } from 'vue';
import { postAction } from 'src/api/manage'
import GqaLanguage from 'src/components/GqaLanguage/index.vue'
import GqaAvatar from 'src/components/GqaAvatar/index.vue'
import { useUserStore } from 'src/stores/user'
import { useRouter, useRoute } from 'vue-router';

const { gqaFrontend } = useCommon()

const isPwd = ref(true)
const form = ref({
    username: '',
    password: '',
    captcha: '',
    captcha_id: '',
})
const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const rememberMe = ref(true)
const captchaImage = ref('')
const loading = ref(false)
const getCaptcha = () => {
    postAction('public/get-captcha').then((res) => {
        captchaImage.value = res.data.captcha_image
        form.value.captcha_id = res.data.captcha_id
    })
}
const changeRememberMe = (value) => {
    userStore.ChangeRememberMe(value)
}

onMounted(() => {
    getCaptcha()
})

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
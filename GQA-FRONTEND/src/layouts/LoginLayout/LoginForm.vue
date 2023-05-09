<template>
    <q-card-section>
        <q-card-section>
            <div class="text-center">
                <gqa-avatar size="xl" :src="gqaFrontend.logo" />
            </div>
            <div class="text-h4 text-center text-bold">
                {{ gqaFrontend.subTitle }}
            </div>
            <div class="text-h6 text-center q-mt-md q-mb-xs">
                {{ $t('WelcomeBack') }}
            </div>
            <q-form @submit="onSubmit" class="q-mt-lg gqa-form">
                <q-input :disable="loading" outlined dense no-error-icon v-model.trim="form.username" rounded color="black"
                    :placeholder="$t('Username')" :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]" />
                <q-input :disable="loading" outlined dense no-error-icon :type="isPwd ? 'password' : 'text'"
                    v-model.trim="form.password" :placeholder="$t('Password')" rounded color="black"
                    :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]">
                    <template v-slot:append>
                        <q-icon :name="isPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer"
                            @click="isPwd = !isPwd" />
                    </template>
                </q-input>
                <q-input :disable="loading" outlined dense no-error-icon v-model.trim="form.captcha" rounded color="black"
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
                    <q-btn rounded glossy :label="$t('Login')" type="submit" color="primary" :loading="loading"
                        style="width: 100%" />
                </div>
            </q-form>
            <q-inner-loading :showing="loading" style="background-color: rgba(0, 0, 0, 0);">
                <q-spinner-hourglass color="primary" size="3em" />
            </q-inner-loading>
        </q-card-section>
    </q-card-section>
</template>

<script setup>
import useCommon from 'src/composables/useCommon'
import { ref, onMounted } from 'vue';
import { postAction } from 'src/api/manage'
import { useUserStore } from 'src/stores/user'
import { usePermissionStore } from 'src/stores/permission';
import { useRouter, useRoute } from 'vue-router';
import { useQuasar } from 'quasar';

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
const $q = useQuasar()
const userStore = useUserStore()
const permissionStore = usePermissionStore()
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
    if (res) {
        // The Get User menu is added here to trigger the get user default page. 
        // After successful access here, the boot/permission.js file determines that permissionStore.userMenu.length exists.
        // Go to next()
        const userMenu = await permissionStore.GetUserMenu()
        if (userMenu && userMenu.length) {
            userMenu.forEach(item => {
                router.addRoute(item)
            })
            router.push({ name: permissionStore.defaultPage[0] })
        } else {
            form.value.captcha = ''
            getCaptcha()
            loading.value = false
            $q.notify({
                type: 'negative',
                message: userMenu,
            })
        }
    } else {
        form.value.captcha = ''
        loading.value = false
        getCaptcha()
    }
}
</script>
<template>
    <q-dialog v-model="passwordDialog" persistent>
        <q-card style="min-width: 350px">
            <q-card-section>
                <div class="text-h6">
                    {{ $t('Change') + ' ' + $t('Password') }}
                </div>
            </q-card-section>

            <q-card-section class="q-pt-none">
                <q-form class="q-gutter-md" ref="passwordFormRef">
                    <q-input filled v-model.trim="passwordForm.old_password" autocomplete="off"
                        :label="$t('Old') + ' ' + $t('Password')" :type="isPwd ? 'password' : 'text'"
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]">
                        <template v-slot:append>
                            <q-icon :name="isPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer"
                                @click="isPwd = !isPwd" />
                        </template>
                    </q-input>

                    <q-input filled v-model.trim="passwordForm.new_password_1" autocomplete="off"
                        :label="$t('New') + ' ' + $t('Password')" :type="isPwd ? 'password' : 'text'"
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]">
                        <template v-slot:append>
                            <q-icon :name="isPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer"
                                @click="isPwd = !isPwd" />
                        </template>
                    </q-input>

                    <q-input filled v-model.trim="passwordForm.new_password_2" autocomplete="off"
                        :label="$t('New') + ' ' + $t('Password')" :type="isPwd ? 'password' : 'text'"
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]">
                        <template v-slot:append>
                            <q-icon :name="isPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer"
                                @click="isPwd = !isPwd" />
                        </template>
                    </q-input>

                </q-form>
            </q-card-section>

            <q-card-actions align="right" class="text-primary">
                <q-btn flat :label="$t('Cancel')" v-close-popup />
                <q-btn flat :label="$t('Save')" @click="handleChangePasswrod" />
            </q-card-actions>
        </q-card>
    </q-dialog>
</template>

<script setup>
import { postAction } from 'src/api/manage'
import { useUserStore } from 'src/stores/user'
import { ref } from 'vue';
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';

const $q = useQuasar()
const { t } = useI18n()
const userStore = useUserStore()
const router = useRouter()
const passwordDialog = ref(false)
const passwordForm = ref({
    old_password: '',
    new_password_1: '',
    new_password_2: '',
})
const isPwd = ref(true)
const changePasswordUrl = 'user/change-password'

const show = () => {
    passwordDialog.value = true
}
defineExpose({
    show
})
const passwordFormRef = ref(null)
const handleChangePasswrod = async () => {
    const success = await passwordFormRef.value.validate()
    if (success) {
        if (passwordForm.value.new_password_1 !== passwordForm.value.new_password_2) {
            $q.notify({
                type: 'negative',
                message: t('TwoPasswordsError'),
            })
        } else {
            const res = await postAction(changePasswordUrl, passwordForm.value)
            if (res.code === 1) {
                $q.notify({
                    type: 'positive',
                    message: t('Change') + ' ' + t('Password') + ' ' + t('Success') + ',' + t('Relogin'),
                })
                passwordDialog.value = false
                userStore.HandleLogout()
                router.push({ name: 'login' })
            }
        }
    } else {
        $q.notify({
            type: 'negative',
            message: t('FixForm'),
        })
    }
}
</script>

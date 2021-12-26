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
                    <q-input filled v-model.trim="passwordForm.oldPassword" autocomplete="off"
                        :label="$t('Old') + ' ' + $t('Password')" :type="isPwd ? 'password' : 'text'"
                        :rules="[ val => val && val.length > 0 || $t('NeedInput')]">
                        <template v-slot:append>
                            <q-icon :name="isPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer"
                                @click="isPwd = !isPwd" />
                        </template>
                    </q-input>

                    <q-input filled v-model.trim="passwordForm.newPassword1" autocomplete="off"
                        :label="$t('New') + ' ' + $t('Password')" :type="isPwd ? 'password' : 'text'"
                        :rules="[ val => val && val.length > 0 || $t('NeedInput')]">
                        <template v-slot:append>
                            <q-icon :name="isPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer"
                                @click="isPwd = !isPwd" />
                        </template>
                    </q-input>

                    <q-input filled v-model.trim="passwordForm.newPassword2" autocomplete="off"
                        :label="$t('New') + ' ' + $t('Password')" :type="isPwd ? 'password' : 'text'"
                        :rules="[ val => val && val.length > 0 || $t('NeedInput')]">
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

<script>
import { postAction } from 'src/api/manage'
import { mapActions } from 'vuex'

export default {
    name: 'ChangePasswordDialog',
    data() {
        return {
            passwordDialog: false,
            passwordForm: {
                oldPassword: '',
                newPassword1: '',
                newPassword2: '',
            },
            changePasswordUrl: 'user/user-change-password',
            isPwd: true,
        }
    },
    methods: {
        ...mapActions('user', ['HandleLogout']),
        show() {
            this.passwordDialog = true
        },
        async handleChangePasswrod() {
            const success = await this.$refs.passwordFormRef.validate()
            if (success) {
                if (this.passwordForm.newPassword1 !== this.passwordForm.newPassword2) {
                    this.$q.notify({
                        type: 'negative',
                        message: this.$t('TwoPasswordsCheck'),
                    })
                } else {
                    const res = await postAction(this.changePasswordUrl, this.passwordForm)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: this.$t('Change') + ' ' + this.$t('Password') + ' ' + this.$t('Success') + ',' + this.$t('Relogin'),
                        })
                        this.passwordDialog = false
                        this.HandleLogout()
                        this.$router.push({ name: 'login' })
                    }
                }
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: this.$t('FixForm'),
                })
            }
        },
    },
}
</script>
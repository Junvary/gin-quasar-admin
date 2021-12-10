<template>
    <q-dialog v-model="showProfile" full-height full-width>
        <q-card class="column full-height full-width">
            <q-card-section class="row items-center q-pb-none">
                <div class="text-h6">
                    {{ $t('UserProfile') }}
                </div>
                <q-space />
                <q-btn icon="close" flat round dense v-close-popup />
            </q-card-section>
            <q-card-section>
                <div class="row q-pa-md items-center justify-around">
                    <div class="q-pa-md col-4 column items-center">
                        <GqaAvatar loginUser size="200px" />

                        <div class="text-subtitle1 q-mt-md q-mb-md">
                            <GqaShowName showMyName />
                        </div>

                        <q-list>
                            <q-item clickable v-ripple>
                                <q-item-section avatar>
                                    <q-icon size="lg" name="star" class="text-warning" />
                                </q-item-section>
                                <q-item-section>等级:23</q-item-section>
                            </q-item>
                            <q-item clickable v-ripple>
                                <q-item-section avatar>
                                    <q-icon size="lg" name="star" class="text-warning" />
                                </q-item-section>
                                <q-item-section>积分:88888</q-item-section>
                            </q-item>
                            <q-item clickable v-ripple>
                                <q-item-section avatar>
                                    <q-icon size="lg" name="star" class="text-warning" />
                                </q-item-section>
                                <q-item-section>成就:888</q-item-section>
                            </q-item>
                        </q-list>

                        <q-separator inset spaced />

                        <div class="row q-gutter-md">
                            <q-btn color="primary" :label="$t('Change') + ' ' + $t('Password')"
                                @click="showPasswordDialog"></q-btn>
                        </div>
                    </div>
                    <div class="q-pa-md col-8 column items-center">
                        <profileDetail />
                    </div>
                </div>
            </q-card-section>
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
                        <q-btn flat label="取消" v-close-popup />
                        <q-btn flat label="确定更改" @click="handleChangePasswrod" />
                    </q-card-actions>
                </q-card>
            </q-dialog>
        </q-card>
    </q-dialog>
</template>

<script>
import GqaShowName from 'src/components/GqaShowName'
import GqaAvatar from 'src/components/GqaAvatar'
import profileDetail from './modules/profileDetail'
import { postAction } from 'src/api/manage'
import { mapActions } from 'vuex'

export default {
    name: 'UserProfile',
    components: {
        GqaShowName,
        GqaAvatar,
        profileDetail,
    },
    data() {
        return {
            showProfile: false,
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
        show() {
            this.showProfile = true
        },
        showPasswordDialog() {
            this.passwordDialog = true
        },
        ...mapActions('user', ['HandleLogout']),
        async handleChangePasswrod() {
            const success = await this.$refs.passwordFormRef.validate()
            if (success) {
                if (this.passwordForm.newPassword1 !== this.passwordForm.newPassword2) {
                    this.$q.notify({
                        type: 'negative',
                        message: '两次新密码不一致！',
                    })
                } else {
                    const res = await postAction(this.changePasswordUrl, this.passwordForm)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: '修改密码成功，请重新登录！',
                        })
                        this.passwordDialog = false
                        this.HandleLogout()
                        this.$router.push({ name: 'login' })
                    }
                }
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: '请完善表格信息！',
                })
            }
        },
    },
}
</script>
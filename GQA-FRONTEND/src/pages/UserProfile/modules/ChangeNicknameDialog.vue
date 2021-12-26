<template>
    <q-dialog v-model="nicknameDialog" persistent>
        <q-card style="min-width: 450px">
            <q-card-section>
                <div class="text-h6">
                    {{ $t('Change') + ' ' + $t('Nickname') }}
                </div>
            </q-card-section>

            <q-card-section class="q-pt-none">
                <q-form class="q-gutter-md" ref="nicknameFormRef">
                    <q-input filled v-model.trim="nicknameForm.nickname" autocomplete="off"
                        :label="$t('New') + ' ' + $t('Nickname')"
                        :rules="[ val => val && val.length > 0 || $t('NeedInput')]">
                        <template v-slot:hint>
                            {{ $t('NicknameMessage') }}
                        </template>
                    </q-input>
                </q-form>
            </q-card-section>

            <q-card-actions align="right" class="text-primary">
                <q-btn flat :label="$t('Cancel')" v-close-popup />
                <q-btn flat :label="$t('Save')" @click="handleChangeNickname" />
            </q-card-actions>
        </q-card>
    </q-dialog>
</template>

<script>
import { postAction } from 'src/api/manage'
import { mapActions } from 'vuex'

export default {
    name: 'ChangeNicknameDialog',
    data() {
        return {
            nicknameDialog: false,
            nicknameForm: {
                nickname: this.$q.cookies.get('gqa-nickname'),
            },
            changeNicknameUrl: 'user/user-change-nickname',
        }
    },
    methods: {
        ...mapActions('user', ['HandleLogout']),
        show() {
            this.nicknameDialog = true
        },
        async handleChangeNickname() {
            const success = await this.$refs.nicknameFormRef.validate()
            if (success) {
                const res = await postAction(this.changeNicknameUrl, this.nicknameForm)
                if (res.code === 1) {
                    this.$q.notify({
                        type: 'positive',
                        message: this.$t('Change') + ' ' + this.$t('Nickname') + ' ' + this.$t('Success') + ',' + this.$t('Relogin'),
                    })
                    this.nicknameDialog = false
                    this.HandleLogout()
                    this.$router.push({ name: 'login' })
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
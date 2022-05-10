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
                        :rules="[val => val && val.length > 0 || $t('NeedInput')]">
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

<script setup>
import { useUserStore } from 'src/stores/user'
import { postAction } from 'src/api/manage'
import { ref } from 'vue';
import { useQuasar } from 'quasar';
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';

const $q = useQuasar()
const { t } = useI18n()
const router = useRouter()
const userStore = useUserStore()
const nicknameDialog = ref(false)
const nicknameForm = ref({
    nickname: $q.cookies.get('gqa-nickname')
})
const changeNicknameUrl = 'user/change-nickname'
const show = () => {
    nicknameDialog.value = true
}
defineExpose({
    show
})
const nicknameFormRef = ref(null)
const handleChangeNickname = async () => {
    const success = await nicknameFormRef.value.validate()
    if (success) {
        const res = await postAction(changeNicknameUrl, nicknameForm.value)
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: t('Change') + ' ' + t('Nickname') + ' ' + t('Success') + ',' + t('Relogin'),
            })
            nicknameDialog.value = false
            userStore.HandleLogout()
            router.push({ name: 'login' })
        }
    } else {
        $q.notify({
            type: 'negative',
            message: t('FixForm'),
        })
    }
}
</script>

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
            <q-card-section style="max-width: 95vw; height: 85vh;">
                <div class="row items-center" style="height: 100%">
                    <div class="col-4 column items-center justify-center">
                        <gqa-avatar loginUser size="200px" />

                        <div class="text-subtitle1 q-mt-md q-mb-md">
                            <GqaShowName showMyName />
                        </div>

                        <div class="row q-gutter-md">
                            <q-btn color="primary" :label="$t('Change') + ' ' + $t('Nickname')"
                                @click="showNicknameDialog"></q-btn>
                            <q-btn color="primary" :label="$t('Change') + ' ' + $t('Password')"
                                @click="showPasswordDialog"></q-btn>
                        </div>
                    </div>
                    <div class="col-8 column items-center justify-start" style="height: 100%;">
                        <NoticeTab ref="noticeTab" />
                    </div>
                </div>
            </q-card-section>
            <ChangePasswordDialog ref="changePasswordDialog" />
            <ChangeNicknameDialog ref="changeNicknameDialog" />
        </q-card>
    </q-dialog>
</template>

<script setup>
import GqaShowName from 'src/components/GqaShowName/index.vue'
import NoticeTab from './modules/NoticeTab.vue'
import ChangePasswordDialog from './modules/ChangePasswordDialog.vue'
import ChangeNicknameDialog from './modules/ChangeNicknameDialog.vue'
import { ref, nextTick } from 'vue'

const showProfile = ref(false)
const noticeTab = ref(null)
const show = (type) => {
    showProfile.value = true
    if (type) {
        nextTick(() => {
            noticeTab.value.changeMessageType(type)
        })
    }
}
defineExpose({
    show
})
const changePasswordDialog = ref(null)
const showPasswordDialog = () => {
    changePasswordDialog.value.show()
}
const changeNicknameDialog = ref(null)
const showNicknameDialog = () => {
    changeNicknameDialog.value.show()
}
</script>

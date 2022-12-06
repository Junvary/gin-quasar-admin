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
                <div class="row q-pa-md items-center justify-around" style="max-width: 95vw">
                    <div class="col-4 column items-center">
                        <GqaAvatar loginUser size="200px" />

                        <div class="text-subtitle1 q-mt-md q-mb-md">
                            <GqaShowName showMyName />
                        </div>

                        <!-- <q-list>
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
                        </q-list> -->

                        <q-separator inset spaced />

                        <div class="row q-gutter-md">
                            <q-btn color="primary" :label="$t('Change') + ' ' + $t('Nickname')"
                                @click="showNicknameDialog"></q-btn>
                            <q-btn color="primary" :label="$t('Change') + ' ' + $t('Password')"
                                @click="showPasswordDialog"></q-btn>
                        </div>
                    </div>
                    <div class="q-pa-md col-8 column items-center">
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
import GqaShowName from 'src/components/GqaShowName'
import GqaAvatar from 'src/components/GqaAvatar'
import NoticeTab from './modules/NoticeTab'
import ChangePasswordDialog from './modules/ChangePasswordDialog'
import ChangeNicknameDialog from './modules/ChangeNicknameDialog'
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

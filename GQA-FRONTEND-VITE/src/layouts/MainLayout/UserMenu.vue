<template>
    <q-btn dense round flat>
        <GqaAvatar loginUser size="26px" />
        <q-menu class="row items-center justify-around q-pa-md">
            <div class="column">
                <div class="text-h6">摘要</div>
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
            </div>

            <q-separator vertical inset class="q-mx-lg" />

            <div class="column items-center">
                <GqaAvatar loginUser size="88px" />

                <div class="text-subtitle1 q-mt-md q-mb-md">
                    <GqaShowName showMyName />
                </div>

                <div class="row q-gutter-md">
                    <q-btn icon="person_pin" color="primary" :label="$t('UserProfile')" push size="sm" v-close-popup
                        @click="showProfile" />
                    <q-btn icon="logout" color="primary" :label="$t('Logout')" push size="sm" v-close-popup
                        @click="logout" />
                </div>
            </div>
        </q-menu>
    </q-btn>
</template>

<script setup>
import { useUserStore } from 'src/stores/user'
import GqaShowName from 'src/components/GqaShowName/index.vue'
import GqaAvatar from 'src/components/GqaAvatar/index.vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import { useRouter } from 'vue-router'

const userStore = useUserStore()
const $q = useQuasar()
const { t } = useI18n()
const router = useRouter()

const logout = () => {
    $q.dialog({
        title: t('Logout'),
        message: t('Confirm') + t('Logout') + '?',
        cancel: true,
        persistent: true,
    }).onOk(() => {
        userStore.HandleLogout()
        router.push({ path: '/login' })

    })
}

const emit = defineEmits(['showProfile'])
const showProfile = () => {
    emit('showProfile')
}
</script>

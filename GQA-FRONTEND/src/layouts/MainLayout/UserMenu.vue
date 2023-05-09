<template>
    <q-btn dense flat rounded>
        <gqa-avatar rounded loginUser size="28px" />
        <q-menu class="row no-wrap items-center justify-around q-pa-md">
            <div class="column items-center" style="width: 240px;">
                <gqa-avatar loginUser size="88px" />
                <div class="text-subtitle1 q-mt-md q-mb-md">
                    <GqaShowName showMyName />
                </div>
                <div class="row no-wrap q-gutter-md">
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

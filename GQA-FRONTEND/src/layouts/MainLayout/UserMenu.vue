<template>
    <q-btn-dropdown dense glossy push color="primary">
        <template v-slot:label>
            <GqaAvatar loginUser size="26px" />
            <span style="margin-left:5px">
                {{ $t('LayoutMainUserMenuWelcome') }}
                <GqaShowName showMyName />
            </span>
        </template>
        <div class="row no-wrap q-pa-md">
            <div class="column">
                <div class="text-h6 q-mb-md">{{ $t('LayoutMainUserMenuSettings')}}</div>
                <q-toggle v-model="mobileData" :label="$t('LayoutMainUserMenuSettingsUseMobileData')" />
                <q-toggle v-model="bluetooth" :label="$t('LayoutMainUserMenuSettingsBluetooth')" />
            </div>

            <q-separator vertical inset class="q-mx-lg" />

            <div class="items-center column">
                <GqaAvatar loginUser size="72px" />

                <div class="text-subtitle1 q-mt-md q-mb-xs">
                    <GqaShowName />
                </div>

                <q-btn color="primary" :label="$t('LayoutMainUserMenuLogout')" push size="sm" v-close-popup @click="logout" />
            </div>
        </div>
    </q-btn-dropdown>
</template>

<script>
import { mapActions } from 'vuex'
import GqaShowName from 'src/components/GqaShowName'
import GqaAvatar from 'src/components/GqaAvatar'

export default {
    name: 'UserMenu',
    components: {
        GqaShowName,
        GqaAvatar,
    },
    data() {
        return {
            mobileData: false,
            bluetooth: true,
        }
    },
    methods: {
        ...mapActions('user', ['HandleLogout']),
        logout() {
            this.$q
                .dialog({
                    title: $t('LayoutMainUserMenuLogoutTitle'),
                    message: $t('LayoutMainUserMenuLogoutMessage'),
                    cancel: true,
                    persistent: true,
                })
                .onOk(() => {
                    this.HandleLogout().then(() => {
                        this.$router.push({ name: 'login' })
                    })
                })
        },
    },
}
</script>

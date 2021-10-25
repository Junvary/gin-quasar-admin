<template>
    <q-btn-dropdown flat no-caps stretch auto-close>
        <template v-slot:label>
            <gqa-avatar loginUser size="26px" />
            <span style="margin-left:5px">
                欢迎你，
                <gqa-show-name />
            </span>
        </template>
        <div class="row no-wrap q-pa-md">
            <div class="column" style="min-width: 200px">
                <div class="text-h6 q-mb-md">Settings</div>
                <q-toggle dense v-model="mobileData" label="Use Mobile Data" />
                <q-toggle dense v-model="bluetooth" label="Bluetooth" class="q-mt-md" />
                <div class="q-mt-md">
                    <gqa-theme />
                </div>
            </div>

            <q-separator vertical inset class="q-mx-lg" />

            <div class="column items-center">
                <gqa-avatar loginUser size="72px" />

                <div class="text-subtitle1 q-mt-md q-mb-xs">
                    <gqa-show-name />
                </div>

                <q-btn color="primary" label="退出登录" push size="sm" v-close-popup @click="logout" />
            </div>
        </div>
    </q-btn-dropdown>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import GqaTheme from 'src/components/GqaTheme'
import GqaShowName from 'src/components/GqaShowName'
import GqaAvatar from 'src/components/GqaAvatar'

export default {
    name: 'UserMenu',
    components: {
        GqaTheme,
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
                    title: '确定退出？',
                    message: '你真的要退出系统吗？',
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
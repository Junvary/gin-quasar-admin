<template>
    <q-btn-dropdown dense glossy push color="primary">
        <template v-slot:label>
            <GqaAvatar loginUser size="26px" />
            <span style="margin-left:5px">
                {{ $t('Welcome') }}
                &nbsp;
                <GqaShowName showMyName />
            </span>
        </template>
        <div class="row items-center justify-around q-pa-md">
            <div class="column">
                <div class="text-h6">摘要</div>
                <q-list>
                    <q-item clickable v-ripple>
                        <q-item-section avatar>
                            <q-icon size="lg" name="star" class="text-warning" />
                        </q-item-section>
                        <q-item-section>等级：23</q-item-section>
                    </q-item>
                    <q-item clickable v-ripple>
                        <q-item-section avatar>
                            <q-icon size="lg" name="star" class="text-warning" />
                        </q-item-section>
                        <q-item-section>积分：88888</q-item-section>
                    </q-item>
                    <q-item clickable v-ripple>
                        <q-item-section avatar>
                            <q-icon size="lg" name="star" class="text-warning" />
                        </q-item-section>
                        <q-item-section>成就：888</q-item-section>
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
                    title: this.$t('LogoutTitle'),
                    message: this.$t('LogoutMessage'),
                    cancel: true,
                    persistent: true,
                })
                .onOk(() => {
                    this.HandleLogout().then(() => {
                        this.$router.push({ name: 'login' })
                    })
                })
        },
        showProfile() {
            this.$emit('showProfile')
        },
    },
}
</script>
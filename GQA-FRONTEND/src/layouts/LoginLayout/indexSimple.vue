<template>
    <div class="row gin-quasar-admin-init-db-page">
        <div class="col column justify-center items-center text-center" style="margin-bottom: 100px;">
            <GqaAvatar class="gin-quasar-admin-logo" :src="gqaFrontend.logo" size="100px" />
            <br />
            <span class="text-weight-bold text-h3 text-white">
                {{ gqaFrontend.subTitle }}
            </span>
            <span class="text-white text-subtitle1" style="margin-top: 20px">
                {{ gqaFrontend.webDescribe }}
            </span>
            <span class="q-gutter-md">
                <q-btn push glossy color="primary" @click="openLink('https://github.com/Junvary/gin-quasar-admin')">
                    Github
                </q-btn>
                <q-btn push glossy color="primary">
                    {{ $t('Version') }}{{ $t('Info') }}
                    <GqaVersion />
                </q-btn>
                <q-btn push glossy color="primary" @click="openLink('https://gitee.com/junvary/gin-quasar-admin')">
                    Gitee
                </q-btn>
            </span>
        </div>
        <div class="col row column justify-center items-center" style="margin-bottom: 80px;">
            <q-card style="width: 60%; background: rgba(255, 255, 255, 0.2); padding: 10px;">
                <LoginForm />
            </q-card>
        </div>
    </div>
</template>

<script setup>
import useCommon from 'src/composables/useCommon'
import GqaVersion from 'src/components/GqaVersion/index.vue';
import GqaAvatar from 'src/components/GqaAvatar/index.vue';
import LoginForm from './LoginForm.vue';

const { gqaFrontend, openLink } = useCommon()
</script>
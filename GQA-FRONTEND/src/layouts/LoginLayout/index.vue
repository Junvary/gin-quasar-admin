<template>
    <q-layout style="overflow-x: hidden">
        <q-page-container>

            <page-banner :checkDbStatus="checkDbStatus" ref="pageBanner" />

            <GqaScrollDown id="login-layout-scroll-down" class="login-layout-scroll-down"
                v-if="pluginCurrent && pluginComponent && showScrollDown" @click="scrollDown" />

            <component v-if="pluginCurrent && pluginComponent" :key="pluginCurrent" :is="pluginComponent" />

            <q-card v-else class="row items-center justify-center" style="padding: 20px 0;">
                <q-btn color="primary" :label="$t('LoginLayoutWithoutPlugin')"></q-btn>
            </q-card>

            <page-footer style="z-index: 0" />

            <q-page-scroller position="bottom-right" :scroll-offset="150" :offset="[3, 18]">
                <q-btn dense fab push icon="keyboard_arrow_up" color="primary" />
            </q-page-scroller>
        </q-page-container>
        <q-dialog v-model="initDbVisible" persistent>
            <q-card style="width: 100%; max-width: 50vw">
                <q-toolbar>
                    <q-avatar class="gin-quasar-admin-logo">
                        <img src="gqa128.png" />
                    </q-avatar>

                    <q-toolbar-title class="row items-center">
                        <span class="text-weight-bold">
                            {{ $t('WelcomeTo')}}
                            Gin-Quasar-Admin
                        </span>
                        <q-space />
                        <GqaLanguage style="width: 15%" />
                        <q-btn dense push rounded glossy color="primary"
                            @click="openLink('https://gitee.com/junvary/gin-quasar-admin')">
                            Gitee
                        </q-btn>
                        <q-btn dense push rounded glossy color="primary" style="margin: 0 5px"
                            @click="openLink('https://github.com/Junvary/gin-quasar-admin')">
                            Github
                        </q-btn>

                        <q-btn dense push rounded glossy color="primary">
                            {{ $t('Versions') }}
                            <GqaVersion />
                        </q-btn>
                    </q-toolbar-title>
                </q-toolbar>

                <GqaPluginList />

                <q-card-section>
                    <q-form class="text-center" ref="initDbForm">
                        <q-stepper v-model="step" ref="stepper" color="primary" animated>

                            <q-step :name="1" :title="$t('Welcome')" icon="home" :done="step > 1">
                                <div class="text-h5 column text-center q-gutter-md">
                                    <span class="col">
                                        {{ $t('WelcomeTo')}}
                                        Gin-Quasar-Admin
                                    </span>
                                    <span class="col">
                                        {{ $t('LoginLayoutWizardFirstTimeText1') }}
                                        <span class="text-negative">
                                            {{ $t('LoginLayoutWizardFirstTimeText2') }}
                                        </span>
                                    </span>
                                    <span class="col">
                                        {{ $t('LoginLayoutWizardFirstTimeText3') }}
                                        <span class="text-negative">
                                            {{ $t('LoginLayoutWizardFirstTimeText4') }}
                                        </span>
                                    </span>
                                    <span class="col text-subtitle1 text-negative">
                                        {{ $t('LoginLayoutWizardFirstTimeText5')}}
                                    </span>
                                </div>
                            </q-step>

                            <q-step :name="2" :title="$t('LoginLayoutWizardFirstTimeText2')" icon="settings"
                                :done="step > 2">
                                <div class="q-gutter-y-md column">
                                    <span class="col text-negative text-subtitle1">
                                        ** {{ $t('LoginLayoutWizardInitDBText1') }} **
                                    </span>
                                    <q-input outlined bottom-slots v-model.trim="form.dbType"
                                        :label="$t('LoginLayoutWizardInitDBFormDBType')" disable
                                        :rules="[(val) =>(val && val.length > 0) || $t('LoginLayoutWizardInitDBFormDBTypeRule'),]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>

                                    <div class="row">
                                        <q-input class="col" outlined v-model.trim="form.host"
                                            :label="$t('LoginLayoutWizardInitDBFormHost')"
                                            :rules="[(val) =>(val && val.length > 0) || $t('LoginLayoutWizardInitDBFormHostRule'),]">
                                            <template v-slot:before>
                                                <q-avatar class="gin-quasar-admin-logo">
                                                    <img src="gqa128.png" />
                                                </q-avatar>
                                            </template>
                                        </q-input>
                                        <q-input class="col" outlined v-model.trim="form.port"
                                            :label="$t('LoginLayoutWizardInitDBFormPort')"
                                            :rules="[(val) =>(val && val.length > 0) || $t('LoginLayoutWizardInitDBFormPortRule'),]">
                                            <template v-slot:before>
                                                <q-avatar class="gin-quasar-admin-logo">
                                                    <img src="gqa128.png" />
                                                </q-avatar>
                                            </template>
                                        </q-input>
                                    </div>

                                    <div class="row">
                                        <q-input class="col" outlined v-model.trim="form.username"
                                            :label="$t('LoginLayoutWizardInitDBFormUsername')"
                                            :rules="[(val) =>(val && val.length > 0) || $t('LoginLayoutWizardInitDBFormUsernameRule'),]">
                                            <template v-slot:before>
                                                <q-avatar class="gin-quasar-admin-logo">
                                                    <img src="gqa128.png" />
                                                </q-avatar>
                                            </template>
                                        </q-input>
                                        <q-input class="col" outlined v-model.trim="form.password"
                                            :label="$t('LoginLayoutWizardInitDBFormPassword')"
                                            :rules="[(val) =>(val && val.length > 0) || $t('LoginLayoutWizardInitDBFormPasswordRule'),]">
                                            <template v-slot:before>
                                                <q-avatar class="gin-quasar-admin-logo">
                                                    <img src="gqa128.png" />
                                                </q-avatar>
                                            </template>
                                        </q-input>
                                    </div>

                                    <q-input outlined v-model.trim="form.dbName"
                                        :label="$t('LoginLayoutWizardInitDBFormSchema')"
                                        :rules="[(val) =>(val && val.length > 0) || $t('LoginLayoutWizardInitDBFormSchemaRule'),]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>
                                </div>
                            </q-step>

                            <q-step :name="3" :title="$t('LoginLayoutWizardFirstTimeText6')" icon="settings"
                                :header-nav="step > 3">
                                <div class="text-h5 column text-center q-gutter-md">
                                    <span class="col">
                                        {{ $t('LoginLayoutWizardInitDBSummaryDBUser') }}
                                        <span class="text-negative">{{form.username}}</span>
                                        {{ $t('LoginLayoutWizardInitDBSummaryDBPassword') }}
                                        <span class="text-negative">{{form.password}}</span>
                                        {{ $t('LoginLayoutWizardInitDBSummaryDBPasswordSave') }}
                                    </span>

                                    <span class="col">
                                        {{ $t('LoginLayoutWizardInitDBSummaryDBHost') }}
                                        <span class="text-negative">{{form.host}}:{{form.port}}</span>
                                        {{ $t('LoginLayoutWizardInitDBSummaryDBType') }}
                                        <span class="text-negative">{{form.dbType}}</span>
                                        {{ $t('LoginLayoutWizardInitDBSummaryDBName') }}
                                        <span class="text-negative">{{form.dbName}}</span>
                                    </span>

                                    <span class="col">
                                        {{ $t('LoginLayoutWizardInitDBAfterText1') }}
                                    </span>
                                    <span class="col">
                                        {{ $t('LoginLayoutWizardFirstTimeText6') }}？
                                    </span>
                                </div>
                            </q-step>
                            <template v-slot:navigation>
                                <q-stepper-navigation>
                                    <q-btn @click="onInitDb" color="primary" :label="initLabel" />
                                    <q-btn v-if="step > 1" flat color="primary" @click="$refs.stepper.previous()"
                                        :label="$t('LoginLayoutWizardStepBack')" class="q-ml-sm" />
                                </q-stepper-navigation>
                            </template>

                        </q-stepper>
                        <q-inner-loading :showing="initLoading">
                            <q-spinner-gears size="50px" color="primary" />
                        </q-inner-loading>
                    </q-form>
                </q-card-section>

            </q-card>
        </q-dialog>
    </q-layout>
</template>

<script>
import { mapActions } from 'vuex'
import GqaLanguage from 'src/components/GqaLanguage'
import PageBanner from './PageBanner'
import PageFooter from './PageFooter'
import { getAction, postAction } from 'src/api/manage'
import { checkDbUrl, initDbUrl } from 'src/api/url'
import GqaVersion from 'src/components/GqaVersion'
import GqaPluginList from 'src/components/GqaPluginList'
import { markRaw, defineAsyncComponent } from 'vue'
import GqaScrollDown from 'src/components/GqaScrollDown'

export default {
    components: {
        GqaLanguage,
        PageBanner,
        PageFooter,
        GqaVersion,
        GqaPluginList,
        GqaScrollDown,
    },
    computed: {
        initLabel() {
            if (this.step === 1) {
                return this.$t('LoginLayoutWizardStepStart')
            } else if (this.step === 2) {
                return this.$t('LoginLayoutWizardStepNext')
            } else if (this.step === 3) {
                return this.$t('LoginLayoutWizardStepInitializingStart')
            } else {
                return 'Next'
            }
        },
    },
    data() {
        return {
            checkDbStatus: true,
            initDbVisible: false,
            step: 1,
            form: {
                dbType: 'Mysql',
                host: '127.0.0.1',
                port: '3306',
                username: 'root',
                password: '',
                dbName: 'gin-quasar-admin',
            },
            initLoading: false,
            pluginCurrent: null,
            pluginComponent: null,
            showScrollDown: true,
        }
    },
    created() {
        this.checkDb()
        window.addEventListener('scroll', this.documentTop)
    },
    unmounted() {
        window.removeEventListener('scroll', this.documentTop)
    },
    methods: {
        documentTop() {
            let top = document.documentElement.scrollTop
            if (top > 200) {
                this.showScrollDown = false
            } else {
                this.showScrollDown = true
            }
        },
        ...mapActions('storage', ['SetGqaGoVersion', 'SetGqaGinVersion', 'SetGqaPluginList']),
        checkDb() {
            this.checkDbStatus = true
            getAction(checkDbUrl).then((res) => {
                if (res.code === 1) {
                    this.SetGqaGoVersion(res.data.goVersion)
                    this.SetGqaGinVersion(res.data.ginVersion)
                    this.SetGqaPluginList(res.data.pluginList)
                    if (res.data.needInit === false) {
                        this.getPublic()
                        this.checkDbStatus = false
                        this.pluginCurrent = res.data.pluginLoginLayout
                        if (this.pluginCurrent) {
                            try {
                                this.pluginComponent = markRaw(defineAsyncComponent(() => import(`src/layouts/LoginLayout/${this.pluginCurrent}/index.vue`)))
                            } catch (error) {
                                this.$q.notify({
                                    type: 'negative',
                                    message: this.$t('LoginLayoutWizardInitDBAfterText2'),
                                })
                            }
                        }
                    }
                    if (res.data.needInit === true) {
                        this.$q.notify({
                            type: 'warning',
                            // message: res.message,
                            message: this.$t('DbNeedInit'),
                        })
                        this.initDbVisible = true
                    }
                }
            })
        },
        onInitDb() {
            if (this.step === 1) {
                this.$refs.stepper.next()
            } else if (this.step === 2) {
                this.$refs.initDbForm.validate().then((success) => {
                    if (success) {
                        this.$refs.stepper.next()
                    }
                })
            } else if (this.step === 3) {
                // 再次验证
                this.$refs.initDbForm.validate().then((success) => {
                    if (success) {
                        this.initLoading = true
                        postAction(initDbUrl, this.form)
                            .then((res) => {
                                if (res.code === 1) {
                                    this.$refs.stepper.next()
                                    this.$q.notify({
                                        type: 'positive',
                                        message: this.$t('DbInitSuccess'),
                                    })
                                    this.initDbVisible = false
                                    this.checkDbStatus = false
                                    this.getPublic()
                                    this.$refs.pageBanner.showLoginForm()
                                }
                            })
                            .finally(() => {
                                this.initLoading = false
                            })
                    } else {
                        this.step = 2
                    }
                })
            }
        },
        ...mapActions('storage', ['GetGqaDict', 'GetGqaFrontend', 'GetGqaBackend']),
        getPublic() {
            this.GetGqaDict()
            this.GetGqaFrontend()
            this.GetGqaBackend()
        },
        openLink(url) {
            window.open(url)
        },
        changeSuccess() {
            this.$router.go(0)
        },
        scrollDown() {
            document.getElementById('login-layout-scroll-down').nextElementSibling.scrollIntoView({
                behavior: 'smooth',
            })
        },
    },
}
</script>

<style lang="scss" scoped>
.login-layout-scroll-down {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: -60px;
    z-index: 100;
}
</style>
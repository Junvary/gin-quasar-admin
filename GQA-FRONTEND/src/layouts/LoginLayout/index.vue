<template>
    <q-layout>
        <q-page-container>
            <page-header />
            <page-banner :checkDbStatus="checkDbStatus" ref="pageBanner" />
            <page-news />
            <page-project />
            <!-- <page-honour /> -->
            <page-weapon />
            <page-document />
            <page-download />
            <page-footer />
        </q-page-container>
        <q-dialog v-model="initDbVisible" persistent>
            <q-card style="width: 100%; max-width: 50vw">
                <q-toolbar>
                    <q-avatar class="gin-quasar-admin-logo">
                        <img src="gqa128.png" />
                    </q-avatar>

                    <q-toolbar-title>
                        <span class="text-weight-bold">
                            {{ $t('LoginLayoutWizardTitle') }}
                        </span>
                    </q-toolbar-title>

                </q-toolbar>
                <q-card-section>
                    <q-form class="text-center" ref="initDbForm">
                        <q-stepper v-model="step" ref="stepper" color="primary" animated>

                            <q-step :name="1" :title="$t('LoginLayoutWizardWelcome')" icon="home" :done="step > 1">
                                <div class="text-center text-h5 column q-gutter-md">
                                    <span class="col">
                                        {{ $t('LoginLayoutWizardFirstTime') }}
                                    </span>
                                    <span class="col">
                                        {{ $t('LoginLayoutWizardFirstTimeText1') }}
                                        <span class="text-red">{{ $t('LoginLayoutWizardFirstTimeText2') }}</span>，
                                        {{ $t('LoginLayoutWizardFirstTimeText3') }}
                                        <span class="text-red">{{ $t('LoginLayoutWizardFirstTimeText4') }}</span>
                                    </span>
                                    <span class="col">
                                       {{ $t('LoginLayoutWizardFirstTimeText5') }}
                                    </span>
                                </div>
                            </q-step>

                            <q-step :name="2" :title="$t('LoginLayoutWizardInitDB')" icon="settings" :done="step > 2">
                                <div class="q-gutter-y-md column">
                                    <span class="col text-red text-h6">
                                        {{ $t('LoginLayoutWizardInitDBText1') }}
                                    </span>
                                    <span class="col text-red text-h7">
                                        {{ $t('LoginLayoutWizardInitDBText2') }}
                                    </span>
                                    <q-input outlined bottom-slots v-model.trim="form.dbType" :label="$t('LoginLayoutWizardInitDBFormDDBType')" disable
                                        :rules="[(val) =>(val && val.length > 0) || $t('LoginLayoutWizardInitDBFormDDBTypeRule'),]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                        <template v-slot:hint>
                                            {{ $t('LoginLayoutWizardInitDBFormDDBMysqlOnly') }}
                                        </template>
                                    </q-input>
                                    <q-input outlined v-model.trim="form.host" :label="$t('LoginLayoutWizardInitDBFormHost')"
                                        :rules="[(val) =>(val && val.length > 0) || $t('LoginLayoutWizardInitDBFormHostRule'),]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>
                                    <q-input outlined v-model.trim="form.port" :label="$t('LoginLayoutWizardInitDBFormPort')"
                                        :rules="[(val) =>(val && val.length > 0) || $t('LoginLayoutWizardInitDBFormPortRule'),]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>
                                    <q-input outlined v-model.trim="form.username" :label="$t('LoginLayoutWizardInitDBFormUsername')"
                                        :rules="[(val) =>(val && val.length > 0) || $t('LoginLayoutWizardInitDBFormUsernameRule'),]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>
                                    <q-input outlined v-model.trim="form.password" :label="$t('LoginLayoutWizardInitDBFormPassword')"
                                        :rules="[(val) =>(val && val.length > 0) || $t('LoginLayoutWizardInitDBFormPasswordRule'),]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>
                                    <q-input outlined v-model.trim="form.dbName" :label="$t('LoginLayoutWizardInitDBFormSchema')"
                                        :rules="[(val) =>(val && val.length > 0) || $t('LoginLayoutWizardInitDBFormSchemaRule'),]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>
                                </div>
                            </q-step>

                            <q-step :name="3" :title="$t('LoginLayoutWizardInitDBSummary')" icon="settings" :header-nav="step > 3">
                                <div class="text-center text-h5 column q-gutter-md">
                                    <span class="col">
                                        {{ $t('LoginLayoutWizardInitDBSummaryDBUser') }}
                                        <span class="text-red">{{form.username}}</span>
                                        {{ $t('LoginLayoutWizardInitDBSummaryDBPassword') }}
                                        <span class="text-red">{{form.password}}</span>
                                        {{ $t('LoginLayoutWizardInitDBSummaryDBPasswordSave') }}
                                    </span>

                                    <span class="col">
                                        {{ $t('LoginLayoutWizardInitDBSummaryDBHost') }}
                                        <span class="text-red">{{form.host}}:{{form.port}}</span>
                                        {{ $t('LoginLayoutWizardInitDBSummaryDBType') }}
                                        <span class="text-red">{{form.dbType}}</span>
                                        {{ $t('LoginLayoutWizardInitDBSummaryDBName') }}
                                        <span class="text-red">{{form.dbName}}</span>
                                    </span>

                                    <span class="col">
                                        {{ $t('LoginLayoutWizardInitDBAfterText1') }}
                                        <span class="text-red">{{ $t('LoginLayoutWizardSuperAdminLogin') }}</span>
                                        {{ $t('LoginLayoutWizardInitDBAfterText2') }}
                                        <span class="text-red">{{ $t('LoginLayoutWizardPersonalConfigurations') }}</span>
                                    </span>
                                    <span class="col">
                                        {{ $t('LoginLayoutWizardNextSuperAdmin') }}
                                        <span class="text-red">{{ $t('LoginLayoutWizardSuperAdmin') }}</span>
                                        {{ $t('LoginLayoutWizardMenuFound') }}
                                        <span class="text-red">{{ $t('LoginLayoutWizardPersonalConfigurations') }}</span>
                                        {{ $t('LoginLayoutWizardEntry') }}
                                    </span>
                                    <span class="col">
                                        {{ $t('LoginLayoutWizardDefaults') }}
                                    </span>
                                    <span class="col">
                                        {{ $t('LoginLayoutWizardReady')}}
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
import PageHeader from './PageHeader'
import PageBanner from './PageBanner'
import PageNews from './PageNews'
import PageProject from './PageProject'
// import PageHonour from './PageHonour'
import PageWeapon from './PageWeapon'
import PageDocument from './PageDocument'
import PageDownload from './PageDownload'
import PageFooter from './PageFooter'

import { getAction, postAction } from 'src/api/manage'
import { checkDbUrl, initDbUrl, dictDetailUrl } from 'src/api/url'

import { ArrayToTree } from 'src/utils/arrayAndTree'

export default {
    components: {
        PageHeader,
        PageBanner,
        PageNews,
        PageProject,
        // PageHonour,
        PageWeapon,
        PageDocument,
        PageDownload,
        PageFooter,
    },
    computed: {
        initLabel() {
            if (this.step === 1) {
                return $t('LoginLayoutWizardStepStart')
            } else if (this.step === 2) {
                return $t('LoginLayoutWizardStepNext')
            } else if (this.step === 3) {
                return $t('LoginLayoutWizardStepInitializingStart')
            } else {
                return $t('LoginLayoutWizardStepNext')
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
        }
    },
    created() {
        this.checkDb()
    },
    methods: {
        checkDb() {
            this.checkDbStatus = true
            getAction(checkDbUrl).then((res) => {
                if (res.code === 1) {
                    if (res.data.needInit === false) {
                        this.getDictDetail()
                        this.checkDbStatus = false
                    }
                    if (res.data.needInit === true) {
                        this.$q.notify({
                            type: 'warning',
                            message: res.message,
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
                                        message: res.message,
                                    })
                                    this.initDbVisible = false
                                    this.checkDbStatus = false
                                    this.getDictDetail()
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
        getDictDetail() {
            postAction(dictDetailUrl).then((res) => {
                if (res.code === 1) {
                    const dictDetail = res.data.records
                    const dictList = ArrayToTree(dictDetail)
                    let dict = {}
                    for (let d of dictList) {
                        dict[d.value] = d.children
                    }
                    this.$q.cookies.set('gqa-dict', dict)
                }
            })
        },
    },
}
</script>

<style lang="scss" scoped>
</style>

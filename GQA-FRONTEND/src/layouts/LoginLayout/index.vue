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
                            欢迎使用 Gin-Quasar-Admin！
                        </span>
                    </q-toolbar-title>

                </q-toolbar>
                <q-card-section>
                    <q-form class="text-center" ref="initDbForm">
                        <q-stepper v-model="step" ref="stepper" color="primary" animated>

                            <q-step :name="1" title="欢迎" icon="home" :done="step > 1">
                                <div class="text-h5 column text-center q-gutter-md">
                                    <span class="col">
                                        嘿,看起来这是你第一次使用 Gin-Quasar-Admin
                                    </span>
                                    <span class="col">
                                        首先你需要
                                        <span class="text-red">数据库初始化</span>，
                                        后续根据提示进行网站
                                        <span class="text-red">个性化配置</span>
                                    </span>
                                    <span class="col">
                                        准备好了吗？让我们开始吧！
                                    </span>
                                </div>
                            </q-step>

                            <q-step :name="2" title="数据库初始化" icon="settings" :done="step > 2">
                                <div class="q-gutter-y-md column">
                                    <span class="col text-red text-h6">
                                        首先，你需要对数据库进行初始化，请确保已经安装并启动了数据库
                                    </span>
                                    <q-input outlined bottom-slots v-model.trim="form.dbType" label="数据库类型" disable
                                        :rules="[(val) =>(val && val.length > 0) || '请输入数据库类型',]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                        <template v-slot:hint>
                                            只支持Mysql
                                        </template>
                                    </q-input>
                                    <q-input outlined v-model.trim="form.host" label="数据库地址"
                                        :rules="[(val) =>(val && val.length > 0) || '请输入数据库地址',]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>
                                    <q-input outlined v-model.trim="form.port" label="数据库端口"
                                        :rules="[(val) =>(val && val.length > 0) || '请输入数据库端口',]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>
                                    <q-input outlined v-model.trim="form.username" label="数据库用户名"
                                        :rules="[(val) =>(val && val.length > 0) || '请输入数据库用户名',]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>
                                    <q-input outlined v-model.trim="form.password" label="数据库密码"
                                        :rules="[(val) =>(val && val.length > 0) || '请输入数据库密码',]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>
                                    <q-input outlined v-model.trim="form.dbName" label="数据库/Schema名称"
                                        :rules="[(val) =>(val && val.length > 0) || '请输入数据库/Schema名',]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>
                                </div>
                            </q-step>

                            <q-step :name="3" title="准备就绪" icon="settings" :header-nav="step > 3">
                                <div class="text-h5 column text-center q-gutter-md">
                                    <span class="col">
                                        数据库用户名
                                        <span class="text-red">{{form.username}}</span>
                                        密码
                                        <span class="text-red">{{form.password}}</span>
                                        将保存在后台中
                                    </span>

                                    <span class="col">
                                        系统将为你建立
                                        <span class="text-red">{{form.host}}:{{form.port}}</span>
                                        的
                                        <span class="text-red">{{form.dbType}}</span>
                                        数据库
                                        <span class="text-red">{{form.dbName}}</span>
                                    </span>

                                    <span class="col">
                                        完成此步后，你需要使用
                                        <span class="text-red">超级管理员登录</span>
                                        进行网站
                                        <span class="text-red">个性化配置</span>
                                    </span>
                                    <span class="col">
                                        后续你也可以在
                                        <span class="text-red">超级管理员</span>
                                        的相关菜单中找到
                                        <span class="text-red">个性化配置</span>
                                        入口
                                    </span>
                                    <span class="col">
                                        默认超级管理员：admin，密码：123456
                                    </span>
                                    <span class="col">
                                        准备就绪？
                                    </span>
                                </div>
                            </q-step>
                            <template v-slot:navigation>
                                <q-stepper-navigation>
                                    <q-btn @click="onInitDb" color="primary" :label="initLabel" />
                                    <q-btn v-if="step > 1" flat color="primary" @click="$refs.stepper.previous()"
                                        label="上一步" class="q-ml-sm" />
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
                return '开始'
            } else if (this.step === 2) {
                return '下一步'
            } else if (this.step === 3) {
                return '开始初始化'
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

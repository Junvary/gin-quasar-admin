<template>
    <q-layout style="overflow-x: hidden">
        <q-page-container>
            <page-banner :checkDbStatus="checkDbStatus" ref="pageBanner" />

            <component v-if="pluginCurrent && pluginComponent" :key="pluginCurrent" :is="pluginComponent" />
            <q-card v-else class="row items-center justify-center" style="padding: 20px 0;">
                <q-btn color="primary" label="未安装任何登录页插件，请联系管理员！"></q-btn>
            </q-card>

            <page-footer />
            <q-page-scroller position="bottom-right" :scroll-offset="150" :offset="[18, 18]">
                <q-btn dense fab push icon="keyboard_arrow_up" color="primary" />
            </q-page-scroller>
        </q-page-container>
        <q-dialog v-model="initDbVisible" persistent>
            <q-card style="width: 100%; max-width: 50vw">
                <q-toolbar>
                    <q-avatar class="gin-quasar-admin-logo">
                        <img src="gqa128.png" />
                    </q-avatar>

                    <q-toolbar-title class="row justify-between items-center">
                        <span class="text-weight-bold">
                            欢迎使用 Gin-Quasar-Admin！
                        </span>
                        <span class="q-gutter-md">
                            <q-btn dense push rounded glossy color="primary"
                                @click="openLink('https://gitee.com/junvary/gin-quasar-admin')">
                                访问Gitee
                            </q-btn>
                            <q-btn dense push rounded glossy color="primary"
                                @click="openLink('https://github.com/Junvary/gin-quasar-admin')">
                                访问Github
                            </q-btn>

                            <q-btn dense push rounded glossy color="primary">
                                查看版本信息
                                <GqaVersion />
                            </q-btn>
                        </span>
                    </q-toolbar-title>
                </q-toolbar>

                <GqaPluginList />

                <q-card-section>
                    <q-form class="text-center" ref="initDbForm">
                        <q-stepper v-model="step" ref="stepper" color="primary" animated>

                            <q-step :name="1" title="欢迎" icon="home" :done="step > 1">
                                <div class="text-h5 column text-center q-gutter-md">
                                    <span class="col">
                                        欢迎使用 Gin-Quasar-Admin
                                    </span>
                                    <span class="col">
                                        请跟随本向导完成
                                        <span class="text-negative">数据库初始化</span>
                                    </span>
                                    <span class="col">
                                        完成后可使用管理员账户登录进行
                                        <span class="text-negative">网站个性化配置</span>
                                    </span>
                                    <span class="col text-subtitle1 text-negative">
                                        ** 可通过上方 “已装插件” 选择登录首页组件 **
                                    </span>
                                </div>
                            </q-step>

                            <q-step :name="2" title="初始化数据库" icon="settings" :done="step > 2">
                                <div class="q-gutter-y-md column">
                                    <span class="col text-negative text-subtitle1">
                                        ** 请确保已经安装并启动了Mysql，但你无须建库，向导会根据下方配置为你自动创建数据库，并导入初始数据 **
                                    </span>
                                    <q-input outlined bottom-slots v-model.trim="form.dbType" label="数据库类型" disable
                                        :rules="[(val) =>(val && val.length > 0) || '请输入数据库类型',]">
                                        <template v-slot:before>
                                            <q-avatar class="gin-quasar-admin-logo">
                                                <img src="gqa128.png" />
                                            </q-avatar>
                                        </template>
                                    </q-input>

                                    <div class="row">
                                        <q-input class="col" outlined v-model.trim="form.host" label="数据库地址"
                                            :rules="[(val) =>(val && val.length > 0) || '请输入数据库地址',]">
                                            <template v-slot:before>
                                                <q-avatar class="gin-quasar-admin-logo">
                                                    <img src="gqa128.png" />
                                                </q-avatar>
                                            </template>
                                        </q-input>
                                        <q-input class="col" outlined v-model.trim="form.port" label="数据库端口"
                                            :rules="[(val) =>(val && val.length > 0) || '请输入数据库端口',]">
                                            <template v-slot:before>
                                                <q-avatar class="gin-quasar-admin-logo">
                                                    <img src="gqa128.png" />
                                                </q-avatar>
                                            </template>
                                        </q-input>
                                    </div>

                                    <div class="row">
                                        <q-input class="col" outlined v-model.trim="form.username" label="数据库用户名"
                                            :rules="[(val) =>(val && val.length > 0) || '请输入数据库用户名',]">
                                            <template v-slot:before>
                                                <q-avatar class="gin-quasar-admin-logo">
                                                    <img src="gqa128.png" />
                                                </q-avatar>
                                            </template>
                                        </q-input>
                                        <q-input class="col" outlined v-model.trim="form.password" label="数据库密码"
                                            :rules="[(val) =>(val && val.length > 0) || '请输入数据库密码',]">
                                            <template v-slot:before>
                                                <q-avatar class="gin-quasar-admin-logo">
                                                    <img src="gqa128.png" />
                                                </q-avatar>
                                            </template>
                                        </q-input>
                                    </div>

                                    <q-input outlined v-model.trim="form.dbName" label="自动创建数据库/Schema名称"
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
                                        <span class="text-negative">{{form.username}}</span>
                                        密码
                                        <span class="text-negative">{{form.password}}</span>
                                        将保存在后台中
                                    </span>

                                    <span class="col">
                                        系统将为你建立
                                        <span class="text-negative">{{form.host}}:{{form.port}}</span>
                                        的
                                        <span class="text-negative">{{form.dbType}}</span>
                                        数据库
                                        <span class="text-negative">{{form.dbName}}</span>
                                    </span>

                                    <span class="col">
                                        网站默认超级管理员账户：admin / 123456 ，请尽早修改
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
import { mapActions } from 'vuex'
import PageBanner from './PageBanner'
import PageFooter from './PageFooter'
import { getAction, postAction } from 'src/api/manage'
import { checkDbUrl, initDbUrl } from 'src/api/url'
import GqaVersion from 'src/components/GqaVersion'
import GqaPluginList from 'src/components/GqaPluginList'
import { markRaw, defineAsyncComponent } from 'vue'

export default {
    components: {
        PageBanner,
        PageFooter,
        GqaVersion,
        GqaPluginList,
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
            pluginCurrent: null,
            pluginComponent: null,
        }
    },
    created() {
        this.checkDb()
    },
    methods: {
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
                                    message: '所选插件还未支持登录页面，请联系管理员！',
                                })
                            }
                        }
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
    },
}
</script>

<style lang="scss" scoped>
</style>

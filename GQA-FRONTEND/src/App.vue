<template>
    <router-view />
</template>

<script>
import { defineComponent } from 'vue'
import { gqaFrontendMixin } from 'src/mixins/gqaFrontendMixin'
import { GqaFrontendDefault } from 'src/settings'
import { mapGetters } from 'vuex'
import { useI18n } from 'vue-i18n'
import { Quasar } from 'quasar'

export default defineComponent({
    name: 'App',
    mixins: [gqaFrontendMixin],
    watch: {
        gqaFrontend: {
            handler(val) {
                document.title = val.gqaSubTitle
                this.createLink()
            },
            deep: true,
        },
    },
    computed: {
        ...mapGetters({
            language: 'user/language',
        }),
    },
    async mounted() {
        document.title = this.gqaFrontend.gqaSubTitle || GqaFrontendDefault.gqaSubTitle
        this.createLink()
        // 初始化语言
        const { locale } = useI18n({ useScope: 'global' })
        locale.value = this.language
        await import(
            /* webpackInclude: /(zh-CN|en-US)\.js$/ */
            'quasar/lang/' + this.language
        ).then((lang) => {
            // ! NOTICE ssrContext param:
            Quasar.lang.set(lang.default)
        })
        console.info('欢迎使用Gin-Quasar-Admin!')
        console.info('项目地址: https://github.com/Junvary/gin-quasar-admin ')
        console.info('欢迎交流，感谢Star!')
    },
    methods: {
        createLink() {
            const toDelete = document.getElementsByName('gqa-link-href')
            if (toDelete && toDelete.length) {
                document.getElementsByTagName('head')[0].removeChild(toDelete[0])
            }
            const gqaLink = document.createElement('link')
            gqaLink.type = 'image/ico'
            gqaLink.rel = 'icon'
            gqaLink.setAttribute('name', 'gqa-link-href')
            if (this.gqaFrontend.gqaHeaderLogo && this.gqaFrontend.gqaHeaderLogo !== '') {
                const gqaHeaderLogo = process.env.API + this.gqaFrontend.gqaHeaderLogo.substring(11)
                gqaLink.href = gqaHeaderLogo
            } else {
                gqaLink.href = 'favicon.ico'
            }
            document.getElementsByTagName('head')[0].appendChild(gqaLink)
        },
    },
})
</script>

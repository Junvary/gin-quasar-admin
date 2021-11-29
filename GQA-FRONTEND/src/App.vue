<template>
    <router-view />
</template>

<script>
import { defineComponent } from 'vue'
import { gqaFrontendMixin } from 'src/mixins/gqaFrontendMixin'
import { GqaFrontendDefault } from 'src/settings'

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
    mounted() {
        document.title = this.gqaFrontend.gqaSubTitle || GqaFrontendDefault.gqaSubTitle
        this.createLink()
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
                const gqaHeaderLogo = '/gqa-api/' + this.gqaFrontend.gqaHeaderLogo.substring(11)
                gqaLink.href = gqaHeaderLogo
            } else {
                gqaLink.href = 'favicon.ico'
            }
            document.getElementsByTagName('head')[0].appendChild(gqaLink)
        },
    },
})
</script>

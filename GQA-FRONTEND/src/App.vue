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
        gqaFrontend() {
            document.title = this.gqaFrontend.gqaSubTitle
        },
    },
    beforeMount() {
        document.title = this.gqaFrontend.gqaSubTitle || GqaFrontendDefault.gqaSubTitle
    },
})
</script>

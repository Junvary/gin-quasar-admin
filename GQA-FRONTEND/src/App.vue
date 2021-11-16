<template>
    <router-view />
</template>

<script>
import { defineComponent } from 'vue'
import { gqaFrontendMixin } from 'src/mixins/gqaFrontendMixin'

export default defineComponent({
    name: 'App',
    mixins: [gqaFrontendMixin],
    watch: {
        gqaFrontend() {
            document.title = this.gqaFrontend.gqaSubTitle
        },
    },
})
</script>

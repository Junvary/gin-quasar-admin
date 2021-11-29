<template>
    <footer>
        <div class="gin-quasar-admin-banner" style="padding-top: 50px; padding-bottom: 30px;">
            <div class="row justify-center">

                <div class="column items-center q-gutter-y-md">
                    <div class="col">
                        <GqaAvatar size="xl" :src="gqaFrontend.gqaWebLogo" />
                    </div>
                    <div class="col" style="font-size: 35px; font-weight: bold;letter-spacing: 5px;">
                        {{ gqaFrontend.gqaSubTitle }}
                    </div>
                    <span class="col" style="text-transform: capitalize;">
                        {{ gqaFrontend.gqaDescribe }}
                    </span>
                </div>
            </div>

            <q-separator />

            <div class="column items-center" style="margin-top: 20px">
                <span class="col">
                    技术支持：{{thisYear}}
                </span>
                <q-separator />
                <span class="col">
                    Powered by
                    <a href="https://github.com/Junvary/gin-quasar-admin" target="_blank" style="color: white">
                        Gin-Quasar-Admin
                    </a>
                </span>
            </div>
        </div>
    </footer>
</template>

<script>
import { gqaFrontendMixin } from 'src/mixins/gqaFrontendMixin'
import GqaAvatar from 'src/components/GqaAvatar'

export default {
    name: 'PageFooter',
    mixins: [gqaFrontendMixin],
    components: {
        GqaAvatar,
    },
    data() {
        return {
            thisYear: new Date(),
        }
    },
    mounted() {
        this.getThisYear()
    },
    methods: {
        getThisYear() {
            const year = new Date()
            this.thisYear = year.getFullYear()
        },
    },
}
</script>

<style lang="scss" scoped>
footer {
    position: relative;
    z-index: 1;
    overflow: hidden;
    color: #ffffff;
}
</style>
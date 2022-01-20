<template>
    <q-page>
        <q-img :src="bannerImage" class="absolute-full">
            <div class="absolute-full text-subtitle2 flex flex-center">
                <span style="font-size: 40px">
                    {{ $t('WelcomeTo') }}
                    {{ gqaFrontend.gqaSubTitle }}
                </span>
            </div>
        </q-img>
    </q-page>
</template>

<script>
import { gqaFrontendMixin } from 'src/mixins/gqaFrontendMixin'

export default {
    name: 'Dashboard',
    mixins: [gqaFrontendMixin],
    computed: {
        bannerImage() {
            if (this.gqaFrontend.gqaBannerImage && this.gqaFrontend.gqaBannerImage.substring(0, 11) === 'gqa-upload:') {
                return process.env.API + this.gqaFrontend.gqaBannerImage.substring(11)
            }
            return this.randomImg
        },
    },
    data() {
        return {
            randomImg: 'https://api.ixiaowai.cn/api/api.php',
        }
    },
}
</script>

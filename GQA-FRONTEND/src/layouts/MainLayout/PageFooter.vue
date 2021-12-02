<template>
    <div class="row justify-center items-center ">
        {{ gqaFrontend.gqaSubTitle }}
    </div>
</template>

<script>
import { gqaFrontendMixin } from 'src/mixins/gqaFrontendMixin'

export default {
    name: 'PageFooter',
    mixins: [gqaFrontendMixin],
}
</script>
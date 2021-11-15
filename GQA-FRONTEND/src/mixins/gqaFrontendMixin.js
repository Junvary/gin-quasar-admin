import { mapGetters } from 'vuex'

export const gqaFrontendMixin = {
    computed: {
        ...mapGetters({
            gqaFrontend: 'storage/gqaFrontend',
        }),
    }
}
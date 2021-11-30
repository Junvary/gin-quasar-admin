import { mapGetters } from 'vuex'

export const gqaBackendMixin = {
    computed: {
        ...mapGetters({
            gqaBackend: 'storage/gqaBackend',
        }),
    }
}
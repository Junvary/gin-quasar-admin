import { useStorageStore } from 'src/stores/storage'

export const gqaBackendMixin = {
    computed: {
        gqaBackend() {
            const storageStore = useStorageStore()
            return storageStore.GetGqaBackend()
        }
    }
}
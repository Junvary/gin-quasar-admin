import { useStorageStore } from 'src/stores/storage'

export const gqaFrontendMixin = {
    computed: {
        gqaFrontend() {
            const storageStore = useStorageStore()
            return storageStore.GetGqaFrontend()
        }
    }
}

// export default function gqaFrontendMixin() {
//     const storageStore = useStorageStore()
//     const config = storageStore.GetGqaFrontend()
//     return { config }
// }
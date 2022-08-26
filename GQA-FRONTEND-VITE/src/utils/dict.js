import { LocalStorage } from "quasar"
import { useStorageStore } from "src/stores/storage"


export async function DictOptions() {
    const detailLocal = LocalStorage.getItem('gqa-dict')
    const storageStore = useStorageStore()
    if (detailLocal) {
        return detailLocal
    } else {
        await storageStore.SetGqaDict().then(() => {
            return LocalStorage.getItem('gqa-dict')
        })

    }
}
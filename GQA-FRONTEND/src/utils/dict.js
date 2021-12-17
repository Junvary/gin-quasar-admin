import { LocalStorage } from "quasar"
import Store from 'src/store'


export async function DictOptions() {
    const detailLocal = LocalStorage.getItem('gqa-dict')
    if (detailLocal) {
        return detailLocal
    } else {
        await Store().dispatch('storage/GetGqaDict').then(() => {
            return LocalStorage.getItem('gqa-dict')
        })

    }
}
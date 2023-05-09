import { Quasar } from 'quasar';
import { computed } from 'vue';

const langList = import.meta.glob('../../node_modules/quasar/lang/*.mjs')

export default async () => {
    const langIso = computed(() => settingStore.GetLanguage());
    try {
        langList[`../../node_modules/quasar/lang/${langIso}.mjs`]().then(lang => {
            Quasar.lang.set(lang.default)
        })
    }
    catch (err) {
        // Requested Quasar Language Pack does not exist,
        // let's not break the app, so catching error
    }
}
import { LocalStorage } from 'quasar'
import { GqaFrontendDefault } from 'src/settings'

export function gqaDict(state) {
    const dict = LocalStorage.getItem("gqa-dict")
    if (state.gqaDict) {
        return state.gqaDict
    } else {
        return dict
    }

}

export function gqaFrontend(state) {
    const frontend = LocalStorage.getItem("gqa-frontend")
    if (state.gqaFrontend) {
        return state.gqaFrontend
    } else if (frontend) {
        return frontend
    } else {
        return GqaFrontendDefault
    }
}

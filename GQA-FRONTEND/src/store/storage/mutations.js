import { LocalStorage } from 'quasar'


export function SET_GQA_DICT(state, dict) {
    state.gqaDict = dict
    LocalStorage.set('gqa-dict', dict)
}

export function SET_GQA_FRONTEND(state, frontend) {
    state.gqaFrontend = frontend
    LocalStorage.set('gqa-frontend', frontend)
}
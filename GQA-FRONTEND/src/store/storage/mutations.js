import { LocalStorage } from 'quasar'


export function SET_GQA_DICT(state, dict) {
    state.gqaDict = dict
    LocalStorage.set('gqa-dict', dict)
}

export function SET_GQA_BACKEND(state, backend) {
    state.gqaBackend = backend
    LocalStorage.set('gqa-backend', backend)
}

export function SET_GQA_FRONTEND(state, frontend) {
    state.gqaFrontend = frontend
    LocalStorage.set('gqa-frontend', frontend)
}

export function SET_GQA_GO_VERSION(state, goVersion) {
    state.goVersion = goVersion
    LocalStorage.set('gqa-goVersion', goVersion)
}

export function SET_GQA_GIN_VERSION(state, ginVersion) {
    state.ginVersion = ginVersion
    LocalStorage.set('gqa-ginVersion', ginVersion)
}

export function SET_GQA_PLUGIN_LIST(state, pluginList) {
    state.pluginList = pluginList
    LocalStorage.set('gqa-pluginList', pluginList)
}
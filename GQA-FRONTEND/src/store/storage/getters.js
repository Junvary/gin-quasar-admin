import { LocalStorage } from 'quasar'
import { GqaFrontendDefault, GqaBackendDefault } from 'src/settings'

export function gqaDict(state) {
    const dict = LocalStorage.getItem("gqa-dict")
    if (state.gqaDict) {
        return state.gqaDict
    } else {
        return dict
    }
}

export function gqaBackend(state) {
    const backend = LocalStorage.getItem("gqa-backend")
    if (state.gqaBackend) {
        return state.gqaBackend
    } else if (backend) {
        return backend
    } else {
        return GqaBackendDefault
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

export function gqaGoVersion(state) {
    const goversion = LocalStorage.getItem("gqa-goVersion")
    if (state.goVersion) {
        return state.goVersion
    } else {
        return goversion
    }
}

export function gqaGinVersion(state) {
    const ginversion = LocalStorage.getItem("gqa-ginVersion")
    if (state.ginversion) {
        return state.ginversion
    } else {
        return ginversion
    }
}

export function gqaPluginList(state) {
    const pluginList = LocalStorage.getItem("gqa-pluginList")
    if (state.pluginList) {
        return state.pluginList
    } else {
        return pluginList
    }
}

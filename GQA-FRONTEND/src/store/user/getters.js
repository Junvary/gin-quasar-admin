import { Cookies, SessionStorage } from 'quasar'

export function token(state) {
    if (SessionStorage.getItem('gqa-token')) {
        return SessionStorage.getItem('gqa-token')
    } else if (Cookies.get('gqa-token')) {
        return Cookies.get('gqa-token')
    } else {
        return state.token
    }
}

export function username(state) {
    if (state.username) {
        return state.username
    } else if (Cookies.get('gqa-username')) {
        return Cookies.get('gqa-username')
    } else {
        return ""
    }
}

export function nickname(state) {
    if (state.nickname) {
        return state.nickname
    } else if (Cookies.get('gqa-nickname')) {
        return Cookies.get('gqa-nickname')
    } else {
        return ""
    }
}

export function realName(state) {
    if (state.realName) {
        return state.realName
    } else if (Cookies.get('gqa-realName')) {
        return Cookies.get('gqa-realName')
    } else {
        return ""
    }
}

export function avatar(state) {
    if (state.avatar) {
        return state.avatar
    } else if (Cookies.get('gqa-avatar')) {
        return Cookies.get('gqa-avatar')
    } else {
        return ""
    }
}

export function rememberMe(state) {
    return state.rememberMe
}

export function language(state) {
    if (state.language) {
        return state.language
    } else if (Cookies.get('gqa-language')) {
        return Cookies.get('gqa-language')
    } else {
        return 'zh-CN'
    }

}


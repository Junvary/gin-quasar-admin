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

export function nickname(state) {
    return state.nickname
}

export function realName(state) {
    return state.realName
}

export function avatar(state) {
    return state.avatar
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


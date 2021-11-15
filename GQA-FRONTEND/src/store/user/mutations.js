import { Cookies, SessionStorage, LocalStorage } from 'quasar'

export function SET_TOKEN(state, token) {
    state.token = token
    if (state.rememberMe) {
        Cookies.set('gqa-token', token)
    } else {
        SessionStorage.set('gqa-token', token)
    }
}

export function SET_NICKNAME(state, nickname) {
    state.nickname = nickname
    Cookies.set('gqa-nickname', nickname)
}

export function SET_REALNAME(state, realName) {
    state.realName = realName
    Cookies.set('gqa-realName', realName)
}

export function SET_AVATAR(state, avatar) {
    state.avatar = avatar
    Cookies.set('gqa-avatar', avatar)
}

export function CHANGE_REMEMBER_ME(state, type) {
    state.rememberMe = type
}

export function LOGOUT(state) {
    SessionStorage.remove('gqa-token')
    Cookies.remove('gqa-token')
    Cookies.remove('gqa-nickname')
    Cookies.remove('gqa-realName')
    Cookies.remove('gqa-avatar')
    // 字典不删除
    // LocalStorage.remove('gqa-dict')
    state.token = undefined
    state.nickname = undefined
    state.realName = undefined
    state.avatar = undefined
}
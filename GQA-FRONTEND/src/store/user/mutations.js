import { Cookies } from 'quasar'

export function SET_TOKEN(state, token) {
    state.token = token
    Cookies.set('gqa-token', token)
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

export function LOGOUT(state) {
    Cookies.remove('gqa-token')
    Cookies.remove('gqa-nickname')
    Cookies.remove('gqa-realName')
    Cookies.remove('gqa-avatar')
    state.token = undefined
    state.nickname = undefined
    state.realName = undefined
    state.avatar = undefined
}
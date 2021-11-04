import { postAction } from 'src/api/manage'
import { loginUrl } from 'src/api/url'

export async function HandleLogin({ commit, dispatch }, loginForm) {
    const res = await postAction(loginUrl, loginForm)
    if (res.code === 1) {
        const token = res.data.token
        const nickname = res.data.nickname
        const realName = res.data.realName
        const avatar = res.data.avatar
        dispatch('SetToken', token)
        commit('SET_NICKNAME', nickname)
        commit('SET_REALNAME', realName)
        commit('SET_AVATAR', avatar)
        return true
    } else {
        return
    }
}

export function SetToken({ commit }, token) {
    commit('SET_TOKEN', token)
}

export function ChangeRememberMe({ commit }, type) {
    commit('CHANGE_REMEMBER_ME', type)
}

export async function HandleLogout({ commit, dispatch }) {
    await dispatch('permission/ClearMenu', null, { root: true })
    commit('LOGOUT')
}


import { postAction } from 'src/api/manage'
import { loginUrl } from 'src/api/url'

export async function HandleLogin({ commit }, loginForm) {
    const res = await postAction(loginUrl, loginForm)
    if (res.code === 1) {
        const token = res.data.token
        const nickname = res.data.nickname
        const realName = res.data.realName
        const avatar = res.data.avatar
        commit('SET_TOKEN', token)
        commit('SET_NICKNAME', nickname)
        commit('SET_REALNAME', realName)
        commit('SET_AVATAR', avatar)
        return true
    } else {
        return
    }
}

export async function HandleLogout({ commit, dispatch }) {
    await dispatch('permission/ClearMenu', null, { root: true })
    commit('LOGOUT')
}

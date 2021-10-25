import { HomeMenu } from 'src/settings'

export function AddTabMenu({ commit, state }, tab) {
    if (!state.tabMenus.length) {
        commit('INIT_TAB_MENU')
    }
    // 判断当前菜单数组中是否存在要加入的菜单，如果存在返回false
    if (!state.tabMenus.some(item => item.fullPath === tab.fullPath)) {
        commit("ADD_TAB_MENU", tab)
        commit('CHANGE_CURRENT_TAB', tab)
    }
}

export function ChangeCurrentTab({ commit }, tab) {
    commit('CHANGE_CURRENT_TAB', tab)
}

export function RemoveTab({ commit }, tab) {
    // 多一次判断，如果要关闭的是首页，返回false
    if (tab.fullPath === HomeMenu[0].fullPath) {
        return false
    }
    commit('REMOVE_TAB', tab)
    return true
}

export function RemoveTabMenu({ commit }) {
    commit('REMOVE_TAB_MENU')
}
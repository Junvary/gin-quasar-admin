export function AddTabMenu({ commit, state, rootState }, tab) {
    const base = rootState.permission.userMenu[0]
    // 退出时userMenu被清空，不走下面逻辑
    if (base) {
        // 如果没有仪表盘，则加入仪表盘
        if (state.tabMenus.filter(item => item.path === base.path).length === 0) {
            commit("ADD_TAB_MENU", base)
            commit('CHANGE_CURRENT_TAB', base)
        }
        // 判断tab是否存在，是为了关闭所有菜单时，不会传递tab
        if (tab && !state.tabMenus.some(item => item.path === tab.path)) {
            commit("ADD_TAB_MENU", tab)
            commit('CHANGE_CURRENT_TAB', tab)
        }
    }
}

export function ChangeCurrentTab({ commit }, tab) {
    commit('CHANGE_CURRENT_TAB', tab)
}

export function RemoveTab({ commit }, tab) {
    commit('REMOVE_TAB', tab)
}

export function RemoveRightTab({ commit }, tab) {
    commit('REMOVE_RIGHT_TAB', tab)
}

export function RemoveLeftTab({ commit, rootState }, tab) {
    const base = rootState.permission.userMenu[0]
    commit('REMOVE_LEFT_TAB', { base, tab })
}

export function DestroyTabMenu({ commit }) {
    commit('DESTROY_TAB_MENU')
}
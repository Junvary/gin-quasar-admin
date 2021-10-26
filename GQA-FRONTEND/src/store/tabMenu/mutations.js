export function ADD_TAB_MENU(state, tab) {
    state.tabMenus = state.tabMenus.concat([tab])
}

export function CHANGE_CURRENT_TAB(state, tab) {
    state.currentTab = tab
}

export function REMOVE_TAB(state, tab) {
    const removeIndex = state.tabMenus.indexOf(tab)
    state.tabMenus = state.tabMenus.filter(item => item.path !== tab.path)
    state.currentTab = state.tabMenus[removeIndex - 1]
}

export function REMOVE_RIGHT_TAB(state, tab) {
    const removeIndex = state.tabMenus.indexOf(tab)
    state.tabMenus = state.tabMenus.slice(0, removeIndex + 1)
}

export function REMOVE_LEFT_TAB(state, arg) {
    const removeIndex = state.tabMenus.indexOf(arg.tab)
    const rightMenu = state.tabMenus.slice(removeIndex)
    state.tabMenus = [arg.base].concat(rightMenu)
}

export function DESTROY_TAB_MENU(state) {
    state.currentTab = ''
    state.tabMenus = []
}
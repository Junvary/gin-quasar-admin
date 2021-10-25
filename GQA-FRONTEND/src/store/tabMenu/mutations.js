import { HomeMenu } from 'src/settings'

export function INIT_TAB_MENU(state) {
    // tabMenus数组初始化
    state.tabMenus = HomeMenu
    // 当前页改为首页
    state.currentTab = HomeMenu[0]
}

export function ADD_TAB_MENU(state, tab) {
    state.tabMenus = state.tabMenus.concat([tab])
}

export function CHANGE_CURRENT_TAB(state, tab) {
    state.currentTab = tab
}

export function REMOVE_TAB(state, tab) {
    const removeIndex = state.tabMenus.indexOf(tab)
    // 如果被关闭的就是第一个标签页，那么初始化
    if (removeIndex === 0) {
        state.tabMenus = HomeMenu
        state.currentTab = HomeMenu[0]
    } else {
        state.tabMenus = state.tabMenus.filter(item => item.fullPath !== tab.fullPath)
        state.currentTab = state.tabMenus[removeIndex - 1]
    }

}

export function REMOVE_TAB_MENU(state) {
    state.currentTab = ''
    state.tabMenus = ''
}
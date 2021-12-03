import { getAction } from 'src/api/manage'
import { userMenuUrl } from 'src/api/url'
import { HandleRouter } from 'src/utils/router'
import { HandleAsideMenu, TreeToArray } from 'src/utils/arrayAndTree'

export async function GetUserMenu({ commit, state, dispatch }) {
    const res = await getAction(userMenuUrl)
    if (res.code === 1) {
        const data = res.data.records

        // 拿到鉴权路由表（用户自己的所有菜单），整理称路由
        const userMenu = HandleRouter(data)
        // 设置所有菜单
        commit('INIT_USER_MENU', userMenu)

        // 重组搜索菜单
        const searchMenu = data.filter(value => value.hidden === "no")
        // 设置搜索菜单
        commit('INIT_SEARCH_MENU', searchMenu)

        // 重组顶部菜单分组详情
        const topMenu = []
        HandleAsideMenu(data, "name", "parentCode").forEach(item => {
            if (item.children && item.children.length) {
                topMenu.push({
                    top: item,
                    treeChildren: item.children.filter(item => item.hidden !== 'yes'),
                    arrayChildren: TreeToArray(item.children)
                })
            } else {
                topMenu.push({
                    top: item
                })
            }
        })
        // 设置顶部分组菜单详情
        commit('INIT_TOP_MENU', topMenu)

        // 返回鉴权路由表
        return userMenu
    } else {
        return
    }
}

export async function ClearMenu({ commit }) {
    commit('CLEAR_MENU')
}
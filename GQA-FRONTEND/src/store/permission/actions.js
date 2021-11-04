import { getAction } from 'src/api/manage'
import { userMenuUrl } from 'src/api/url'
import { HandleRouter } from 'src/utils/router'
import { HandleAsideMenu } from 'src/utils/arrayAndTree'

export async function GetUserMenu({ commit, state, dispatch }) {
    const res = await getAction(userMenuUrl)
    if (res.code === 1) {
        const data = res.data.menu
        // 拿到鉴权路由表（用户自己的所有菜单），整理称路由
        const routes = HandleRouter(data)
        // 设置所有菜单
        commit('INIT_USER_MENU', routes)
        // 重组搜索
        const search = HandleAsideMenu(data)
        // 重组侧边栏
        const aside = HandleAsideMenu(data.filter(value => value.hidden === "no"))
        // 设置搜索
        commit('SEARCH_MENU', search)
        // 设置侧边栏菜单
        commit('ASIDE_MENU', aside)
        // 返回鉴权路由表
        return routes
    } else {
        return
    }
}

export async function ClearMenu({ commit }) {
    commit('CLEAR_MENU')
}
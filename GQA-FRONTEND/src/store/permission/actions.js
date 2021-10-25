import { getAction } from 'src/api/manage'
import { userMenuUrl, deptUrl, roleUrl } from 'src/api/url'
import { HandleRouter } from 'src/utils/router'
import { HandleAsideMenu } from 'src/utils/arrayAndTree'

export async function GetUserMenu({ commit, state, dispatch }) {
    const res = await getAction(userMenuUrl)
    const data = res.data.menu
    // 拿到鉴权路由表（用户自己的所有菜单），整理称路由
    const routes = HandleRouter(data)
    // 设置所有菜单
    commit('INIT_USER_MENU', routes)
    // 重组搜索
    const search = HandleAsideMenu(data)
    // 重组侧边栏
    const aside = HandleAsideMenu(data.filter(value => value.hidden === false))
    // 设置搜索
    commit('SEARCH_MENU', search)
    // 设置侧边栏菜单
    commit('ASIDE_MENU', aside)
    // 借助这里获取全部部门列表，TODO-部门列表缓存等内容
    // dispatch('GetAllDeptList')
    // // 借助这里获取全部角色列表，TODO-角色列表缓存等内容
    // dispatch('GetAllRoleList')

    // 返回鉴权路由表
    return routes
}

// export async function GetAllDeptList({ commit }) {
//     const res = await getAction(deptUrl)
//     const deptList = res.data.data
//     commit('SET_ALL_DEPT_LIST', deptList)
// }

// export async function GetAllRoleList({ commit }) {
//     const res = await getAction(roleUrl)
//     const roleList = res.data.data
//     commit('SET_ALL_ROLE_LIST', roleList)
// }

import { defineStore } from 'pinia';
import { postAction } from 'src/api/manage';
import { HandleRouter } from 'src/utils/router';
import { ArrayToTree, HandleAsideMenu, TreeToArray } from 'src/utils/arrayAndTree';

export const usePermissionStore = defineStore('permission', {
    state: () => ({
        userMenu: [],
        searchMenu: [],
        topMenu: []
    }),
    getters: {},
    actions: {
        async GetUserMenu() {
            const res = await postAction('user/get-user-menu')
            if (res.code === 1) {
                const data = res.data.records

                // 拿到鉴权路由表（用户自己的所有菜单），整理称路由
                const userMenu = HandleRouter(data)
                // 加入404界面
                userMenu.push({
                    path: '/:catchAll(.*)*',
                    name: 'notFound',
                    component: () => import('pages/Error404.vue')
                })
                // 设置所有菜单
                this.InitUserMenu(userMenu)

                // 重组搜索菜单
                const searchMenu = data.filter(value => value.hidden === "no")
                // 设置搜索菜单
                this.InitSearchMenu(searchMenu)

                this.InitTopMenu(ArrayToTree(searchMenu, "name", "parent_code"))

                // 重组顶部菜单分组详情
                // const topMenu = []
                // HandleAsideMenu(data, "name", "parent_code").forEach(item => {
                //     if (item.children && item.children.length) {
                //         topMenu.push({
                //             top: item,
                //             treeChildren: item.children.filter(item => item.hidden !== 'yes'),
                //             arrayChildren: TreeToArray(item.children)
                //         })
                //     } else {
                //         topMenu.push({
                //             top: item
                //         })
                //     }
                // })
                // // 设置顶部分组菜单详情
                // this.InitTopMenu(topMenu)

                // 返回鉴权路由表
                return userMenu
            } else {
                return
            }
        },
        InitUserMenu(routes) {
            const menu = []
            const push = function (routes) {
                routes.forEach(route => {
                    if (route.children && route.children.length > 0) {
                        push(route.children)
                    } else {
                        // if (route.meta.hidden === "no") {
                        const { meta, name, path } = route
                        menu.push({ meta, name, path })
                        // }
                    }
                })
            }
            push(routes)
            this.userMenu = menu
        },
        InitSearchMenu(searchMenu) {
            this.searchMenu = searchMenu
        },
        InitTopMenu(topMenu) {
            this.topMenu = topMenu
        },
        ClearMenu() {
            this.userMenu = []
            this.searchMenu = []
            this.topMenu = []
        }
    },
});

import { defineStore } from 'pinia';
import { postAction } from 'src/api/manage';
import { HandleRouter } from 'src/utils/router';
import { ArrayToTree } from 'src/utils/arrayAndTree';

export const usePermissionStore = defineStore('permission', {
    state: () => ({
        userMenu: [],
        searchMenu: [],
        topMenu: [],
        userButton: [],
        defaultPage: [],
    }),
    getters: {},
    actions: {
        async GetUserMenu() {
            try {
                const res = await postAction('user/get-user-menu')
                if (res.code === 1) {
                    const data = res.data.records
                    // role default page
                    const dp = res.data.default_page_list
                    const redirect = data.filter(item => item.name === dp[0] && item.redirect !== '')
                    if (redirect.length) {
                        this.InitUserDefautlPage([data.filter(item => item.path === redirect[0].redirect)[0].name])
                    } else {
                        this.InitUserDefautlPage(dp)
                    }
                    // get user button permisstion
                    this.InitUserButton(res.data.buttons)
                    // Get the authentication route table (all the user's own menus) and handle it into routes
                    const userMenu = HandleRouter(data)
                    // add 404 page
                    userMenu.push({
                        path: '/:catchAll(.*)*',
                        name: 'notFound',
                        component: () => import('pages/Error404.vue')
                    })
                    // set all menus
                    this.InitUserMenu(userMenu)
                    // drop hidden menus
                    const noHiddenMenu = data.filter(value => value.hidden === "yesNo_no")
                    const searchMenu = noHiddenMenu.filter(value => value.parent_code !== "")
                    // set search menus
                    this.InitSearchMenu(searchMenu)
                    // Deep copy to avoid affecting other data
                    const deepTopMenu = JSON.parse(JSON.stringify(noHiddenMenu))
                    const topMenu = ArrayToTree(deepTopMenu, "name", "parent_code")
                    this.InitTopMenu(topMenu)

                    // Return authentication routes
                    return userMenu
                } else {
                    return []
                }
            } catch (error) {
                return error
            }
        },
        InitUserMenu(routes) {
            const menu = []
            const push = function (routes) {
                routes.forEach(route => {
                    if (route.children && route.children.length > 0) {
                        push(route.children)
                    } else {
                        const { meta, name, path } = route
                        menu.push({ meta, name, path })
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
            this.userButton = []
            this.defaultPage = []
        },
        InitUserButton(buttons) {
            this.userButton = buttons
        },
        InitUserDefautlPage(defaultPageList) {
            this.defaultPage = defaultPageList
        }
    },
});

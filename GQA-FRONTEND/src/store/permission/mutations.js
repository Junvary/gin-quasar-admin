export function INIT_USER_MENU(state, routes) {
    const menu = []
    const push = function (routes) {
        routes.forEach(route => {
            if (route.children && route.children.length > 0) {
                push(route.children)
            } else {
                if (!route.hidden) {
                    const { meta, name, path } = route
                    menu.push({ meta, name, path })
                }
            }
        })
    }
    push(routes)
    state.userMenu = menu
}

export function ASIDE_MENU(state, aside) {
    state.aside = aside
}

export function SEARCH_MENU(state, menu) {
    const search = []
    const push = function (menu, titlePrefix = []) {
        menu.forEach(m => {
            if (m.children) {
                push(m.children, [...titlePrefix, m.title])
            } else {
                search.push({
                    ...m,
                    fullTitle: [...titlePrefix, m.title].join(' / ')
                })
            }
        })
    }
    push(menu)
    state.search = search
}

export function SET_ALL_DEPT_LIST(state, deptList) {
    state.allDeptList = deptList
}

export function SET_ALL_ROLE_LIST(state, roleList) {
    state.allRoleList = roleList
}
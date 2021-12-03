export function INIT_USER_MENU(state, routes) {
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
    state.userMenu = menu
}

export function INIT_SEARCH_MENU(state, searchMenu) {
    state.searchMenu = searchMenu
}

export function INIT_TOP_MENU(state, topMenu) {
    state.topMenu = topMenu
}


export function CLEAR_MENU(state) {
    state.userMenu = []
    state.searchMenu = []
    state.topMenu = []
}
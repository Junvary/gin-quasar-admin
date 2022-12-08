import { PrivateRoutes } from 'src/router/routes'

export const HandleRouter = (menuData) => {
    const result = []
    for (let item of menuData) {
        if (item.path !== '') {
            const obj = {
                path: item.path,
                name: item.name,
                component: pageImporter(item.component),
                meta: {
                    hidden: item.hidden,
                    keepAlive: item.keep_alive,
                    title: item.title,
                    icon: item.icon,
                    parent_code: item.parent_code,
                },
                redirect: item.redirect,
            }
            result.push(obj)
        } else {
            if (item.is_link === "yesNo_yes") {
                delete item.path
            }
        }
    }
    // Replace the children that need authentication routing (empty by default) with the collated routes from the backend.
    PrivateRoutes[0].children = [...result]
    // return authentication routes
    return PrivateRoutes
}

const pageImporter = (component) => {
    // Quasar2:
    return () => Promise.resolve(require(`src/${component}.vue`).default)
    // Quasar1:
    // return (resolve) => require([`src/pages/${component}`], resolve)
}
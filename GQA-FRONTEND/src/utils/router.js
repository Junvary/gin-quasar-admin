import { PrivateRoutes } from 'src/router/routes'
const pagesFile = import.meta.glob('../pages/**/*.vue')
const pluginsFile = import.meta.glob('../plugins/**/*.vue')

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
                    keep_alive: item.keep_alive,
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
    // Vite 版本：
    let fileKey = []
    let resultKey = ''
    if (component.split('/')[0] === 'pages') {
        fileKey = Object.keys(pagesFile)
        resultKey = getResultKey(fileKey, component)
        return pagesFile[resultKey[0]]
    } else if (component.split('/')[0] === 'plugins') {
        fileKey = Object.keys(pluginsFile)
        resultKey = getResultKey(fileKey, component)
        return pluginsFile[resultKey[0]]
    } else {
        return Promise.resolve()
    }
    // Quasar2 Webpack版本:
    // return () => Promise.resolve(require(`src/${component}.vue`).default)
    // Quasar1 Webpack版本:
    // return (resolve) => require([`src/pages/${component}`], resolve)
}

const getResultKey = (fileKey, component) => {
    const resultKey = fileKey.filter(key => {
        return key.replace('../', '').replace('.vue', '') === component
    })
    return resultKey
}
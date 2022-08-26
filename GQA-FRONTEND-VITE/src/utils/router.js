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
                    keepAlive: item.keep_alive,
                    title: item.title,
                    icon: item.icon,
                },
                redirect: item.redirect,
            }
            result.push(obj)
        } else {
            if (item.is_link === "yes") {
                delete item.path
            }
        }
    }
    // 将需要鉴权路由的children（默认是空的）替换成后台传过来的整理后的路由表
    PrivateRoutes[0].children = [...result]
    // 返回鉴权路由
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
import { PrivateRoutes } from 'src/router/routes'

export const HandleRouter = (menuData) => {
    const result = []
    for (let item of menuData) {
        if (item.path !== '' && item.component) {
            const obj = {
                path: item.path,
                name: item.name,
                component: pageImporter(item.component),
                meta: {
                    hidden: item.hidden,
                    keepAlive: item.keepAlive,
                    title: item.title,
                    icon: item.icon,
                }
            }
            result.push(obj)
        } else {
            if (item.isLink === "yes") {
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
    // Quasar2 版本：
    return () => Promise.resolve(require(`src/pages${component}.vue`).default)
    // Quasar1 版本：
    // return (resolve) => require([`src/pages/${component}`], resolve)
}
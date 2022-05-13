// 公共路由，无须鉴权
export const PublicRoutes = [
    {
        path: '/login',
        name: 'login',
        component: () => import('layouts/LoginLayout/index.vue'),
        children: []
    },
    {
        path: '/init-db',
        name: 'init-db',
        component: () => import('layouts/LoginLayout/initDb.vue'),
    },
    {
        path: '/new-tab',
        component: () => import('layouts/NewTabLayout/index.vue'),
        children: []
    }

    // // Always leave this as last one,
    // // but you can also remove it

    // 以下内容在动态路由中添加，解决刷新404的问题:store-permission-actions
    // {
    //     path: '/:catchAll(.*)*',
    //     name: 'notFound',
    //     component: () => import('pages/Error404.vue')
    // }
]

export const PrivateRoutes = [
    {
        path: '/',
        redirect: { path: '/dashboard' },
        component: () => import('layouts/MainLayout/index.vue'),
        children: []
    }
]

export default [...PublicRoutes, ...PrivateRoutes]

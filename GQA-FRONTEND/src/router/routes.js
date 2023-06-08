// Public route without authentication
export const PublicRoutes = [
    {
        path: '/login',
        name: 'login',
        component: () => import('layouts/LoginLayout/index.vue'),
        children: []
    },
    {
        path: '/new-tab',
        component: () => import('layouts/NewTabLayout/index.vue'),
        children: [
            // 非鉴权、新tab页面 在下面配置，例如下面的路由（插件内容）:
            // {
            //     path: 'plugin-bp/bp/screen',
            //     component: () => import('src/plugins/Bp/Screen/index.vue')
            // },
            // {
            //     path: 'add-ywd',
            //     component: () => import('src/plugins/YWD/Ywd/modules/addYwd.vue')
            // },
        ]
    }

    // // Always leave this as last one,
    // // but you can also remove it

    // The following content is added to the dynamic routing
    // and is commented out here to solve the problem of refreshing 404:store-permission-actions

    // {
    //     path: '/:catchAll(.*)*',
    //     name: 'notFound',
    //     component: () => import('pages/Error404.vue')
    // }
]

const MainLayoutFile = import.meta.glob('../layouts/MainLayout/index.vue')

export const PrivateRoutes = [
    {
        path: '/',
        redirect: { path: '/dashboard' },
        component: MainLayoutFile['../layouts/MainLayout/index.vue'],
        // component: () => import('layouts/MainLayout/index.vue'),
        children: []
    }
]

export default [...PublicRoutes, ...PrivateRoutes]

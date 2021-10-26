// 公共路由，无须鉴权
const PublicRoutes = [
    {
        path: '/login',
        name: 'login',
        component: () => import('layouts/LoginLayout/index.vue'),
        children: []
    },

    // // Always leave this as last one,
    // // but you can also remove it
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

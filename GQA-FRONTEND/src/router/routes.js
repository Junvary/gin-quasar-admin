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
        children: []
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

export const PrivateRoutes = [
    {
        path: '/',
        redirect: { path: '/dashboard' },
        component: () => import('layouts/MainLayout/index.vue'),
        children: []
    }
]

export default [...PublicRoutes, ...PrivateRoutes]

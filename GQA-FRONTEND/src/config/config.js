// 演示模式开关
export const DemoMode = true
// 演示模式禁止的API
export const ForbiddenUrl = [
    'user/edit-user',
    'user/delete-user-by-id',
    'user/reset-password',
    'user/change-password',
    'role/edit-role',
    'role/delete-role-by-id',
    'role/edit-role-menu',
    'role/edit-role-api',
    'role/remove-role-user',
    'role/edit-role-dept-data-permission',
    'menu/edit-menu',
    'menu/delete-menu-by-id',
    'dept/edit-dept',
    'dept/delete-dept-by-id',
    'dept/remove-dept-user',
    'dict/edit-dict',
    'dict/delete-dict-by-id',
    'api/edit-api',
    'api/delete-api-by-id',
    'config-backend/edit-config-backend',
    'config-backend/delete-config-backend-by-id',
    'config-frontend/edit-config-frontend',
    'config-frontend/delete-config-frontend-by-id',
    'log/delete-log-login-by-id',
    'log/delete-log-operation-by-id',
    'notice/delete-notice-by-id',
    'todo/edit-todo',
    'todo/delete-todo-by-id',
    'user-online/kick-online-user',
    'cron/start-cron',
    'cron/stop-cron'
]
// 路由白名单
export const AllowList = [
    '/login',
    '/new-tab/add-ywd'
]
// 控制台打印内容
export const GqaConsoleLogo = () => {
    console.info('Welcome to Gin-Quasar-Admin!')
    console.info('Github: https://github.com/Junvary/gin-quasar-admin ')
    console.info('Expecting Your Star!')
}

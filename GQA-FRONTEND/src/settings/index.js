// 首页等允许无token的白名单
export const Allowlist = ['/login']
    // 没有用户名的时候使用这个名字
export const GqaUsername = function() { return this.$t('gqaUserName') }
    // 没有头像配置的时候使用这个头像
export const GqaLoginAvatar = "gqa128.png"
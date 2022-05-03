// 首页等允许无token的白名单
export const Allowlist = ['/login']

// 没有用户名的时候使用这个名字
export const GqaDefaultUsername = function () { return this.$t('GqaDefaultUsername') }
// 没有头像配置的时候使用这个头像
export const GqaDefaultAvatar = "gqa128.png"

// 没有网站前台配置的时候用这个配置
export const GqaFrontendDefault = {
    mainTitle: "Gin-Quasar-Admin",
    subTitle: "Gin-Quasar-Admin",
    webDescribe: "Be the change you want to see in the world.",
    showGit: "yes"
}
// 没有网站后台配置的时候用这个配置
export const GqaBackendDefault = {}


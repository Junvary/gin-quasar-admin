import { useI18n } from 'vue-i18n';

export default function useConfig() {
    const { t } = useI18n
    // 浏览器console打印内容
    const GqaConsoleLogo = () => {
        console.info('欢迎使用Gin-Quasar-Admin!')
        console.info('项目地址: https://github.com/Junvary/gin-quasar-admin ')
        console.info('欢迎交流, 感谢Star!')
    }
    // 首页等允许无token的白名单
    const AllowList = ['/login',]
    // 没有用户名的时候使用这个名字
    const GqaDefaultUsername = () => t('GqaDefaultUsername')
    // 没有头像配置的时候使用这个头像
    const GqaDefaultAvatar = "gqa128.png"
    // 没有网站前台配置的时候用这个配置
    const GqaFrontendDefault = {
        mainTitle: "Gin-Quasar-Admin",
        subTitle: "Gin-Quasar-Admin",
        webDescribe: "Be the change you want to see in the world.",
        showGit: "yesNo_yes"
    }
    // 没有网站后台配置的时候用这个配置
    const GqaBackendDefault = {}
    // Quasar主题配色
    const ThemeStyleQuasar = {
        primary: '#1976D2',
        secondary: '#26A69A',
        accent: '#9C27B0',
        positive: '#21BA45',
        negative: '#C10015',
        info: '#31CCEC',
        warning: '#F2C037',
        light: '#FFFFFF',
        dark: '#1D1D1D',
    }
    // Element主题配色
    const ThemeStyleElement = {
        primary: '#409EFF',
        secondary: '#26A69A',
        accent: '#9C27B0',
        positive: '#67C23A',
        negative: '#F56C6C',
        info: '#8896b3',
        warning: '#e6a23c',
        light: '#FFFFFF',
        dark: '#1D1D1D',
    }
    // Ant Design主题配色
    const ThemeStyleAnt = {
        primary: '#1890ff',
        secondary: '#26A69A',
        accent: '#9C27B0',
        positive: '#52c41a',
        negative: '#f5222d',
        info: '#fafafa',
        warning: '#faad14',
        light: '#FFFFFF',
        dark: '#1D1D1D',
    }
    return {
        GqaConsoleLogo,
        AllowList,
        GqaDefaultUsername,
        GqaDefaultAvatar,
        GqaFrontendDefault,
        GqaBackendDefault,
        ThemeStyleQuasar,
        ThemeStyleElement,
        ThemeStyleAnt,
    }
}
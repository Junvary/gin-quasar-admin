export default function useConfig() {
    // console logo
    const GqaConsoleLogo = () => {
        console.info('Welcome to Gin-Quasar-Admin!')
        console.info('Github: https://github.com/Junvary/gin-quasar-admin ')
        console.info('Expecting Your Star!')
    }
    // allow list without token
    const AllowList = ['/login',]
    // use this name if there is no username
    const GqaDefaultUsername = 'GQA'
    // use this avatar if there is no avatar
    const GqaDefaultAvatar = "gqa128.png"
    // use this frontend config if there is no frontend config
    const GqaFrontendDefault = {
        mainTitle: "Gin-Quasar-Admin",
        subTitle: "Gin-Quasar-Admin",
        webDescribe: "Be the change you want to see in the world.",
        showGit: "yesNo_yes"
    }
    // use this backend config if there is no backend config
    const GqaBackendDefault = {}
    // Quasar Theme Color
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
    // Element Theme Color
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
    // Ant Design Theme Color
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
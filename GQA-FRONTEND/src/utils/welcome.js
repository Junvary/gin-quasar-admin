import { i18n } from 'src/boot/i18n'

export const timeWelcome = () => {
    const time = new Date()
    const hour = time.getHours()
    return hour <= 11
        ? i18n.global.t('GoodMorning') :
        (hour <= 13 ? i18n.global.t('GoodNoon') : (hour <= 17 ? i18n.global.t('GoodAfternoon') : i18n.global.t('GoodEvening')))
}

export const randomWelcome = () => {
    const w = [
        i18n.global.t('WelcomeTo') + ' Gin-Quasar-Admin!',
        i18n.global.t('GoOutForAWalk'),
        i18n.global.t('WantSomeFeatures'),
        i18n.global.t('NeedSomeCoffee'),
        i18n.global.t('CleanTheWindows'),
        i18n.global.t('StarOnGithub'),
    ]
    return w[Math.floor(Math.random() * w.length)]
}
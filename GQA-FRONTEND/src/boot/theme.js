import { boot } from 'quasar/wrappers'
import { Cookies, setCssVar, getCssVar } from 'quasar'


export default boot(({ app }) => {
    const primary = Cookies.get('gqa-theme-primary') || getCssVar('primary')
    const secondary = Cookies.get('gqa-theme-secondary') || getCssVar('secondary')
    const accent = Cookies.get('gqa-theme-accent') || getCssVar('accent')
    const dark = Cookies.get('gqa-theme-dark') || getCssVar('dark')
    const positive = Cookies.get('gqa-theme-positive') || getCssVar('positive')
    const negative = Cookies.get('gqa-theme-negative') || getCssVar('negative')
    const info = Cookies.get('gqa-theme-info') || getCssVar('info')
    const warning = Cookies.get('gqa-theme-warning') || getCssVar('warning')
    setCssVar("primary", primary)
    setCssVar("secondary", secondary)
    setCssVar("accent", accent)
    setCssVar("dark", dark)
    setCssVar("positive", positive)
    setCssVar("negative", negative)
    setCssVar("info", info)
    setCssVar("warning", warning)
})

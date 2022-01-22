import { boot } from 'quasar/wrappers'
import { Cookies, setCssVar, getCssVar } from 'quasar'

export default boot(({ app }) => {
    const primary = Cookies.get('gqa-theme-primary') || getCssVar('primary') || '#1976D2'
    const secondary = Cookies.get('gqa-theme-secondary') || getCssVar('secondary') || '#26A69A'
    const accent = Cookies.get('gqa-theme-accent') || getCssVar('accent') || '#9C27B0'
    const positive = Cookies.get('gqa-theme-positive') || getCssVar('positive') || '#21BA45'
    const negative = Cookies.get('gqa-theme-negative') || getCssVar('negative') || '#C10015'
    const info = Cookies.get('gqa-theme-info') || getCssVar('info') || '#31CCEC'
    const warning = Cookies.get('gqa-theme-warning') || getCssVar('warning') || '#F2C037'
    const light = Cookies.get('gqa-theme-light') || getCssVar('light') || '#FFFFFF'
    const dark = Cookies.get('gqa-theme-dark') || getCssVar('dark') || '#1D1D1D'

    setCssVar("primary", primary)
    setCssVar("secondary", secondary)
    setCssVar("accent", accent)
    setCssVar("positive", positive)
    setCssVar("negative", negative)
    setCssVar("info", info)
    setCssVar("warning", warning)
    setCssVar("light", light)
    setCssVar("dark", dark)
})

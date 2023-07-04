import { boot } from 'quasar/wrappers'
import { Cookies, setCssVar, getCssVar } from 'quasar'
import { ThemeStyleQuasar } from 'src/config/default'
import 'uno.css'

export default boot(({ app }) => {
    const primary = Cookies.get('gqa-theme-primary') || getCssVar('primary') || ThemeStyleQuasar.primary
    const secondary = Cookies.get('gqa-theme-secondary') || getCssVar('secondary') || ThemeStyleQuasar.secondary
    const accent = Cookies.get('gqa-theme-accent') || getCssVar('accent') || ThemeStyleQuasar.accent
    const positive = Cookies.get('gqa-theme-positive') || getCssVar('positive') || ThemeStyleQuasar.positive
    const negative = Cookies.get('gqa-theme-negative') || getCssVar('negative') || ThemeStyleQuasar.negative
    const info = Cookies.get('gqa-theme-info') || getCssVar('info') || ThemeStyleQuasar.info
    const warning = Cookies.get('gqa-theme-warning') || getCssVar('warning') || ThemeStyleQuasar.warning
    const light = Cookies.get('gqa-theme-light') || getCssVar('light') || ThemeStyleQuasar.light
    const dark = Cookies.get('gqa-theme-dark') || getCssVar('dark') || ThemeStyleQuasar.dark

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

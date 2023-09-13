import { defineStore } from 'pinia';
import { Cookies, LocalStorage } from 'quasar';

export const useSettingStore = defineStore('setting', {
    state: () => ({
        language: undefined,
        sidebarWidth: undefined,
        darkTheme: undefined,
        sidebarDense: undefined,
        themeStyle: ''
    }),
    getters: {},
    actions: {
        SetThemeStyle(style) {
            this.themeStyle = style
            LocalStorage.set('gqa-theme-style', style)
        },
        GetThemeStyle() {
            if (this.themeStyle) {
                return this.themeStyle
            } else if (LocalStorage.getItem('gqa-theme-style')) {
                return LocalStorage.getItem('gqa-theme-style')
            } else {
                return 'Gin-Quasar-Admin'
            }
        },
        SetDarkTheme(val) {
            this.darkTheme = val
            LocalStorage.set('gqa-dark-theme', val)
        },
        GetDarkTheme() {
            if (this.darkTheme) {
                return this.darkTheme
            } else if (LocalStorage.getItem('gqa-dark-theme')) {
                return LocalStorage.getItem('gqa-dark-theme')
            } else {
                return false
            }
        },
        SetSidebarDense(val) {
            this.sidebarDense = val
            LocalStorage.set('gqa-sidebar-dense', val)
        },
        GetSidebarDense() {
            if (this.sidebarDense) {
                return this.sidebarDense
            } else if (LocalStorage.getItem('gqa-sidebar-dense')) {
                return LocalStorage.getItem('gqa-sidebar-dense')
            } else {
                return false
            }
        },
        ChangeLanguage(lang) {
            this.language = lang
            Cookies.set('gqa-language', lang)
        },
        GetLanguage() {
            if (this.language) {
                return this.language
            } else if (Cookies.get('gqa-language')) {
                return Cookies.get('gqa-language')
            } else {
                return 'zh-CN'
            }
        },
        SetSidebarWidth(width) {
            this.sidebarWidth = width
            LocalStorage.set('gqa-sidebar-width', width)
        },
        GetSidebarWidth() {
            if (this.sidebarWidth) {
                return this.sidebarWidth
            } else if (LocalStorage.getItem('gqa-sidebar-width')) {
                return LocalStorage.getItem('gqa-sidebar-width')
            } else {
                return 200
            }
        },
    }
})
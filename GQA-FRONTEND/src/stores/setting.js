import { defineStore } from 'pinia';
import { Cookies, LocalStorage } from 'quasar';

export const useSettingStore = defineStore('setting', {
    state: () => ({
        language: undefined,
        sideDrawerWidth: undefined,
        darkTheme: undefined,
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
        SetSideDrawerWidth(width) {
            this.sideDrawerWidth = width
            LocalStorage.set('gqa-side-drawer-width', width)
        },
        GetSideDrawerWidth() {
            if (this.sideDrawerWidth) {
                return this.sideDrawerWidth
            } else if (LocalStorage.getItem('gqa-side-drawer-width')) {
                return LocalStorage.getItem('gqa-side-drawer-width')
            } else {
                return 200
            }
        },
    }
})
import { defineStore } from 'pinia';
import { Cookies, LocalStorage } from 'quasar';

export const useSettingStore = defineStore('setting', {
    state: () => ({
        language: undefined,
        sideDrawerWidth: 220,
        darkTheme: false,
        sideDrawerTop: 'h',
        sideDrawerBottom: 'l'
    }),
    getters: {},
    actions: {
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
                return 220
            }
        },
        SetSideDrawerTop(val) {
            this.sideDrawerTop = val
            LocalStorage.set('gqa-side-drawer-top', val)
        },
        GetSideDrawerTop() {
            if (this.sideDrawerTop) {
                return this.sideDrawerTop
            } else if (LocalStorage.getItem('gqa-side-drawer-top')) {
                return LocalStorage.getItem('gqa-side-drawer-top')
            } else {
                return 'h'
            }
        },
        SetSideDrawerBottom(val) {
            this.sideDrawerBottom = val
            LocalStorage.set('gqa-side-drawer-bottom', val)
        },
        GetSideDrawerBottom() {
            if (this.sideDrawerBottom) {
                return this.sideDrawerBottom
            } else if (LocalStorage.getItem('gqa-side-drawer-bottom')) {
                return LocalStorage.getItem('gqa-side-drawer-bottom')
            } else {
                return 'l'
            }
        },
        GetLayoutView() {
            const top = this.GetSideDrawerTop() + "Hh"
            const center = 'LpR'
            const bottom = this.GetSideDrawerBottom() + 'Fr'
            return top + ' ' + center + ' ' + bottom
        }
    }
})
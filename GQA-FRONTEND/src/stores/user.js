import { defineStore } from 'pinia';
import { Cookies, SessionStorage } from 'quasar';
import { usePermissionStore } from './permission';
import { postAction } from 'src/api/manage';

export const useUserStore = defineStore('user', {
    state: () => ({
        token: undefined,
        username: undefined,
        nickname: undefined,
        realName: undefined,
        avatar: undefined,
        rememberMe: true,
    }),
    getters: {},
    actions: {
        async HandleLogin(loginForm) {
            const res = await postAction('public/login', loginForm)
            if (res.code === 1) {
                const token = res.data.token
                const username = res.data.username
                const nickname = res.data.nickname
                const realName = res.data.real_name
                const avatar = res.data.avatar
                this.SetToken(token)
                this.username = username
                Cookies.set('gqa-username', username)
                this.nickname = nickname
                Cookies.set('gqa-nickname', nickname)
                this.realName = realName
                Cookies.set('gqa-realName', realName)
                this.avatar = avatar
                Cookies.set('gqa-avatar', avatar)
                return true
            } else {
                return
            }
        },
        SetToken(token) {
            this.token = token
            if (this.rememberMe) {
                Cookies.set('gqa-token', token)
            } else {
                SessionStorage.set('gqa-token', token)
            }
        },
        ChangeRememberMe(type) {
            this.rememberMe = type
        },
        HandleLogout() {
            const permissionStore = usePermissionStore()
            permissionStore.ClearMenu()
            SessionStorage.remove('gqa-token')
            Cookies.remove('gqa-token')
            Cookies.remove('gqa-username')
            Cookies.remove('gqa-nickname')
            Cookies.remove('gqa-realName')
            Cookies.remove('gqa-avatar')
            // dont delete dict
            // LocalStorage.remove('gqa-dict')
            this.token = undefined
            this.username = undefined
            this.nickname = undefined
            this.realName = undefined
            this.avatar = undefined
        },
        GetToken() {
            if (SessionStorage.getItem('gqa-token')) {
                return SessionStorage.getItem('gqa-token')
            } else if (Cookies.get('gqa-token')) {
                return Cookies.get('gqa-token')
            } else {
                return this.token
            }
        },
        GetUsername() {
            if (this.username) {
                return this.username
            } else if (Cookies.get('gqa-username')) {
                return Cookies.get('gqa-username')
            } else {
                return ""
            }
        },
        GetNickname() {
            if (this.nickname) {
                return this.nickname
            } else if (Cookies.get('gqa-nickname')) {
                return Cookies.get('gqa-nickname')
            } else {
                return ""
            }
        },
        GetRealName() {
            if (this.realName) {
                return this.realName
            } else if (Cookies.get('gqa-realName')) {
                return Cookies.get('gqa-realName')
            } else {
                return ""
            }
        },
        GetAvatar() {
            if (this.avatar) {
                return this.avatar
            } else if (Cookies.get('gqa-avatar')) {
                return Cookies.get('gqa-avatar')
            } else {
                return ""
            }
        },
    },
});

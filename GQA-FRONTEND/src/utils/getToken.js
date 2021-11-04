import { Cookies, SessionStorage } from 'quasar'

export const GetToken = () => {
    if (SessionStorage.getItem('gqa-token')) {
        return SessionStorage.getItem('gqa-token')
    } else if (Cookies.get('gqa-token')) {
        return Cookies.get('gqa-token')
    } else {
        return
    }
}

export const RemoveToken = () => {
    SessionStorage.remove('gqa-token')
    Cookies.remove('gqa-token')
}
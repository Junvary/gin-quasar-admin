import { date } from 'quasar'

export const FormatDateTime = (datetime) => {
    return date.formatDate(datetime, "YYYY-MM-DD HH:mm:ss")
}

export const FormatDate = (datetime) => {
    return date.formatDate(datetime, "YYYY-MM-DD")
}

export const FormatDateTimeShort = (datetime) => {
    return date.formatDate(datetime, "YYYYMMDDHHmmss")
}
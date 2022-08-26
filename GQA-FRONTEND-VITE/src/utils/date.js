import { date } from 'quasar'

export const FormatDateTime = (datetime) => {
    return date.formatDate(datetime, "YYYY-MM-DD HH:mm:ss")
}
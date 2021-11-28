import { date } from 'quasar'

export const FormatDataTime = (datetime) => {
    return date.formatDate(datetime, "YYYY-MM-DD HH:mm:ss")
}
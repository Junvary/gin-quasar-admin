export const timeWelcome = () => {
    const time = new Date()
    const hour = time.getHours()
    return hour < 5 ? '凌晨好,'
        : (hour <= 7 ? '早上好,' : (hour <= 11 ? '上午好,' : (hour <= 13 ? '中午好,' : (hour <= 17 ? '下午好,' : (hour <= 22 ? '晚上好,' : '夜深了,')))))
}

export const randomWelcome = () => {
    const w = [
        '欢迎来到Gin-Quasar-Admin!',
        '要不要出去走走?',
        '也许你想要一些更多的功能',
        '我想我需要喝点咖啡',
        '窗户该擦一擦了',
        '喂,不要太卷了',
        '可不可以去Github帮忙点个star呢?',
        '哦,你也来一趟?',
        '巴拉巴拉,你好啊',
    ]
    return w[Math.floor(Math.random() * w.length)]
}
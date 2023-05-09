import languages from "./languages"

export const Translator = () => {
    const temp = {
        'zh-CN': {},
        'en-US': {},
        'ru': {}
    }
    for (let key in languages) {
        temp['zh-CN'][key] = languages[key]['zh-CN']
        temp['en-US'][key] = languages[key]['en-US']
        temp['ru'][key] = languages[key]['ru']
    }
    return temp
}

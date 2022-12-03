import languages from "./languages"

const Translator = (lang) => {
    const temp = {}
    for (let key in languages) {
        temp[key] = languages[key][lang]
    }
    return temp
}

export default {
    'zh-CN': Translator('zh-CN'),
    'en-US': Translator('en-US'),
    'ru': Translator('ru'),
}

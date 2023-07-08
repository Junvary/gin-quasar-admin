import { markRaw, defineAsyncComponent } from 'vue'
import { UpperFirst } from './convert'

const getPluginKey = (fileKey, pluginCode, cutNum) => {
    const resultKey = fileKey.filter(key => {
        return key.split('/')[cutNum] === pluginCode
    })
    return resultKey
}

export const CheckComponent = (pluginCurrent, pluginsFile, cutNum) => {
    let pluginCode, pluginComponent
    try {
        if (pluginCurrent.slice(0, 7) === 'plugin-') {
            pluginCode = UpperFirst(pluginCurrent.slice(7))
        } else {
            pluginCode = UpperFirst(pluginCurrent)
        }
        const pluginCodeComponent = getPluginKey(Object.keys(pluginsFile), pluginCode, cutNum)
        pluginComponent = markRaw(defineAsyncComponent(pluginsFile[pluginCodeComponent[0]]))
        return [pluginCode, pluginComponent, null]
    } catch (err) {
        console.log('引入组件失败', err)
        return [null, null, err]
    }
}
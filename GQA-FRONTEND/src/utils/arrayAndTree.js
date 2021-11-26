import XEUtils from 'xe-utils'
import { uniqueId } from 'lodash'

export const HandleAsideMenu = function (menuData, key, parentKey) {
    // 将列表数据转换为树形数据
    // 处理菜单成树，key值为name，parentKey为parentCode
    const menu = ArrayToTree(menuData, key, parentKey)
    return checkPathAndChildren(menu)
}

function checkPathAndChildren(menu) {
    return menu.map(m => ({
        ...m, path: m.path || uniqueId('gqa-null-path-'), ...m.children
            ? { children: checkPathAndChildren(m.children) }
            : {}
    }))
}

export const ArrayToTree = (arrayData, key, parentKey) => {
    const data = XEUtils.toArrayTree(arrayData, {
        key: key,
        parentKey: parentKey,
        strict: true
    })
    return data
}

export const TreeToArray = (treeData) => {
    const data = XEUtils.toTreeArray(treeData)
    return data
}
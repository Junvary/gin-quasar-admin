import XEUtils from 'xe-utils'
import { uniqueId } from 'lodash'

export const HandleAsideMenu = function (menuData) {
    // 将列表数据转换为树形数据
    const data = XEUtils.toArrayTree(menuData, {
        parentKey: 'parentId',
        strict: true
    })
    const menu = [
        ...data
    ]
    return checkPathAndChildren(menu)
}

function checkPathAndChildren(menu) {
    return menu.map(e => ({ ...e, path: e.path || uniqueId('gqa-null-path-'), ...e.children ? { children: checkPathAndChildren(e.children) } : {} }))
}

export const ArrayToTree = (arrayData) => {
    const data = XEUtils.toArrayTree(arrayData, {
        parentKey: 'parentId',
        strict: true
    })
    return data
}

export const TreeToArray = (treeData) => {
    const data = XEUtils.toTreeArray(treeData)
    return data
}
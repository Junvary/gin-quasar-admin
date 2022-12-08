import XEUtils from 'xe-utils'
import { uniqueId } from 'lodash'

export const HandleAsideMenu = function (menuData, key, parentKey) {
    // change list to tree
    // menu to tree: key => name，parentKey => parentCode
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
    const treeData = XEUtils.toArrayTree(arrayData, {
        key: key,
        parentKey: parentKey,
        // strict: true
    })
    return treeData
}

export const TreeToArray = (treeData) => {
    const data = XEUtils.toTreeArray(treeData)
    return data
}

export const ChangeNullChildren2Array = (data) => {
    const index = data?.length
    for (let i = index - 1; i >= 0; i--) {
        if (data[i].children === null) {
            data[i].children = []
        } else {
            ChangeNullChildren2Array(data[i].children)
        }
    }
    return data
}
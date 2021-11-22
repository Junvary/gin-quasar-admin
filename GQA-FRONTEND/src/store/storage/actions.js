import { getAction } from "src/api/manage"
import { pubDictUrl, pubFrontendUrl } from 'src/api/url'
import { ArrayToTree } from 'src/utils/arrayAndTree'


export async function GetGqaDict({ commit, state, dispatch }) {
    const res = await getAction(pubDictUrl)
    if (res.code === 1) {
        const dictDetail = res.data.records
        for (let i of dictDetail) {
            i.value = i.dictCode
            i.label = i.dictLabel
        }
        const dictList = ArrayToTree(dictDetail, "dictCode", "parentCode")
        let dict = {}
        for (let d of dictList) {
            dict[d.dictCode] = d.children
        }
        commit("SET_GQA_DICT", dict)
    }
}

export async function GetGqaFrontend({ commit, state, dispatch }) {
    const res = await getAction(pubFrontendUrl)
    if (res.code === 1) {
        const frontend = {}
        res.data.records.forEach(item => {
            frontend[item.gqaOption] = item.custom ? item.custom : item.default
        })
        commit("SET_GQA_FRONTEND", frontend)
    }
}

export async function SetGqaFrontend({ commit, state, dispatch }, records) {
    const frontend = {}
    records.forEach(item => {
        frontend[item.gqaOption] = item.custom ? item.custom : item.default
    })
    commit("SET_GQA_FRONTEND", frontend)
}

export function SetGqaGoVersion({ commit, state, dispatch }, goVersion) {
    commit("SET_GQA_GO_VERSION", goVersion)
}

export function SetGqaGinVersion({ commit, state, dispatch }, ginVersion) {
    commit("SET_GQA_GIN_VERSION", ginVersion)
}

export function SetGqaPluginList({ commit, state, dispatch }, pluginList) {
    commit("SET_GQA_PLUGIN_LIST", pluginList)
}
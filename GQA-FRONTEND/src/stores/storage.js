import { defineStore } from 'pinia';
import { LocalStorage } from 'quasar';
import { postAction } from 'src/api/manage';
import { ArrayToTree } from 'src/utils/arrayAndTree'
import { GqaFrontendDefault, GqaBackendDefault } from "src/config/default"

export const useStorageStore = defineStore('storage', {
    state: () => ({
        gqaDict: undefined,
        gqaBackend: undefined,
        gqaFrontend: undefined,
        goVersion: undefined,
        ginVersion: undefined,
        pluginList: undefined
    }),
    getters: {},
    actions: {
        async SetGqaDict() {
            const res = await postAction('public/get-dict-all')
            if (res.code === 1) {
                const dictDetail = res.data.records
                for (let i of dictDetail) {
                    i.value = i.dict_code
                    i.label = i.dict_label
                }
                const dictList = ArrayToTree(dictDetail, "dict_code", "parent_code")
                let dict = {}
                for (let d of dictList) {
                    dict[d.dict_code] = d.children
                }
                this.gqaDict = dict
                LocalStorage.set('gqa-dict', dict)
            }
        },

        async SetGqaBackend() {
            const res = await postAction('public/get-config-backend-all')
            if (res.code === 1) {
                const backend = {}
                res.data.records.forEach(item => {
                    backend[item.config_item] = item.item_custom ? item.item_custom : item.item_default
                })
                this.gqaBackend = backend
                LocalStorage.set('gqa-backend', backend)
            }
        },

        async SetGqaFrontend() {
            const res = await postAction('public/get-config-frontend-all')
            if (res.code === 1) {
                const frontend = {}
                res.data.records.forEach(item => {
                    frontend[item.config_item] = item.item_custom ? item.item_custom : item.item_default
                })
                this.gqaFrontend = frontend
                LocalStorage.set('gqa-frontend', frontend)
            }
        },

        SetGqaGoVersion(goVersion) {
            this.goVersion = goVersion
            LocalStorage.set('gqa-goVersion', goVersion)
        },

        SetGqaGinVersion(ginVersion) {
            this.ginVersion = ginVersion
            LocalStorage.set('gqa-ginVersion', ginVersion)
        },

        SetGqaPluginList(pluginList) {
            this.pluginList = pluginList
            LocalStorage.set('gqa-pluginList', pluginList)
        },
        GetGqaDict() {
            const dict = LocalStorage.getItem("gqa-dict")
            if (this.gqaDict) {
                return this.gqaDict
            } else {
                return dict
            }
        },
        GetGqaBackend() {
            const backend = LocalStorage.getItem("gqa-backend")
            if (this.gqaBackend) {
                return this.gqaBackend
            } else if (backend) {
                return backend
            } else {
                return GqaBackendDefault
            }
        },
        GetGqaFrontend() {
            const frontend = LocalStorage.getItem("gqa-frontend")
            if (this.gqaFrontend) {
                return this.gqaFrontend
            } else if (frontend) {
                return frontend
            } else {
                return GqaFrontendDefault
            }
        },
        GetGqaGoVersion() {
            const goversion = LocalStorage.getItem("gqa-goVersion")
            if (this.goVersion) {
                return this.goVersion
            } else {
                return goversion
            }
        },
        GetGqaGinVersion() {
            const ginversion = LocalStorage.getItem("gqa-ginVersion")
            if (this.ginVersion) {
                return this.ginVersion
            } else {
                return ginversion
            }
        },
        GetGqaPluginList() {
            const pluginList = LocalStorage.getItem("gqa-pluginList")
            if (this.pluginList) {
                return this.pluginList
            } else {
                return pluginList
            }
        },
    },
});

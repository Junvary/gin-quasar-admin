import { defineStore } from 'pinia';
import { usePermissionStore } from './permission';
const permissionStore = usePermissionStore()
const base = permissionStore.userMenu.filter(item => item.name === 'dashboard')[0]


export const useTabMenuStore = defineStore('tabMenu', {
    state: () => ({
        tabMenus: [],
        currentTab: {},
    }),
    getters: {},
    actions: {
        AddTabMenu(tab) {

            // 退出时userMenu被清空，不走下面逻辑
            if (base) {
                // 如果没有仪表盘，则加入仪表盘
                if (this.tabMenus.filter(item => item.path === base.path).length === 0) {
                    this.tabMenus = this.tabMenus.concat([base])
                    this.currentTab = base
                }
                // 判断tab是否存在，是为了关闭所有菜单时，不会传递tab
                if (tab && !this.tabMenus.some(item => item.path === tab.path)) {
                    this.tabMenus = this.tabMenus.concat([tab])
                    this.currentTab = tab
                }
            }
        },
        ChangeCurrentTab(tab) {
            this.currentTab = tab
        },
        RemoveTab(tab) {
            const removeIndex = this.tabMenus.indexOf(tab)
            this.tabMenus = this.tabMenus.filter(item => item.path !== tab.path)
            this.currentTab = this.tabMenus[removeIndex - 1]
        },
        RemoveRightTab(tab) {
            const removeIndex = this.tabMenus.indexOf(tab)
            this.tabMenus = this.tabMenus.slice(0, removeIndex + 1)
        },
        RemoveLeftTab(tab) {
            const removeIndex = this.tabMenus.indexOf(tab)
            const rightMenu = this.tabMenus.slice(removeIndex)
            this.tabMenus = [base].concat(rightMenu)
        },
        DestroyTabMenu() {
            this.currentTab = ''
            this.tabMenus = []
        }
    },
});

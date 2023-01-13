import { defineStore } from 'pinia';
import { usePermissionStore } from './permission';
const permissionStore = usePermissionStore()

export const useTabMenuStore = defineStore('tabMenu', {
    state: () => ({
        tabMenus: [],
        currentTab: {},
    }),
    getters: {
        base() {
            return permissionStore.userMenu.filter(item => item.name === permissionStore.defaultPage[0])[0]
        }
    },
    actions: {
        AddTabMenu(tab) {
            // When exiting, the userMenu is cleared
            // pass it
            if (this.base) {
                // If there is no default page, add it
                if (this.tabMenus.filter(item => item.path === this.base.path).length === 0) {
                    this.tabMenus = this.tabMenus.concat([this.base])
                    this.currentTab = this.base
                }
                // To determine whether a tab exists, the tab will not be passed when all menus are closed
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
            this.tabMenus = [this.base].concat(rightMenu)
        },
        DestroyTabMenu() {
            this.currentTab = ''
            this.tabMenus = []
        }
    },
});

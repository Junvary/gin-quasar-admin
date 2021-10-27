package system

import (
	"gin-quasar-admin/api"
	"github.com/gin-gonic/gin"
)

type RouterMenu struct{}

func (r *RouterMenu) InitRouterMenu(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuGroup := Router.Group("menu")
	apiMenu := api.GroupApiApp.ApiSystem.ApiMenu
	{
		// 获取菜单列表
		menuGroup.POST("menu-list", apiMenu.GetMenuList)
		// 编辑菜单信息
		menuGroup.PUT("menu-edit", apiMenu.EditMenu)
		// 新增菜单
		menuGroup.POST("menu-add", apiMenu.AddMenu)
		// 删除菜单
		menuGroup.DELETE("menu-delete", apiMenu.DeleteMenu)
		// 根据ID查找菜单
		menuGroup.POST("menu-id", apiMenu.QueryMenuById)
	}
	return Router
}

package system

import (
	"gin-quasar-admin/middleware"
	"github.com/gin-gonic/gin"
)

type RouterMenu struct{}

func (r *RouterMenu) InitRouterMenu(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuGroup := Router.Group("menu").Use(middleware.LogOperationHandler())
	{
		//获取菜单列表
		menuGroup.POST("menu-list", ApiMenu.GetMenuList)
		//编辑菜单信息
		menuGroup.PUT("menu-edit", ApiMenu.EditMenu)
		//新增菜单
		menuGroup.POST("menu-add", ApiMenu.AddMenu)
		//删除菜单
		menuGroup.DELETE("menu-delete", ApiMenu.DeleteMenu)
		//根据ID查找菜单
		menuGroup.POST("menu-id", ApiMenu.QueryMenuById)
	}
	return Router
}

package private

import "github.com/gin-gonic/gin"

type RouterMenu struct{}

func (r *RouterMenu) InitRouterRouterMenu(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("menu")
	{
		privateGroup.POST("get-menu-list", apiPrivate.ApiMenu.GetMenuList)
		privateGroup.POST("edit-menu", apiPrivate.ApiMenu.EditMenu)
		privateGroup.POST("add-menu", apiPrivate.ApiMenu.AddMenu)
		privateGroup.POST("delete-menu-by-id", apiPrivate.ApiMenu.DeleteMenuById)
		privateGroup.POST("query-menu-by-id", apiPrivate.ApiMenu.QueryMenuById)
	}
	return privateGroup
}

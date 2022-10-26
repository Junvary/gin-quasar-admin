package private

import "github.com/gin-gonic/gin"

type RouterRole struct{}

func (r *RouterRole) InitRouterRouterRole(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("role")
	{
		privateGroup.POST("get-role-list", apiPrivate.ApiRole.GetRoleList)
		privateGroup.POST("edit-role", apiPrivate.ApiRole.EditRole)
		privateGroup.POST("add-role", apiPrivate.ApiRole.AddRole)
		privateGroup.POST("delete-role-by-id", apiPrivate.ApiRole.DeleteRoleById)
		privateGroup.POST("query-role-by-id", apiPrivate.ApiRole.QueryRoleById)

		privateGroup.POST("query-user-by-role", apiPrivate.ApiRole.QueryUserByRole)
		privateGroup.POST("remove-role-user", apiPrivate.ApiRole.RemoveRoleUser)
		privateGroup.POST("add-role-user", apiPrivate.ApiRole.AddRoleUser)

		privateGroup.POST("get-role-menu-list", apiPrivate.ApiRole.GetRoleMenuList)
		privateGroup.POST("edit-role-menu", apiPrivate.ApiRole.EditRoleMenu)

		privateGroup.POST("get-role-api-list", apiPrivate.ApiRole.GetRoleApiList)
		privateGroup.POST("edit-role-api", apiPrivate.ApiRole.EditRoleApi)

		privateGroup.POST("edit-role-dept-data-permission", apiPrivate.ApiRole.EditRoleDeptDataPermission)

		privateGroup.POST("get-role-button-list", apiPrivate.ApiRole.GetRoleButtonList)
	}
	return privateGroup
}

package system

import (
	"github.com/gin-gonic/gin"
)

type RouterRole struct {}

func (r *RouterRole) InitRouterRole(Router *gin.RouterGroup) (R gin.IRoutes) {
	roleGroup := Router.Group("role")
	{
		//获取角色列表
		roleGroup.POST("role-list", ApiRole.GetRoleList)
		//编辑角色信息
		roleGroup.PUT("role-edit", ApiRole.EditRole)
		//新增角色
		roleGroup.POST("role-add", ApiRole.AddRole)
		//删除角色
		roleGroup.DELETE("role-delete", ApiRole.DeleteRole)
		//根据ID查找角色
		roleGroup.POST("role-id", ApiRole.QueryRoleById)
		//获取角色菜单
		roleGroup.POST("role-menu", ApiRole.GetRoleMenuList)
		//编辑角色菜单
		roleGroup.PUT("role-menu-edit", ApiRole.EditRoleMenu)
		//获取角色API
		roleGroup.POST("role-api", ApiRole.GetRoleApiList)
		//编辑角色Api
		roleGroup.PUT("role-api-edit", ApiRole.EditRoleApi)
		//根据角色查找用户
		roleGroup.POST("role-user", ApiRole.QueryUserByRole)
		//从角色中移除某个用户
		roleGroup.POST("role-user-remove", ApiRole.RemoveRoleUser)
		//添加用户到某个角色
		roleGroup.POST("role-user-add", ApiRole.AddRoleUser)
	    //修改角色部门数据权限
		roleGroup.PUT("role-dept-data-permission-edit", ApiRole.EditRoleDeptDataPermission)
	}
	return Router
}

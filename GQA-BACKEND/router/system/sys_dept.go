package system

import (
	"gin-quasar-admin/api"
	"github.com/gin-gonic/gin"
)

type RouterDept struct {

}

func (r *RouterDept) InitRouterDept(Router *gin.RouterGroup) (R gin.IRoutes) {
	deptGroup := Router.Group("dept")
	apiDept := api.GroupApiApp.ApiSystem.ApiDept
	{
		// 获取部门列表
		deptGroup.POST("dept-list", apiDept.GetDeptList)
		// 编辑部门信息
		deptGroup.PUT("dept-edit", apiDept.EditDept)
		// 新增部门
		deptGroup.POST("dept-add", apiDept.AddDept)
		// 删除部门
		deptGroup.DELETE("dept-delete", apiDept.DeleteDept)
		// 根据ID查找部门
		deptGroup.POST("dept-id", apiDept.QueryDeptById)
	}
	return Router
}

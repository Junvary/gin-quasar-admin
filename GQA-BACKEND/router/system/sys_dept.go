package system

import (
	"github.com/gin-gonic/gin"
)

type RouterDept struct {}

func (r *RouterDept) InitRouterDept(Router *gin.RouterGroup) (R gin.IRoutes) {
	deptGroup := Router.Group("dept")
	{
		//获取部门列表
		deptGroup.POST("dept-list", ApiDept.GetDeptList)
		//编辑部门信息
		deptGroup.PUT("dept-edit", ApiDept.EditDept)
		//新增部门
		deptGroup.POST("dept-add", ApiDept.AddDept)
		//删除部门
		deptGroup.DELETE("dept-delete", ApiDept.DeleteDept)
		//根据ID查找部门
		deptGroup.POST("dept-id", ApiDept.QueryDeptById)
	}
	return Router
}

package system

import (
	"gin-quasar-admin/middleware"
	"github.com/gin-gonic/gin"
)

type RouterDept struct {}

func (r *RouterDept) InitRouterDept(Router *gin.RouterGroup) (R gin.IRoutes) {
	deptGroup := Router.Group("dept").Use(middleware.LogOperationHandler())
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
		//根据部门查找用户
		deptGroup.POST("dept-user", ApiDept.QueryUserByDept)
		//从部门中移除某个用户
		deptGroup.POST("dept-user-remove", ApiDept.RemoveDeptUser)
		//添加用户到某个部门
		deptGroup.POST("dept-user-add", ApiDept.AddDeptUser)
	}
	return Router
}

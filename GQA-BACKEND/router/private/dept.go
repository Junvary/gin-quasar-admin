package private

import "github.com/gin-gonic/gin"

type RouterDept struct{}

func (r *RouterDept) InitRouterRouterDept(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("dept")
	{
		privateGroup.POST("get-dept-list", apiPrivate.ApiDept.GetDeptList)
		privateGroup.POST("edit-dept", apiPrivate.ApiDept.EditDept)
		privateGroup.POST("add-dept", apiPrivate.ApiDept.AddDept)
		privateGroup.POST("delete-dept-by-id", apiPrivate.ApiDept.DeleteDeptById)
		privateGroup.POST("query-dept-by-id", apiPrivate.ApiDept.QueryDeptById)
		privateGroup.POST("query-user-by-dept", apiPrivate.ApiDept.QueryUserByDept)
		privateGroup.POST("remove-dept-user", apiPrivate.ApiDept.RemoveDeptUser)
		privateGroup.POST("add-dept-user", apiPrivate.ApiDept.AddDeptUser)
	}
	return privateGroup
}

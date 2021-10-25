package system

import (
	"gin-quasar-admin/api"
	"github.com/gin-gonic/gin"
)

type RouterDict struct {
}

func (r *RouterDict) InitRouterDict(Router *gin.RouterGroup) (R gin.IRoutes) {
	dictGroup := Router.Group("dict")
	apiDict := api.GroupApiApp.ApiSystem.ApiDict
	{
		// 获取父级字典列表
		dictGroup.POST("dict-list", apiDict.GetDictList)
		// 编辑字典信息
		dictGroup.PUT("dict-edit", apiDict.EditDict)
		// 新增字典
		dictGroup.POST("dict-add", apiDict.AddDict)
		// 删除字典
		dictGroup.DELETE("dict-delete", apiDict.DeleteDict)
		// 根据ID查找字典
		dictGroup.POST("dict-id", apiDict.QueryDictById)
		// 根据ParentId查找字典
		dictGroup.POST("dict-parent-id", apiDict.QueryDictByParentId)
	}
	return Router
}

package system

import (
	"gin-quasar-admin/middleware"
	"github.com/gin-gonic/gin"
)

type RouterDict struct {}

func (r *RouterDict) InitRouterDict(Router *gin.RouterGroup) (R gin.IRoutes) {
	dictGroup := Router.Group("dict").Use(middleware.LogOperationHandler())
	{
		//获取根字典列表
		dictGroup.POST("dict-list", ApiDict.GetDictList)
		//编辑字典信息
		dictGroup.PUT("dict-edit", ApiDict.EditDict)
		//新增字典
		dictGroup.POST("dict-add", ApiDict.AddDict)
		//删除字典
		dictGroup.DELETE("dict-delete", ApiDict.DeleteDict)
		//根据ID查找字典
		dictGroup.POST("dict-id", ApiDict.QueryDictById)
	}
	return Router
}

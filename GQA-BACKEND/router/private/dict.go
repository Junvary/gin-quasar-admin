package private

import "github.com/gin-gonic/gin"

type RouterDict struct{}

func (r *RouterDict) InitRouterRouterDict(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("dict")
	{
		privateGroup.POST("get-dict-list", apiPrivate.ApiDict.GetDictList)
		privateGroup.POST("edit-dict", apiPrivate.ApiDict.EditDict)
		privateGroup.POST("add-dict", apiPrivate.ApiDict.AddDict)
		privateGroup.POST("delete-dict-by-id", apiPrivate.ApiDict.DeleteDictById)
		privateGroup.POST("query-dict-by-id", apiPrivate.ApiDict.QueryDictById)
	}
	return privateGroup
}

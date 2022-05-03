package public

import "github.com/gin-gonic/gin"

type RouterDict struct{}

func (r *RouterDict) InitRouterDict(publicGroup *gin.RouterGroup) (R gin.IRoutes) {
	{
		publicGroup.POST("get-dict-all", apiPublic.ApiDict.GetDictAll)
	}
	return publicGroup
}

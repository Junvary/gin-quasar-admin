package system

import (
	"gin-quasar-admin/api"
	"github.com/gin-gonic/gin"
)

type RouterUpload struct {
}

func (r *RouterUpload) InitRouterUpload(Router *gin.RouterGroup) (R gin.IRoutes) {
	uploadGroup := Router.Group("upload")
	apiUpload := api.GroupApiApp.ApiSystem.ApiUpload
	{
		// 上传头像
		uploadGroup.POST("avatar", apiUpload.UploadAvatar)
		// 上传文件
		uploadGroup.POST("file", apiUpload.UploadFile)
	}
	return Router
}

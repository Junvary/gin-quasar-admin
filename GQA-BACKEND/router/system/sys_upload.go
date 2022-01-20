package system

import (
	"github.com/gin-gonic/gin"
)

type RouterUpload struct {}

func (r *RouterUpload) InitRouterUpload(Router *gin.RouterGroup) (R gin.IRoutes) {
	uploadGroup := Router.Group("upload")/*.Use(middleware.LogOperationHandler())*/
	{
		//上传头像
		uploadGroup.POST("avatar", ApiUpload.UploadAvatar)
		//上传文件
		uploadGroup.POST("file", ApiUpload.UploadFile)
		//上传网站Logo
		uploadGroup.POST("banner-image", ApiUpload.UploadBannerImage)
		//上传网站Logo
		uploadGroup.POST("web-logo", ApiUpload.UploadWebLogo)
		//上传标签页图标
		uploadGroup.POST("header-logo", ApiUpload.UploadHeaderLogo)
	}
	return Router
}

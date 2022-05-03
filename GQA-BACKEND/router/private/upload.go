package private

import "github.com/gin-gonic/gin"

type RouterUpload struct{}

func (r *RouterUpload) InitRouterRouterUpload(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("upload")
	{
		privateGroup.POST("upload-avatar", apiPrivate.ApiUpload.UploadAvatar)
		privateGroup.POST("upload-file", apiPrivate.ApiUpload.UploadFile)
		privateGroup.POST("upload-banner-image", apiPrivate.ApiUpload.UploadBannerImage)
		privateGroup.POST("upload-logo", apiPrivate.ApiUpload.UploadLogo)
		privateGroup.POST("upload-favicon", apiPrivate.ApiUpload.UploadFavicon)
	}
	return privateGroup
}

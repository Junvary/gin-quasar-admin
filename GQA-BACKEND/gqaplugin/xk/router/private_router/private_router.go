package private_router

import (
	"gin-quasar-admin/gqaplugin/xk/api/private_api"
	"gin-quasar-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	xkRouter := privateGroup.Use(middleware.LogOperationHandler())
	{
		//获取news列表
		xkRouter.POST("news-list", private_api.GetNewsList)
		//编辑news信息
		xkRouter.PUT("news-edit", private_api.EditNews)
		//新增news
		xkRouter.POST("news-add", private_api.AddNews)
		//删除news
		xkRouter.DELETE("news-delete", private_api.DeleteNews)
		//根据ID查找news
		xkRouter.POST("news-id", private_api.QueryNewsById)
	}
	{
		//获取project列表
		xkRouter.POST("project-list", private_api.GetProjectList)
		//编辑project信息
		xkRouter.PUT("project-edit", private_api.EditProject)
		//编辑project详情
		xkRouter.PUT("project-edit-detail", private_api.EditProjectDetail)
		//新增project
		xkRouter.POST("project-add", private_api.AddProject)
		//删除project
		xkRouter.DELETE("project-delete", private_api.DeleteProject)
		//根据ID查找project
		xkRouter.POST("project-id", private_api.QueryProjectById)
	}
	{
		//获取honour列表
		xkRouter.POST("honour-list", private_api.GetHonourList)
		//编辑honour信息
		xkRouter.PUT("honour-edit", private_api.EditHonour)
		//新增honour
		xkRouter.POST("honour-add", private_api.AddHonour)
		//删除honour
		xkRouter.DELETE("honour-delete", private_api.DeleteHonour)
		//根据ID查找honour
		xkRouter.POST("honour-id", private_api.QueryHonourById)
	}
	{
		//获取document列表
		xkRouter.POST("document-list", private_api.GetDocumentList)
		//编辑document信息
		xkRouter.PUT("document-edit", private_api.EditDocument)
		//新增document
		xkRouter.POST("document-add", private_api.AddDocument)
		//删除document
		xkRouter.DELETE("document-delete", private_api.DeleteDocument)
		//根据ID查找document
		xkRouter.POST("document-id", private_api.QueryDocumentById)
	}
	{
		//获取download列表
		xkRouter.POST("download-list", private_api.GetDownloadList)
		//编辑download信息
		xkRouter.PUT("download-edit", private_api.EditDownload)
		//新增download
		xkRouter.POST("download-add", private_api.AddDownload)
		//删除download
		xkRouter.DELETE("download-delete", private_api.DeleteDownload)
		//根据ID查找download
		xkRouter.POST("download-id", private_api.QueryDownloadById)
	}
}

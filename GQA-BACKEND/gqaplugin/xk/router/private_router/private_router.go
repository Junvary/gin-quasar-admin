package private_router

import (
	"gin-quasar-admin/gqaplugin/xk/api/private_api"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	{
		//获取news列表
		privateGroup.POST("news-list", private_api.GetNewsList)
		//编辑news信息
		privateGroup.PUT("news-edit", private_api.EditNews)
		//新增news
		privateGroup.POST("news-add", private_api.AddNews)
		//删除news
		privateGroup.DELETE("news-delete", private_api.DeleteNews)
		//根据ID查找news
		privateGroup.POST("news-id", private_api.QueryNewsById)
	}
	{
		//获取project列表
		privateGroup.POST("project-list", private_api.GetProjectList)
		//编辑project信息
		privateGroup.PUT("project-edit", private_api.EditProject)
		//新增project
		privateGroup.POST("project-add", private_api.AddProject)
		//删除project
		privateGroup.DELETE("project-delete", private_api.DeleteProject)
		//根据ID查找project
		privateGroup.POST("project-id", private_api.QueryProjectById)
	}
	{
		//获取honour列表
		privateGroup.POST("honour-list", private_api.GetHonourList)
		//编辑honour信息
		privateGroup.PUT("honour-edit", private_api.EditHonour)
		//新增honour
		privateGroup.POST("honour-add", private_api.AddHonour)
		//删除honour
		privateGroup.DELETE("honour-delete", private_api.DeleteHonour)
		//根据ID查找honour
		privateGroup.POST("honour-id", private_api.QueryHonourById)
	}
	{
		//获取document列表
		privateGroup.POST("document-list", private_api.GetDocumentList)
		//编辑document信息
		privateGroup.PUT("document-edit", private_api.EditDocument)
		//新增document
		privateGroup.POST("document-add", private_api.AddDocument)
		//删除document
		privateGroup.DELETE("document-delete", private_api.DeleteDocument)
		//根据ID查找document
		privateGroup.POST("document-id", private_api.QueryDocumentById)
	}
	{
		//获取download列表
		privateGroup.POST("download-list", private_api.GetDownloadList)
		//编辑download信息
		privateGroup.PUT("download-edit", private_api.EditDownload)
		//新增download
		privateGroup.POST("download-add", private_api.AddDownload)
		//删除download
		privateGroup.DELETE("download-delete", private_api.DeleteDownload)
		//根据ID查找download
		privateGroup.POST("download-id", private_api.QueryDownloadById)
	}
}

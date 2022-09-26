package privaterouter

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/achievement/api/privateapi"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。

	{

		CategoryRouter := privateGroup.Use()

		//获取Category列表
		CategoryRouter.POST("get-category-list", privateapi.GetCategoryList)
		//编辑Category信息
		CategoryRouter.POST("edit-category", privateapi.EditCategory)
		//新增Category
		CategoryRouter.POST("add-category", privateapi.AddCategory)
		//删除Category
		CategoryRouter.POST("delete-category-by-id", privateapi.DeleteCategoryById)
		//根据ID查找Category
		CategoryRouter.POST("query-category-by-id", privateapi.QueryCategoryById)
	}

	{

		ObtainRouter := privateGroup.Use()

		//获取Obtain列表
		ObtainRouter.POST("obtain-find", privateapi.ObtainFind)
		ObtainRouter.POST("get-obtain-list", privateapi.GetObtainList)
	}

}

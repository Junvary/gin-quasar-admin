package privaterouter

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/api/privateapi"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	ExampleRouter := privateGroup.Use()
	{
		//获取test-data列表
		ExampleRouter.POST("get-test-data-list", privateapi.GetTestDataList)
		//编辑test-data信息
		ExampleRouter.POST("edit-test-data", privateapi.EditTestData)
		//新增test-data
		ExampleRouter.POST("add-test-data", privateapi.AddTestData)
		//删除test-data
		ExampleRouter.POST("delete-test-data-by-id", privateapi.DeleteTestDataById)
		//根据ID查找test-data
		ExampleRouter.POST("query-test-data-by-id", privateapi.QueryTestDataById)
		//下载模板
		ExampleRouter.POST("download-template-test-data", privateapi.DownloadTemplateTestData)
		//导出数据
		ExampleRouter.POST("export-test-data", privateapi.ExportTestData)
		//导入数据
		ExampleRouter.POST("import-test-data", privateapi.ImportTestData)
	}
}

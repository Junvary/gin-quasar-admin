package privaterouter

import (
    "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/api/privateapi"
    "github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	{{.PluginCode}}Router := privateGroup.Use()
	{{ range .PluginModel }}
	{
	    //获取{{.ModelName}}列表
        {{$.PluginCode}}Router.POST("get-{{.ModelName}}-list", privateapi.Get{{.ModelName}}List)
        //编辑{{.ModelName}}信息
        {{$.PluginCode}}Router.POST("edit-{{.ModelName}}", privateapi.Edit{{.ModelName}})
        //新增{{.ModelName}}
        {{$.PluginCode}}Router.POST("add-{{.ModelName}}", privateapi.Add{{.ModelName}})
        //删除{{.ModelName}}
        {{$.PluginCode}}Router.POST("delete-{{.ModelName}}-by-id", privateapi.Delete{{.ModelName}}ById)
        //根据ID查找{{.ModelName}}
        {{$.PluginCode}}Router.POST("query-{{.ModelName}}-by-id", privateapi.Query{{.ModelName}}ById)
	}
	{{ end }}
}

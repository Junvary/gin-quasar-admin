package privaterouter

import (
    "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/api/privateapi"
    gqaMiddleware "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/middleware"
    "github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	{{ range .PluginModel }}
	{
	    {{ if .WithLogOperation }}
	    {{.ModelName}}Router := privateGroup.Use(gqaMiddleware.LogOperationHandler())
	    {{ else }}
	    {{.ModelName}}Router := privateGroup.Use()
	    {{ end }}
	    //获取{{.ModelName}}列表
        {{.ModelName}}Router.POST("get-{{.ModelName}}-list", privateapi.Get{{.ModelName}}List)
        //编辑{{.ModelName}}信息
        {{.ModelName}}Router.POST("edit-{{.ModelName}}", privateapi.Edit{{.ModelName}})
        //新增{{.ModelName}}
        {{.ModelName}}Router.POST("add-{{.ModelName}}", privateapi.Add{{.ModelName}})
        //删除{{.ModelName}}
        {{.ModelName}}Router.POST("delete-{{.ModelName}}-by-id", privateapi.Delete{{.ModelName}}ById)
        //根据ID查找{{.ModelName}}
        {{.ModelName}}Router.POST("query-{{.ModelName}}-by-id", privateapi.Query{{.ModelName}}ById)
	}
	{{ end }}
}

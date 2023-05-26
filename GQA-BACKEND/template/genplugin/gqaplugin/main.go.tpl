package {{.PluginCode}}

import (
    "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/data"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/router/privaterouter"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/router/publicrouter"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
)

var Plugin{{.PluginCode}} = new({{.PluginCode}})

type {{.PluginCode}} struct{}

func (*{{.PluginCode}}) PluginCode() string { //实现接口方法，插件编码。返回值：请使用 "plugin-"前缀开头。
	return "plugin-{{.PluginCode}}"
}

func (*{{.PluginCode}}) PluginName() string { //实现接口方法，插件名称
	return "{{.PluginName}}"
}

func (*{{.PluginCode}}) PluginVersion() string { //实现接口方法，插件版本
	return "v0.0.1"
}

func (*{{.PluginCode}}) PluginMemo() string { //实现接口方法，插件描述
	return "这是{{.PluginName}}插件"
}

func (*{{.PluginCode}}) PluginRouterPublic(publicGroup *gin.RouterGroup) { //实现接口方法，公开路由初始化
    publicrouter.InitPublicRouter(publicGroup)
}

func (*{{.PluginCode}}) PluginRouterPrivate(privateGroup *gin.RouterGroup) { //实现接口方法，鉴权路由初始化
	privaterouter.InitPrivateRouter(privateGroup)
}

func (*{{.PluginCode}}) PluginMigrate() []interface{} { //实现接口方法，迁移插件数据表
	var ModelList = []interface{}{
	    {{ range .PluginModel }}
        model.GqaPlugin{{$.PluginCode}}{{ .ModelName }}{},
        {{ end }}
    }
    return ModelList
}

func (*{{.PluginCode}}) PluginData() []interface{ LoadData() (err error) } { //实现接口方法，初始化数据
	var DataList = []interface{ LoadData() (err error) }{
        data.Plugin{{ .PluginCode }}SysApi,
        data.Plugin{{ .PluginCode }}SysRoleApi,
        data.Plugin{{ .PluginCode }}SysMenu,
        data.Plugin{{ .PluginCode }}SysRoleMenu,
    }
    return DataList
}

func (*{{.PluginCode}}) PluginCron() ([]gqaModel.SysCron, map[uuid.UUID]func()) {
	return nil, nil
}

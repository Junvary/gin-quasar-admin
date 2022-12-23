package example

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/data"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/router/privaterouter"
	"github.com/gin-gonic/gin"
)

var PluginExample = new(example)

type example struct{}

func (*example) PluginCode() string { //实现接口方法，插件编码。返回值：请使用 "plugin-"前缀开头。
	return "plugin-example"
}

func (*example) PluginName() string { //实现接口方法，插件名称
	return "示例插件"
}

func (*example) PluginVersion() string { //实现接口方法，插件版本
	return "v1.0.0"
}

func (*example) PluginMemo() string { //实现接口方法，插件描述
	return "这是示例插件"
}

func (p *example) PluginRouterPublic(publicGroup *gin.RouterGroup) { //实现接口方法，公开路由初始化
}

func (p *example) PluginRouterPrivate(privateGroup *gin.RouterGroup) { //实现接口方法，鉴权路由初始化
	privaterouter.InitPrivateRouter(privateGroup)
}

func (p *example) PluginMigrate() []interface{} { //实现接口方法，迁移插件数据表
	var ModelList = []interface{}{
		model.GqaPluginExampleTestData{},
	}
	return ModelList
}

func (p *example) PluginData() []interface{ LoadData() (err error) } { //实现接口方法，初始化数据
	var DataList = []interface{ LoadData() (err error) }{
		data.PluginExampleSysApi,
		data.PluginExampleSysRoleApi,
		data.PluginExampleSysMenu,
		data.PluginExampleSysRoleMenu,
		data.PluginExampleTestData,
	}
	return DataList
}

func (p *example) PluginCron() map[string]func() {
	return nil
}

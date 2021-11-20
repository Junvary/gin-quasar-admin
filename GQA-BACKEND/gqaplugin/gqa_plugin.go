package gqaplugin

import (
	"gin-quasar-admin/gqaplugin/example"
	"github.com/gin-gonic/gin"
)

/*
	1.import插件
	2.插件填入以下三个方法
*/

func RegisterPluginRouter(PublicGroup, PrivateGroup *gin.RouterGroup) { //注册插件路由
	PluginRouter(PublicGroup, PrivateGroup, example.PluginExample)
	//顺序扩展其他插件
	//PluginRouter(PublicGroup, PrivateGroup, ??.Plugin??)
}

func MigratePluginModel() []interface{} { //迁移插件数据库
	var model []interface{}
	model = append(model,
		example.PluginExample.PluginMigrate()...,
	//顺序添加其他插件
	//??.Plugin??.PluginMigrate(),
	)
	return model
}

func LoadPluginData() []interface{ LoadData() (err error) } { //初始化插件数据
	var data []interface{ LoadData() (err error) }
	data = append(data,
		example.PluginExample.PluginData()...,
	//顺序添加其他插件
	//??.Plugin??.PluginData()...,
	)
	return data
}

/*
	插件需实现的接口
*/

type GqaPlugin interface {
	PluginCode() string                                //插件代号，用于路由分组名
	PluginName() string                                //插件名称
	PluginRouterPublic(publicGroup *gin.RouterGroup)   //公开路由
	PluginRouterPrivate(privateGroup *gin.RouterGroup) //鉴权路由
	PluginMigrate() []interface{}                      //迁移数据库模型
	PluginData() []interface{ LoadData() (err error) } //初始化数据
}

func PluginRouter(publicGroup, privateGroup *gin.RouterGroup, Plugin ...GqaPlugin) {
	for i := range Plugin {
		PublicGroup := publicGroup.Group(Plugin[i].PluginCode())
		Plugin[i].PluginRouterPublic(PublicGroup)
		PrivateGroup := privateGroup.Group(Plugin[i].PluginCode())
		Plugin[i].PluginRouterPrivate(PrivateGroup)
	}
}

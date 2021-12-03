package public

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"runtime"
)

type ApiCheckAndInitDb struct{}

func (a *ApiCheckAndInitDb) CheckDb(c *gin.Context) {
	goVersion := runtime.Version()
	ginVersion := gin.Version
	type plugin struct {
		PluginName    string
		PluginCode    string
		PluginVersion string
		PluginRemark  string
	}
	var pluginList []plugin
	for _, v := range gqaplugin.PluginList {
		pluginList = append(pluginList, plugin{
			PluginName:    v.PluginName(),
			PluginCode:    v.PluginCode(),
			PluginVersion: v.PluginVersion(),
			PluginRemark:  v.PluginRemark(),
		})
	}
	if global.GqaDb != nil {
		pluginLoginLayout := utils.GetConfigFrontend("gqaPluginLoginLayout")
		global.GqaLog.Info("数据库无需初始化")
		global.SuccessMessageData(gin.H{
			"needInit":          false,
			"goVersion":         goVersion,
			"ginVersion":        ginVersion,
			"pluginList":        pluginList,
			"pluginLoginLayout": pluginLoginLayout}, "数据库无需初始化", c)
		return
	} else {
		global.GqaLog.Info("数据库需要初始化")
		global.SuccessMessageData(gin.H{
			"needInit":   true,
			"goVersion":  goVersion,
			"ginVersion": ginVersion,
			"pluginList": pluginList}, "数据库需要初始化", c)
		return
	}
}

func (a *ApiCheckAndInitDb) InitDb(c *gin.Context) {
	if global.GqaDb != nil {
		global.GqaLog.Error("已存在数据库配置，无须再次初始化！")
		global.ErrorMessage("已存在数据库配置，无须再次初始化！", c)
		return
	}

	var initDbInfo system.RequestInitDb
	if err := c.ShouldBindJSON(&initDbInfo); err != nil {
		global.GqaLog.Error("参数校验不通过！", zap.Any("err", err))
		global.ErrorMessage("参数校验不通过，"+err.Error(), c)
		return
	}
	if err := ServiceCheckAndInitDb.CheckAndInitDb(initDbInfo); err != nil {
		global.GqaLog.Error("自动创建数据库失败！", zap.Any("err", err))
		global.ErrorMessage("自动创建数据库失败，"+err.Error(), c)
		return
	}
	global.GqaCasbin = utils.Casbin(global.GqaDb)
	global.SuccessMessage("系统初始化成功，请登录！", c)
}

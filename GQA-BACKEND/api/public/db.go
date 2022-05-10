package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"runtime"
)

type ApiDb struct{}

func (a *ApiDb) CheckDb(c *gin.Context) {
	goVersion := runtime.Version()
	ginVersion := gin.Version

	var pluginList []model.Plugin
	for _, v := range gqaplugin.PluginList {
		pluginList = append(pluginList, model.Plugin{
			PluginName:    v.PluginName(),
			PluginCode:    v.PluginCode(),
			PluginVersion: v.PluginVersion(),
			PluginMemo:    v.PluginMemo(),
		})
	}
	if global.GqaDb != nil {
		pluginLoginLayout := utils.GetConfigFrontend("gqaPluginLoginLayout")
		global.GqaLogger.Info("数据库无需初始化")
		model.ResponseSuccessMessageData(gin.H{
			"need_init":           false,
			"go_version":          goVersion,
			"gin_version":         ginVersion,
			"plugin_list":         pluginList,
			"plugin_login_layout": pluginLoginLayout}, "数据库无需初始化", c)
		return
	} else {
		global.GqaLogger.Info("数据库需要初始化")
		model.ResponseSuccessMessageData(gin.H{
			"need_init":   true,
			"go_version":  goVersion,
			"gin_version": ginVersion,
			"plugin_list": pluginList}, "数据库需要初始化", c)
		return
	}
}

func (a *ApiDb) InitDb(c *gin.Context) {
	if global.GqaDb != nil {
		global.GqaLogger.Error("已存在数据库配置，无须再次初始化！")
		model.ResponseErrorMessage("已存在数据库配置，无须再次初始化！", c)
		return
	}

	var initDbInfo model.RequestDbInit
	if err := model.RequestShouldBindJSON(c, &initDbInfo); err != nil {
		return
	}
	if err := servicePublic.ServiceDb.InitDb(initDbInfo); err != nil {
		global.GqaLogger.Error("创建数据库失败！", zap.Any("err", err))
		model.ResponseErrorMessage("创建数据库失败，"+err.Error(), c)
		return
	}
	model.ResponseSuccessMessage("系统初始化成功，请登录！", c)
}

package utils

import (
	"fmt"
	"gin-quasar-admin/global"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Casbin(db *gorm.DB) *casbin.Enforcer {
	a, _ := gormadapter.NewAdapterByDB(db)
	model := GetConfig("casbinModel")
	if model == "" {
		global.GqaLog.Error("启动失败，找不到Casbin模型！")
		panic(fmt.Errorf("启动失败，找不到Casbin模型！"))
	}
	e, err := casbin.NewEnforcer(model, a)
	if err != nil {
		global.GqaLog.Error("启动失败，初始化Casbin失败！", zap.Any("err", err))
		panic(fmt.Errorf("启动失败，初始化Casbin失败：%s \n", err.Error()))
	}
	_ = e.LoadPolicy()
	return e
}

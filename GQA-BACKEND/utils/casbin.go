package utils

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Casbin(db *gorm.DB) *casbin.Enforcer {
	a, _ := gormadapter.NewAdapterByDB(db)
	model := GetConfigBackend("casbinModel")
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

func GetUserRole(c *gin.Context) (err error, role []system.SysRole) {
	username := GetUsername(c)
	var user system.SysUser
	err = global.GqaDb.Preload("Role").Where("username=?", username).First(&user).Error
	if err != nil {
		return err, nil
	}
	var userRole []system.SysRole
	err = global.GqaDb.Model(&user).Association("Role").Find(&userRole)
	if err != nil {
		return err, nil
	}
	return nil, userRole
}

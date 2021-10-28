package system

import (
	"gin-quasar-admin/global"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type ServiceCasbin struct {
}

func (s *ServiceCasbin) Casbin() *casbin.Enforcer {
	a, _ := gormadapter.NewAdapterByDB(global.GqaDb)
	e, _ := casbin.NewEnforcer(global.GqaConfig.Casbin.ModelPath, a)
	_ = e.LoadPolicy()
	return e
}


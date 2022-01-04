package main

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/boot"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
)

func main() {
	global.GqaViper = boot.Viper()
	global.GqaLog = boot.Zap()
	global.GqaDb = boot.Mysql()
	if global.GqaDb != nil {
		boot.Migrate(global.GqaDb)
		global.GqaCasbin = utils.Casbin(global.GqaDb)
		db, _ := global.GqaDb.DB()
		defer db.Close()
	}
	boot.Boot()
}
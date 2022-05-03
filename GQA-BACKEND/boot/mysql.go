package boot

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/config"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql() *gorm.DB {
	gqaMysqlConfig := global.GqaConfig.Mysql
	if gqaMysqlConfig.Database == "" {
		return nil
	}
	mysqlConfig := config.MysqlConfig(gqaMysqlConfig)
	gormConfig := config.GormConfig()
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig); err != nil {
		return nil
	} else {
		mysqlDb, _ := db.DB()
		mysqlDb.SetMaxIdleConns(gqaMysqlConfig.MaxIdle)
		mysqlDb.SetMaxOpenConns(gqaMysqlConfig.MaxOpen)
		return db
	}
}

package public

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/config"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/data"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceDb struct{}

type dataInterface interface {
	LoadData() (err error)
}

var migrateList = []interface{}{
	model.SysUser{},
	model.SysRole{},
	model.SysUserRole{},
	model.SysButton{},
	model.SysRoleButton{},
	model.SysMenu{},
	model.SysRoleMenu{},
	model.SysApi{},
	model.SysRoleApi{},
	model.SysDept{},
	model.SysDeptUser{},
	model.SysDict{},
	model.SysConfigBackend{},
	model.SysConfigFrontend{},
	model.SysLogLogin{},
	model.SysLogOperation{},
	model.SysNotice{},
	model.SysNoticeToUser{},
	model.SysTodo{},
	model.SysUserOnline{},
}

var dataList = []dataInterface{
	data.SysUser,
	data.SysRole,
	data.SysUserRole,
	data.SysButton,
	data.SysRoleButton,
	data.SysMenu,
	data.SysRoleMenu,
	data.SysApi,
	data.SysRoleApi,
	data.SysDept,
	data.SysDeptUser,
	data.SysDict,
	data.SysConfigBackend,
	data.SysConfigFrontend,
}

func (s *ServiceDb) InitDb(initDbInfo model.RequestDbInit) error {
	if err := s.createDatabase(initDbInfo); err != nil {
		return err
	}
	MysqlConfig := config.Mysql{
		Host:     initDbInfo.DbHost,
		Port:     initDbInfo.DbPort,
		User:     initDbInfo.DbUser,
		Password: initDbInfo.DbPassword,
		Database: initDbInfo.DbSchema,
	}
	if MysqlConfig.Database == "" {
		return nil
	}
	mysqlConfig := config.MysqlConfig(MysqlConfig)
	gormConfig := config.GormConfig()
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig); err != nil {
		return nil
	} else {
		mysqlDb, _ := db.DB()
		mysqlDb.SetMaxIdleConns(MysqlConfig.MaxIdle)
		mysqlDb.SetMaxOpenConns(MysqlConfig.MaxOpen)
		global.GqaDb = db
	}
	//迁移Gin-Quasar-Admin数据库
	err := global.GqaDb.AutoMigrate(migrateList...)
	if err != nil {
		global.GqaDb = nil
		global.GqaLogger.Error("Merge Gin-Quasar-Admin database failed！", zap.Any("err", err))
		return errors.New("Merge Gin-Quasar-Admin database failed：" + err.Error())
	}
	//迁移Gin-Quasar-Admin插件数据库
	err = global.GqaDb.AutoMigrate(gqaplugin.MigratePluginModel()...)
	if err != nil {
		global.GqaDb = nil
		global.GqaLogger.Error("Merge Gin-Quasar-Admin database failed！", zap.Any("err", err))
		return errors.New("Merge Gin-Quasar-Admin database failed：" + err.Error())
	}
	// Init Gin-Quasar-Admin data
	for _, v := range dataList {
		err = v.LoadData()
		if err != nil {
			return err
		}
	}
	// Init Gin-Quasar-Admin plugin data
	for _, v := range gqaplugin.LoadPluginData() {
		err = v.LoadData()
		if err != nil {
			return err
		}
	}
	if err = s.writeConfig(global.GqaViper, MysqlConfig); err != nil {
		return err
	}
	return nil
}

func (s *ServiceDb) createDatabase(initDbInfo model.RequestDbInit) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", initDbInfo.DbUser, initDbInfo.DbPassword, initDbInfo.DbHost, initDbInfo.DbPort)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", initDbInfo.DbSchema)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func (s *ServiceDb) writeConfig(viper *viper.Viper, mysql config.Mysql) error {
	global.GqaConfig.Mysql = mysql
	cs := utils.Struct2Map(global.GqaConfig)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

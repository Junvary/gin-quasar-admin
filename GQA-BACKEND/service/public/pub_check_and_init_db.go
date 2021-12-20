package public

import (
	"database/sql"
	"errors"
	"fmt"
	"gin-quasar-admin/boot/data"
	"gin-quasar-admin/config/config"
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"

	//"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceCheckAndInitDb struct {}

func (s *ServiceCheckAndInitDb) CheckAndInitDb(initDbInfo system.RequestInitDb) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", initDbInfo.Username, initDbInfo.Password, initDbInfo.Host, initDbInfo.Port)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", initDbInfo.DbName)
	if err := s.createTable(dsn, "mysql", createSql); err != nil {
		return err
	}
	MysqlConfig := config.Mysql{
		Path:     fmt.Sprintf("%s:%s", initDbInfo.Host, initDbInfo.Port),
		DbName:   initDbInfo.DbName,
		Username: initDbInfo.Username,
		Password: initDbInfo.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}
	if MysqlConfig.DbName == "" {
		return nil
	}
	linkDns := MysqlConfig.Username + ":" + MysqlConfig.Password + "@tcp(" + MysqlConfig.Path + ")/" + MysqlConfig.DbName + "?" + MysqlConfig.Config
	mysqlConfig := mysql.Config{
		DSN:                       linkDns, // DSN data source name
		DefaultStringSize:         255,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true, NamingStrategy: schema.NamingStrategy{SingularTable: true}}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(MysqlConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(MysqlConfig.MaxOpenConns)
		global.GqaDb = db
	}
	//迁移Gin-Quasar-Admin数据库
	err := global.GqaDb.AutoMigrate(
		system.SysUser{},
		system.SysRole{},
		system.SysUserRole{},
		system.SysMenu{},
		system.SysRoleMenu{},
		system.SysDept{},
		system.SysDeptUser{},
		system.SysDict{},
		system.SysApi{},
		gormadapter.CasbinRule{},
		system.SysConfigBackend{},
		system.SysConfigFrontend{},
		system.SysLogLogin{},
	)
	if err != nil {
		global.GqaDb = nil
		global.GqaLog.Error("迁移【Gin-Quasar-admin】数据库失败！", zap.Any("err", err))
		return errors.New("迁移【Gin-Quasar-admin】数据库失败：" + err.Error())
	}
	//迁移GQA-Plugin数据库
	err = global.GqaDb.AutoMigrate(gqaplugin.MigratePluginModel()...)
	if err != nil {
		global.GqaDb = nil
		global.GqaLog.Error("迁移【GQA-Plugin】数据库失败！", zap.Any("err", err))
		return errors.New("迁移【GQA-Plugin】数据库失败：" + err.Error())
	}
	//初始化Gin-Quasar-Admin数据
	err = s.loadData(
		data.SysUser,
		data.SysRole,
		data.SysUserRole,
		data.SysMenu,
		data.SysRoleMenu,
		data.SysDept,
		data.SysDeptUser,
		data.SysDict,
		data.SysApi,
		data.SysCasbin,
		data.SysConfigBackend,
		data.SysConfigFrontend,
	)
	if err != nil {
		global.GqaDb = nil
		global.GqaLog.Error("初始化【Gin-Quasar-admin】数据失败！", zap.Any("err", err))
		return errors.New("初始化【Gin-Quasar-admin】数据失败：" + err.Error())
	}
	//初始化GQA-Plugin数据
	err = s.loadData(gqaplugin.LoadPluginData()...)
	if err != nil {
		global.GqaDb = nil
		global.GqaLog.Error("初始化【GQA-Plugin】数据失败！", zap.Any("err", err))
		return errors.New("初始化【GQA-Plugin】数据失败：" + err.Error())
	}

	if err = s.writeConfig(global.GqaViper, MysqlConfig); err != nil {
		return err
	}
	return nil
}

func (s *ServiceCheckAndInitDb) createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
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

func (s *ServiceCheckAndInitDb) loadData(InitDBFunctions ...interface{ LoadData() (err error) }) (err error) {
	for _, v := range InitDBFunctions {
		err = v.LoadData()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ServiceCheckAndInitDb) writeConfig(viper *viper.Viper, mysql config.Mysql) error {
	global.GqaConfig.Mysql = mysql
	cs := utils.StructToMap(global.GqaConfig)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysConfig = new(sysConfig)

type sysConfig struct{}

var sysConfigData = []system.SysConfig{
	{GqaModel: global.GqaModel{Status: "on", Sort: 1, Remark: "测试配置1", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Option: "test1", Default: "a",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 2, Remark: "测试配置2", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Option: "test2", Default: "b",
	},
}

func (s *sysConfig) Init() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysConfig{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_config 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("sys_config 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysConfigData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_config 表初始数据成功！")
		global.GqaLog.Error("sys_config 表初始数据成功！")
		return nil
	})
}

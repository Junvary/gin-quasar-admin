package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginExampleSysApi = new(sysApi)

type sysApi struct{}

var sysApiData = []system.SysApi{
	{GqaModel: global.GqaModel{Stable: "no", Status: "on", Sort: 47, Remark: "插件：获取news-list", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-example", ApiPath: "/plugin-example/news-list", ApiMethod: "POST",
	},
}

func (s *sysApi) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysApi{}).Where("api_group = ?", "plugin-example").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_api 表中example插件数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("[GQA-Plugin] --> sys_api 表中example插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> example插件数据进入 sys_api 表初始数据成功！")
		global.GqaLog.Error("[GQA-Plugin] --> example插件数据进入 sys_api 表初始数据成功！")
		return nil
	})
}

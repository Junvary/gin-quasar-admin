package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysDict = new(sysDict)

type sysDict struct{}

var sysDictData = []system.SysDict{
	// 父级设置ID
	{GqaModel: global.GqaModel{Id: 1, Status: "on", Sort: 1, Remark: "这是性别字典", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 0, Value: "gender", Label: "性别",
	},
	// 父级设置ID
	{GqaModel: global.GqaModel{Id: 2, Status: "on", Sort: 2, Remark: "这是是否状态", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 0, Value: "status", Label: "是否状态",
	},

	// 子级内容：
	{GqaModel: global.GqaModel{Status: "on", Sort: 1, Remark: "这是男", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 1, Value: "m", Label: "男",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 2, Remark: "这是女", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 1, Value: "f", Label: "女",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 3, Remark: "这是保密", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 1, Value: "u", Label: "保密",
	},

	{GqaModel: global.GqaModel{Status: "on", Sort: 1, Remark: "这是是", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 2, Value: "on", Label: "是",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 2, Remark: "这是否", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 2, Value: "off", Label: "否",
	},
}

func (s *sysDict) Init() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysDict{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_dict 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("sys_dict 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysDictData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_dict 表初始数据成功！")
		global.GqaLog.Error("sys_dict 表初始数据成功！")
		return nil
	})
}

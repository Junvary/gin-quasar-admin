package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginVoteSysDict = new(sysDict)

type sysDict struct{}

var sysDictData = []system.SysDict{
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 701, Remark: "投票类型", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "voteType", DictLabel: "投票类型",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteType", DictCode: "v1", DictLabel: "党员投票",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteType", DictCode: "v2", DictLabel: "管理人员投票",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 702, Remark: "投票细类", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "voteTypeDetail", DictLabel: "投票细类",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetail", DictCode: "base_1", DictLabel: "德",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetail", DictCode: "base_2", DictLabel: "能",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetail", DictCode: "base_3", DictLabel: "勤",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 4, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetail", DictCode: "base_4", DictLabel: "绩",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 5, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetail", DictCode: "base_5", DictLabel: "廉",
	},
	//{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 6, CreatedAt: time.Now(), CreatedBy: "admin"},
	//	ParentCode: "voteTypeDetail", DictCode: "other_1", DictLabel: "笔记",
	//},
	//{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 7, CreatedAt: time.Now(), CreatedBy: "admin"},
	//	ParentCode: "voteTypeDetail", DictCode: "other_2", DictLabel: "笔记2",
	//},
}

func (s *sysDict) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysDict{}).Where("parent_code = ?", "voteType").Or("dict_code = ?", "voteTypeDetail").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_dict 表中vote插件菜单已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[GQA-Plugin] --> sys_dict 表中vote插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysDictData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> vote插件初始数据进入 sys_dict 表成功！")
		global.GqaLog.Info("[GQA-Plugin] --> vote插件初始数据进入 sys_dict 表成功！")
		return nil
	})
}

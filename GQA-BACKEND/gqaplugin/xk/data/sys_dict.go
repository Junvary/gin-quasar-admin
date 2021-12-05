package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginXkSysDict = new(sysDict)

type sysDict struct{}

var sysDictData = []system.SysDict{
	{GqaModel: global.GqaModel{Status: "on", Sort: 801, Remark: "编程语言/框架", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "codeLanguage", DictLabel: "编程语言/框架",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 1, Remark: "Java", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "Java", DictLabel: "Java",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 2, Remark: "Python", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "Python", DictLabel: "Python",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 3, Remark: "C#", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "C#", DictLabel: "C#",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 4, Remark: "Go", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "Go", DictLabel: "Go",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 5, Remark: "Vue", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "Vue", DictLabel: "Vue",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 6, Remark: "React", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "React", DictLabel: "React",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 7, Remark: "Bootstrap", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "Bootstrap", DictLabel: "Bootstrap",
	},

	{GqaModel: global.GqaModel{Status: "on", Sort: 802, Remark: "项目节点", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "projectNode", DictLabel: "项目节点",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 1, Remark: "需求接报", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "demand", DictLabel: "需求接报",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 2, Remark: "开发中", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "develop", DictLabel: "开发中",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 3, Remark: "测试中", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "test", DictLabel: "测试中",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 4, Remark: "交付上线", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "online", DictLabel: "交付上线",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 5, Remark: "移交运维", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "maintain", DictLabel: "移交运维",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 6, Remark: "产品下线", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "offline", DictLabel: "产品下线",
	},
}

func (s *sysDict) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysDict{}).Where("parent_code = ?", "projectNode").Or("dict_code = ?", "projectNode").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_dict 表中xk插件菜单已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("[GQA-Plugin] --> sys_dict 表中xk插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysDictData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xk插件初始数据进入 sys_dict 表成功！")
		global.GqaLog.Error("[GQA-Plugin] --> xk插件初始数据进入 sys_dict 表成功！")
		return nil
	})
}

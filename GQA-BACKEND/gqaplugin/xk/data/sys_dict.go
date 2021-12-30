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
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 801, Remark: "编程语言/框架", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "codeLanguage", DictLabel: "编程语言/框架",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "Java", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "Java", DictLabel: "Java",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "Python", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "Python", DictLabel: "Python",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, Remark: "C#", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "C#", DictLabel: "C#",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 4, Remark: "Go", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "Go", DictLabel: "Go",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 5, Remark: "Vue", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "Vue", DictLabel: "Vue",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 6, Remark: "React", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "React", DictLabel: "React",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 7, Remark: "Bootstrap", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "codeLanguage", DictCode: "Bootstrap", DictLabel: "Bootstrap",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 802, Remark: "项目节点", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "projectNode", DictLabel: "项目节点",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1,
		Remark: "需求申请文档、需求规格说明书、需求确认单等", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "p1", DictLabel: "需求申请",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2,
		Remark: "技术方案、系统设计、数据库设计等", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "p2", DictLabel: "系统设计",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3,
		Remark: "进度计划、进度汇报等", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "p3", DictLabel: "系统开发",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 4,
		Remark: "测试计划、问题反馈、测试用例、测试记录等", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "p4", DictLabel: "系统测试",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 5,
		Remark: "上线计划、操作手册、环境配置、代码位置等", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "p5", DictLabel: "交付上线",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 6,
		Remark: "运维手册、开关机手册、日常点检手册等", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "p6", DictLabel: "产品运维",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 7,
		Remark: "下线原因、替代产品等", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "p7", DictLabel: "产品下线",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 21,
		Remark: "会议纪要、培训计划、培训记录、会议签到表等", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "m1", DictLabel: "会议材料",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 22,
		Remark: "立项申请表、立项会确认表、功能确认表、验收表、代码统计等", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "m2", DictLabel: "立项材料",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 23,
		Remark: "各类软著申请材料", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "projectNode", DictCode: "m3", DictLabel: "软著材料",
	},
}

func (s *sysDict) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysDict{}).Where("parent_code = ?", "projectNode").Or("dict_code = ?", "projectNode").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_dict 表中xk插件菜单已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[GQA-Plugin] --> sys_dict 表中xk插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysDictData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xk插件初始数据进入 sys_dict 表成功！")
		global.GqaLog.Info("[GQA-Plugin] --> xk插件初始数据进入 sys_dict 表成功！")
		return nil
	})
}

package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginXkSysApi = new(sysApi)

type sysApi struct{}

var sysApiData = []system.SysApi{
	{GqaModel: global.GqaModel{Status: "on", Sort: 48, Remark: "插件xk：获取news-list", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/news-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 49, Remark: "插件xk：编辑news信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/news-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 50, Remark: "插件xk：新增news", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/news-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 51, Remark: "插件xk：删除news", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/news-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 52, Remark: "插件xk：根据ID查找news", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/news-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 53, Remark: "插件xk：获取project-list", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/project-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 54, Remark: "插件xk：编辑project信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/project-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 55, Remark: "插件xk：新增project", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/project-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 56, Remark: "插件xk：删除project", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/project-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 57, Remark: "插件xk：根据ID查找project", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/project-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 58, Remark: "插件xk：获取honour-list", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/honour-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 59, Remark: "插件xk：编辑honour信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/honour-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 60, Remark: "插件xk：新增honour", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/honour-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 61, Remark: "插件xk：删除honour", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/honour-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 62, Remark: "插件xk：根据ID查找honour", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/honour-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 63, Remark: "插件xk：获取document-list", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/document-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 64, Remark: "插件xk：编辑document信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/document-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 65, Remark: "插件xk：新增document", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/document-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 66, Remark: "插件xk：删除document", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/document-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 67, Remark: "插件xk：根据ID查找document", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/document-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 68, Remark: "插件xk：获取download-list", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/download-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 69, Remark: "插件xk：编辑download信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/download-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 70, Remark: "插件xk：新增download", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/download-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 71, Remark: "插件xk：删除download", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/download-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 72, Remark: "插件xk：根据ID查找download", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-xk", ApiPath: "/plugin-xk/download-id", ApiMethod: "POST",
	},
}

func (s *sysApi) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysApi{}).Where("api_group = ?", "plugin-xk").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_api 表中xk插件数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("[GQA-Plugin] --> sys_api 表中xk插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xk插件初始数据进入 sys_api 表成功！")
		global.GqaLog.Error("[GQA-Plugin] --> xk插件初始数据进入 sys_api 表成功！")
		return nil
	})
}

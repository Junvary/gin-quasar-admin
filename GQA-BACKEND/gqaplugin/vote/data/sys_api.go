package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginVoteSysApi = new(sysApi)

type sysApi struct{}

var sysApiData = []system.SysApi{
	{GqaModel: global.GqaModel{Status: "on", Sort: 701, Remark: "插件vote：随机和固定投票人列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-random-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 702, Remark: "插件vote：获取随机投票人", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-random-get", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 703, Remark: "插件vote：新增随机和固定投票人", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-random-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 704, Remark: "插件vote：删除随机和固定投票人", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-random-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 705, Remark: "插件vote：投票", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/vote-handle", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 706, Remark: "插件vote：检查是否可以投票-党员", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/can-vote-dy", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 707, Remark: "插件vote：检查是否可以投票-管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/can-vote-gl", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 708, Remark: "插件vote：投票结果列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/vote-result-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 709, Remark: "插件vote：投票结果图形", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/vote-result-chart", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 710, Remark: "插件vote：固定投票人列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 711, Remark: "插件vote：新增固定投票人", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 712, Remark: "插件vote：删除固定投票人", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 713, Remark: "插件vote：候选人列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/candidate-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 714, Remark: "插件vote：新增候选人", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/candidate-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 715, Remark: "插件vote：删除候选人", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/candidate-delete", ApiMethod: "DELETE",
	},
}

func (s *sysApi) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysApi{}).Where("api_group = ?", "plugin-vote").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_api 表中vote插件数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[GQA-Plugin] --> sys_api 表中vote插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> vote插件初始数据进入 sys_api 表成功！")
		global.GqaLog.Info("[GQA-Plugin] --> vote插件初始数据进入 sys_api 表成功！")
		return nil
	})
}

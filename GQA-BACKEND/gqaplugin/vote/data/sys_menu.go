package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginVoteSysMenu = new(sysMenu)

type sysMenu struct{}

var sysMenuData = []system.SysMenu{
	{GqaModel: global.GqaModel{Status: "on", Sort: 701, Remark: "投票系统", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "GqaPluginVote", Title: "投票系统", Icon: "how_to_vote",
		Path: "", Component: "",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 1, Remark: "民主评议党员投票", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-vote-vote-dy", ParentCode: "GqaPluginVote", Title: "民主评议党员投票", Icon: "how_to_vote",
		Path: "/plugin-vote/vote/vote-dy", Component: "/Plugin/Vote/VoteDy/index",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 2, Remark: "民主评议业务骨干及管理人员投票", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-vote-vote-gl", ParentCode: "GqaPluginVote", Title: "民主评议业务骨干及管理人员投票", Icon: "how_to_vote",
		Path: "/plugin-vote/vote/vote-gl", Component: "/Plugin/Vote/VoteGl/index",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 3, Remark: "随机抽取投票人", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-vote-voter-random", ParentCode: "GqaPluginVote", Title: "随机抽取投票人", Icon: "refresh",
		Path: "/plugin-vote/vote/voter-random", Component: "/Plugin/Vote/VoterRandom/index",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 4, Remark: "固定投票人配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-vote-voter", ParentCode: "GqaPluginVote", Title: "固定投票人配置", Icon: "emoji_people",
		Path: "/plugin-vote/vote/voter", Component: "/Plugin/Vote/Voter/index",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 5, Remark: "候选人配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-vote-candidate", ParentCode: "GqaPluginVote", Title: "候选人配置", Icon: "people",
		Path: "/plugin-vote/vote/candidate", Component: "/Plugin/Vote/Candidate/index",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 6, Remark: "投票结果", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-vote-score-result", ParentCode: "GqaPluginVote", Title: "投票结果分析", Icon: "dvr",
		Path: "/plugin-vote/vote/score-result", Component: "/Plugin/Vote/ScoreResult/index",
	},
}

func (s *sysMenu) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysMenu{}).Where("parent_code = ?", "GqaPluginVote").Or("name = ?", "GqaPluginVote").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_menu 表中vote插件菜单已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[GQA-Plugin] --> sys_menu 表中vote插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> vote插件初始数据进入 sys_menu 表成功！")
		global.GqaLog.Info("[GQA-Plugin] --> vote插件初始数据进入 sys_menu 表成功！")
		return nil
	})
}

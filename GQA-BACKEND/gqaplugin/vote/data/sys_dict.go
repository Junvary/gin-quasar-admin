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
		ParentCode: "voteType", DictCode: "dy", DictLabel: "民主评议党员投票",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteType", DictCode: "gl", DictLabel: "民主评议业务骨干及管理人员投票",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 702, Remark: "民主评议党员投票细类", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "voteTypeDetailDy", DictLabel: "民主评议党员投票细类",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetailDy", DictCode: "dy_1", DictLabel: "政治立场",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetailDy", DictCode: "dy_2", DictLabel: "服务群众",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetailDy", DictCode: "dy_3", DictLabel: "工作作风",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 4, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetailDy", DictCode: "dy_4", DictLabel: "组织纪律",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 5, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetailDy", DictCode: "dy_5", DictLabel: "履行职责",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 6, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetailDy", DictCode: "dy_p_zhenggong", DictLabel: "政工干事评价",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 7, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetailDy", DictCode: "dy_p_jijian", DictLabel: "纪检委员评价",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 703, Remark: "民主评议业务骨干及管理人员投票细类", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "voteTypeDetailGl", DictLabel: "民主评议业务骨干及管理人员投票细类",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetailGl", DictCode: "gl_1", DictLabel: "德",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetailGl", DictCode: "gl_2", DictLabel: "能",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetailGl", DictCode: "gl_3", DictLabel: "勤",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 4, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetailGl", DictCode: "gl_4", DictLabel: "绩",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 5, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteTypeDetailGl", DictCode: "gl_5", DictLabel: "廉",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 704, Remark: "民主评议党员投票权重", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "voteDyRatio", DictLabel: "民主评议党员投票权重",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteDyRatio", DictCode: "dy_zhenggong", DictLabel: "政工干事", DictExt1: "10",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteDyRatio", DictCode: "dy_jijian", DictLabel: "纪检委员", DictExt1: "10",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteDyRatio", DictCode: "dy_zhiweihui", DictLabel: "支委会成员", DictExt1: "40",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 4, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteDyRatio", DictCode: "dy_qita", DictLabel: "其他评议人员", DictExt1: "40", DictExt2: "dy_base",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 705, Remark: "民主评议业务骨干及管理人员投票权重", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "voteGlRatio", DictLabel: "民主评议业务骨干及管理人员投票权重",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteGlRatio", DictCode: "gl_zhigong", DictLabel: "职工代表/普通职工", DictExt1: "30", DictExt2: "gl_base",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteGlRatio", DictCode: "gl_keduanzhang", DictLabel: "科段长", DictExt1: "30",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "voteGlRatio", DictCode: "gl_banzi", DictLabel: "班子成员", DictExt1: "40",
	},


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

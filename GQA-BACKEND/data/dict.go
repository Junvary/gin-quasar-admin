package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysDict = new(sysDict)

type sysDict struct{}

func (s *sysDict) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysDict{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_dict 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLogger.Warn("[Gin-Quasar-Admin] --> sys_dict 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysDictData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_dict 表初始数据成功！")
		global.GqaLogger.Info("[Gin-Quasar-Admin] --> sys_dict 表初始数据成功！")
		return nil
	})
}

var sysDictData = []model.SysDict{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "性别字典",
	}}, DictCode: "gender", DictLabel: "性别"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "男",
	}}, DictCode: "m", DictLabel: "男", ParentCode: "gender"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "女",
	}}, DictCode: "f", DictLabel: "女", ParentCode: "gender"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 3, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "保密",
	}}, DictCode: "u", DictLabel: "保密", ParentCode: "gender"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "启用状态",
	}}, DictCode: "statusOnOff", DictLabel: "启用状态"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "启用",
	}}, DictCode: "on", DictLabel: "启用", ParentCode: "statusOnOff"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "禁用",
	}}, DictCode: "off", DictLabel: "禁用", ParentCode: "statusOnOff"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 3, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "是否状态",
	}}, DictCode: "statusYesNo", DictLabel: "是否状态"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "是",
	}}, DictCode: "yes", DictLabel: "是", ParentCode: "statusYesNo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "否",
	}}, DictCode: "no", DictLabel: "否", ParentCode: "statusYesNo"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 4, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "部门数据权限",
	}}, DictCode: "deptDataPermissionType", DictLabel: "部门数据权限"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "全部部门数据权限",
	}}, DictCode: "all", DictLabel: "全部部门数据权限", ParentCode: "deptDataPermissionType"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户本人数据权限",
	}}, DictCode: "user", DictLabel: "用户本人数据权限", ParentCode: "deptDataPermissionType"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 3, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户所在部门数据权限",
	}}, DictCode: "dept", DictLabel: "用户所在部门数据权限", ParentCode: "deptDataPermissionType"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 4, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户所在部门及子部门数据权限",
	}}, DictCode: "deptAndChildren", DictLabel: "用户所在部门及子部门数据权限", ParentCode: "deptDataPermissionType"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 5, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "自定义部门数据权限",
	}}, DictCode: "custom", DictLabel: "自定义部门数据权限", ParentCode: "deptDataPermissionType"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 5, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "消息类型",
	}}, DictCode: "noticeType", DictLabel: "消息类型"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "系统消息",
	}}, DictCode: "system", DictLabel: "系统消息", ParentCode: "noticeType"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "消息提示",
	}}, DictCode: "message", DictLabel: "消息提示", ParentCode: "noticeType"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 6, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "显示风格",
	}}, DictCode: "displayStyle", DictLabel: "显示风格"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "简约类型",
	}}, DictCode: "simple", DictLabel: "简约风格", ParentCode: "displayStyle"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "复杂",
	}}, DictCode: "complex", DictLabel: "复杂风格", ParentCode: "displayStyle"},
}

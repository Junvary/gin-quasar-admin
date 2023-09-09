package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var SysDict = new(sysDict)

type sysDict struct{}

func (s *sysDict) LoadData(c *gin.Context) error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysDict{}).Count(&count)
		if count != 0 {
			fmt.Println(utils.GqaI18nWithData(c, "SkipInsertWithData", "sys_dict"), count)
			global.GqaSLogger.Warn(utils.GqaI18nWithData(c, "SkipInsertWithData", "sys_dict"), "has_count", count)
			return nil
		}
		if err := tx.Create(&sysDictData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println(utils.GqaI18nWithData(c, "TableInitSuccess", "sys_dict"))
		global.GqaSLogger.Info(utils.GqaI18nWithData(c, "TableInitSuccess", "sys_dict"))
		return nil
	})
}

var sysDictData = []model.SysDict{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "性别字典",
	}}, DictCode: "gender", DictLabel: "性别"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "男",
	}}, DictCode: "gender_male", DictLabel: "Male", ParentCode: "gender"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "女",
	}}, DictCode: "gender_female", DictLabel: "Female", ParentCode: "gender"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 3, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "保密",
	}}, DictCode: "gender_unknown", DictLabel: "Unknown", ParentCode: "gender"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "启用状态",
	}}, DictCode: "onOff", DictLabel: "启用状态"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "启用",
	}}, DictCode: "onOff_on", DictLabel: "On", ParentCode: "onOff"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "禁用",
	}}, DictCode: "onOff_off", DictLabel: "Off", ParentCode: "onOff"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 3, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "是否状态",
	}}, DictCode: "yesNo", DictLabel: "是否状态"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "是",
	}}, DictCode: "yesNo_yes", DictLabel: "Yes", ParentCode: "yesNo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "否",
	}}, DictCode: "yesNo_no", DictLabel: "No", ParentCode: "yesNo"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 4, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "部门数据权限",
	}}, DictCode: "deptDataPermissionType", DictLabel: "部门数据权限"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "全部部门数据权限",
	}}, DictCode: "deptDataPermissionType_all", DictLabel: "全部部门数据权限", ParentCode: "deptDataPermissionType"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户本人数据权限",
	}}, DictCode: "deptDataPermissionType_user", DictLabel: "用户本人数据权限", ParentCode: "deptDataPermissionType"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 3, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户所在部门数据权限",
	}}, DictCode: "deptDataPermissionType_dept", DictLabel: "用户所在部门数据权限", ParentCode: "deptDataPermissionType"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 4, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户所在部门及子部门数据权限",
	}}, DictCode: "deptDataPermissionType_deptAndChildren", DictLabel: "用户所在部门及子部门数据权限", ParentCode: "deptDataPermissionType"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 5, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "自定义部门数据权限",
	}}, DictCode: "deptDataPermissionType_custom", DictLabel: "自定义部门数据权限", ParentCode: "deptDataPermissionType"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 5, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "消息类型",
	}}, DictCode: "noticeType", DictLabel: "消息类型"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "系统消息",
	}}, DictCode: "noticeType_system", DictLabel: "系统消息", ParentCode: "noticeType"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "消息提示",
	}}, DictCode: "noticeType_message", DictLabel: "消息提示", ParentCode: "noticeType"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 6, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "登录页风格",
	}}, DictCode: "loginLayoutStyle", DictLabel: "登录页风格"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "登录页",
	}}, DictCode: "loginLayoutStyle_login", DictLabel: "登录页", ParentCode: "loginLayoutStyle"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "门户页",
	}}, DictCode: "loginLayoutStyle_portal", DictLabel: "门户页", ParentCode: "loginLayoutStyle"},
}

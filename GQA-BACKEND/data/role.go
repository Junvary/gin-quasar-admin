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

var SysRole = new(sysRole)

type sysRole struct{}

func (s *sysRole) LoadData(c *gin.Context) error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysRole{}).Count(&count)
		if count != 0 {
			fmt.Println(utils.GqaI18nWithData(c, "SkipInsertWithData", "sys_role"), count)
			global.GqaSLogger.Warn(utils.GqaI18nWithData(c, "SkipInsertWithData", "sys_role"), "has_count", count)
			return nil
		}
		if err := tx.Create(&sysRoleData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println(utils.GqaI18nWithData(c, "TableInitSuccess", "sys_role"))
		global.GqaSLogger.Info(utils.GqaI18nWithData(c, "TableInitSuccess", "sys_role"))
		return nil
	})
}

var sysRoleData = []model.SysRole{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是超级管理员组，拥有所有权限，请不要编辑！",
	}}, RoleCode: "super-admin", RoleName: "超级管理员组", DeptDataPermissionType: "deptDataPermissionType_all"},
}

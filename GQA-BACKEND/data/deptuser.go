package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var SysDeptUser = new(sysDeptUser)

type sysDeptUser struct{}

func (s *sysDeptUser) LoadData(c *gin.Context) error {
	return global.GqaDb.Table("sys_dept_user").Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysDeptUser{}).Count(&count)
		if count != 0 {
			fmt.Println(utils.GqaI18nWithData(c, "SkipInsertWithData", "sys_dept_user"), count)
			global.GqaSLogger.Warn(utils.GqaI18nWithData(c, "SkipInsertWithData", "sys_dept_user"), "has_count", count)
			return nil
		}
		if err := tx.Create(&sysDeptUserData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println(utils.GqaI18nWithData(c, "TableInitSuccess", "sys_dept_user"))
		global.GqaSLogger.Info(utils.GqaI18nWithData(c, "TableInitSuccess", "sys_dept_user"))
		return nil
	})
}

var sysDeptUserData = []model.SysDeptUser{
	{"gin-quasar-admin", "admin"},
}

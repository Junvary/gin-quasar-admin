package datas

import (
	"gin-quasar-admin/global"
	"github.com/gookit/color"
	"os"
	"time"

	"gin-quasar-admin/model"
	"gorm.io/gorm"
)

func InitSysDictionaryDetail(db *gorm.DB) {
	status := new(bool)
	*status = true
	DictionaryDetail := []model.SysDictionaryDetail{
		{global.GqaModel{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "smallint", 1, status, 1, 2},
		{global.GqaModel{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumint", 2, status, 2, 2},
		{global.GqaModel{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "int", 3, status, 3, 2},
		{global.GqaModel{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "bigint", 4, status, 4, 2},
		{global.GqaModel{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "date", 0, status, 0, 3},
		{global.GqaModel{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "time", 1, status, 1, 3},
		{global.GqaModel{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "year", 2, status, 2, 3},
		{global.GqaModel{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "datetime", 3, status, 3, 3},
		{global.GqaModel{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "timestamp", 5, status, 5, 3},
		{global.GqaModel{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "float", 0, status, 0, 4},
		{global.GqaModel{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "double", 1, status, 1, 4},
		{global.GqaModel{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "decimal", 2, status, 2, 4},
		{global.GqaModel{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "char", 0, status, 0, 5},
		{global.GqaModel{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "varchar", 1, status, 1, 5},
		{global.GqaModel{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinyblob", 2, status, 2, 5},
		{global.GqaModel{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinytext", 3, status, 3, 5},
		{global.GqaModel{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "text", 4, status, 4, 5},
		{global.GqaModel{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "blob", 5, status, 5, 5},
		{global.GqaModel{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumblob", 6, status, 6, 5},
		{global.GqaModel{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumtext", 7, status, 7, 5},
		{global.GqaModel{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "longblob", 8, status, 8, 5},
		{global.GqaModel{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "longtext", 9, status, 9, 5},
		{global.GqaModel{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinyint", 0, status, 0, 6},
	}
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 23}).Find(&[]model.SysDictionaryDetail{}).RowsAffected == 2 {
			color.Danger.Println("sys_dictionary_details???????????????????????????!")
			return nil
		}
		if err := tx.Create(&DictionaryDetail).Error; err != nil { // ???????????????????????????
			return err
		}
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> sys_dictionary_details ????????????????????????,err: %v\n", err)
		os.Exit(0)
	}
}

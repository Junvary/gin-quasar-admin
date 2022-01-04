package public_service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
)

func GetWeaponLanguage() (err error, language []interface{}) {
	var languageList []system.SysDict
	if err = global.GqaDb.Where("parent_code = ?", "codeLanguage").Find(&languageList).Error; err !=nil{
		return err, nil
	}
	type record struct {
		Value int64 `json:"value"`
		Name string `json:"name"`
	}
	for _, v := range languageList{
		var total int64
		projectDb := global.GqaDb.Model(&model.GqaPluginXkProject{})
		projectDb.Where("language like ?", "%" + v.DictCode + "%").Count(&total)
		if total != 0{
			language = append(language, record{
				Value: total,
				Name: v.DictCode,
			})
		}
	}
	return nil, language
}

func GetWeaponNode() (err error, node []interface{}) {
	var nodeList []system.SysDict
	if err = global.GqaDb.Where("parent_code = ?", "projectNode").Find(&nodeList).Error; err !=nil{
		return err, nil
	}
	type record struct {
		Value int64 `json:"value"`
		Name string `json:"name"`
	}
	for _, v := range nodeList{
		var total int64
		projectDb := global.GqaDb.Model(&model.GqaPluginXkProject{})
		projectDb.Where("node like ?", "%" + v.DictCode + "%").Count(&total)
		if total != 0{
			node = append(node, record{
				Value: total,
				Name: v.DictLabel,
			})
		}
	}
	return nil, node
}

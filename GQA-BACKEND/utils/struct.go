package utils

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"reflect"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func GlobalModelToMap(model *global.GqaModel) map[string]interface{} {
	var data = make(map[string]interface{})
	data["memo"] = model.Memo
	data["updated_by"] = model.UpdatedBy
	return data
}

package utils

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

func GlobalModelToMap(model *global.GqaModel) map[string]interface{} {
	var data = make(map[string]interface{})
	data["sort"] = model.Sort
	data["status"] = model.Status
	data["remark"] = model.Remark
	data["updated_by"] = model.UpdatedBy
	return data
}

func MergeMap(mObj ...map[string]interface{}) map[string]interface{} {
	newObj := map[string]interface{}{}
	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}

package system

import "gin-quasar-admin/global"

type RequestQueryById struct {
	Id uint `json:"id"`
}

type RequestPageByParentId struct {
	global.RequestPage
	ParentId uint   `json:"parentId" form:"parentId"`
}

type RequestQueryByValue struct {
	Value string `json:"value"`
}

type ResponsePage struct {
	Records  interface{} `json:"records"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type ResponsePageWithParentId struct {
	Records  interface{} `json:"records"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	ParentId uint        `json:"parentId"`
}

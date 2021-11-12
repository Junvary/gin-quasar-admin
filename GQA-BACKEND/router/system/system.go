package system

import "gin-quasar-admin/api"

type RouterSystem struct {
	RouterMenu
	RouterUser
	RouterRole
	RouterDept
	RouterDict
	RouterApi
	RouterUpload
	RouterConfig
}

var ApiSystem = api.GroupApiApp.ApiSystem
var ApiMenu = ApiSystem.ApiMenu
var ApiUser = ApiSystem.ApiUser
var ApiRole = ApiSystem.ApiRole
var ApiDept = ApiSystem.ApiDept
var ApiDict = ApiSystem.ApiDict
var ApiApi = ApiSystem.ApiApi
var ApiUpload = ApiSystem.ApiUpload
var ApiConfig = ApiSystem.ApiConfig

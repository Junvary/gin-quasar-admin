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
	RouterConfigBackend
	RouterConfigFrontend
	RouterLog
}

var ApiSystem = api.GroupApiApp.ApiSystem
var ApiMenu = ApiSystem.ApiMenu
var ApiUser = ApiSystem.ApiUser
var ApiRole = ApiSystem.ApiRole
var ApiDept = ApiSystem.ApiDept
var ApiDict = ApiSystem.ApiDict
var ApiApi = ApiSystem.ApiApi
var ApiUpload = ApiSystem.ApiUpload
var ApiConfigBackend = ApiSystem.ApiConfigBackend
var ApiConfigFrontend = ApiSystem.ApiConfigFrontend
var ApiLogLogin = ApiSystem.ApiLogLogin
var ApiLogOperation = ApiSystem.ApiLogOperation

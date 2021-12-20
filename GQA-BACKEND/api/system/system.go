package system

import "gin-quasar-admin/service"

type ApiSystem struct {
	ApiMenu
	ApiUser
	ApiRole
	ApiDept
	ApiDict
	ApiApi
	ApiUpload
	ApiConfigBackend
	ApiConfigFrontend
	ApiLogLogin
}

var ServiceSystem = service.GroupServiceApp.ServiceSystem
var ServiceMenu = ServiceSystem.ServiceMenu
var ServiceUser = ServiceSystem.ServiceUser
var ServiceRole = ServiceSystem.ServiceRole
var ServiceDept = ServiceSystem.ServiceDept
var ServiceDict = ServiceSystem.ServiceDict
var ServiceApi = ServiceSystem.ServiceApi
var ServiceUpload = ServiceSystem.ServiceUpload
var ServiceConfigBackend = ServiceSystem.ServiceConfigBackend
var ServiceConfigFrontend = ServiceSystem.ServiceConfigFrontend
var ServiceLogLogin = ServiceSystem.ServiceLogLogin

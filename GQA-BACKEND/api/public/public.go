package public

import "gin-quasar-admin/service"

type ApiPublic struct {
	ApiCheckAndInitDb
	ApiCaptcha
	ApiLogin
	ApiGetDict
	ApiGetFrontend
}

var ServicePublic = service.GroupServiceApp.ServicePublic
var ServiceCheckAndInitDb = ServicePublic.ServiceCheckAndInitDb
var ServiceLogin = ServicePublic.ServiceLogin
var ServiceGetDict = ServicePublic.ServiceGetDict
var ServiceGetFrontend = ServicePublic.ServiceGetFrontend

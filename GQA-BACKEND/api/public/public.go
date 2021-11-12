package public

import "gin-quasar-admin/service"

type ApiPublic struct {
	ApiCheckAndInitDb
	ApiCaptcha
	ApiLogin
	ApiDictDetail
}

var ServicePublic = service.GroupServiceApp.ServicePublic
var ServiceCheckAndInitDb = ServicePublic.ServiceCheckAndInitDb
var ServiceServiceDictDetail = ServicePublic.ServiceDictDetail
var ServiceLogin = ServicePublic.ServiceLogin

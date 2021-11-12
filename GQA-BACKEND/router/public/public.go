package public

import "gin-quasar-admin/api"

type RouterPublic struct {
	RouterCheckDb
	RouterCaptcha
	RouterLogin
	RouterDictDetail
}

var ApiPublic = api.GroupApiApp.ApiPublic
var ApiCaptcha = ApiPublic.ApiCaptcha
var ApiCheckAndInitDb = ApiPublic.ApiCheckAndInitDb
var ApiDictDetail = ApiPublic.ApiDictDetail
var ApiLogin = ApiPublic.ApiLogin

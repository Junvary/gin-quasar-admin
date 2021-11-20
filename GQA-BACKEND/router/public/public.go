package public

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/api"

type RouterPublic struct {
	RouterCheckDb
	RouterCaptcha
	RouterLogin
	RouterGetDict
	RouterGetFrontend
}

var ApiPublic = api.GroupApiApp.ApiPublic
var ApiCaptcha = ApiPublic.ApiCaptcha
var ApiCheckAndInitDb = ApiPublic.ApiCheckAndInitDb
var ApiLogin = ApiPublic.ApiLogin
var ApiGetDict = ApiPublic.ApiGetDict
var ApiGetFrontend = ApiPublic.ApiGetFrontend

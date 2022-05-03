package public

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/api"

type RouterPublic struct {
	RouterCaptcha        RouterCaptcha
	RouterConfigBackend  RouterConfigBackend
	RouterConfigFrontend RouterConfigFrontend
	RouterDb             RouterDb
	RouterDict           RouterDict
	RouterLogin          RouterLogin
	RouterWebsocket      RouterWebsocket
}

var apiPublic = api.GqaApi.ApiPublic

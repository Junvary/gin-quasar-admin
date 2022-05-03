package public

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service"

type ApiPublic struct {
	ApiCaptcha        ApiCaptcha
	ApiConfigBackend  ApiConfigBackend
	ApiConfigFrontend ApiConfigFrontend
	ApiDb             ApiDb
	ApiDict           ApiDict
	ApiLogin          ApiLogin
	ApiWebSocket      ApiWebSocket
}

var servicePublic = service.GqaService.ServicePublic

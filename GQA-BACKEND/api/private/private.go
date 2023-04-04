package private

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service"

type ApiPrivate struct {
	ApiApi            ApiApi
	ApiConfigBackend  ApiConfigBackend
	ApiConfigFrontend ApiConfigFrontend
	ApiCron           ApiCron
	ApiDept           ApiDept
	ApiDict           ApiDict
	ApiGenPlugin      ApiGenPlugin
	ApiLogLogin       ApiLogLogin
	ApiLogOperation   ApiLogOperation
	ApiMenu           ApiMenu
	ApiTodo           ApiTodo
	ApiNotice         ApiNotice
	ApiRole           ApiRole
	ApiUpload         ApiUpload
	ApiUser           ApiUser
	ApiUserOnline     ApiUserOnline
}

var servicePrivate = service.GqaService.ServicePrivate

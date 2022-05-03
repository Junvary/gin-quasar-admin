package private

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service"

type ApiPrivate struct {
	ApiApi            ApiApi
	ApiConfigBackend  ApiConfigBackend
	ApiConfigFrontend ApiConfigFrontend
	ApiDept           ApiDept
	ApiDict           ApiDict
	ApiGenPlugin      ApiGenPlugin
	ApiLogLogin       ApiLogLogin
	ApiLogOperation   ApiLogOperation
	ApiMenu           ApiMenu
	ApiNoteTodo       ApiNoteTodo
	ApiNotice         ApiNotice
	ApiRole           ApiRole
	ApiUpload         ApiUpload
	ApiUser           ApiUser
}

var servicePrivate = service.GqaService.ServicePrivate

package private

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/api"

type RouterPrivate struct {
	RouterApi            RouterApi
	RouterConfigBackend  RouterConfigBackend
	RouterConfigFrontend RouterConfigFrontend
	RouterCron           RouterCron
	RouterDept           RouterDept
	RouterDict           RouterDict
	RouterGenPlugin      RouterGenPlugin
	RouterLog            RouterLog
	RouterMenu           RouterMenu
	RouterTodo           RouterTodo
	RouterNotice         RouterNotice
	RouterRole           RouterRole
	RouterUpload         RouterUpload
	RouterUser           RouterUser
	RouterUserOnline     RouterUserOnline
}

var apiPrivate = api.GqaApi.ApiPrivate

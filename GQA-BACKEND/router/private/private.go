package private

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/api"

type RouterPrivate struct {
	RouterApi            RouterApi
	RouterConfigBackend  RouterConfigBackend
	RouterConfigFrontend RouterConfigFrontend
	RouterDept           RouterDept
	RouterDict           RouterDict
	RouterLog            RouterLog
	RouterMenu           RouterMenu
	RouterNoteTodo       RouterNoteTodo
	RouterNotice         RouterNotice
	RouterRole           RouterRole
	RouterUpload         RouterUpload
	RouterUser           RouterUser
}

var apiPrivate = api.GqaApi.ApiPrivate

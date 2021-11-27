package api

import (
	"gin-quasar-admin/api/public"
	"gin-quasar-admin/api/system"
)

type GroupApi struct {
	ApiPublic public.ApiPublic
	ApiSystem system.ApiSystem
}

var GroupApiApp = new(GroupApi)

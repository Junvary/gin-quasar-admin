package api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/api/public"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/api/system"
)

type GroupApi struct {
	ApiPublic public.ApiPublic
	ApiSystem system.ApiSystem
}

var GroupApiApp = new(GroupApi)

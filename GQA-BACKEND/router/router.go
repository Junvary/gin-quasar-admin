package router

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/router/public"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/router/system"
)

type GroupRouter struct {
	RouterPublic public.RouterPublic
	RouterSystem system.RouterSystem
}

var GroupRouterApp = new(GroupRouter)

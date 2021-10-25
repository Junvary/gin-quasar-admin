package router

import (
	"gin-quasar-admin/router/public"
	"gin-quasar-admin/router/system"
)

type GroupRouter struct {
	RouterPublic public.RouterPublic
	RouterSystem system.RouterSystem
}

var GroupRouterApp = new(GroupRouter)

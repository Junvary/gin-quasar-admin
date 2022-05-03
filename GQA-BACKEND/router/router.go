package router

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/router/private"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/router/public"
)

var GqaRouter = new(struct {
	RouterPublic  public.RouterPublic
	RouterPrivate private.RouterPrivate
})

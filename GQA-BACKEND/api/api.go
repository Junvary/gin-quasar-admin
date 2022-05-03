package api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/api/private"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/api/public"
)

var GqaApi = new(struct {
	ApiPublic  public.ApiPublic
	ApiPrivate private.ApiPrivate
})

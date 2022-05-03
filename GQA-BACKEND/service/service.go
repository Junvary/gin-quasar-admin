package service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/private"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/public"
)

var GqaService = new(struct {
	ServicePublic  public.ServicePublic
	ServicePrivate private.ServicePrivate
})

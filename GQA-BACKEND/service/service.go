package service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/public"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/system"
)

type GroupService struct {
	public.ServicePublic
	system.ServiceSystem
}

var GroupServiceApp = new(GroupService)

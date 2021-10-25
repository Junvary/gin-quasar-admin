package service

import (
	"gin-quasar-admin/service/public"
	"gin-quasar-admin/service/system"
)

type GroupService struct {
	public.ServicePublic
	system.ServiceSystem
}

var GroupServiceApp = new(GroupService)

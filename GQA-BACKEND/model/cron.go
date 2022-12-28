package model

import (
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
)

type SysCron struct {
	UUID uuid.UUID    `json:"uuid"`
	Id   cron.EntryID `json:"id"`
	Name string       `json:"name"`
	Spec string       `json:"spec"`
}

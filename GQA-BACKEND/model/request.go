package model

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestPage struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"pageSize"`
}

type RequestSort struct {
	SortBy string `json:"sort_by" form:"sortBy"`
	Desc   bool   `json:"desc" form:"desc"`
}

type RequestPageAndSort struct {
	RequestPage
	RequestSort
}

type RequestAdd struct {
	Sort   uint   `json:"sort"`
	Status string `json:"status"`
	Memo   string `json:"memo"`
}

type RequestQueryById struct {
	Id uint `json:"id"`
}

type RequestQueryByUUID struct {
	UUID uuid.UUID `json:"uuid"`
}

type RequestQueryByUsername struct {
	Username string `json:"username"`
}

func RequestShouldBindJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		ResponseErrorMessageWithLog(global.GqaConfig.System.BindError+err.Error(), c)
		return err
	}
	return nil
}

package global

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/config"
	"github.com/casbin/casbin/v2"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

var Store = base64Captcha.DefaultMemStore

var (
	GqaConfig config.Config
	GqaViper  *viper.Viper
	GqaLog    *zap.Logger
	GqaDb     *gorm.DB
	GqaCasbin *casbin.Enforcer
)

type GqaModel struct {
	Id        uint           `json:"id" gorm:"comment:id;primarykey;autoIncrement;"`
	Stable    string         `json:"stable" gorm:"comment:系统内置;default:no;"`
	Status    string         `json:"status" gorm:"comment:状态;default:on;"`
	Sort      uint           `json:"sort" gorm:"comment:排序;default:1;"`
	Remark    string         `json:"remark" gorm:"comment:备注描述;type:text;"`
	CreatedAt time.Time      `json:"createdAt"`
	CreatedBy string         `json:"createdBy"`
	UpdatedAt time.Time      `json:"updatedAt"`
	UpdatedBy string         `json:"updatedBy"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func OrderByColumn(sortBy string, desc bool) interface{} {
	return clause.OrderByColumn{Column: clause.Column{Name: sortBy}, Desc: desc}
}

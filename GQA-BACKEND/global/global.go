package global

import (
	"gin-quasar-admin/config"
	"github.com/casbin/casbin/v2"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

var Store = base64Captcha.DefaultMemStore

var (
	GqaConfig       config.Config
	GqaViper        *viper.Viper
	GqaLog          *zap.Logger
	GqaDb           *gorm.DB
	GqaSingleFlight = &singleflight.Group{}
	GqaCasbin       *casbin.Enforcer
)

type GqaModel struct {
	Id       uint           `json:"id" gorm:"comment:id;primarykey;autoIncrement;"`
	Stable   string         `json:"stable" gorm:"comment:系统内置;"`
	Status   string         `json:"status" gorm:"comment:状态;default:on;"`
	Sort     uint           `json:"sort" gorm:"comment:排序;default:1;"`
	Remark   string         `json:"remark" gorm:"comment:备注描述;type:text;"`
	CreateAt time.Time      `json:"createAt"`
	CreateBy string         `json:"createBy"`
	UpdateAt time.Time      `json:"updateAt"`
	UpdateBy string         `json:"updateBy"`
	DeleteAt gorm.DeletedAt `json:"-" gorm:"index"`
	DeleteBy string         `json:"-"`
}

func OrderByColumn(sortBy string, desc bool) interface{} {
	return clause.OrderByColumn{Column: clause.Column{Name: sortBy}, Desc: desc}
}

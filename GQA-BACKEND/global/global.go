package global

import (
	"gin-quasar-admin/config"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"time"
)

var Store = base64Captcha.DefaultMemStore

var (
	GqaConfig       config.Config
	GqaViper        *viper.Viper
	GqaLog          *zap.Logger
	GqaDb           *gorm.DB
	GqaSingleFlight = &singleflight.Group{}
)

type GqaModel struct {
	Id       uint           `json:"id" gorm:"comment:id;primarykey;"`
	Status   string         `json:"status" gorm:"comment:状态;default:on;"`
	Sort     int            `json:"sort" gorm:"comment:排序;default:1;"`
	Desc     string         `json:"desc" gorm:"comment:备注描述;type:text;"`
	CreateAt time.Time      `json:"createAt"`
	CreateBy string         `json:"createBy"`
	UpdateAt time.Time      `json:"updateAt"`
	UpdateBy string         `json:"updateBy"`
	DeleteAt gorm.DeletedAt `json:"-" gorm:"index"`
	DeleteBy string         `json:"-"`
}


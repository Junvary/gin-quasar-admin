package global

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/config"
	"github.com/mojocn/base64Captcha"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log/slog"
	"runtime"
	"time"
)

var Store = base64Captcha.DefaultMemStore

var (
	GqaConfig      config.Config
	GqaViper       *viper.Viper
	GqaSLogger     *slog.Logger
	GqaDb          *gorm.DB
	GqaCron        *cron.Cron
	GqaServeUpload = "upload"
	GqaServeAvatar = "avatar"
	GqaServeFile   = "file"
	GqaServeLogo   = "logo"
	GqaServeBanner = "banner"
	GqaOsType      = runtime.GOOS
)

type GqaModel struct {
	Id        uint           `json:"id" gorm:"comment:id;primarykey;autoIncrement;"`
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy string         `json:"created_by" gorm:"index;"`
	UpdatedAt time.Time      `json:"updated_at"`
	UpdatedBy string         `json:"updated_by"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Sort      uint           `json:"sort" gorm:"comment:排序;default:1;"`
	Stable    string         `json:"stable" gorm:"comment:系统内置;default:yesNo_no;type:varchar(10);"`
	Status    string         `json:"status" gorm:"comment:状态;default:onOff_on;type:varchar(10);"`
	Memo      string         `json:"memo" gorm:"comment:备注描述;type:text;"`
}

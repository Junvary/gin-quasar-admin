package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/middleware"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"time"
)

type ServiceUserOnline struct{}

func (s *ServiceUserOnline) GetUserOnlineList(requestUserOnlineList model.RequestGetUserOnlineList) (err error, user interface{}, total int64) {
	_ = s.CheckUserOnline()
	pageSize := requestUserOnlineList.PageSize
	offset := requestUserOnlineList.PageSize * (requestUserOnlineList.Page - 1)
	db := global.GqaDb.Model(&model.SysUserOnline{})
	var userList []model.SysUserOnline
	// Search
	if requestUserOnlineList.Username != "" {
		db = db.Where("username like ?", "%"+requestUserOnlineList.Username+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(requestUserOnlineList.SortBy, requestUserOnlineList.Desc)).
		Preload("User").Find(&userList).Error

	return err, userList, total
}

func (s *ServiceUserOnline) CheckUserOnline() error {
	var userOnlineList []model.SysUserOnline
	err := global.GqaDb.Model(&model.SysUserOnline{}).Error
	for _, v := range userOnlineList {
		token := middleware.Jwt{
			SigningKey: []byte(v.Token),
		}
		claims, err := token.ParseToken(v.Token)
		if err != nil {
			if err.Error() == "checkRefresh" {
				if claims.RefreshAt < time.Now().Unix() {
					_ = s.KickOnlineUserByUsername(v.Username)
				}
			}
		}
	}
	return err
}

func (s *ServiceUserOnline) KickOnlineUserByUsername(username string) (err error) {
	var sysUserOnline model.SysUserOnline
	err = global.GqaDb.Where("username = ?", username).Delete(&sysUserOnline).Error
	return err
}

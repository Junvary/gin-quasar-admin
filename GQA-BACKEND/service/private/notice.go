package private

import (
	"encoding/json"
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServiceNotice struct{}

func (s *ServiceNotice) GetNoticeList(requestNoticeList model.RequestGetNoticeList) (err error, notice interface{}, total int64) {
	pageSize := requestNoticeList.PageSize
	offset := requestNoticeList.PageSize * (requestNoticeList.Page - 1)
	var db *gorm.DB
	var noticeList []model.SysNotice
	//配置搜索
	if requestNoticeList.NoticeToUser != "" {
		//通知栏调用时候的搜索内容
		//先去SysNoticeToUser表中查找to_user是这个人，并且未读的内容
		var noticeToUser []model.SysNoticeToUser
		if requestNoticeList.NoticeRead != "" {
			//顶部消息通知图标会传递NoticeRead
			if err = global.GqaDb.Where("to_user = ? and user_read = ?", requestNoticeList.NoticeToUser, requestNoticeList.NoticeRead).
				Find(&noticeToUser).Error; err != nil {
				return err, noticeList, 0
			}
		} else {
			//个人中心不会传递NoticeRead
			if err = global.GqaDb.Where("to_user = ? ", requestNoticeList.NoticeToUser).Find(&noticeToUser).Error; err != nil {
				return err, noticeList, 0
			}
		}
		//取出noticeId，组合notice_to_user_type为all的内容返回
		var noticeIdList []uuid.UUID
		for _, v := range noticeToUser {
			noticeIdList = append(noticeIdList, v.NoticeId)
		}
		db = global.GqaDb.Where("notice_id in ?", noticeIdList).Model(&model.SysNotice{})
	} else {
		//管理员查询
		db = global.GqaDb.Model(&model.SysNotice{})
	}
	if requestNoticeList.NoticeTitle != "" {
		db = db.Where("notice_title like ?", "%"+requestNoticeList.NoticeTitle+"%")
	}
	if requestNoticeList.NoticeType != "" {
		db = db.Where("notice_type = ?", requestNoticeList.NoticeType)
	}
	if requestNoticeList.NoticeSent != "" {
		db = db.Where("notice_sent = ?", requestNoticeList.NoticeSent)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(requestNoticeList.SortBy, requestNoticeList.Desc)).
		Preload("NoticeToUser").
		Find(&noticeList).Error
	return err, noticeList, total
}

func (s *ServiceNotice) AddNotice(toAddNotice model.RequestAddNotice, username string) (err error) {
	noticeId := uuid.New()
	var noticeToUser []model.SysNoticeToUser
	if toAddNotice.NoticeToUserType == "some" {
		//发送给部分用户，那么循环发来的NoticeToUser字段添加内容到noticeToUser
		for _, v := range toAddNotice.NoticeToUser {
			noticeToUser = append(noticeToUser, model.SysNoticeToUser{
				NoticeId: noticeId,
				ToUser:   v,
			})
		}
	} else if toAddNotice.NoticeToUserType == "all" {
		var allUserList []model.SysUser
		if err = global.GqaDb.Find(&allUserList).Error; err != nil {
			return err
		}
		for _, v := range allUserList {
			noticeToUser = append(noticeToUser, model.SysNoticeToUser{
				NoticeId: noticeId,
				ToUser:   v.Username,
			})
		}
	} else {
		return errors.New("未获取到[发送给]内容")
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: username,
		},
	}
	addNotice := &model.SysNotice{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		NoticeId:                          noticeId,
		NoticeTitle:                       toAddNotice.NoticeTitle,
		NoticeContent:                     toAddNotice.NoticeContent,
		NoticeType:                        toAddNotice.NoticeType,
		NoticeToUserType:                  toAddNotice.NoticeToUserType,
		NoticeToUser:                      noticeToUser,
	}
	err = global.GqaDb.Create(&addNotice).Error
	return err
}

func (s *ServiceNotice) DeleteNoticeById(id uint) (err error) {
	var sysNotice model.SysNotice
	if err = global.GqaDb.Where("id = ?", id).First(&sysNotice).Error; err != nil {
		return err
	}
	err = global.GqaDb.Select("NoticeToUser").Where("id = ?", id).Unscoped().Delete(&sysNotice).Error
	return err
}

func (s *ServiceNotice) QueryNoticeById(id uint) (err error, noticeInfo model.SysNotice) {
	var notice model.SysNotice
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").Preload("NoticeToUser").
		First(&notice, "id = ?", id).Error
	return err, notice
}

func (s *ServiceNotice) QueryNoticeReadById(id uint, username string) (err error, noticeInfo model.SysNotice) {
	var sysNotice model.SysNotice
	if err = global.GqaDb.Where("id = ?", id).First(&sysNotice).Error; err != nil {
		return err, sysNotice
	}
	if err = global.GqaDb.Model(&model.SysNoticeToUser{}).
		Where("notice_id = ? and to_user = ?", sysNotice.NoticeId, username).
		Update("user_read", "yesNo_yes").Error; err != nil {
		return err, sysNotice
	}
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").Preload("NoticeToUser").
		First(&sysNotice, "id = ?", id).Error
	return err, sysNotice
}

func (s *ServiceNotice) SendNotice(toSendNotice model.SysNotice) (err error) {
	var sysNotice model.SysNotice
	if err = global.GqaDb.Where("id = ?", toSendNotice.Id).First(&sysNotice).Error; err != nil {
		return err
	}
	if sysNotice.NoticeSent == "yesNo_yes" {
		return errors.New("这条消息已被发送过！")
	}
	//还没发送的，就置为发送
	toSendNotice.NoticeSent = "yesNo_yes"
	//发送字段同步到表
	if err = global.GqaDb.Omit("NoticeToUser").Updates(&toSendNotice).Error; err != nil {
		return err
	}
	if toSendNotice.NoticeToUserType == "all" {
		//如果对全体发送notice，ToUser保持为空，这里不填
		var byteMessage, _ = json.Marshal(model.WsMessage{
			Text:              toSendNotice.NoticeTitle,
			MessageToUserType: "all",
			MessageType:       "notice",
		})
		model.BroadcastMsg <- byteMessage
		return nil
	} else if toSendNotice.NoticeToUserType == "some" {
		var userList []string
		for _, v := range toSendNotice.NoticeToUser {
			userList = append(userList, v.ToUser)
		}
		//如果对部分用户发送notice，ToUser字段加入NoticeToUser内容，让websocket判断
		var byteMessage, _ = json.Marshal(model.WsMessage{
			Text:              toSendNotice.NoticeTitle,
			MessageToUserType: "some",
			MessageType:       "notice",
			ToUser:            userList,
		})
		model.BroadcastMsg <- byteMessage
		return nil
	} else {
		return nil
	}
}

package utils

import (
	"encoding/json"
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/google/uuid"
)

func CreateAndSendNotice(noticeTitle string, noticeContent string, noticeType string, noticeToUserType string, toUser []string, username string) (err error) {
	toAddNotice := model.RequestAddNotice{
		NoticeTitle:      noticeTitle,
		NoticeContent:    noticeContent,
		NoticeType:       noticeType,
		NoticeToUserType: noticeToUserType,
		NoticeToUser:     toUser,
	}
	noticeId := uuid.New()
	var noticeToUser []model.SysNoticeToUser
	if toAddNotice.NoticeToUserType == "some" {
		// If it is sent to some users, the circularly sent NoticeToUser field adds content to the noticeToUser
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
	var sysNotice model.SysNotice
	if err = global.GqaDb.Where("id = ?", addNotice.Id).First(&sysNotice).Error; err != nil {
		return err
	}
	if sysNotice.NoticeSent == "yesNo_yes" {
		return errors.New("这条消息已被发送过！")
	}
	// If it has not been sent, it will be set as sending
	addNotice.NoticeSent = "yesNo_yes"
	//发送字段同步到表
	if err = global.GqaDb.Omit("NoticeToUser").Updates(&addNotice).Error; err != nil {
		return err
	}
	if addNotice.NoticeToUserType == "all" {
		// If a notice is sent to all, ToUser remains blank. It is not filled in here
		var byteMessage, _ = json.Marshal(model.WsMessage{
			Text:              addNotice.NoticeTitle,
			MessageToUserType: addNotice.NoticeToUserType,
			MessageType:       toAddNotice.NoticeType,
		})
		model.BroadcastMsg <- byteMessage
		return nil
	} else if addNotice.NoticeToUserType == "some" {
		var userList []string
		for _, v := range addNotice.NoticeToUser {
			userList = append(userList, v.ToUser)
		}
		// If you send a notice to some users, add the NoticeToUser content to the ToUser field and let websocket judge
		var byteMessage, _ = json.Marshal(model.WsMessage{
			Text:              addNotice.NoticeTitle,
			MessageToUserType: addNotice.NoticeToUserType,
			MessageType:       toAddNotice.NoticeType,
			ToUser:            userList,
		})
		model.BroadcastMsg <- byteMessage
		return nil
	} else {
		return nil
	}
}

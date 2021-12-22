package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
)

type ServiceNotice struct {}

func (s *ServiceNotice) GetNoticeList(requestNoticeList system.RequestNoticeList) (err error, notice interface{}, total int64) {
	pageSize := requestNoticeList.PageSize
	offset := requestNoticeList.PageSize * (requestNoticeList.Page - 1)
	db := global.GqaDb.Model(&system.SysNotice{})
	var noticeList []system.SysNotice
	//配置搜索
	if requestNoticeList.NoticeTitle != ""{
		db = db.Where("notice_title like ?", "%" + requestNoticeList.NoticeTitle + "%")
	}
	if requestNoticeList.NoticeType != ""{
		db = db.Where("notice_type = ?", requestNoticeList.NoticeType)
	}
	if requestNoticeList.NoticeRead != ""{
		db = db.Where("notice_read = ?", requestNoticeList.NoticeRead)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(requestNoticeList.SortBy, requestNoticeList.Desc)).
		Find(&noticeList).Error
	return err, noticeList, total
}

func (s *ServiceNotice) EditNotice(toEditNotice system.SysNotice) (err error) {
	var sysNotice system.SysNotice
	if err = global.GqaDb.Where("id = ?", toEditNotice.Id).First(&sysNotice).Error; err != nil {
		return err
	}
	err = global.GqaDb.Updates(&toEditNotice).Error
	return err
}

func (s *ServiceNotice) AddNotice(toAddNotice system.SysNotice) (err error) {
	err = global.GqaDb.Create(&toAddNotice).Error
	return err
}

func (s *ServiceNotice) DeleteNotice(id uint) (err error) {
	var sysNotice system.SysNotice
	if err = global.GqaDb.Where("id = ?", id).First(&sysNotice).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysNotice).Error
	return err
}

func (s *ServiceNotice) QueryNoticeById(id uint) (err error, noticeInfo system.SysNotice) {
	var notice system.SysNotice
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&notice, "id = ?", id).Error
	return err, notice
}

func (s *ServiceNotice) QueryNoticeByIdRead(id uint) (err error, noticeInfo system.SysNotice) {
	var sysNotice system.SysNotice
	if err = global.GqaDb.Where("id = ?", id).First(&sysNotice).Error; err != nil {
		return err, sysNotice
	}
	err = global.GqaDb.Updates(&sysNotice).Error
	var notice system.SysNotice
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&notice, "id = ?", id).Error
	return err, notice
}

func (s *ServiceNotice) SendNotice(toSendNotice system.SysNotice) (err error) {
	var sysNotice system.SysNotice
	if err = global.GqaDb.Where("id = ?", toSendNotice.Id).First(&sysNotice).Error; err != nil {
		return err
	}
	toSendNotice.NoticeSent = "yes"
	err = global.GqaDb.Updates(&toSendNotice).Error
	return err
}


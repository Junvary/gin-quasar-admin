package private_service

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"gorm.io/gorm"
)

func VoterList(getVoterList model.RequestVoterList, username string) (err error, voter []model.GqaPluginVoteVoter, total int64) {
	pageSize := getVoterList.PageSize
	offset := getVoterList.PageSize * (getVoterList.Page - 1)
	var voterList []model.GqaPluginVoteVoter
	var db *gorm.DB
	db = global.GqaDb.Model(&model.GqaPluginVoteVoter{})
	//配置搜索
	if getVoterList.Voter != "" {
		db = db.Where("title like ?", "%"+getVoterList.Voter+"%")
	}
	if getVoterList.VoteType != "" {
		db = db.Where("vote_type like ?", "%"+getVoterList.VoteType+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn("voter", false)).Preload("VoterByUser").Find(&voterList).Error
	return err, voterList, total
}

func VoterAdd(toAddVoter *model.RequestAddVoter, username string) (err error) {
	var addVoter []model.GqaPluginVoteVoter
	for _, v := range toAddVoter.Voter {
		av := model.GqaPluginVoteVoter{
			Voter:   v,
			VoteType: toAddVoter.VoteType,
			VoterByUser: system.SysUser{Username: v},
			CreatedBy: username,
		}
		addVoter = append(addVoter, av)
	}
	if len(addVoter) != 0 {
		err = global.GqaDb.Model(&model.GqaPluginVoteVoter{}).Save(&addVoter).Error
		return err
	}else{
		return errors.New("本次操作没有影响！")
	}
}

func VoterDelete(id uint, username string) (err error) {
	db := global.GqaDb.Model(&model.GqaPluginVoteVoter{})
	var candidate model.GqaPluginVoteVoter
	if err = db.Where("id = ?", id).First(&candidate).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&candidate).Error
	return err
}

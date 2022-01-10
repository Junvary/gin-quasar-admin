package private_service

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"gorm.io/gorm"
)

func CandidateList(getCandidateList model.RequestCandidateList, username string) (err error, candidate []model.GqaPluginVoteCandidate, total int64) {
	pageSize := getCandidateList.PageSize
	offset := getCandidateList.PageSize * (getCandidateList.Page - 1)
	var candidateList []model.GqaPluginVoteCandidate
	var db *gorm.DB
	db = global.GqaDb.Model(&model.GqaPluginVoteCandidate{})
	//配置搜索
	if getCandidateList.Candidate != "" {
		db = db.Where("candidate like ?", "%"+getCandidateList.Candidate+"%")
	}
	if getCandidateList.VoteType != "" {
		db = db.Where("vote_type like ?", "%"+getCandidateList.VoteType+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn("candidate", false)).Preload("CandidateByUser").Find(&candidateList).Error
	return err, candidateList, total
}

func CandidateAdd(toAddCandidate *model.RequestAddCandidate, username string) (err error) {
	var addCandidate []model.GqaPluginVoteCandidate
	for _, v := range toAddCandidate.Candidate {
		ac := model.GqaPluginVoteCandidate{
			Candidate:   v,
			VoteType: toAddCandidate.VoteType,
			CandidateByUser: system.SysUser{Username: v},
			CreatedBy: username,
		}
		addCandidate = append(addCandidate, ac)
	}
	if len(addCandidate) != 0 {
		err = global.GqaDb.Model(&model.GqaPluginVoteCandidate{}).Save(&addCandidate).Error
		return err
	}else{
		return errors.New("本次操作没有影响！")
	}
}

func CandidateDelete(id uint, username string) (err error) {
	db := global.GqaDb.Model(&model.GqaPluginVoteCandidate{})
	var candidate model.GqaPluginVoteCandidate
	if err = db.Where("id = ?", id).First(&candidate).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&candidate).Error
	return err
}
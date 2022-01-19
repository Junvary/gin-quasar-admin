package private_service

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"gorm.io/gorm"
	"time"
)

func VoteHandle(toAddScore *model.RequestAddScore, username string) (err error) {
	voteType := toAddScore.RequestAddScoreDetail[0].VoteType
	voteMonth := time.Now().Format("200601")
	var alreadyVote int64
	if err = global.GqaDb.Where("vote_type = ? and vote_month = ? and vote_from = ?", voteType, voteMonth, username).
		Find(&model.GqaPluginVoteScoreResult{}).Count(&alreadyVote).Error; err != nil {
		return err
	}
	if alreadyVote != 0 {
		return errors.New("本轮你已经投过票了！")
	}
	var addScore []model.GqaPluginVoteScoreResult
	for _, v := range toAddScore.RequestAddScoreDetail {
		as := model.GqaPluginVoteScoreResult{
			Candidate:       v.Candidate,
			CandidateByUser: system.SysUser{Username: v.Candidate},
			VoteType:        v.VoteType,
			VoteTypeDetail:  v.VoteTypeDetail,
			VoteScore:       v.VoteScore,
			VoteFrom:        username,
			VoteMonth:       time.Now().Format("200601"),
		}
		addScore = append(addScore, as)
	}
	if len(addScore) != 0 {
		err = global.GqaDb.Model(&model.GqaPluginVoteScoreResult{}).Save(&addScore).Error
		return err
	} else {
		return errors.New("本次操作没有影响！")
	}
}

func CanVote(username string) (err error, canVoteResultList interface{}) {
	voteMonth := time.Now().Format("200601")
	type canVote struct {
		V1 bool `json:"v1"`
		V2 bool `json:"v2"`
	}
	var canVoteResult canVote
	var alreadyVote int64
	if err = global.GqaDb.Where("vote_type = ? and vote_month = ? and vote_from = ?", "v1", voteMonth, username).
		Find(&model.GqaPluginVoteScoreResult{}).Count(&alreadyVote).Error; err != nil {
		return err, nil
	}
	var userInVoter bool
	userInVoter = checkUserInVoterRandom(username, "v1")
	if alreadyVote != 0 || !userInVoter{
		canVoteResult.V1 = false
	} else {
		canVoteResult.V1 = true
	}
	if err = global.GqaDb.Where("vote_type = ? and vote_month = ? and vote_from = ?", "v2", voteMonth, username).
		Find(&model.GqaPluginVoteScoreResult{}).Count(&alreadyVote).Error; err != nil {
		return err, nil
	}
	userInVoter = checkUserInVoterRandom(username, "v2")
	if alreadyVote != 0 || !userInVoter {
		canVoteResult.V2 = false
	} else {
		canVoteResult.V2 = true
	}

	return err, canVoteResult
}

func VoteResultList(getVoteScoreList model.RequestScoreResultList, username string) (err error, score []model.GqaPluginVoteScoreResult, total int64) {
	pageSize := getVoteScoreList.PageSize
	offset := getVoteScoreList.PageSize * (getVoteScoreList.Page - 1)
	var voteScoreList []model.GqaPluginVoteScoreResult
	var db *gorm.DB
	db = global.GqaDb.Model(&model.GqaPluginVoteScoreResult{})
	//配置搜索
	if getVoteScoreList.VoteMonth != "" {
		db = db.Where("vote_month = ?", getVoteScoreList.VoteMonth)
	}
	if getVoteScoreList.VoteType != "" {
		db = db.Where("vote_type like ?", "%"+getVoteScoreList.VoteType+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn("vote_month", true)).
		Order(global.OrderByColumn("candidate", false)).
		Preload("CandidateByUser").Preload("VoteFromByUser").Find(&voteScoreList).Error
	return err, voteScoreList, total
}

func VoteResultChart(getVoteScoreList model.RequestScoreResultList, username string) (err error, score []model.GqaPluginVoteScoreResult, total int64) {
	pageSize := getVoteScoreList.PageSize
	offset := getVoteScoreList.PageSize * (getVoteScoreList.Page - 1)
	var voteScoreList []model.GqaPluginVoteScoreResult
	var db *gorm.DB
	db = global.GqaDb.Model(&model.GqaPluginVoteScoreResult{})
	//配置搜索
	if getVoteScoreList.VoteMonth != "" {
		db = db.Where("vote_month = ?", getVoteScoreList.VoteMonth)
	}
	if getVoteScoreList.VoteType != "" {
		db = db.Where("vote_type like ?", "%"+getVoteScoreList.VoteType+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Select("id, candidate, vote_type, vote_type_detail, avg(vote_score) as vote_score, vote_from, vote_month, created_at").
		Group("candidate").Group("vote_type_detail").
		Limit(pageSize).Offset(offset).Order(global.OrderByColumn("vote_month", true)).
		Order(global.OrderByColumn("candidate", false)).
		Preload("CandidateByUser").Preload("VoteFromByUser").Find(&voteScoreList).Error
	return err, voteScoreList, total
}

func checkUserInVoterRandom(username string, voteType string) (ifIn bool) {
	var userInVoter model.GqaPluginVoteVoterRandom
	var userInVoterList []model.GqaPluginVoteVoterRandom
	global.GqaDb.Where("vote_type = ?", voteType).Order(global.OrderByColumn("created_at", true)).First(&userInVoter)
	voteVersion := userInVoter.VoteVersion
	global.GqaDb.Where("vote_version = ?", voteVersion).Find(&userInVoterList)
	for _, v := range userInVoterList {
		if v.Voter == username {
			return true
		}
	}
	return false
}

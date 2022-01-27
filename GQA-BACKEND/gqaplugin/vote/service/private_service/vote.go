package private_service

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"gorm.io/gorm"
	"strconv"
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
	var voter model.GqaPluginVoteVoter
	//global.GqaDb.
	//	Where("voter = ? and vote_type = ? and vote_ratio != ? and vote_ratio != ?", username, voteType, "dy_zhenggong", "dy_jijian").
	//	First(&voter)
	global.GqaDb.
		Where("voter = ? and vote_type = ? ", username, voteType).
		First(&voter)
	ratio := voter.VoteRatio
	var dictRatio system.SysDict
	result := global.GqaDb.Where("dict_code = ?", ratio).First(&dictRatio)
	var voteRatio string
	var ratio8 uint8
	//如果没有找到记录，说明是随机选取的人员，那么按照投票分类的指定投票权重分配
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		var ramdomVoterResult system.SysDict
		if voteType == "dy" {
			if err = global.GqaDb.Where("dict_ext2 = ?", "dy_base").First(&ramdomVoterResult).Error; err!=nil{
				return err
			}
		} else {
			if err = global.GqaDb.Where("dict_ext2 = ?", "gl_base").First(&ramdomVoterResult).Error; err !=nil{
				return err
			}
		}
		voteRatio = ramdomVoterResult.DictCode
		trueRatio := ramdomVoterResult.DictExt1
		trueRatio64, _ := strconv.ParseUint(trueRatio, 10, 64)
		ratio8 = uint8(trueRatio64)
	} else {
		//如果找到了，那么去字典查一下，这个人的ratio类型，获取投票权重
		voteRatio = dictRatio.DictCode
		trueRatio := dictRatio.DictExt1
		trueRatio64, _ := strconv.ParseUint(trueRatio, 10, 64)
		ratio8 = uint8(trueRatio64)
	}
	var addScore []model.GqaPluginVoteScoreResult
	for _, v := range toAddScore.RequestAddScoreDetail {
		as := model.GqaPluginVoteScoreResult{
			Candidate:       v.Candidate,
			CandidateByUser: system.SysUser{Username: v.Candidate},
			VoteType:        v.VoteType,
			VoteTypeDetail:  v.VoteTypeDetail,
			VoteScore:       v.VoteScore,
			VoteRatio:       voteRatio,
			Ratio:           ratio8,
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
		Dy bool `json:"dy"`
		Gl bool `json:"gl"`
	}
	var canVoteResult canVote
	var alreadyVote int64
	if err = global.GqaDb.Where("vote_type = ? and vote_month = ? and vote_from = ?", "dy", voteMonth, username).
		Find(&model.GqaPluginVoteScoreResult{}).Count(&alreadyVote).Error; err != nil {
		return err, nil
	}
	var userInVoter bool
	userInVoter = checkUserInVoterRandom(username, "dy")
	if alreadyVote != 0 || !userInVoter {
		canVoteResult.Dy = false
	} else {
		canVoteResult.Dy = true
	}
	if err = global.GqaDb.Where("vote_type = ? and vote_month = ? and vote_from = ?", "gl", voteMonth, username).
		Find(&model.GqaPluginVoteScoreResult{}).Count(&alreadyVote).Error; err != nil {
		return err, nil
	}
	userInVoter = checkUserInVoterRandom(username, "gl")
	if alreadyVote != 0 || !userInVoter {
		canVoteResult.Gl = false
	} else {
		canVoteResult.Gl = true
	}
	return err, canVoteResult
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

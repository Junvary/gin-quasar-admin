package private_service

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func VoterRandomList(getVoterRandomList model.RequestVoterRandomList, username string) (err error, voter []model.GqaPluginVoteVoterRandom, total int64) {
	pageSize := getVoterRandomList.PageSize
	offset := getVoterRandomList.PageSize * (getVoterRandomList.Page - 1)
	var voterRandomList []model.GqaPluginVoteVoterRandom
	var db *gorm.DB
	db = global.GqaDb.Model(&model.GqaPluginVoteVoterRandom{})
	//配置搜索
	if getVoterRandomList.Voter != "" {
		db = db.Where("voter like ?", "%"+getVoterRandomList.Voter+"%")
	}
	if getVoterRandomList.VoteType != "" {
		db = db.Where("vote_type like ?", "%"+getVoterRandomList.VoteType+"%")
	}
	if getVoterRandomList.VoteVersion.String() != "00000000-0000-0000-0000-000000000000" {
		db = db.Where("vote_version like ?", "%"+getVoterRandomList.VoteVersion.String()+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn("created_at", true)).Preload("VoterByUser").Find(&voterRandomList).Error
	return err, voterRandomList, total
}

func VoterRandomGet(getVoterRandomNumber model.RequestGetVoterRandom, username string) (err error, voter []system.SysUser, total int64) {
	//选出这个投票类型的固定投票人
	var baseVoterList []model.GqaPluginVoteVoter
	if err = global.GqaDb.Where("vote_type = ?", getVoterRandomNumber.VoteType).Find(&baseVoterList).Error; err != nil {
		return err, nil, 0
	}
	//计算固定投票人数量
	baseVoterNumber := len(baseVoterList)
	//如果随机数为0，那么直接返回
	if getVoterRandomNumber.RandomNumber == 0 {
		return nil, nil, 0
	} else {
		//找出所有用户
		var allUserList []system.SysUser
		if err = global.GqaDb.Where("username != ?", "admin").Find(&allUserList).Error; err != nil {
			return err, nil, 0
		}
		//计算所有用户数量
		totalUser := len(allUserList)
		//如果所有用户数量-固定投票人数量<想要随机的数量，报错
		if totalUser-baseVoterNumber < int(getVoterRandomNumber.RandomNumber) {
			return errors.New("随机选择用户数量大于剩余用户数量了！"), nil, 0
		}
		//取出固定投票人的用户名
		var baseVoterUsernameList []string
		for _, v := range baseVoterList {
			baseVoterUsernameList = append(baseVoterUsernameList, v.Voter)
		}
		//取出剩余用户
		db := global.GqaDb.Where("username != ? and username not in ?", "admin", baseVoterUsernameList).Find(&system.SysUser{})
		//按数量随机选取用户
		var surplusUser []system.SysUser
		err = db.Order("rand()").Limit(int(getVoterRandomNumber.RandomNumber)).Find(&surplusUser).Error
		return err, surplusUser, int64(getVoterRandomNumber.RandomNumber)
	}
}

func VoterRandomAdd(toAddVoterRandom *model.RequestAddVoterRandom, username string) (err error) {
	var addVoterRandom []model.GqaPluginVoteVoterRandom
	var randomId = uuid.New()
	for _, v := range toAddVoterRandom.Voter {
		avr := model.GqaPluginVoteVoterRandom{
			Voter:       v,
			VoterByUser: system.SysUser{Username: v},
			VoteType:    toAddVoterRandom.VoteType,
			VoteVersion: randomId,
			Memo:        toAddVoterRandom.Memo,
			CreatedBy:   username,
		}
		addVoterRandom = append(addVoterRandom, avr)
	}
	if len(addVoterRandom) != 0 {
		err = global.GqaDb.Model(&model.GqaPluginVoteVoterRandom{}).Save(&addVoterRandom).Error
		return err
	} else {
		return errors.New("本次操作没有影响！")
	}
}

func VoterRandomDelete(id uint, username string) (err error) {
	db := global.GqaDb.Model(&model.GqaPluginVoteVoterRandom{})
	var candidate model.GqaPluginVoteVoterRandom
	if err = db.Where("id = ?", id).First(&candidate).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&candidate).Error
	return err
}

package model

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/google/uuid"
	"time"
)

type GqaPluginVoteVoterRandom struct {
	Id            uint           `json:"id" gorm:"comment:id;autoIncrement;index"`
	Voter         string         `json:"voter" gorm:"primaryKey;comment:随机加固定投票人;not null;index"`
	VoterByUser   system.SysUser `json:"voterByUser" gorm:"foreignKey:Voter;references:Username"`
	VoteType      string         `json:"voteType" gorm:"primaryKey;comment:投票类型;index"`
	VoteVersion   uuid.UUID      `json:"voteVersion" gorm:"primaryKey;comment:投票版本"`
	Memo          string         `json:"memo" gorm:"comment:投票说明"`
	CreatedAt     time.Time      `json:"createdAt"`
	CreatedBy     string         `json:"createdBy" gorm:"comment:创建人"`
	CreatedByUser system.SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
}

type RequestGetVoterRandom struct {
	RandomNumber uint   `json:"randomNumber"`
	VoteType     string `json:"voteType"`
}

type RequestAddVoterRandom struct {
	Voter    []string `json:"voter"`
	VoteType string   `json:"voteType"`
	Memo     string   `json:"memo"`
}

type RequestVoterRandomList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddVoterRandom 中的字段
	Voter       string    `json:"voter"`
	VoteType    string    `json:"voteType"`
	VoteVersion uuid.UUID `json:"voteVersion"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//GqaPluginVoteVoterRandom
}

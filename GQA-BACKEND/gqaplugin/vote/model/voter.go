package model

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
)

type GqaPluginVoteVoter struct {
	Id          uint           `json:"id" gorm:"comment:id;autoIncrement;index"`
	Voter       string         `json:"voter" gorm:"primaryKey;comment:固定投票人;not null;index"`
	VoterByUser system.SysUser `json:"voterByUser" gorm:"foreignKey:Voter;references:Username"`
	VoteType    string         `json:"voteType" gorm:"primaryKey;comment:投票类型;index"`
	VoteRatio   string         `json:"voteRatio" gorm:"primaryKey;comment:投票权重"`
	CreatedBy   string         `json:"createdBy" gorm:"comment:创建人"`
}

type RequestAddVoter struct {
	Voter     []string `json:"voter"`
	VoteType  string   `json:"voteType"`
	VoteRatio string   `json:"voteRatio"`
}

type RequestVoterList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddVoter 中的字段
	Voter     string `json:"voter"`
	VoteType  string `json:"voteType"`
	VoteRatio string `json:"voteRatio"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//GqaPluginVoteVoter
}

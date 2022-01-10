package model

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
)

type GqaPluginVoteCandidate struct {
	Id              uint           `json:"id" gorm:"comment:id;autoIncrement;index"`
	Candidate       string         `json:"candidate" gorm:"primaryKey;comment:候选人;index;not null"`
	CandidateByUser system.SysUser `json:"candidateByUser" gorm:"foreignKey:Candidate;references:Username"`
	VoteType        string         `json:"voteType" gorm:"primaryKey;comment:投票类型;index"`
	CreatedBy       string         `json:"createdBy" gorm:"comment:创建人"`
}

type RequestAddCandidate struct {
	Candidate []string `json:"candidate"`
	VoteType  string   `json:"voteType"`
}

type RequestCandidateList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddCandidate 中的字段
	Candidate string `json:"candidate"`
	VoteType  string `json:"voteType"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//GqaPluginVoteCandidate
}

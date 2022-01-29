package model

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"time"
)

type GqaPluginVoteScoreResult struct {
	Id              uint           `json:"id" gorm:"comment:id;autoIncrement;index"`
	Candidate       string         `json:"candidate" gorm:"primaryKey;comment:候选人;not null;index"`
	CandidateByUser system.SysUser `json:"candidateByUser" gorm:"foreignKey:Candidate;references:Username"`
	VoteType        string         `json:"voteType" gorm:"primaryKey;comment:投票类型;type:varchar(10)"`
	VoteTypeDetail  string         `json:"voteTypeDetail" gorm:"primaryKey;comment:投票细类;type:varchar(20)"`
	VoteScore       float64        `json:"voteScore" gorm:"comment:得分"`
	VoteRatio       string         `json:"voteRatio" gorm:"comment:投票权重;not null"`
	Ratio           uint8          `json:"ratio" gorm:"comment:权重值;not null"`
	VoteFrom        string         `json:"voteFrom" gorm:"primaryKey;comment:投票人;type:varchar(15)"`
	VoteFromByUser  system.SysUser `json:"voteFromByUser" gorm:"foreignKey:VoteFrom;references:Username"`
	VoteMonth       string         `json:"voteMonth" gorm:"primaryKey;comment:投票时段;index;type:varchar(10)"`
	CreatedAt       time.Time      `json:"createdAt"`
}

type RequestAddScore struct {
	RequestAddScoreDetail []RequestAddScoreDetail `json:"requestAddScoreDetail"`
}

type RequestAddScoreDetail struct {
	Candidate      string  `json:"candidate"`
	VoteType       string  `json:"voteType"`
	VoteTypeDetail string  `json:"voteTypeDetail"`
	VoteScore      float64 `json:"voteScore"`
}

type RequestScoreResultList struct {
	global.RequestPageAndSort
	VoteType  string `json:"voteType"`
	VoteMonth string `json:"voteMonth"`
}

package wf

import "time"

type Sys_wfccinfo struct {
	WaiteID    string    `xorm:"WaiteID" json:"WaiteID"`
	Pid        int       `xorm:"Pid" json:"Pid"`
	StepID     string    `xorm:"stepid" json:"stepid"`
	CCUserID   string    `xorm:"ccuserid" json:"ccuserid"`
	ActionName string    `xorm:"actionname" json:"actionname"`
	OpUid      int       `xorm:"OpUid" json:"OpUid"`
	FlowID     string    `xorm:"FlowID" json:"FlowID"`
	OpDate     time.Time `xorm:"OpDate" json:"OpDate"`
	OpInputMsg string    `xorm:"OpInputMsg" json:"OpInputMsg"`
	NewGuid    string    `xorm:"NewGuid" json:"NewGuid"`
	EntId      int       `xorm:"entid" json:"entid"`
	IsEnd      int       `xorm:"isend" json:"isend"`
}

func (*Sys_wfccinfo) TableName() string {
	return "sys_wfccinfo"
}

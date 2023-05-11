package wf

type Sys_wfstepaccessrule2 struct {
	Acc2ID     int    `xorm:"acc2id" json:"acc2id"`
	AccID      string `xorm:"accid" json:"accid"`
	AccPivotID string `xorm:"accpivotid" json:"accpivotid"`
	Pid        int    `xorm:"pid" json:"pid"`
	GridId     int    `xorm:"gridid" json:"gridid"`
	IsMain     int    `xorm:"ismain" json:"ismain"`
	MainCol    string `xorm:"maincol" json:"maincol"` // 例：main.userid,main.xxuserid
	DSPid      int    `xorm:"dspid" json:"dspid"`
	DSSqlTable string `xorm:"dssqltable" json:"dssqltable"`
	DSCol      string `xorm:"dscol" json:"dscol"`
	StepId     string `xorm:"stepid" json:"stepid"`
	CType      string `xorm:"ctype" json:"ctype"`
	EntId      int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfstepaccessrule2) TableName() string {
	return "sys_wfstepaccessrule2"
}

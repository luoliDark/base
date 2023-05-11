package sysmodel

type Sys_wfccuser struct {
	AccID      int    `xorm:"AccID" json:"AccID"`
	StepID     string `xorm:"StepID" json:"StepID"`
	CCUserID   int    `xorm:"CCUserID" json:"CCUserID"`
	EntId      int    `xorm:"entid" json:"entid"`
	ActionName string `xorm:"ActionName" json:"ActionName"`
}

func (*Sys_wfccuser) TableName() string {
	return "sys_wfccuser"
}

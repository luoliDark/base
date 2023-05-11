package wf

type Sys_wfccuser struct {
	AccID    int    `xorm:"accid" json:"accid"`
	StepID   string `xorm:"stepid" json:"stepid"`
	CCUserID string `xorm:"ccuserid" json:"ccuserid"`
	Flowid   string `xorm:"flowid" json:"flowid"`
	EntId    int    `xorm:"entid" json:"entid"`
	Pid      int    `xorm:"pid" json:"pid"`
}

func (*Sys_wfccuser) TableName() string {
	return "sys_wfccuser"
}

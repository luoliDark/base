package wf

type Sys_wfneeduserchoice struct {
	Accobjid string `xorm:"accobjid" json:"accobjid"`
	Acctype  string `xorm:"acctype" json:"acctype"`
	Stepid   string `xorm:"stepid" json:"stepid"`
	Pid      int    `xorm:"pid" json:"pid"`
}

func (*Sys_wfneeduserchoice) TableName() string {
	return "sys_wfneeduserchoice"
}

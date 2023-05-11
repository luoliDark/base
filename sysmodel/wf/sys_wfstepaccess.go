package wf

type Sys_wfstepaccess struct {
	Accid    int    `xorm:"accid" json:"accid"`
	Stepid   string `xorm:"stepid" json:"stepid"`
	Accobjid string `xorm:"accobjid" json:"accobjid"`
	Acctype  string `xorm:"acctype" json:"acctype"`
	EntId    int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfstepaccess) TableName() string {
	return "Sys_wfstepaccess"
}

package wf

type Sys_wfexbillnomap struct {
	Id        string `xorm:"id" json:"id"`
	Pid       int    `xorm:"pid" json:"pid"`
	BillId    string `xorm:"billid" json:"billid"`
	ExWaiteId string `xorm:"exwaiteid" json:"exwaiteid"` // 外部流程实例ID
	EntId     int    `xorm:"entid" json:"entid"`
	Ver       string `xorm:"ver" json:"ver"`
}

func (*Sys_wfexbillnomap) TableName() string {
	return "sys_wfexbillnomap"
}

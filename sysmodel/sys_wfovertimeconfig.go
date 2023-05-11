package sysmodel

type Sys_wfovertimeconfig struct {
	Id              int    `xorm:"id" json:"id"`
	EntID           string `xorm:"entid" json:"entid"`
	StepAttr        string `xorm:"stepattr" json:"stepattr"`
	OverTime_Number int    `xorm:"overtime_number" json:"overtime_number"`
}

func (this *Sys_wfovertimeconfig) TableName() string {
	return "sys_wfovertimeconfig"
}

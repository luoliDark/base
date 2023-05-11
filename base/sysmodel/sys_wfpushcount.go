package sysmodel

import "time"

type Sys_wfpushcount struct {
	Id          int       `xorm:"id" json:"id"`
	EntID       string    `xorm:"entid" json:"entid"`
	StepAttr    string    `xorm:"stepattr" json:"stepattr"`
	Pid         int       `xorm:"pid" json:"pid"`
	BillID      string    `xorm:"billid" json:"billid"`
	WaiteID     string    `xorm:"waiteid" json:"waiteid"`
	StepId      string    `xorm:"stepid" json:"stepid"`
	NewGuid     string    `xorm:"newguid" json:"newguid"`
	Expeditor   string    `xorm:"expeditor" json:"expeditor"`
	Urgedperson string    `xorm:"urgedperson" json:"urgedperson"`
	Create_time time.Time `xorm:"create_time" json:"create_time"`
	Modify_time time.Time `xorm:"modify_time" json:"modify_time"`
	Count       int       `xorm:"count" json:"count"`
}

func (this *Sys_wfpushcount) TableName() string {
	return "sys_wfpushcount"
}

package wf

import "time"

type Sys_wfstepdynamic_active struct {
	DynamicID   int       `xorm:"dynamicid pk autoincr " json:"dynamicid"`
	StepID      string    `xorm:"stepid" json:"stepid"`
	IsActive    int       `xorm:"isactive" json:"isactive"`
	FlowID      string    `xorm:"flowid" json:"flowid"`
	WaiteID     string    `xorm:"Waiteid" json:"Waiteid"`
	Create_Date time.Time `xorm:"create_date" json:"create_date"`
	NewGuid     string    `xorm:"newguid" json:"newguid"`
	EntId       int       `xorm:"entid" json:"entid"`
}

func (*Sys_wfstepdynamic_active) TableName() string {
	return "Sys_wfstepdynamic_active"
}

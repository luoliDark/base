package sysmodel

import (
	"time"
)

type Sys_PayPush struct {
	Id          int       `xorm:"id" json:"id"`
	Status      int       `xorm:"status" json:"status"`
	Oldtime     time.Time `xorm:"oldtime" json:"oldtime"`
	Create_time time.Time `xorm:"create_time" json:"create_time"`
}

func (this *Sys_PayPush) TableName() string {
	return "sys_paypush"
}

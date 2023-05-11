package wf

import "time"

/**
代理填单
*/
type Sys_AgentMaker struct {
	ConsinID          int       `xorm:"consinid" json:"consinid"`
	AppUid            string    `xorm:"appuid" json:"appuid"`
	ConsignTargentUid string    `xorm:"consigntargentuid" json:"consigntargentuid"`
	ConsignStatus     string    `xorm:"consignstatus" json:"consignstatus"`
	ConsignBeginTime  time.Time `xorm:"consignbegintime" json:"consignbegintime"`
	ConsignEndTime    time.Time `xorm:"consignendtime" json:"consignendtime"`
	EntId             int       `xorm:"entid" json:"entid"`
}

func (*Sys_AgentMaker) TableName() string {
	return "Sys_AgentMaker"
}

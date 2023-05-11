package wf

import "time"

type Sys_wfstepdynamic struct {
	DynamicID          int       `xorm:"DynamicID pk autoincr " json:"DynamicID"`
	WaiteID            string    `xorm:"WaiteID" json:"WaiteID"`
	FlowID             string    `xorm:"FlowID" json:"FlowID"` // 等同于 targetFlowID
	Pid                int       `xorm:"Pid" json:"Pid"`
	BillID             string    `xorm:"BillID" json:"BillID"`
	SourceStepID       string    `xorm:"SourceStepID" json:"SourceStepID"`
	TargetStepID       string    `xorm:"TargetStepID" json:"TargetStepID"`
	SourceFlowID       string    `xorm:"SourceFlowID" json:"SourceFlowID"`
	TargetFlowID       string    `xorm:"TargetFlowID" json:"TargetFlowID"`
	NewGuid            string    `xorm:"NewGuid" json:"NewGuid"`
	AppUsers           string    `xorm:"AppUsers" json:"AppUsers"`
	TargetIsChildStep  int       `xorm:"TargetIsChildStep" json:"TargetIsChildStep"`
	TargetIsSubmitSkip int       `xorm:"TargetIsSubmitSkip" json:"TargetIsSubmitSkip"`
	InsertDate         time.Time `xorm:"InsertDate" json:"InsertDate"`
	EntId              int       `xorm:"entid" json:"entid"`
	SourceIsUser       int       `xorm:"sourceisuser" json:"sourceisuser"`
	TargetIsUser       int       `xorm:"targetisuser" json:"targetisuser"`
}

func (*Sys_wfstepdynamic) TableName() string {
	return "sys_wfstepdynamic"
}

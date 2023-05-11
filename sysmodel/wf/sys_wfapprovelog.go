package wf

import "time"

type Sys_wfapprovelog struct {
	LogID        int       `xorm:"logid pk autoincr " json:"logid"`
	WaiteID      string    `xorm:"waiteid" json:"waiteid"`
	FlowID       string    `xorm:"flowid" json:"flowid"`
	Pid          int       `xorm:"pid" json:"pid"`
	BillID       string    `xorm:"billid" json:"billid"`
	StepID       string    `xorm:"stepid" json:"stepid"`
	ResultType   string    `xorm:"resulttype" json:"resulttype"`
	AppOpinion   string    `xorm:"appopinion" json:"appopinion"`
	FlowStatus   int       `xorm:"flowstatus" json:"flowstatus"`
	ActionName   string    `xorm:"actionname" json:"actionname"`
	NewGuid      string    `xorm:"newguid" json:"newguid"`
	AppTerminal  string    `xorm:"appterminal" json:"appterminal"`
	ApproveUid   string    `xorm:"approve_uid" json:"approve_uid"`
	BeginDate    time.Time `xorm:"begindate" json:"begindate"`
	ApproveDate  time.Time `xorm:"approve_date" json:"approve_date"`
	Times        int       `xorm:"times" json:"times"`
	IsInvalid    int       `xorm:"isinvalid" json:"isinvalid"`
	DynamicID    int       `xorm:"dynamicid" json:"dynamicid"`
	IsSystem     int       `xorm:"issystem" json:"issystem"`
	OriginAppUid int       `xorm:"originappuid" json:"originappuid"`
	SourceAction string    `xorm:"sourceaction" json:"sourceaction"`
	EntId        int       `xorm:"entid" json:"entid"`
	TargetIsUser int       `xorm:"targetisuser" json:"targetisuser"`
	StepAttr     string    `xorm:"stepattr" json:"stepattr"`
}

func (*Sys_wfapprovelog) TableName() string {
	return "sys_wfapprovelog"
}

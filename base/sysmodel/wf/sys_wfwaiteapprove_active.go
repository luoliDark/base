package wf

type Sys_wfwaiteapprove_active struct {
	DynamicID      int    `xorm:"dynamicid pk autoincr " json:"dynamicid"`
	WaiteID        string `xorm:"waiteid" json:"waiteid"`
	FlowStatus     int    `xorm:"flowstatus" json:"flowstatus"`
	IsReSubmitGoto int    `xorm:"isresubmitgoto" json:"IsReSubmitGoto"`
	ReturnStepID   string `xorm:"returnstepid" json:"ReturnStepID"`
	EntId          int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfwaiteapprove_active) TableName() string {
	return "Sys_wfwaiteapprove_active"
}

package wf

type Sys_wfsubmitflow struct {
	Id             int    `xorm:"id pk autoincr " json:"id"`
	WaiteID        string `xorm:"WaiteID" json:"WaiteID"`
	Pid            int    `xorm:"Pid" json:"Pid"`
	BillID         string `xorm:"BillID" json:"BillID"`
	IsActive       int    `xorm:"IsActive" json:"IsActive"`
	SourceStepID   string `xorm:"SourceStepID" json:"SourceStepID"`
	TargetStepID   string `xorm:"TargetStepID" json:"TargetStepID"`
	SourceStepName string `xorm:"SourceStepName" json:"SourceStepName"`
	TargetStepName string `xorm:"TargetStepName" json:"TargetStepName"`
	Wherestr       string `xorm:"Wherestr" json:"Wherestr"`
	Formdata       string `xorm:"Formdata" json:"Formdata"`
	NewGuid        string `xorm:"NewGuid" json:"NewGuid"`
	EntId          string `xorm:"entid" json:"entid"`
}

func (*Sys_wfsubmitflow) TableName() string {
	return "Sys_wfsubmitflow"
}

func (this *Sys_wfsubmitflow) GetStepFlow(sourceStep, targetStep Sys_wfstep) {
	this.SourceStepName = sourceStep.StepName
	this.SourceStepID = sourceStep.StepID
	this.TargetStepName = targetStep.StepName
	this.TargetStepID = targetStep.StepID
	this.IsActive = 0
}

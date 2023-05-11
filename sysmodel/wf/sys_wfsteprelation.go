package wf

type Sys_wfsteprelation struct {
	RefID             string `xorm:"refid" json:"refid"`
	FlowID            string `xorm:"flowid" json:"flowid"`
	Pid               int    `xorm:"pid" json:"pid"`
	SourceStepID      string `xorm:"sourcestepid" json:"sourcestepid"`
	TargetStepID      string `xorm:"targetstepid" json:"targetstepid"`
	TargetIsChildStep int    `xorm:"targetischildstep" json:"targetischildstep"`
	TargetIsEnd       int    `xorm:"targetisend" json:"targetisend"`
	InWhereSQL        string `xorm:"inwheresql" json:"inwheresql"`
	InWhereText       string `xorm:"inwheretext" json:"inwheretext"`
	InWhereStr        string `xorm:"inwherestr" json:"inwherestr"`
	TargetHtmlDivId   string `xorm:"targethtmldivid" json:"targethtmldivid"`
	SourceHtmlDivId   string `xorm:"sourcehtmldivid" json:"sourcehtmldivid"`
	EntId             int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfsteprelation) TableName() string {
	return "sys_wfsteprelation"
}

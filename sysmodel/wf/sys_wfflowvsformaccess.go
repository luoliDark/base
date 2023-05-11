package wf

//流程vs表单审批权限
type Sys_wfflowvsformaccess struct {
	FlowVsFormAccID string `xorm:"flowvsformaccid" json:"flowvsformaccid"`
	FlowVsFormID    string `xorm:"flowvsformid" json:"flowvsformid"`
	Pid             int    `xorm:"pid" json:"pid"`
	FlowID          string `xorm:"flowid" json:"flowid"`
	AccValue        string `xorm:"accvalue" json:"accvalue"`
	AccType         string `xorm:"acctype" json:"acctype"`
	EntId           int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfflowvsformaccess) TableName() string {
	return "sys_wfflowvsformaccess"
}

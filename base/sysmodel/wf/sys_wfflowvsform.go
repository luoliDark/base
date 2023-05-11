package wf

type Sys_wfflowvsform struct {
	FlowVsFormID  string `xorm:"flowvsformid" json:"flowvsformid"`
	Pid           int    `xorm:"pid" json:"pid"`
	FlowID        string `xorm:"flowid" json:"flowid"`
	FlowName      string `xorm:"flowname" json:"flowname"`
	IsCompPrivate int    `xorm:"iscompprivate" json:"iscompprivate"`
	CompID        string `xorm:"compid" json:"compid"`
	CheckType     string `xorm:"checktype" json:"checktype"`
	FlowAdminUid  string `xorm:"flowadminuid" json:"flowadminuid"`
	SortId        int    `xorm:"sortid" json:"sortid"`
	EntId         int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfflowvsform) TableName() string {
	return "sys_wfflowvsform"
}

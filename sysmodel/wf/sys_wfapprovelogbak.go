package wf

/**
 * @Author: weiyg
 * @Date: 2020/3/18 23:19
 * @describe:审批日志备份表 -结构体
 */
type Sys_wfapprovelog_bak struct {
	LogID        int    `xorm:"logid" json:"logid"`
	WaiteID      int    `xorm:"waiteid" json:"waiteid"`
	FlowID       string `xorm:"flowid" json:"flowid"`
	Pid          int    `xorm:"pid" json:"pid"`
	BillID       string `xorm:"billid" json:"billid"`
	StepID       string `xorm:"stepid" json:"stepid"`
	ResultType   string `xorm:"resulttype" json:"resulttype"`
	AppOpinion   string `xorm:"appopinion" json:"appopinion"`
	FlowStatus   int    `xorm:"flowstatus" json:"flowstatus"`
	ActionName   string `xorm:"actionname" json:"actionname"`
	NewGuid      string `xorm:"newguid" json:"newguid"`
	AppTerminal  string `xorm:"appterminal" json:"appterminal"`
	ApproveUid   int    `xorm:"approve_uid" json:"approve_uid"`
	BeginDate    string `xorm:"begindate" json:"begindate"`
	ApproveDate  string `xorm:"approve_date" json:"approve_date"`
	Times        int    `xorm:"times" json:"times"`
	IsInvalid    int    `xorm:"isinvalid" json:"isinvalid"`
	DynamicID    int    `xorm:"dynamicid" json:"dynamicid"`
	IsSystem     int    `xorm:"issystem" json:"issystem"`
	Originappuid int    `xorm:"originappuid" json:"originappuid"`
	SourceAction string `xorm:"sourceaction" json:"sourceaction"`
	EntId        int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfapprovelog_bak) TableName() string {
	return "sys_wfapprovelog_bak"
}

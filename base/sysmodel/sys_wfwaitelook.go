package sysmodel

type Sys_wfwaitelook struct {
	ID       int    `xorm:"ID" json:"ID"`
	WaiteID  int    `xorm:"WaiteID" json:"WaiteID"`
	Pid      int    `xorm:"Pid" json:"Pid"`
	BillID   int    `xorm:"BillID" json:"BillID"`
	LookUid  int    `xorm:"LookUid" json:"LookUid"`
	LookDate string `xorm:"LookDate" json:"LookDate"`
}

func (*Sys_wfwaitelook) TableName() string {
	return "sys_wfwaitelook"
}

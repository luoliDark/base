package event

type Sys_eventtype struct {
	EventTypeId   string `xorm:"eventtypeid" json:"eventtypeid"`
	EventTypeName string `xorm:"eventtypename" json:"eventtypename"`
	Memo          string `xorm:"memo" json:"memo"`
}

func (*Sys_eventtype) TableName() string {
	return "sys_eventtype"
}

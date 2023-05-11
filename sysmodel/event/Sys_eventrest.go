package event

type Sys_eventrest struct {
	EventID      string `xorm:"eventid" json:"eventid"`
	EventTypeId  string `xorm:"eventtypeid" json:"eventtypeid"`
	EventCode    string `xorm:"eventcode" json:"eventcode"`
	EventName    string `xorm:"eventname" json:"eventname"`
	EventRefCode string `xorm:"eventrefcode" json:"eventrefcode"` // 例：按钮事件 就写按钮编码，审批事件就写stepID
	Pid          int    `xorm:"pid" json:"pid"`
	RestUrl      string `xorm:"resturl" json:"resturl"`
	IsEnable     int    `xorm:"isenable" json:"isenable"`
	SubSysID     int    `xorm:"subsysid" json:"subsysid"`
	IsCheck      int    `xorm:"ischeck" json:"ischeck"`
	Memo         string `xorm:"memo" json:"memo"`
}

func (*Sys_eventrest) TableName() string {
	return "sys_eventrest"
}

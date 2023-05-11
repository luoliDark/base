package print

type Sys_reportfiels struct {
	RepID     string `xorm:"repid" json:"repid"`
	Pid       int    `xorm:"pid" json:"pid"`
	RepName   string `xorm:"repname" json:"repname"`
	IsDefault int    `xorm:"isdefault" json:"isdefault"`
	IsEnable  int    `xorm:"isenable" json:"isenable"`
	RepPath   string `xorm:"reppath" json:"reppath"`
	IsPrintWF int    `xorm:"isprintwf" json:"isprintwf"`
}

func (*Sys_reportfiels) TableName() string {
	return "sys_reportfiels"
}

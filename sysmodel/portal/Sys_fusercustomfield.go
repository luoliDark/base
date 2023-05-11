package portal

type Sys_fusercustomfield struct {
	Id     string `xorm:"id" json:"id"`
	UserId string `xorm:"userid" json:"userid"`
	Pid    string `xorm:"pid" json:"pid"`
	SqlCol string `xorm:"sqlcol" json:"sqlcol"`
	IsShow string `xorm:"isshow" json:"isshow"`
	Ver    string `xorm:"ver" json:"ver"`
	Sortid string `xorm:"sortid" json:"sortid"`
}

func (*Sys_fusercustomfield) TableName() string {
	return "sys_fusercustomfield"
}

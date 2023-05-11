package invoice

type Sys_mobileocrdefaultcol struct {
	Id         int    `xorm:"id" json:"id"`
	Pid        int    `xorm:"pid" json:"pid"`
	Gridid     int    `xorm:"gridid" json:"gridid"`
	SqlField   string `xorm:"sqlfield" json:"sqlfield"`
	DefaultVal string `xorm:"defaultval" json:"defaultval"`
}

func (*Sys_mobileocrdefaultcol) TableName() string {
	return "sys_mobileocrdefaultcol"
}

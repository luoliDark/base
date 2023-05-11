package uitemplete

type Sys_uitabs struct {
	TabGroupID   string `xorm:"tabgroupid" json:"tabgroupid"`
	TabStyleID   int    `xorm:"tabstyleid" json:"tabstyleid"`
	Pid          int    `xorm:"pid" json:"pid"`
	TabGroupName string `xorm:"tabgroupname" json:"tabgroupname"`
	Entid        string `xorm:"entid" json:"entid"`
}

func (*Sys_uitabs) TableName() string {
	return "sys_ftabs"
}

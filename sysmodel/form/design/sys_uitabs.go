package design

type Sys_uitabs struct {
	TabGroupID   string `xorm:"tabgroupid" json:"tabgroupid"`
	TabStyleID   string `xorm:"tabstyleid" json:"tabstyleid"`
	TabGroupName string `xorm:"tabgroupname" json:"tabgroupname"`
	Pid          int    `xorm:"pid" json:"pid"`
	SortID       int    `xorm:"SortID" json:"SortID"`
	EntId        int    `xorm:"entid" json:"entid"`
}

func (*Sys_uitabs) TableName() string {
	return "sys_uitabs"
}

package design

type Sys_uiform struct {
	TempId   int    `xorm:"tempid" json:"tempid"`
	Pid      int    `xorm:"pid" json:"pid"`
	TempName string `xorm:"tempname" json:"tempname"`
	EntId    int    `xorm:"entid" json:"entid"`
}

func (*Sys_uiform) TableName() string {
	return "sys_uiform"
}

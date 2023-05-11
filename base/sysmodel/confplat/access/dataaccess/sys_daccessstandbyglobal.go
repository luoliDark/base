package dataaccess

type Sys_daccessstandbyglobal struct {
	GlobalID   int    `xorm:"globalid pk" json:"globalid"`
	ConfigCode string `xorm:"configcode" json:"configcode"`
	IsOpen     int    `xorm:"isopen" json:"isopen"`
}

func (*Sys_daccessstandbyglobal) TableName() string {
	return "sys_daccessstandbyglobal"
}

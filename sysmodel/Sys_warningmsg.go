package sysmodel

type Sys_WarningMsg struct {
	Id         int    `xorm:"id" json:"id"`
	Pid        int    `xorm:"pid" json:"pid"`
	Billid     string `xorm:"billid" json:"billid"`
	ErrType    string `xorm:"errtype" json:"errtype"`
	Errmsg     string `xorm:"errmsg" json:"errmsg"`
	Insertdate string `xorm:"insertdate" json:"insertdate"`
	Userid     string `xorm:"userid" json:"userid"`
}

func (*Sys_WarningMsg) TableName() string {
	return "sys_warningmsg"
}

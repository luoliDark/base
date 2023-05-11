package ssomodel

type Sys_LogCheckLogin struct {
	Logid       string `xorm:"logid" json:"logid"`
	LogMsg      string `xorm:"logmsg" json:"logmsg"`
	UserId      string `xorm:"userid" json:"userid"`
	InsertDate  string `xorm:"insertdate" json:"insertdate"`
	Sid         string `xorm:"sid" json:"sid"`
	Appterminal string `xorm:"appterminal" json:"appterminal"`
}

func (*Sys_LogCheckLogin) TableName() string {
	return "Sys_LogCheckLogin"
}

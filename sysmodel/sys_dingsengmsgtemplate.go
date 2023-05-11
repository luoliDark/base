package sysmodel

type Sys_dingSengMsgTemplate struct {
	Id              int    `xorm:"id" json:"id"`
	Templatename    string `xorm:"templatename" json:"templatename"`
	Templatecontent string `xorm:"templatecontent" json:"templatecontent"`
	Templatecontype string `xorm:"templatecontype" json:"templatecontype"`
	Code            string `xorm:"code" json:"code"`
	Create_time     string `xorm:"create_time" json:"create_time"`
}

func (*Sys_dingSengMsgTemplate) TableName() string {
	return "sys_dingsengmsgtemplate"
}

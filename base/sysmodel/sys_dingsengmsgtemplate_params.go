package sysmodel

type Sys_dingSengMsgTemplate_Params struct {
	Id              int    `xorm:"id" json:"id"`
	Templatename_id string `xorm:"templatename_id" json:"templatename_id"`
	Param_name      string `xorm:"param_name" json:"param_name"`
	Param_title     string `xorm:"param_title" json:"param_title"`
	Param_type      string `xorm:"param_type" json:"param_type"`
	IsRequired      int    `xorm:"isrequired" json:"isrequired"`
	Remark          string `xorm:"remark" json:"remark"`
	Create_time     string `xorm:"create_time" json:"create_time"`
}

func (*Sys_dingSengMsgTemplate_Params) TableName() string {
	return "sys_dingsengmsgtemplate_params"
}

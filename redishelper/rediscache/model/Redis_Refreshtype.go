package model

type Redis_Refreshtype struct {
	Id      string `xorm:"id" json:"id"`
	CType   string `xorm:"ctype" json:"ctype"`     //刷新类型 WF表示工作流 Form表示表单(用formid) formCopy表示拷备关系 Other 表示其它
	IsByPid int    `xorm:"IsByPid" json:"IsByPid"` //是否按Pid分版本
	IsByEnt int    `xorm:"IsByEnt" json:"IsByEnt"` //是否按客户分版本
}

func (*Redis_Refreshtype) TableName() string {
	return "Redis_Refreshtype"
}

package sysmodel

type Eb_waitask_log struct {
	Id                  int    `xorm:"id" json:"id"`
	Pid                 string `xorm:"pid" json:"pid"`
	PrimaryKey          string `xorm:"primaryKey" json:"primaryKey"`
	Stepid              string `xorm:"stepid" json:"stepid"`
	Taskid              string `xorm:"taskid" json:"taskid"`
	Entid               int    `xorm:"entid" json:"entid"`
	Process_instance_id string `xorm:"process_instance_id" json:"process_instance_id"` //实例id
	UserId              string `xorm:"userId" json:"userId"`                           //待办消息接收人
	Create_UID          string `xorm:"create_uid" json:"create_uid"`
	Create_Date         string `xorm:"create_Date" json:"create_Date"`
	Update_UID          string `xorm:"update_UID" json:"update_UID"`
	Update_Date         string `xorm:"update_Date" json:"update_Date"`
	IsDiscard           string `xorm:"isDiscard" json:"isDiscard"`
}

func (*Eb_waitask_log) TableName() string {
	return "eb_waitask_log"
}

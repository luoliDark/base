package log

type Sys_log_error struct {
	Id           int    `xorm:"id" json:"id"`
	Logtext      string `xorm:"logtext" json:"logtext"`             // 日志内容
	Logtype      string `xorm:"logtype" json:"logtype"`             // 日志类型
	SourceSystem string `xorm:"source_system" json:"source_system"` // 来源系统
	Userid       string `xorm:"userid" json:"userid"`
	InsertDate   string `xorm:"insert_date" json:"insert_date"`
}

func (*Sys_log_error) TableName() string {
	return "sys_log_error"
}
